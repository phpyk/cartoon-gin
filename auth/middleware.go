package auth

import (
	"cartoon-gin/utils"
	"cartoon-gin/dao"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

const SECRET_KEY = "yuekai"

//自定义生成token所需的clamis，讲UID放入
type MyClaims struct {
	UID uint
	jwt.StandardClaims
}

func GenerateToken(user *dao.User) (string, error) {
	claims := make(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24*7)).Unix()
	claims["iat"] = time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(SECRET_KEY))
	return tokenString, err
}

func ValidateToken(next http.Handler) http.HandlerFunc {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("authorization")
		if tokenStr == "" {
			responseNotAuthorized(w)
			return
		} else {
			/**
			 * tokenStr解析成token对象
			 */
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					responseNotAuthorized(w)
					return
				}
				return []byte(SECRET_KEY), nil
			})
			if !token.Valid {
				responseNotAuthorized(w)
				return
			}
			next.ServeHTTP(w, r)
		}
	})
}

func ValidateTokenV2() gin.HandlerFunc {
	return func(c *gin.Context) {
		cg := utils.Gin{C: c}

		tokenStr := c.Request.Header.Get("Authorization")
		tokenStr = tokenStr[7:]
		if tokenStr == "" {
			cg.UnAuthorized()
		} else {
			/**
			 * tokenStr解析成token对象
			 */
			token, _ := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (i interface{}, e error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					cg.UnAuthorized()
				}
				return []byte(SECRET_KEY), nil
			})
			if !token.Valid {
				cg.UnAuthorized()
			}
			//claims interface 转为 mycliaims
			myClaims := token.Claims.(*MyClaims)
			log.Printf("claims: %+v \n", myClaims)
			c.Set("uid", myClaims.UID)
			c.Next()
		}
	}

}

func responseNotAuthorized(w http.ResponseWriter) {
	response := utils.Response{State: 0, Message: "You are unauthorized"}
	utils.ResponseWithJson(w, http.StatusUnauthorized, response)
	return
}


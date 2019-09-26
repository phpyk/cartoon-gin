package auth

import (
	"cartoon-gin/common"
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
		} else {
			/**
			 * tokenStr解析成token对象
			 */
			token, _ := jwt.Parse(tokenStr, func(token *jwt.Token) (i interface{}, e error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					responseNotAuthorized(w)
				}
				return []byte(SECRET_KEY), nil
			})
			if !token.Valid {
				responseNotAuthorized(w)
			}
			next.ServeHTTP(w, r)
		}
	})
}

func ValidateTokenV2() gin.HandlerFunc {
	return func(c *gin.Context) {
		cg := common.Gin{C: c}

		tokenStr := c.Request.Header.Get("Authorization")
		tokenStr = tokenStr[7:]
		if tokenStr == "" {
			responseNotAuthorizedV2(&cg)
		} else {
			/**
			 * tokenStr解析成token对象
			 */
			token, _ := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (i interface{}, e error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					responseNotAuthorizedV2(&cg)
				}
				return []byte(SECRET_KEY), nil
			})
			if !token.Valid {
				responseNotAuthorizedV2(&cg)
			}
			//这里这样调用时什么意思，我也不知道
			myClaims := token.Claims.(*MyClaims)
			log.Printf("claims: %+v \n", myClaims)
			c.Set("uid", myClaims.UID)
			c.Next()
		}
	}

}

func responseNotAuthorized(w http.ResponseWriter) {
	response := common.Response{State: 0, Message: "You are unauthorized"}
	common.ResponseWithJson(w, http.StatusUnauthorized, response)
}
func responseNotAuthorizedV2(cg *common.Gin) {
	myResponse := common.Response{State: 0, Message: "You are unauthorized", Result: nil}
	cg.Response(http.StatusUnauthorized, myResponse)
}

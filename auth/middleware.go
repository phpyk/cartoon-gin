package auth

import (
	"cartoon-gin/utils"
	"cartoon-gin/dao"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"encoding/json"
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

func GenerateRedisToken(user *dao.User) (token string,err error) {
	tokenKey := utils.RandomString(32,0)
	tokenValue,err := json.Marshal(user)
	expireTime := 30 * time.Hour * time.Duration(24 * 365)
	utils.CheckError(err)

	clt := utils.NewAuthRedisClient()
	sts := clt.Set(tokenKey,tokenValue,expireTime)
	res,err := sts.Result()
	if res != "OK" || err != nil {
		return tokenKey,err
	}
	return tokenKey,nil
}

func CheckRedisToken(token string) (userJson string, err error) {
	clt := utils.NewAuthRedisClient()
	sts := clt.Get(token)
	res,err := sts.Result()
	if err != nil {
		return "",err
	}
	return res,nil
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

func ValidateJWTToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		cg := utils.Gin{C: c}

		tokenStr := c.Request.Header.Get("Authorization")
		tokenStr = tokenStr[7:]
		if tokenStr == "" {
			cg.UnAuthorized()
			return
		} else {
			/**
			 * tokenStr解析成token对象
			 */
			token, _ := jwt.ParseWithClaims(tokenStr, &MyClaims{}, func(token *jwt.Token) (i interface{}, e error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					cg.UnAuthorized()
					return
				}
				return []byte(SECRET_KEY), nil
			})
			if !token.Valid {
				cg.UnAuthorized()
				return
			}
			//claims interface 转为 mycliaims
			myClaims := token.Claims.(*MyClaims)
			log.Printf("claims: %+v \n", myClaims)
			user := dao.UserFindByID(myClaims.UID)
			c.Set("user",&user)
			//c.Set("uid", myClaims.UID)
			c.Next()
		}
	}
}

func ValidateRedisToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		cg := utils.Gin{C: c}
		tokenStr := c.Request.Header.Get("Authorization")
		if tokenStr == "" || len(tokenStr) <= 7 {
			cg.UnAuthorized()
			return
		} else {
			tokenStr = tokenStr[7:]
			userJson,err := CheckRedisToken(tokenStr)
			if err != nil {
				cg.UnAuthorized()
				return
			}
			var user dao.User
			err = json.Unmarshal([]byte(userJson),&user)
			if err != nil {
				cg.UnAuthorized()
				return
			}
			log.Printf("logined user: %+v \n", user)
			c.Set("user",&user)
			c.Next()
		}
	}
}


func responseNotAuthorized(w http.ResponseWriter) {
	response := utils.Response{State: 0, Message: "You are unauthorized"}
	utils.ResponseWithJson(w, http.StatusUnauthorized, response)
	return
}


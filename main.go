package main

import (
	"cartoon-gin/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"net/http"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()


	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK,"hello world")
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	// Authorized group (uses gin.BasicAuth() middleware)
	// Same than:
	// authorized := r.Group("/")
	// authorized.Use(gin.BasicAuth(gin.Credentials{
	//	  "foo":  "bar",
	//	  "manu": "123",
	//}))
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// Parse JSON
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	//var cats []models.Category
	//cats = models.GetAllCategories()
	//fmt.Println(cats)
	//res := models.AddCategory("testtest",1)
	//fmt.Println(res)


	db, _ := gorm.Open("mysql","root:hajgv8t24oA9@(123.206.107.76:3306)/news?charset=utf8&parseTime=True")
	//db.AutoMigrate(&models.Category{})
	cat := models.Category{CateName:"ttttt",ParentId:0}
	res := db.NewRecord(cat)
	db.Create(&cat)
	//db.Find(&cats)
	fmt.Println(res)


	db2, _ := gorm.Open("mysql","root:hajgv8t24oA9@(123.206.107.76:3306)/cartoon?charset=utf8&parseTime=True")
	var keys []models.Keywords
	db2.Find(&keys)
	fmt.Println(keys)

	db3, _ := gorm.Open("mysql","root:12345678@/cartoon?charset=utf8&parseTime=True")
	var users []models.User
	db3.Find(&users)
	fmt.Println(users)
	var user models.User
	db3.Debug().Where("id = ?",1).First(&user)
	fmt.Println(user.Phone)




	//r := initRouter()
	//err := r.Run(":8080")
	//if err != nil {
	//	log.Fatal("failed to start gin",err)
	//}
}

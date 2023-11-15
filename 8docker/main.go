package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 這邊收json的struct的定義都要使用uppercase，因為回傳給BindJSON所在的package做使用
type LoginInfo struct { //這邊可以打tag去做encode設置
	Username string `json:"Username" binding:"required"`
	Password string `json:"Password" binding:"required"`
}

// 共同使用的error responser
func ErrorResponse(c *gin.Context, status int, message string) {
	c.JSON(status, gin.H{"error": message})
}

func main() {
	//use default middleware
	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		//set response string
		ctx.String(200, "Welcome gin!")
	})

	r.GET("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(200, "Hello, %s!", name)
	})

	r.GET("/docs/*filepath", func(ctx *gin.Context) {
		file := ctx.Param("filepath")
		ctx.String(200, "You are looking at %s", file)
	})

	r.POST("/loginForm", func(ctx *gin.Context) {
		//use post form send request
		username := ctx.PostForm("username")
		//set default password
		password := ctx.DefaultPostForm("password", "$PASSWORD")
		//verify password
		ctx.JSON(200, gin.H{
			"username": username,
			"password": password,
		})
	})

	r.POST("/login", func(ctx *gin.Context) {
		//use post form send request
		user := LoginInfo{}
		if err := ctx.BindJSON(&user); err != nil {
			//因為有使用tag binding required 如果有少資料的話 就會跳錯
			ErrorResponse(ctx, 400, err.Error())
			return
		}
		fmt.Println(user)
		//verify password
		ctx.JSON(200, gin.H{
			"username": user.Username,
			"password": user.Password,
		})
	})
	r.Run(":8080")
}

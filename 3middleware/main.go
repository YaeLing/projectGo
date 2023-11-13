package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// 編寫middleware func
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()

		// 執行其他中間件或路由處理器
		// 這邊會切出去給下一個Second執行 如果沒有用這個的話 就會按照順序執行
		c.Next()

		endTime := time.Now()
		//使用log print會多日期時間
		log.Printf("Request processed in %s", endTime.Sub(startTime))
	}
}

func SecondMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("Request processed in second middleware")
	}
}

func main() {
	//當要release的時候，可以加這行提昇效率！！
	gin.SetMode(gin.ReleaseMode)
	//use default middleware
	r := gin.Default()
	/*
		gin.Default包含 gin.Recovery() gin.Logger()
		gin.Recovery() 將會處理service中的panic，
		並且發出status500，避免service停掉沒有回傳

	*/
	r.Use(LoggerMiddleware(), SecondMiddleware())
	/*
		分開寫也可以
		r.Use(LoggerMiddleware())
		r.Use(SecondMiddleware())
	*/

	r.GET("/hello", func(ctx *gin.Context) {
		//set response string
		ctx.String(200, "Welcome gin!")
	})

	r.Run(":8080")
}

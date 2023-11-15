package main

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

/*
跨來源資源共用（Cross-Origin Resource Sharing (CORS)）
是一種使用額外 HTTP 標頭令目前瀏覽網站的使用者代理 (en-US)
取得存取其他來源（網域）伺服器特定資源權限的機制。
當使用者代理請求一個不是目前文件來源——
例如來自於不同網域（domain）、通訊協定（protocol）或通訊埠（port）的資源時，
會建立一個跨來源 HTTP 請求（cross-origin HTTP request）。
*/
func main() {
	router := gin.Default()
	// CORS for https://foo.com and https://github.com origins, allowing:
	// - PUT and PATCH methods
	// - Origin header
	// - Credentials share
	// - Preflight requests cached for 12 hours
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.Run()
}

func main() {
	router := gin.Default()
	// - No origin allowed by default
	// - GET,POST, PUT, HEAD methods
	// - Credentials share disabled
	// - Preflight requests cached for 12 hours
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://google.com"}
	// config.AllowOrigins = []string{"http://google.com", "http://facebook.com"}
	// config.AllowAllOrigins = true

	router.Use(cors.New(config))
	router.Run()
}

func main() {
	router := gin.Default()
	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	router.Use(cors.Default())
	router.Run()
}

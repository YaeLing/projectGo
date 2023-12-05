package main

import (
	auth "amenBack/authenticate"
	"amenBack/service/adminService"
	"amenBack/service/userService"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	_ "amenBack/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	server := setupRoute()
	server.Run(":8080")
}

func setupRoute() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.GET("/auth/:userID", auth.GenerateToken)
	userRouter := router.Group("/user")
	userRouter.POST("/register", userService.RegisterAPI)
	userRouter.Use(auth.Authenticate)
	{
		userRouter.GET("/info/:key/:value", userService.QueryUserInfosAPI)
		userRouter.GET("/info", userService.QuerySelfInfosAPI)
		userRouter.PUT("/info", userService.UpdateSelfInfoAPI)

		userRouter.GET("/account", userService.QuerySelfAccountAPI)
		userRouter.PUT("/account", userService.UpdateSelfAccountAPI)

		userRouter.DELETE("/", userService.DeleteSelfProfile)
	}

	adminRouter := router.Group("/admin")
	adminRouter.Use(auth.Authenticate)
	adminRouter.Use(auth.Authorize)
	{
		adminRouter.GET("/profile/:key/:value", adminService.QueryUserProfilesAPI)
		adminRouter.DELETE("/profile/:userID", adminService.DeleteUserProfileAPI)

		adminRouter.PUT("/role", adminService.UpdateUserRoleAPI)

		adminRouter.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}

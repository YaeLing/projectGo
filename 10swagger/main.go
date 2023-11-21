package main

import (
	_ "swagger/docs"

	"github.com/gin-gonic/gin" // swagger embed files
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	//use default middleware
	r := gin.Default()

	// @Summary      Show an account
	// @Description  get string by ID
	// @Tags         accounts
	// @Accept       json
	// @Produce      json
	// @Param        id   path      int  true  "Account ID"
	// @Success      200  {object}  model.Account
	// @Failure      400  {object}  httputil.HTTPError
	// @Failure      404  {object}  httputil.HTTPError
	// @Failure      500  {object}  httputil.HTTPError
	// @Router       /accounts/{id} [get]
	r.GET("/hello", func(ctx *gin.Context) {
		//set response string
		ctx.String(200, "Welcome gin!")
	})

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}

//...

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func uploadMulti(c *gin.Context) {
	// Multipart form
	form, _ := c.MultipartForm()
	files := form.File["upload[]"]

	for _, file := range files {
		log.Println(file.Filename)
		dst := "uploads/" + file.Filename
		// Upload the file to specific dst.
		c.SaveUploadedFile(file, dst)
	}
	c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
}

func uploadSingle(c *gin.Context) {
	// single file
	file, _ := c.FormFile("file")
	log.Println(file.Filename)
	dst := "uploads/" + file.Filename
	// Upload the file to specific dst.
	c.SaveUploadedFile(file, dst)

	c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
}
func main() {
	server := setupRoute()
	server.Run(":8080")
}

func setupRoute() *gin.Engine {
	router := gin.Default()
	router.POST("/upload/single", uploadSingle)
	router.POST("/upload/multi", uploadMulti)
	return router
}

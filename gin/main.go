package main

import (
	"fmt"
	"net/http"
	"pbuf/models"

	"github.com/gin-gonic/gin"
)

func sayOk(c *gin.Context) {
	val := 2 + 2
	fmt.Println(val)
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}
func GetDBContentTags(c *gin.Context) {
	contentTags := models.ListContentTags()
	c.JSON(http.StatusOK, contentTags)
}

func main() {
	router := gin.Default()
	router.GET("/", sayOk)
	router.GET("/content_tags", GetDBContentTags)

	router.Run(":8080")
}

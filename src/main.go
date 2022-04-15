package main

import (
	"net/http"

	consts "web-service/src/constants"

	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	var router *gin.Engine = gin.Default()

	router.GET("/testGet", testGet)
	router.POST("/testPost", testPost)
	router.Run(":" + port)
}

func testGet(c *gin.Context) {
	c.String(http.StatusOK, consts.Get("puk"))
}

func testPost(c *gin.Context) {
	c.String(http.StatusOK, "test")
}

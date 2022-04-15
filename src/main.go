package main

import (
	"net/http"

	consts "web-service/src/constants"

	"github.com/gin-gonic/gin"
)

func main() {
	var router *gin.Engine = gin.Default()

	router.GET("/testGet", testGet)
	router.POST("/testPost", testPost)
	router.Run("localhost:8080")
}

func testGet(c *gin.Context) {
	c.String(http.StatusOK, consts.Get("puk"))
}

func testPost(c *gin.Context) {
	c.String(http.StatusOK, "test")
}

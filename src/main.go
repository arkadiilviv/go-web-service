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
	router.GET(consts.Get("main"), mainPage)
	router.GET(consts.Get("get"), get)
	router.POST(consts.Get("post"), post)
	router.Run(":" + port)
}

func mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.tmpl", gin.H{
		"name": consts.Get("name"),
		"urls": getApiUrls(c),
	})
}

func getApiUrls(c *gin.Context) []string {
	var result = [...]string{
		c.Request.Host + consts.Get("get"),
		c.Request.Host + consts.Get("post"),
	}
	return result[:]
}

func get(c *gin.Context) {
	c.String(http.StatusOK, consts.Get("puk"))
}

func post(c *gin.Context) {
	c.String(http.StatusOK, "test")
}

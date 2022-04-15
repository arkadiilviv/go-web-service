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
	//port = "8080"
	if port == "" {
		log.Fatal("$PORT must be set")
	}

	var router *gin.Engine = gin.Default()

	router.Use(gin.Logger())
	router.LoadHTMLGlob("pages/*.html")
	router.Static("/static", "static")

	router.GET(consts.Get("main"), mainPage)
	router.GET(consts.Get("get"), get)
	router.POST(consts.Get("post"), post)
	router.Run(":" + port)
}

func mainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "main.html", gin.H{
		"name": consts.Get("name"),
		"urls": getApiUrls(c),
	})
}

func getApiUrls(c *gin.Context) map[string]string {
	var result = map[string]string{
		consts.Get("get"):  c.Request.Host + consts.Get("get"),
		consts.Get("post"): c.Request.Host + consts.Get("post"),
	}
	return result
}

func get(c *gin.Context) {
	c.String(http.StatusOK, consts.Get("puk"))
}

func post(c *gin.Context) {
	c.String(http.StatusOK, "test")
}

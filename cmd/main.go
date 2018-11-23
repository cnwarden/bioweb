package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	."bioweb/pkg/util"
	"gopkg.in/russross/blackfriday.v2"
	"io/ioutil"
)

func loadmd() []byte  {
	content, _ := ioutil.ReadFile("markdowns/test.md")
	return content
}

func main() {
	fmt.Println("bioweb started")
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.tmpl")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.GET("/python", func(c *gin.Context) {
		Logger.Info("python started")
		c.Data(http.StatusOK, "text/html; charset=utf-8", blackfriday.Run([]byte("## TITLE\n-----")))
	})
	r.GET("/md", func(c *gin.Context) {
		Logger.Info("md started")
		c.Data(http.StatusOK, "text/html; charset=utf-8", blackfriday.Run(loadmd()))
	})
	r.GET("/", func(c *gin.Context) {
		Logger.Info("server started")
		c.HTML(http.StatusOK, "home.tmpl", gin.H{"title":"website"})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

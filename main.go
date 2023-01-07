package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/joshuaven/share-space/config"
	"github.com/joshuaven/share-space/handlers"
)

func main() {

	godotenv.Load()

	r := gin.Default()
	config.ServeStaticContents(r)
	config.ServeTemplates(r)

	r.GET("/", handlers.Home)

	r.GET("/short-urls", handlers.GetShortUrl)
	r.POST("/short-urls", handlers.PostUrl)
	r.GET("/u/:urlid", handlers.OpenShortUrl)
	
	r.Run() // listen and serve on 0.0.0.0:8080
}

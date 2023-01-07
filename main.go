package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joshuaven/share-space/config"
	"github.com/joshuaven/share-space/handlers"
)

func main() {
	r := gin.Default()
	config.ServeStaticContents(r)
	config.ServeTemplates(r)

	r.GET("/", handlers.Home)
	r.Run() // listen and serve on 0.0.0.0:8080
}

package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/joshuaven/share-space/config"
	"github.com/joshuaven/share-space/handlers"
)

func main() {

	godotenv.Load()

	r := gin.Default()
	r.MaxMultipartMemory = 5 << 20
	config.ServeStaticContents(r)
	config.ServeTemplates(r)
	config.ConfigureErrors(r)

	r.GET("/", handlers.Home)

	r.GET("/short-urls", handlers.GetShortUrl)
	r.POST("/short-urls", handlers.PostUrl)
	r.GET("/u/:urlid", handlers.OpenShortUrl)

	r.GET("/quick-share", handlers.GetQuickShare)
	r.POST("/quick-share", handlers.PostQuickShare)
	r.GET("/qs/:fileid", handlers.GetQSItem)
	r.GET("/preview/:fileid", handlers.PreviewQSItem)
	
	port := os.Getenv("PORT")
	r.Run(":" + port) // listen and serve on 0.0.0.0:8080
}

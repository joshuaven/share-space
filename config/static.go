package config

import "github.com/gin-gonic/gin"

func ServeStaticContents(engine *gin.Engine) {
	engine.Static("/assets", "./public")
}
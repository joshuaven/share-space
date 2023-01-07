package config

import "github.com/gin-gonic/gin"

func ServeTemplates(engine *gin.Engine) {
	engine.LoadHTMLGlob("client/templates/**/*")
}
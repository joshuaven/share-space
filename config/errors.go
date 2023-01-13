package config

import (
	"github.com/gin-gonic/gin"
	"github.com/joshuaven/share-space/statuses"
)

func ConfigureErrors(engine *gin.Engine) {
	engine.Use(statuses.NotFound, statuses.AppCrash)
}
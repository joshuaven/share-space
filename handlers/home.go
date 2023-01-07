package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

func Home(c *gin.Context) {
	c.HTML(200, "home/index", pageModel {
		Title: "Share Space",
		Timestamp: time.Now().UnixMilli(),
	})
}
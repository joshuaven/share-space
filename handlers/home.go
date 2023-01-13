package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joshuaven/share-space/models"
)

func Home(c *gin.Context) {
	c.HTML(200, "home/index", models.PageModel {
		Title: "Share Space",
		Timestamp: time.Now().UnixMilli(),
	})
}
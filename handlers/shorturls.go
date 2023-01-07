package handlers

import (
	"time"

	"github.com/gin-gonic/gin"
)

func GetShortUrl(ctx *gin.Context) {
	ctx.HTML(200, "shorturls/index", pageModel {
		Title: "Short URL",
		Timestamp: time.Now().UnixMilli(),
	})
}
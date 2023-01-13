package statuses

import (
	"errors"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joshuaven/share-space/models"
)

func NotFound(ctx *gin.Context) {
	ctx.Next()

	lastErr := ctx.Errors.Last()

	if lastErr == nil { return }

	if errors.Is(ErrNotFound, lastErr.Err) {
		ctx.HTML(404, "shared/notfound", models.PageModel {
			Title: "Not found",
			Timestamp: time.Now().UnixMilli(),
		})
	}
}

func AppCrash(ctx *gin.Context) {
	ctx.Next()

	lastErr := ctx.Errors.Last()

	if lastErr != nil {
		ctx.HTML(404, "shared/crash", models.PageModel {
			Title: "App Crashed",
			Timestamp: time.Now().UnixMilli(),
			Error: lastErr.Error(),
		})
	}
}
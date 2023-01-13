package handlers

import (
	"time"
	"unsafe"

	"math/rand"

	"github.com/joshuaven/share-space/models"
	"github.com/joshuaven/share-space/services"

	"github.com/gin-gonic/gin"
)

func GetShortUrl(ctx *gin.Context) {
	ctx.HTML(200, "shorturls/index", models.PageModel {
		Title: "Short URL",
		Timestamp: time.Now().UnixMilli(),
	})
}

func PostUrl(ctx *gin.Context) {
	url := ctx.PostForm("url")

	randomString := randString(7)

	item := models.ShortUrlItem {
		UrlId: randomString,
		OriginalUrl: url,
		DateAdded: time.Now().Local(),
		IpAddress: ctx.RemoteIP(),
	}

	sus := services.CreateShortUrlService()

	if err := sus.AddItem(item); err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.HTML(200, "shorturls/short", shortUrlPage {
		Title: "Your Short URL",
		Timestamp: time.Now().UnixMilli(),
		Surl: randomString,
	})
}

func OpenShortUrl(ctx *gin.Context) {
	urlId := ctx.Param("urlid")

	sus := services.CreateShortUrlService()

	var item models.ShortUrlItem

	if err := sus.GetItem(urlId, &item); err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.Redirect(301, item.OriginalUrl)
}

func randString(n int) string {
		var src = rand.NewSource(time.Now().UnixNano())
    b := make([]byte, n)
    // A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
    for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
        if remain == 0 {
            cache, remain = src.Int63(), letterIdxMax
        }
        if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
            b[i] = letterBytes[idx]
            i--
        }
        cache >>= letterIdxBits
        remain--
    }

    return *(*string)(unsafe.Pointer(&b))
}
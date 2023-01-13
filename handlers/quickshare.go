package handlers

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joshuaven/share-space/models"
	"github.com/joshuaven/share-space/services"
	"github.com/joshuaven/share-space/statuses"
)

func GetQuickShare(ctx *gin.Context) {
	ctx.HTML(200, "quickshare/index", models.PageModel {
		Title: "Quick Share",
		Timestamp: time.Now().UnixMilli(),
	})
}

func PostQuickShare(ctx *gin.Context) {
	file, _ := ctx.FormFile("file")
	tmpPath := os.TempDir()

	log.Printf("Saving file: %v", file.Header)

	fileId := randString(8)

	filename := fileId + filepath.Ext(file.Filename)

	qsItem := models.QSItem {
		FileId: fileId,
		FileName: file.Filename,
		FileType: file.Header.Get("Content-Type"),
		DateAdded: time.Now(),
	}

	service := services.CreateQuickShareService()
	err := service.AddItem(qsItem)
	if err != nil {
		ctx.HTML(500, "quickshare/index", models.PageModel {
			Title: "Quick Share",
			Timestamp: time.Now().UnixMilli(),
			Error: err.Error(),
		})
	}

	ctx.SaveUploadedFile(file, tmpPath + "/" + filename)
	ctx.Redirect(302, "/qs/" + fileId)
}

func GetQSItem(ctx *gin.Context) {
	fileId := ctx.Param("fileid")
	queryParam := ctx.Query("action")

	service := services.CreateQuickShareService()

	var item models.QSItem
	
	if err := service.GetItem(fileId, &item); err != nil {
		ctx.Redirect(301, "/assets/broken.jpeg")
		return
	}

	tmpPath := os.TempDir()
	filename := fileId + filepath.Ext(item.FileName)
	fileLoc := tmpPath + "/" + filename

		_, err := os.Stat(fileLoc)
	if os.IsNotExist(err) {
		ctx.Error(statuses.ErrNotFound)
		return
	}

	if queryParam == "download" {
		ctx.Writer.Header().Set("Content-Disposition", "attachment; filename=" + item.FileName)
		ctx.File(fileLoc)
		return
	}

	ctx.HTML(200, "quickshare/file", qsItemPage {
		Title: fileId,
		Timestamp: time.Now().UnixMilli(),
		Item: item,
		Preview: filename,
		Scripts: []models.Script{
			models.Script("/assets/js/qs/index.js").WithTimestamp(),
		},
	})
}

func PreviewQSItem(ctx *gin.Context) {
	filename := ctx.Param("filename")
	tmpPath := os.TempDir()
	fileLoc := tmpPath + "/" + filename

	ctx.File(fileLoc)
}
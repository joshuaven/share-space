package handlers

import (
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joshuaven/share-space/models"
	"github.com/joshuaven/share-space/services"
)

func GetQuickShare(ctx *gin.Context) {
	ctx.HTML(200, "quickshare/index", pageModel {
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
		ctx.HTML(500, "quickshare/index", pageModel {
			Title: "Quick Share",
			Timestamp: time.Now().UnixMilli(),
			Error: err.Error(),
		})
	}

	ctx.SaveUploadedFile(file, tmpPath + "/" + filename)
	ctx.HTML(200, "quickshare/index", pageModel {
		Title: "Quick Share",
		Timestamp: time.Now().UnixMilli(),
	})
}

func GetQSItem(ctx *gin.Context) {
	fileId := ctx.Param("fileid")

	ctx.HTML(200, "quickshare/file", qsItemPage {
		Title: fileId,
		Timestamp: time.Now().UnixMilli(),
		Item: fileId,
	})
}

func PreviewQSItem(ctx *gin.Context) {
	fileId := ctx.Param("fileid")

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
		ctx.Redirect(301, "/assets/broken.jpeg")
		return
	}

	ctx.File(fileLoc)
}
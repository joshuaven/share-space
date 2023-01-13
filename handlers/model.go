package handlers

import "github.com/joshuaven/share-space/models"

type shortUrlPage struct {
	Title string
	Timestamp int64
	Surl string
}

type qsItemPage struct {
	Title string
	Timestamp int64
	Item models.QSItem
	Preview string
}
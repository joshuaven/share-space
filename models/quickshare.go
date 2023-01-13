package models

import "time"

type QSItem struct {
	FileId string
	FileName string
	FileType string
	DateAdded time.Time
}
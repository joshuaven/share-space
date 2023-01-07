package models

import "time"

type ShortUrlItem struct {
	UrlId string
	IpAddress string
	OriginalUrl string
	DateAdded time.Time
}
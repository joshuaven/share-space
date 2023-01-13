package models

import (
	"fmt"
	"time"
)

type Script string

func (s Script) WithTimestamp() Script {
	currentTime := time.Now().UnixMilli()
	
	s = s + Script(fmt.Sprintf("?v=%v", currentTime))
	return s
}

type PageModel struct {
	Title string
	Timestamp int64
	Error string
	Scripts []Script
}
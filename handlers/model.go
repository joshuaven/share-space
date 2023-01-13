package handlers

type pageModel struct {
	Title string
	Timestamp int64
	Error string
}

type shortUrlPage struct {
	Title string
	Timestamp int64
	Surl string
}

type qsItemPage struct {
	Title string
	Timestamp int64
	Item string
}
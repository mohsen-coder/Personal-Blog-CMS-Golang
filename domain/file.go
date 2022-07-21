package domain

import "time"

type File struct {
	Id       string
	Name     string
	Size     int64
	MimeType MimType
	CreateAt time.Time
	UpdateAt time.Time
}

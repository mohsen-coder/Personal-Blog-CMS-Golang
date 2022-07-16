package domain

import "time"

type File struct {
	Id       string
	Name     string
	Size     int
	MimType  MimType
	CreateAt *time.Time
	UpdateAt *time.Time
}

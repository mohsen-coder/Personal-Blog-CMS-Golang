package domain

import "time"

type Comment struct {
	Id       string
	Parent   *Comment
	Children []*Comment
	Post     *Post
	Emain    string
	Name     string
	Content  string
	Status   CommentStatus
	CreateAt *time.Time
	UpdateAt *time.Time
}

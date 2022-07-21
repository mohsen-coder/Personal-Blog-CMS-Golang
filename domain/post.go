package domain

import "time"

type Post struct {
	Id          string
	Thumbnail   *File
	Title       string
	Content     string
	FullContent string
	Categories  []*Category
	Tags        []string
	Comments    []*Comment
	Author      *Account
	View        int
	Like        int
	PublishDate time.Time
	Status      PostStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

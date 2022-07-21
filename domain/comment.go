package domain

import "time"

type Comment struct {
	Id        string
	Parent    *Comment
	Children  []*Comment
	Post      *Post
	Email     string
	Name      string
	Content   string
	Status    CommentStatus
	CreatedAt time.Time
	UpdatedAt time.Time
}

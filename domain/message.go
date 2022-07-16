package domain

import "time"

type Message struct {
	Id       string
	Name     string
	Email    string
	WebSite  string
	Title    string
	Content  string
	Status   CommentStatus
	Reply    *Message
	Parent   *Message
	File     *Message
	CreateAt *time.Time
	UpdateAt *time.Time
}

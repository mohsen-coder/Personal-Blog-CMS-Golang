package domain

import "time"

type Message struct {
	Id        string
	Name      string
	Email     string
	WebSite   string
	Title     string
	Content   string
	Status    MessageStatus
	Reply     *Message
	Parent    *Message
	File      *File
	CreatedAt time.Time
	UpdatedAt time.Time
}

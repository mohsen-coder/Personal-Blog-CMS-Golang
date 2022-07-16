package domain

import (
	"time"
)

type Account struct {
	Id        string
	Name      string
	Family    string
	Username  string
	Emial     string
	Password  string
	About     string
	Thumbnail *File
	Role      AccountRole
	CreateAt  *time.Time
	UpdateAt  *time.Time
}

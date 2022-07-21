package domain

import (
	"time"
)

type Account struct {
	Id        string
	Name      string
	Family    string
	Username  string
	Email     string
	Password  string
	About     string
	Thumbnail *File
	Role      AccountRole
	CreatedAt time.Time
	UpdatedAt time.Time
}

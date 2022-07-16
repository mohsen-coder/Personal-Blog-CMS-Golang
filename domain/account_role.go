package domain

type AccountRole uint8

const (
	User AccountRole = iota
	Writer
	Admin
)

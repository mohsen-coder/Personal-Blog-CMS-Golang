package domain

type AccountRole uint8

const (
	AccountUser AccountRole = iota + 1
	AccountWriter
	AccountAdmin
)

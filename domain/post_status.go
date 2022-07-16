package domain

type PostStatus string

const (
	Non     PostStatus = "none"
	Publish            = "publish"
	Suspend            = "suspend"
)

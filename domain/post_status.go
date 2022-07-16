package domain

type PostStatus string

const (
	PostNone    PostStatus = "none"
	PostPublish PostStatus = "publish"
	PostSuspend PostStatus = "suspend"
)

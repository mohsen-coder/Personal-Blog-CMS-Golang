package domain

type CommentStatus string

const (
	CommentNone   CommentStatus = "none"
	CommentAccept CommentStatus = "accept"
	CommentReject CommentStatus = "reject"
)

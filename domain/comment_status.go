package domain

type CommentStatus string

const (
	None   CommentStatus = "none"
	Accept               = "accept"
	Reject               = "reject"
)

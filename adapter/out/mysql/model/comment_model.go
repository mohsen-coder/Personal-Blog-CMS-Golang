package model

import (
	"mohsen-coder/PersonalBlog/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CommentModel struct {
	gorm.Model
	CommentId string
	Parent    *CommentModel
	Children  []*CommentModel
	Post      *PostModel
	Email     string
	Name      string
	Content   string
	Status    string
}

func (cm *CommentModel) BeforeCreate(tx *gorm.DB) (err error) {
	cm.CommentId = uuid.New().String()
	return
}

func (cm *CommentModel) ToDomainModel() *domain.Comment {
	comment := domain.Comment{
		Id:        cm.CommentId,
		Email:     cm.Email,
		Name:      cm.Name,
		Content:   cm.Content,
		CreatedAt: cm.CreatedAt,
		UpdatedAt: cm.UpdatedAt,
	}

	if cm.Parent != nil {
		*comment.Parent = *cm.Parent.ToDomainModel()
	}

	if cm.Children != nil {
		for _, childComment := range cm.Children {
			comment.Children = append(comment.Children, childComment.ToDomainModel())
		}
	}

	switch cm.Status {
	case "accept":
		comment.Status = domain.CommentAccept
	case "reject":
		comment.Status = domain.CommentReject
	default:
		comment.Status = domain.CommentNone
	}

	return &comment
}

func (cm *CommentModel) FromDomainModel(comment domain.Comment) {
	cm.Email = comment.Email
	cm.Name = comment.Name
	cm.Content = comment.Content
	cm.Status = string(comment.Status)

	if comment.Id != "" {
		cm.CommentId = comment.Id
	}

	if comment.Parent != nil {
		cm.Parent.FromDomainModel(*comment.Parent)
	}
}

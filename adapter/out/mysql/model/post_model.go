package model

import (
	"mohsen-coder/PersonalBlog/domain"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PostModel struct {
	gorm.Model
	PostId      string
	Thumbnail   *FileModel
	Title       string
	Content     string
	FullContent string
	Categories  []*CategoryModel
	Tags        []string
	Comments    []*CommentModel
	Author      *AccountModel
	View        int
	Like        int
	PublishDate time.Time
	Status      string
}

func (pm *PostModel) BeforeCreate(tx *gorm.DB) (err error) {
	pm.PostId = uuid.New().String()
	return
}

func (pm *PostModel) ToDomainModel() *domain.Post {
	post := domain.Post{
		Id:          pm.PostId,
		Title:       pm.Title,
		Content:     pm.Content,
		FullContent: pm.FullContent,
		Tags:        pm.Tags,
		View:        pm.View,
		Like:        pm.Like,
		PublishDate: pm.PublishDate,
		CreatedAt:   pm.CreatedAt,
		UpdatedAt:   pm.UpdatedAt,
	}

	if pm.Thumbnail != nil {
		*post.Thumbnail = *pm.Thumbnail.ToDomainModel()
	}

	for _, category := range pm.Categories {
		post.Categories = append(post.Categories, category.ToDomainModel())
	}

	for _, comment := range pm.Comments {
		post.Comments = append(post.Comments, comment.ToDomainModel())
	}

	if pm.Author != nil {
		*post.Author = *pm.Author.ToDomainModel()
	}

	return &post
}

func (pm *PostModel) FromDomainModel(post domain.Post) {
	pm.Title = post.Title
	pm.Content = post.Content
	pm.FullContent = post.FullContent
	pm.Tags = post.Tags
	pm.View = post.View
	pm.Like = post.Like
	pm.PublishDate = post.PublishDate
	pm.Status = string(post.Status)

	if post.Id != "" {
		pm.PostId = post.Id
	}

	if post.Thumbnail != nil {
		pm.Thumbnail.FromDomainModel(*post.Thumbnail)
	}

	if len(post.Categories) > 0 {
		for _, category := range post.Categories {
			var categoryModel CategoryModel
			categoryModel.FromDomainModel(*category)
			pm.Categories = append(pm.Categories, &categoryModel)
		}
	}

	if len(post.Comments) > 0 {
		for _, comment := range post.Comments {
			var commentModel CommentModel
			commentModel.FromDomainModel(*comment)
			pm.Comments = append(pm.Comments, &commentModel)
		}
	}

	if post.Author != nil {
		pm.Author.FromDomainModel(*post.Author)
	}
}

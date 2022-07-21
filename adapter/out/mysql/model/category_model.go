package model

import (
	"mohsen-coder/PersonalBlog/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CategoryModel struct {
	gorm.Model
	CategoryId string
	Parent     *CategoryModel
	Title      string
}

func (cm *CategoryModel) BeforeCreate(tx *gorm.DB) (err error) {
	cm.CategoryId = uuid.New().String()
	return
}

func (cm *CategoryModel) ToDomainModel() *domain.Category {
	category := domain.Category{
		Id:    cm.CategoryId,
		Title: cm.Title,
	}

	if cm.Parent != nil {
		*category.Parent = *cm.Parent.ToDomainModel()
	}

	return &category
}

func (cm *CategoryModel) FromDomainModel(category domain.Category) {
	cm.Title = category.Title

	if category.Id != "" {
		cm.CategoryId = category.Id
	}

	if category.Parent != nil {
		cm.Parent.FromDomainModel(*category.Parent)
	}
}

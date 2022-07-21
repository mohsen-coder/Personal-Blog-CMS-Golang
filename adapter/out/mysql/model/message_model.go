package model

import (
	"mohsen-coder/PersonalBlog/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MessageModel struct {
	gorm.Model
	MessageId string
	Name      string
	Email     string
	WebSite   string
	Title     string
	Content   string
	Status    string
	Reply     *MessageModel
	Parent    *MessageModel
	File      *FileModel
}

func (mm *MessageModel) BeforeCreate(tx *gorm.DB) (err error) {
	mm.MessageId = uuid.New().String()
	return
}

func (mm *MessageModel) ToDomainModel() *domain.Message {
	message := domain.Message{
		Id:        mm.MessageId,
		Name:      mm.Name,
		Email:     mm.Email,
		WebSite:   mm.WebSite,
		Title:     mm.Title,
		Content:   mm.Content,
		CreatedAt: mm.CreatedAt,
		UpdatedAt: mm.UpdatedAt,
	}

	switch mm.Status {
	case "read":
		message.Status = domain.MessageRead
	default:
		message.Status = domain.MessageUnRead
	}

	if mm.Reply != nil {
		*message.Reply = *mm.Reply.ToDomainModel()
	}

	if mm.Parent != nil {
		*message.Parent = *mm.Parent.ToDomainModel()
	}

	if mm.File != nil {
		*message.File = *mm.File.ToDomainModel()
	}

	return &message
}

func (mm *MessageModel) FromDomainModel(message domain.Message) {
	mm.Name = message.Name
	mm.Email = message.Email
	mm.WebSite = message.WebSite
	mm.Title = message.Title
	mm.Content = message.Content
	mm.Status = string(message.Status)

	if message.Id != "" {
		mm.MessageId = message.Id
	}

	if message.Parent != nil {
		mm.Parent.FromDomainModel(*message.Parent)
	}

	if message.Reply != nil {
		mm.Reply.FromDomainModel(*message.Reply)
	}

	if message.File != nil {
		mm.File.FromDomainModel(*message.File)
	}
}

package model

import (
	"mohsen-coder/PersonalBlog/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type FileModel struct {
	gorm.Model
	FileId   string
	Size     int64
	Name     string
	MimeType string
}

func (fm *FileModel) BeforeCreate(tx *gorm.DB) (err error) {
	fm.FileId = uuid.New().String()
	return
}

func (fm *FileModel) ToDomainModel() *domain.File {
	file := domain.File{
		Id:   fm.FileId,
		Size: fm.Size,
		Name: fm.Name,
	}

	switch fm.MimeType {
	case "image/jpeg":
		file.MimeType = domain.FileJPG
	case "image/png":
		file.MimeType = domain.FilePNG
	case "image/gif":
		file.MimeType = domain.FileGIF
	case "application/pdf":
		file.MimeType = domain.FilePDF
	case "application/zip":
		file.MimeType = domain.FileZIP
	default:
		file.MimeType = domain.FileNone
	}

	return &file
}

func (fm *FileModel) FromDomainModel(file domain.File) {
	if file.Id != "" {
		fm.FileId = file.Id
	}
	fm.Size = file.Size
	fm.Name = file.Name
	fm.MimeType = string(file.MimeType)
}

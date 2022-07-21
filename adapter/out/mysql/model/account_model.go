package model

import (
	"mohsen-coder/PersonalBlog/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type AccountModel struct {
	gorm.Model
	AccountId    string
	Name         string
	Family       string
	Username     string
	Email        string
	Password     string
	About        string
	ThumbnailRef int
	Thumbnail    *FileModel `gorm:"foreignKey:ThumbnailRef"`
	Role         int
}

func (am *AccountModel) BeforeCreate(tx *gorm.DB) (err error) {
	am.AccountId = uuid.New().String()
	return
}

func (am *AccountModel) ToDomainModel() *domain.Account {
	account := domain.Account{
		Id:        am.AccountId,
		Name:      am.Name,
		Family:    am.Family,
		Username:  am.Username,
		Email:     am.Email,
		Password:  am.Password,
		About:     am.Password,
		CreatedAt: am.CreatedAt,
		UpdatedAt: am.UpdatedAt,
	}

	if am.Thumbnail != nil {
		*account.Thumbnail = *am.Thumbnail.ToDomainModel()
	}

	switch am.Role {
	case 2:
		account.Role = domain.AccountWriter
	case 3:
		account.Role = domain.AccountAdmin
	default:
		account.Role = domain.AccountUser
	}

	return &account
}

func (am *AccountModel) FromDomainModel(account domain.Account) {
	am.Name = account.Name
	am.Family = account.Family
	am.Username = account.Username
	am.Email = account.Email
	am.Role = int(account.Role)

	if account.Id != "" {
		am.AccountId = account.Id
	}

	if account.Password != "" {
		am.Password = account.Password
	}

	if account.Thumbnail != nil {
		am.Thumbnail.FromDomainModel(*account.Thumbnail)
	}
}

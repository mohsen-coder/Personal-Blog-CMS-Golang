package mysql

import (
	"mohsen-coder/PersonalBlog/adapter/out/mysql/model"
	"mohsen-coder/PersonalBlog/domain"

	"gorm.io/gorm"
)

type AccountPersistence struct {
	db *gorm.DB
}

func (ap *AccountPersistence) CreateAccount(account domain.Account) *domain.Account {
	var accountModel model.AccountModel
	accountModel.FromDomainModel(account)
	ap.db.Create(&accountModel)

	return accountModel.ToDomainModel()
}

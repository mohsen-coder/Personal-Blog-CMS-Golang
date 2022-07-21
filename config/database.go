package config

import (
	"log"
	"mohsen-coder/PersonalBlog/adapter/out/mysql/model"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var Database DbInstance

func ConnectDb() {
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: "root:123456789@tcp(127.0.0.1:3306)/gorm_test?charset=utf8&parseTime=True&loc=Local", // data source name
		// DefaultStringSize: 256, // default size for string fields
		// DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		// DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		// DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		// SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		os.Exit(2)
	}

	log.Println("Connected Successfully to Database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	db.AutoMigrate(&model.AccountModel{}, &model.FileModel{})

	Database = DbInstance{
		Db: db,
	}
}

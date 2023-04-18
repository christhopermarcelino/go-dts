package config

import (
	"challenge10/model"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	HOST     = "localhost"
	PORT     = 5432
	USER     = "postgres"
	PASSWORD = "postgres"
	DBNAME   = "go_dts_9"
	DATABASE = "postgres"
)

var (
	db  *gorm.DB
	err error
)

func GetDB() *gorm.DB {
	return db
}

func InitDatabase() {
	configInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, USER, PASSWORD, DBNAME)

	db, err = gorm.Open(postgres.Open(configInfo), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.Debug().AutoMigrate(&model.Book{})
}

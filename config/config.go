package config

import (
	"../app/models"
	"github.com/jinzhu/gorm"
)

func DBInit() *gorm.DB {
	db, err := gorm.Open("mysql", "root:"+"@tcp(127.0.0.1:3306)/go-db?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect to database")
	}
	db.AutoMigrate(models.Person{})
	return db
}
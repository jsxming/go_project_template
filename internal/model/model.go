package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDbEngine() (*gorm.DB, error) {
	dsn := "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	return db, err
}

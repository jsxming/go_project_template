package model

import "gorm.io/gorm"

type User struct {
	Name string
	Age  int
}

// db  数据库操作相关 写在modal中
func (u User) Create(db *gorm.DB) (*User, error) {
	if err := db.Create(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

package dao

import "gorm.io/gorm"

type Dao struct {
	engine *gorm.DB
}

func New(e *gorm.DB) *Dao {
	return &Dao{engine: e}
}

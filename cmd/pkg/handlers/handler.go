package handlers

import "gorm.io/gorm"

// using dependency injection

type handler struct {
	DB *gorm.DB
}

func NewDB(db *gorm.DB) handler {
	return handler{db}
}

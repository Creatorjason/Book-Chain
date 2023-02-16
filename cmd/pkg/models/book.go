package models

import(
	"gorm.io/gorm"
)

type Book struct {
	ID     int    `json:"id" ;gorm:"primaryKey autoIncrement"`
	Author string `json:"author"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
}

func NewMigration(db *gorm.DB) error{
	err := db.AutoMigrate(&Book{})
	return err
}
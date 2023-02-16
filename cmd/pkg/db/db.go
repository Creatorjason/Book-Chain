package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"main.go/cmd/pkg/models"
)

func InitializeDB() *gorm.DB{
	dbURL := "postgres://pg:pass@localhost:5432/crud"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil{
		log.Fatal(err)
	}
	err = models.NewMigration(db)
	if err != nil{
		log.Fatal(err)
	}
	return db
}
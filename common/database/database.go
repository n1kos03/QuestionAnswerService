package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=postgres user=db password=123 dbname=question_answer port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Can't connect to database: ", err)
	}

	DB = db
}
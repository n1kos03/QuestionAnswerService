package question

import (
	"gorm.io/gorm"
)

type questionHandler struct {
	DB *gorm.DB
}

func InitHandler(db *gorm.DB) *questionHandler {
	return &questionHandler{DB: db}
}
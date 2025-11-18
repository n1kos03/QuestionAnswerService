package answer

import (
	"gorm.io/gorm"
)

type answerHandler struct {
	DB *gorm.DB
}

func InitHandler(db *gorm.DB) *answerHandler {
	return &answerHandler{DB: db}
}
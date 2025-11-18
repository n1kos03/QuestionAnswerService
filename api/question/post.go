package question

import (
	"context"
	"encoding/json"
	"net/http"
	"questionAnswer/common/model"

	"gorm.io/gorm"
)

func (h *questionHandler) PostQuestion(w http.ResponseWriter, r *http.Request) {
	var question model.Question
	err := json.NewDecoder(r.Body).Decode(&question)
	if err != nil {
		http.Error(w, "Incorrect body of request", http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	// h.DB.Create(&question)
	err = gorm.G[model.Question](h.DB).Create(ctx, &question)
	if err != nil {
		http.Error(w, "Error while creating question", http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Question was added",
	})
}
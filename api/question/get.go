package question

import (
	"encoding/json"
	"net/http"
	"questionAnswer/common/model"
	"strings"
)

func (h *questionHandler) GETQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	var questions []model.Question
	// id := r.URL.Query().Get("id")
	id := strings.TrimPrefix(r.URL.Path, "/questions/")
	
	if id != "" {
		h.DB.First(&questions, id)
		// Add finding of all answers
	} else {
		h.DB.Find(&questions)
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(questions)
}
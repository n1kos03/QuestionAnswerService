package question

import (
	"encoding/json"
	"net/http"
	"questionAnswer/common/model"
)

func (h *questionHandler) GETQuestions(w http.ResponseWriter, r *http.Request, id string) {
	var questions []model.Question

	h.DB.Find(&questions)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(questions)
}

func (h *questionHandler) GETQuestionByID(w http.ResponseWriter, r *http.Request, id string) {
	var question model.Question
	var answers []model.Answer

	h.DB.First(&question, id)
	h.DB.Where("question_id = ?", id).Find(&answers)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"question": question,
		"answers": answers,
	})
}
package answer

import (
	"encoding/json"
	"net/http"
	"questionAnswer/common/model"
)

func (h *answerHandler) GETAnswerByID(w http.ResponseWriter, r *http.Request, id string) {
	var answer model.Answer

	h.DB.Find(&answer, id)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(answer)
}
package answer

import (
	"encoding/json"
	"net/http"
	"questionAnswer/common/model"
)

func (h *answerHandler) DELETEAnswerByID(w http.ResponseWriter, r *http.Request, id string) {
	var answer model.Answer

	tx := h.DB.Delete(&answer, id)
	if tx.Error != nil {
		http.Error(w, "Can't delete answer", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Answer successfully deleted",
	})
}
package question

import (
	"context"
	"encoding/json"
	"net/http"
	"questionAnswer/common/model"

	"gorm.io/gorm"
)

func (h *questionHandler) DELETEquestion(w http.ResponseWriter, r *http.Request, id string) {
	var question model.Question

	ctx := context.Background()
	
	gorm.G[model.Answer](h.DB).Where("question_id = ?", id).Delete(ctx)
	h.DB.Delete(&question, id)

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Question succefully deleted",
	})
}
package answer

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"questionAnswer/common/model"
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func (h* answerHandler) PostAnswer(w http.ResponseWriter, r *http.Request) {
	var answer model.Answer
	err := json.NewDecoder(r.Body).Decode(&answer)
	if err != nil {
		http.Error(w, "Incorrect body in request", http.StatusBadRequest)
		log.Println("Incorrect body of request: ", err)
		return
	}

	idTmp := strings.TrimPrefix(r.URL.Path, "/questions/")
	id := strings.TrimSuffix(idTmp, "/answers/")

	ctx := context.Background()

	_, err = gorm.G[model.Question](h.DB).Where("id = ?", id).Find(ctx)
	if err != nil {
		http.Error(w, "This question does not exist", http.StatusBadRequest)
		return
	}

	answer.Question_id, _ = strconv.Atoi(id)
	err = gorm.G[model.Answer](h.DB).Create(ctx, &answer)
	if err != nil {
		http.Error(w, "Error while creating answer", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Answer was added",
	})
}
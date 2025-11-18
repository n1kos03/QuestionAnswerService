package main

import (
	"log"
	"net/http"
	"questionAnswer/api/answer"
	"questionAnswer/api/question"
	"questionAnswer/common/database"
	"questionAnswer/common/migrations"
	"strings"
)

func main() {
	database.InitDB()

	migrations.RunMigrations()

	q := question.InitHandler(database.DB)

	http.HandleFunc("/questions/", func(w http.ResponseWriter, r *http.Request) {
		id := strings.TrimPrefix(r.URL.Path, "/questions/")

		switch r.Method {
		case http.MethodPost:
			q.PostQuestion(w, r)
		case http.MethodGet:
			if id != "" {
				q.GETQuestionByID(w, r, id)
			} else {
				q.GETQuestions(w, r, id)
			}
		case http.MethodDelete:
			q.DELETEquestion(w, r, id)
		}
	})

	a := answer.InitHandler(database.DB)

	http.HandleFunc("/questions/{id}/answers/", a.PostAnswer)

	log.Println("Start server on 8080 port")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
	log.Println("Server is stopped")
}
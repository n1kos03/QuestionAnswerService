package main

import (
	"log"
	"net/http"
	"questionAnswer/api/question"
	"questionAnswer/common/database"
	"questionAnswer/common/migrations"
)

func main() {
	database.InitDB()

	migrations.RunMigrations()

	q := question.InitHandler(database.DB)

	http.HandleFunc("/questions/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			q.PostQuestion(w, r)
		case http.MethodGet:
			q.GETQuestionsHandler(w, r)
		}
	})

	// http.HandleFunc("/questions/:id", q.GETQuestionsHandler)

	// http.HandleFunc("/questions/", q.PostQuestion)
	// http.HandleFunc("/questions/", q.GETQuestionsHandler)

	log.Println("Start server on 8080 port")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
	log.Println("Server is stopped")
}
package main

import (
	"log"
	"net/http"

	"bem.itts.ac.id/projects/handlers"
)

const port = ":8080"

func main() {
	http.HandleFunc("/qa", handlers.QuestionAnswerHandler)
	log.Println("Server is running on port", port)
	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

package handlers

import (
	"encoding/json"
	"net/http"

	"bem.itts.ac.id/projects/config"
	"bem.itts.ac.id/projects/models"
	"bem.itts.ac.id/projects/utils"
)

const apiURL = "https://api-inference.huggingface.co/models/distilbert-base-uncased-distilled-squad"

func QuestionAnswerHandler(w http.ResponseWriter, r *http.Request) {
	var qa models.QuestionAnswer
	err := json.NewDecoder(r.Body).Decode(&qa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	token := config.GetHuggingFaceToken()
	body, err := utils.MakePostRequest(apiURL, token, qa)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var answer models.Answer
	if err := json.Unmarshal(body, &answer); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(answer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

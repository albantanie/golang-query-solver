package models

type QuestionAnswer struct {
	Question string `json:"question"`
	Context  string `json:"context"`
}

type Answer struct {
	Answer string  `json:"answer"`
	Score  float64 `json:"score"`
}

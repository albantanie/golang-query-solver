package config

import (
	"os"
)

func GetHuggingFaceToken() string {
	return os.Getenv("HUGGING_FACE_API_TOKEN")
}

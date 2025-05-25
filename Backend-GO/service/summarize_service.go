package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"backend-summarizer/database"
	"backend-summarizer/model"
)

func ProcessSummary(text string) string {
	summary := callMLModel(text)
	database.SaveSummary(text, summary)
	return summary
}

func callMLModel(text string) string {
	payload := map[string]string{"text": text}
	jsonVal, _ := json.Marshal(payload)

	resp, err := http.Post("http://localhost:5000/generate", "application/json", bytes.NewBuffer(jsonVal))
	if err != nil {
		return "Error contacting ML service"
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]string
	json.Unmarshal(body, &result)
	return result["summary"]
}

func FetchSummaries() []model.Summary {
	return database.GetAllSummaries()
}

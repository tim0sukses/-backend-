package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"backend-summarizer/database"
	"backend-summarizer/model"
)

func SaveFullSummary(s model.SummaryRequest) {
	database.SaveSummary(s)
}

func FetchSummaries() []model.SummaryRequest {
	return database.GetAllSummaries()
}

func CallMLModel(text string) string {
	payload := map[string]string{"text": text}
	jsonVal, _ := json.Marshal(payload)

	resp, err := http.Post("https://api-capstone-kappa.vercel.app/process-text", "application/json", bytes.NewBuffer(jsonVal))
	if err != nil {
		return "Error contacting ML service"
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result map[string]string
	json.Unmarshal(body, &result)
	return result["summary"]
}

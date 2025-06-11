package service

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"

	"backend-summarizer/database"
	"backend-summarizer/model"
)

func ProcessSummary(req model.SummarizeRequest) (model.SummaryResponse, error) {
	jsonVal, err := json.Marshal(req)
	if err != nil {
		return model.SummaryResponse{Status: "error", Summary: "Invalid request payload"}, err
	}

	resp, err := http.Post("https://api-capstone-production-1e95.up.railway.app/", "application/json", bytes.NewBuffer(jsonVal))
	if err != nil {
		return model.SummaryResponse{Status: "error", Summary: "Error contacting ML service"}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return model.SummaryResponse{Status: "error", Summary: "Failed to read ML response"}, err
	}

	var result model.SummaryResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return model.SummaryResponse{Status: "error", Summary: "Invalid ML response format"}, err
	}

	if result.CreatedAt == "" {
		result.CreatedAt = time.Now().Format(time.RFC3339)
	}

	database.SaveSummary(result)

	return result, nil
}

func FetchSummaries() ([]model.SummaryResponse, error) {
	return database.GetAllSummaries(), nil
}

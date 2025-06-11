package service

import (
	"backend-summarizer/database"
	"backend-summarizer/model"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

const mlAPI = "https://api-capstone-production-1e95.up.railway.app"

func callMLAPI(endpoint string, req model.SummarizeRequest) (model.SummaryResponse, error) {
	jsonVal, _ := json.Marshal(req)
	resp, err := http.Post(mlAPI+endpoint, "application/json", bytes.NewBuffer(jsonVal))
	if err != nil {
		return model.SummaryResponse{}, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result model.SummaryResponse
	json.Unmarshal(body, &result)

	result.OriginalText = req.Text
	result.CreatedAt = time.Now().Format(time.RFC3339)
	return result, nil
}

func ProcessText(req model.SummarizeRequest) (model.SummaryResponse, error) {
	return callMLAPI("/process-text", req)
}

func Summarize(req model.SummarizeRequest) (model.SummaryResponse, error) {
	return callMLAPI("/summarize", req)
}

func GenerateQuestion(req model.SummarizeRequest) (model.SummaryResponse, error) {
	return callMLAPI("/generate-question", req)
}

func FetchSummaries() ([]model.SummaryResponse, error) {
	return database.GetAllSummaries(), nil
}

func SaveToDB(res model.SummaryResponse) {
	database.SaveSummary(res)
}

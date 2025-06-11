package model

type SummarizeRequest struct {
	Text string `json:"text" binding:"required"`
}

type SummaryResponse struct {
	OriginalText string `json:"original_text"`
	Summary      string `json:"summary"`
	Questions    string `json:"questions"`
	CreatedAt    string `json:"created_at"`
}

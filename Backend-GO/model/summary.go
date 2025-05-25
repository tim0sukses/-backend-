package model

type Summary struct {
	OriginalText string `json:"original_text"`
	SummaryText  string `json:"summary_text"`
	CreatedAt    string `json:"created_at"`
}

package model

type SummarizeRequest struct {
	Text             string `json:"text" binding:"required"`
	Length           string `json:"length"`
	IncludeQuestions bool   `json:"include_questions"`
	NumQuestions     int    `json:"num_questions"`
}

// MethodsUsed records which methods were used by the ML service
type MethodsUsed struct {
	Questions string `json:"questions"`
	Summary   string `json:"summary"`
}

// SummaryResponse is the full response from the ML endpoint.
type SummaryResponse struct {
	CompressionRatio string      `json:"compression_ratio"`
	MethodsUsed      MethodsUsed `json:"methods_used"`
	OriginalLength   int         `json:"original_length"`
	OriginalText     string      `json:"original_text"`
	ProcessingMode   string      `json:"processing_mode"`
	QuestionCount    int         `json:"question_count"`
	Questions        []string    `json:"questions"`
	Status           string      `json:"status"`
	Summary          string      `json:"summary"`
	SummaryLength    int         `json:"summary_length"`
	CreatedAt        string      `json:"created_at"`
}

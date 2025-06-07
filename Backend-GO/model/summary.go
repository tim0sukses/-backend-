// file: model/summary.go
package model

type MethodsUsed struct {
	Questions string `json:"questions"`
	Summary   string `json:"summary"`
}

type SummaryRequest struct {
	CompressionRatio string      `json:"compression_ratio" binding:"required"`
	MethodsUsed      MethodsUsed `json:"methods_used" binding:"required"`
	OriginalLength   int         `json:"original_length" binding:"required"`
	OriginalText     string      `json:"original_text" binding:"required"`
	ProcessingMode   string      `json:"processing_mode" binding:"required"`
	QuestionCount    int         `json:"question_count" binding:"required"`
	Questions        []string    `json:"questions" binding:"required"`
	Status           string      `json:"status" binding:"required"`
	Summary          string      `json:"summary" binding:"required"`
	SummaryLength    int         `json:"summary_length" binding:"required"`
	CreatedAt        string      `json:"created_at" binding:"required"`
}

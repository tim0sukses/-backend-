package controller

import (
	"fmt"
	"net/http"
	"time"
	"unicode/utf8"

	"backend-summarizer/model"
	"backend-summarizer/service"

	"github.com/gin-gonic/gin"
)

func HandleSummarize(c *gin.Context) {
	var req model.SummaryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	req.CreatedAt = time.Now().Format(time.RFC3339)
	req.OriginalLength = utf8.RuneCountInString(req.OriginalText)
	req.Summary = service.CallMLModel(req.OriginalText)
	req.SummaryLength = utf8.RuneCountInString(req.Summary)
	req.CompressionRatio = computeCompressionRatio(req.OriginalLength, req.SummaryLength)
	req.QuestionCount = len(req.Questions)
	req.Status = "completed" // or some logic to determine status

	service.SaveFullSummary(req)

	c.JSON(http.StatusOK, req)
}

func computeCompressionRatio(originalLen, summaryLen int) string {
	if originalLen == 0 {
		return "0"
	}
	ratio := float64(summaryLen) / float64(originalLen)
	return fmt.Sprintf("%.2f", ratio)
}

func GetSummaries(c *gin.Context) {
	summaries := service.FetchSummaries()
	c.JSON(http.StatusOK, gin.H{"data": summaries})
}

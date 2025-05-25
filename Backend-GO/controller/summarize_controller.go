package controller

import (
	"net/http"

	"backend-summarizer/service"
	"github.com/gin-gonic/gin"
)

type SummarizeRequest struct {
	Text   string `json:"text" binding:"required"`
}

func HandleSummarize(c *gin.Context) {
	var req SummarizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	summary := service.ProcessSummary(req.Text)
	c.JSON(http.StatusOK, gin.H{
		"original_text": req.Text,
		"summary_text":  summary,
		"timestamp":     c.MustGet("timestamp"),
	})
}

func GetSummaries(c *gin.Context) {
	summaries := service.FetchSummaries()
	c.JSON(http.StatusOK, gin.H{"data": summaries})
}

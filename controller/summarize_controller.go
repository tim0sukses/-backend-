// controller/controller.go
package controller

import (
	"net/http"

	"backend-summarizer/model"
	"backend-summarizer/service"

	"github.com/gin-gonic/gin"
)

// HandleSummarize binds the user input, processes via service, and returns ML response.
func HandleSummarize(c *gin.Context) {
	var req model.SummarizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := service.ProcessSummary(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": res.Summary})
		return
	}

	c.JSON(http.StatusOK, res)
}

// GetSummaries returns all stored summaries.
func GetSummaries(c *gin.Context) {
	summaries, _ := service.FetchSummaries()
	c.JSON(http.StatusOK, gin.H{"data": summaries})
}

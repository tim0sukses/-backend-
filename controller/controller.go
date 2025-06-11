package controller

import (
	"backend-summarizer/model"
	"backend-summarizer/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleProcessText(c *gin.Context) {
	var req model.SummarizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := service.ProcessText(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to process text"})
		return
	}

	service.SaveToDB(res)
	c.JSON(http.StatusOK, gin.H{"summary": res.Summary, "questions": res.Questions})
}

func HandleSummarize(c *gin.Context) {
	var req model.SummarizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := service.Summarize(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to summarize"})
		return
	}

	service.SaveToDB(res)
	c.JSON(http.StatusOK, gin.H{"summary": res.Summary})
}

func HandleGenerateQuestion(c *gin.Context) {
	var req model.SummarizeRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := service.GenerateQuestion(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate questions"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"questions": res.Questions})
}

func GetSummaries(c *gin.Context) {
	data, _ := service.FetchSummaries()
	c.JSON(http.StatusOK, gin.H{"data": data})
}

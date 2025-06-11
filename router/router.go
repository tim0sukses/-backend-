package router

import (
	"backend-summarizer/controller"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.Use(func(c *gin.Context) {
		c.Set("timestamp", time.Now())
		c.Next()
	})

	r.POST("/process-text", controller.HandleProcessText)
	r.POST("/summarize", controller.HandleSummarize)
	r.POST("/generate-question", controller.HandleGenerateQuestion)
	r.GET("/history", controller.GetSummaries)

	return r
}

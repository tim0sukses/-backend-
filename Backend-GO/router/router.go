package router

import (
	"backend-summarizer/controller"
	"time"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Middleware to add timestamp
	r.Use(func(c *gin.Context) {
		c.Set("timestamp", time.Now())
		c.Next()
	})

	r.POST("/summarize", controller.HandleSummarize)
	r.GET("/history", controller.GetSummaries)

	return r
}

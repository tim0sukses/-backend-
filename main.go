package main

import (
	"backend-summarizer/database"
	"backend-summarizer/router"
)

func main() {
	database.InitDB()

	r := router.SetupRouter()
	r.Run(":8080")
}

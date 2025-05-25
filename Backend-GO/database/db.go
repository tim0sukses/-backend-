package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver

	"backend-summarizer/model"
)

var DB *sql.DB

func InitDB() {
	const dsn = "postgresql://postgres.rszsrrivgowtwvvqtegq:aGiWh1aWC3qLXfPZ@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres"

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to Supabase DB:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error pinging Supabase DB:", err)
	}

	log.Println("Connected to Supabase DB successfully")
}

func SaveSummary(original, summary string) {
	_, err := DB.Exec("INSERT INTO summaries (original_text, summary_text) VALUES ($1, $2)", original, summary)
	if err != nil {
		log.Println("Error inserting summary:", err)
	}
}

func GetAllSummaries() []model.Summary {
	rows, err := DB.Query("SELECT original_text, summary_text, created_at FROM summaries ORDER BY created_at DESC")
	if err != nil {
		log.Println("Error querying summaries:", err)
		return nil
	}
	defer rows.Close()

	var summaries []model.Summary
	for rows.Next() {
		var s model.Summary
		err = rows.Scan(&s.OriginalText, &s.SummaryText, &s.CreatedAt)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		summaries = append(summaries, s)
	}
	return summaries
}

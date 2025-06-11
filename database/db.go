package database

import (
	"backend-summarizer/model"
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	dsn := "postgresql://postgres.rszsrrivgowtwvvqtegq:aGiWh1aWC3qLXfPZ@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres"
	var err error

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to DB:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Error pinging DB:", err)
	}

	log.Println("Connected to DB successfully")
}

func SaveSummary(res model.SummaryResponse) {
	_, err := DB.Exec(`
		INSERT INTO summaries (original_text, summary, created_at)
		VALUES ($1, $2, $3)`,
		res.OriginalText, res.Summary, res.CreatedAt,
	)
	if err != nil {
		log.Println("Error inserting summary:", err)
	}
}

func GetAllSummaries() []model.SummaryResponse {
	rows, err := DB.Query(`
		SELECT original_text, summary, created_at 
		FROM summaries ORDER BY created_at DESC`)
	if err != nil {
		log.Println("Error querying summaries:", err)
		return nil
	}
	defer rows.Close()

	var summaries []model.SummaryResponse
	for rows.Next() {
		var res model.SummaryResponse
		err := rows.Scan(&res.OriginalText, &res.Summary, &res.CreatedAt)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		summaries = append(summaries, res)
	}
	return summaries
}

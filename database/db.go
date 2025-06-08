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

//nyimpen data ke supabase
func SaveSummary(res model.SummaryResponse) {
	_, err := DB.Exec(
		`INSERT INTO summaries (
			original_text, summary, created_at
		) VALUES ($1, $2, $3)`,
		res.OriginalText,
		res.Summary,
		res.CreatedAt,
	)
	if err != nil {
		log.Println("Error inserting summary:", err)
	}
}

//ngambil data dari supabase
func GetAllSummaries() []model.SummaryResponse {
	rows, err := DB.Query(
		`SELECT original_text, summary, created_at
		FROM summaries ORDER BY created_at DESC`,
	)
	if err != nil {
		log.Println("Error querying summaries:", err)
		return nil
	}
	defer rows.Close()

	var list []model.SummaryResponse
	for rows.Next() {
		var s model.SummaryResponse
		err := rows.Scan(
			&s.OriginalText,
			&s.Summary,
			&s.CreatedAt,
		)
		if err != nil {
			log.Println("Error scanning summary row:", err)
			continue
		}
		list = append(list, s)
	}
	return list
}

// database/database.go
package database

import (
	"database/sql"
	"log"

	"backend-summarizer/model"

	"github.com/lib/pq"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	const dsn = "postgresql://postgres.rszsrrivgowtwvvqtegq:aGiWh1aWC3qLXfPZ@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres"
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

// SaveSummary saves a full SummaryResponse record to the DB.
// SaveSummary saves a full SummaryResponse record to the DB.
func SaveSummary(res model.SummaryResponse) {
	// Convert questions slice to Postgres array
	questionsArray := pq.StringArray(res.Questions)
	_, err := DB.Exec(
		`INSERT INTO summaries (
			compression_ratio, methods_used_questions, methods_used_summary,
			original_length, original_text, processing_mode,
			question_count, questions, status, summary,
			summary_length, created_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`,
		res.CompressionRatio,
		res.MethodsUsed.Questions,
		res.MethodsUsed.Summary,
		res.OriginalLength,
		res.OriginalText,
		res.ProcessingMode,
		res.QuestionCount,
		questionsArray,
		res.Status,
		res.Summary,
		res.SummaryLength,
		res.CreatedAt,
	)
	if err != nil {
		log.Println("Error inserting summary:", err)
	}
}

// GetAllSummaries now returns full SummaryResponse list.
func GetAllSummaries() []model.SummaryResponse {
	rows, err := DB.Query(
		`SELECT compression_ratio, methods_used_questions, methods_used_summary,
		original_length, original_text, processing_mode,
		question_count, questions, status, summary,
		summary_length, created_at
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
		var questions pq.StringArray
		err := rows.Scan(
			&s.CompressionRatio,
			&s.MethodsUsed.Questions,
			&s.MethodsUsed.Summary,
			&s.OriginalLength,
			&s.OriginalText,
			&s.ProcessingMode,
			&s.QuestionCount,
			&questions,
			&s.Status,
			&s.Summary,
			&s.SummaryLength,
			&s.CreatedAt,
		)
		if err != nil {
			log.Println("Error scanning summary row:", err)
			continue
		}
		s.Questions = []string(questions)
		list = append(list, s)
	}
	return list
}

package database

import (
	"database/sql"
	"log"

	"backend-summarizer/model"

	"github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	const dsn = "postgresql://postgres.rszsrrivgowtwvvqtegq:aGiWh1aWC3qLXfPZ@aws-0-ap-southeast-1.pooler.supabase.com:5432/postgres"
	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to Supabase DB:", err)
	}
	if err = DB.Ping(); err != nil {
		log.Fatal("Error pinging Supabase DB:", err)
	}
	log.Println("Connected to Supabase DB successfully")
}

func SaveSummary(s model.SummaryRequest) {
	_, err := DB.Exec(`
		INSERT INTO summaries (
			compression_ratio, methods_used_questions, methods_used_summary,
			original_length, original_text, processing_mode,
			question_count, questions, status, summary,
			summary_length, created_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)`,
		s.CompressionRatio, s.MethodsUsed.Questions, s.MethodsUsed.Summary,
		s.OriginalLength, s.OriginalText, s.ProcessingMode,
		s.QuestionCount, pq.Array(s.Questions), s.Status, s.Summary,
		s.SummaryLength, s.CreatedAt,
	)
	if err != nil {
		log.Println("Error inserting summary:", err)
	}
}

func GetAllSummaries() []model.SummaryRequest {
	rows, err := DB.Query(`
		SELECT compression_ratio, methods_used_questions, methods_used_summary,
		       original_length, original_text, processing_mode,
		       question_count, questions, status, summary,
		       summary_length, created_at
		FROM summaries ORDER BY created_at DESC`)
	if err != nil {
		log.Println("Error querying summaries:", err)
		return nil
	}
	defer rows.Close()

	var summaries []model.SummaryRequest
	for rows.Next() {
		var s model.SummaryRequest
		err := rows.Scan(
			&s.CompressionRatio, &s.MethodsUsed.Questions, &s.MethodsUsed.Summary,
			&s.OriginalLength, &s.OriginalText, &s.ProcessingMode,
			&s.QuestionCount, pq.Array(&s.Questions), &s.Status, &s.Summary,
			&s.SummaryLength, &s.CreatedAt,
		)
		if err != nil {
			log.Println("Error scanning row:", err)
			continue
		}
		summaries = append(summaries, s)
	}
	return summaries
}

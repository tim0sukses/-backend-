# Backend Summarizer

## Struktur Project

- **cmd/main.go**: Entry point aplikasi
- **config/config.go**: Konfigurasi aplikasi (DB URL, dll)
- **controller/**: Handler HTTP
- **service/**: Logika pemrosesan ringkasan
- **database/**: Koneksi & operasi DB
- **model/**: Definisi model data
- **router/**: Setup routing
- **graphql/**: (Opsional) Handler GraphQL

## Cara Menjalankan

1. Siapkan database PostgreSQL dan buat tabel:
   ```sql
   CREATE TABLE summaries (
     id SERIAL PRIMARY KEY,
     original_text TEXT NOT NULL,
     summary_text TEXT NOT NULL,
     created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
   );
   ```
2. Jalankan ML service di `localhost:5000`.
3. Jadikan root project, lalu:
   ```bash
   go mod tidy
   go run cmd/main.go
   ```

#  Backend Summarizer

API backend untuk meringkas teks Bahasa Indonesia secara otomatis menggunakan model Python (TextRank) dan menyimpan hasil ringkasan ke PostgreSQL (Supabase).

---

## Fitur
- 🔁 `POST /summarize` — ringkas teks panjang
- 📜 `GET /history` — lihat semua ringkasan yang tersimpan di database

##  Cara Menjalankan 
Pastikan sudah menginstal Docker Desktop dan mengaktifkannya, lalu jalankan command di bawah pada terminal didalam project ini
```bash
docker build -t backend-summarizer .
docker run -p 8080:8080 backend-summarizer
```
## Cara Menggunakan Endpoint
POST http://localhost:8080/summarize
Request JSON:
```bash
{
  "text": "Teks panjangmu...",
  "length": "short"/"medium"/"long",
  "include_questions": true/false,
  "num_questions": masukan angka
}

```
GET http://localhost:8080/history


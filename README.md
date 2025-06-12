#  Backend Summarizer

API backend untuk meringkas teks Bahasa Indonesia secara otomatis menggunakan model Python dan menyimpan hasil ringkasan ke PostgreSQL (Supabase).

---

## Fitur
- 🔁 `POST /summarize` — Meringkas teks panjang.
- ❓ `POST /generate-question` — Menghasilkan pertanyaan dari teks.
- ⚙️ `POST /process-text` — Menghasilkan ringkasan & pertanyaan sekaligus dari satu input teks.
- 📜 `GET /history` — Melihat semua ringkasan yang tersimpan di database (Questions tidak termasuk).

##  Cara Menjalankan 
Pastikan sudah menginstal Docker Desktop dan mengaktifkannya, lalu jalankan command di bawah pada terminal didalam project ini
```bash
docker build -t backend-summarizer .
docker run -p 8080:8080 backend-summarizer
```
## Cara Menggunakan Endpoint
POST http://localhost:8080/summarize
POST http://localhost:8080/generate-question
POST http://localhost:8080/process-text
Request JSON:
```bash
{
  "text": "Teks panjangmu...",
}

```
GET http://localhost:8080/history


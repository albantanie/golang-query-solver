# Golang Query Solver

![License](https://img.shields.io/badge/license-MIT-blue.svg)
![Go Version](https://img.shields.io/badge/go-%3E%3D1.16-blue.svg)

Golang Query Solver adalah sebuah proyek yang menyediakan endpoint API untuk menjawab pertanyaan berdasarkan konteks yang diberikan menggunakan model dari Hugging Face. Proyek ini ditulis dalam bahasa Go.

## Fitur

- Endpoint `/qa` untuk menerima pertanyaan dan konteks.
- Menggunakan model `distilbert-base-uncased-distilled-squad` dari Hugging Face untuk memproses dan menjawab pertanyaan.
- Token API Hugging Face diambil dari variabel lingkungan.

## Prasyarat

- Go 1.16 atau versi lebih baru
- Token API dari Hugging Face

## Instalasi

1. Clone repository ini:
    ```sh
    git clone https://github.com/albantanie/golang-query-solver.git
    cd golang-query-solver
    ```

2. Install dependencies:
    ```sh
    go mod tidy
    ```

3. Setel variabel lingkungan untuk token API Hugging Face:
    ```sh
    export HUGGING_FACE_API_TOKEN=your_hugging_face_api_token
    ```

## Penggunaan

1. Jalankan server:
    ```sh
    go run main.go
    ```

2. Kirim permintaan POST ke endpoint `/qa`:
    ```sh
    curl -X POST http://localhost:8080/qa \
         -H "Content-Type: application/json" \
         -d '{
               "question": "Apa ibu kota Indonesia?",
               "context": "Indonesia adalah sebuah negara di Asia Tenggara yang terdiri dari lebih dari 17.000 pulau. Negara ini memiliki populasi lebih dari 270 juta orang. Ibu kota Indonesia adalah Jakarta, yang terletak di pulau Jawa."
             }'
    ```

3. Anda akan menerima respons dalam format JSON yang berisi jawaban dari model:

    ```json
    [
        {
            "score": 0.9989,
            "answer": "Jakarta"
        }
    ]
    ```

## Struktur Proyek

- `main.go`: Entry point dari aplikasi.
- `handlers/handlers.go`: Mengandung handler untuk endpoint `/qa`.
- `utils/utils.go`: Mengandung fungsi utilitas untuk membuat permintaan POST ke API Hugging Face.
- `models/models.go`: Mengandung definisi struct untuk Question-Answer dan Answer.
- `config/config.go`: Mengandung fungsi untuk mengambil token API dari variabel lingkungan.

## Kontribusi

Kontribusi sangat diterima! Silakan fork repository ini dan buat pull request dengan perubahan Anda. Pastikan untuk mengikuti panduan kontribusi kami.

## Lisensi

Proyek ini dilisensikan di bawah lisensi MIT. Lihat file [LICENSE](LICENSE) untuk informasi lebih lanjut.

## Kontak

Jika Anda memiliki pertanyaan atau masalah, silakan buat issue di repository ini atau hubungi saya di syahrulfaiz30102003@gmail.com.


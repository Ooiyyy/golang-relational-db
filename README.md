📘 Golang Relational DB API
Deskripsi

Project ini adalah REST API berbasis Golang (Gin) untuk mengelola data:

Students
Classes
Subjects
Grades

Fitur utama:

CRUD dasar
Pagination, search, dan sorting
Relasi antar tabel (JOIN)
Perhitungan rata-rata nilai siswa
⚙️ Setup Awal (Prerequisites)

Pastikan sudah menginstall:

1. Golang

Minimal versi 1.20+

Cek:

go version

Jika belum:
https://go.dev/dl/

2. MySQL / MariaDB

Pastikan database server aktif.

Cek:

mysql --version
3. golang-migrate (CLI)

Tool ini digunakan untuk migration database.

Install:

go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest

Tambahkan ke PATH (jika belum):

export PATH=$PATH:$HOME/go/bin

Cek:

migrate -version
🚀 Cara Menjalankan Project
1. Clone Repository
git clone https://github.com/username/golang-relational-db.git
cd golang-relational-db
2. Install Dependency
go mod tidy
3. Konfigurasi Environment

Copy file .env.example menjadi .env lalu sesuaikan dengan konfigurasi lokal

4. Setup Database
Buat database:
CREATE DATABASE golang_study;
Jalankan migration:
migrate -path db/migrations -database "mysql://root:password@tcp(localhost:3306)/golang_study" up
5. Jalankan Server
go run cmd/api/main.go

Server akan berjalan di:

http://localhost:8080
📌 Contoh Endpoint
1. Create Student
POST /api/v1/students

Request body:

{
  "name": "Budi",
  "email": "budi@mail.com",
  "class_id": 1
}
2. Get Students (Pagination + Search)
GET /api/v1/students?page=1&limit=10&search=budi
3. Get Student Detail
GET /api/v1/students/1
4. Student dengan Relasi Class & Teacher
GET /api/v1/student/classes
5. Rata-rata Nilai Student
GET /api/v1/student/avg
📁 Struktur Project
cmd/
internal/
  ├── handler/
  ├── service/
  ├── repository/
  ├── model/
  ├── dto/
  ├── routes/
db/
  └── migrations/
Catatan
Gunakan Postman untuk testing endpoint
Pastikan database sudah aktif sebelum menjalankan server
Migration harus dijalankan sebelum API digunakan
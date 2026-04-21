# Golang Relational DB API

REST API sederhana menggunakan Golang (Gin) untuk mengelola data siswa, kelas, mata pelajaran, dan nilai.
Mendukung relasi antar tabel (JOIN), pagination, search, dan perhitungan rata-rata nilai.

## Fitur

- CRUD dasar (students, classes, subjects, grades)
- Pagination, search, dan sorting
- Relasi data (student + class + teacher)
- Aggregate (average nilai siswa)

## Prerequisites

Pastikan sudah terinstall:

1. **Golang**
   ```bash
   go version
   ```
2. **MySQL / MariaDB**
   ```bash
   mysql --version
   ```
3. **golang-migrate**
   ```bash
   go install github.com/golang-migrate/migrate/v4/cmd/migrate@latest
   ```

Tambahkan ke PATH jika perlu:

```bash
export PATH=$PATH:$HOME/go/bin
```

## Setup Project

**1. Clone repository**
```bash
git clone https://github.com/username/golang-relational-db.git
cd golang-relational-db
```

**2. Install dependency**
```bash
go mod tidy
```

**3. Setup environment**

Copy file:

```bash
cp .env.example .env
```

Edit sesuai konfigurasi:

```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=golang_study
APP_PORT=8080
```

**4. Setup database**

Buat database:

```sql
CREATE DATABASE golang_study;
```

Jalankan migration:

```bash
migrate -path db/migrations -database "mysql://root:password@tcp(localhost:3306)/golang_study" up
```

**5. Jalankan server**
```bash
go run cmd/api/main.go
```

Server berjalan di:

`http://localhost:8080`

## Contoh Endpoint

### Create Student
`POST /api/v1/students`

Request:

```json
{
  "name": "Budi",
  "email": "budi@mail.com",
  "class_id": 1
}
```

### Get Students
`GET /api/v1/students?page=1&limit=10&search=budi`

### Get Student Detail
`GET /api/v1/students/1`

### Student dengan Class & Teacher
`GET /api/v1/student/classes`

### Rata-rata Nilai Student
`GET /api/v1/student/avg`

## Struktur Project

```text
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
```

## Catatan
- Gunakan Postman untuk testing endpoint
- Pastikan database aktif sebelum menjalankan server
- Jalankan migration sebelum menggunakan API
- File `.env` tidak disertakan di repository (gunakan `.env.example`)
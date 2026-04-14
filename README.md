# 🏫 Sistem Informasi Akademik API (Golang)

Selamat datang di repositori Sistem Informasi Akademik (Studi Kasus Sekolah). Aplikasi ini merupakan RESTful API yang ditulis dalam bahasa **Go (Golang)** menggunakan arsitektur **Clean Architecture**.

Awalnya merupakan *clone* dari proyek Todo-List, proyek ini telah berkembang pesat menjadi sistem akademik yang mendukung Manajemen Pengguna (Guru & Siswa), Mata Pelajaran, dan Nilai Siswa beserta kontrol akses berbasis Role (RBAC).

## 🌟 Fitur Utama
- **Autentikasi & Otorisasi JWT**: Sistem pendaftaran dan login yang aman menggunakan sandi *bcrypt* dan token JWT.
- **Role-Based Access Control (RBAC)**: Pembatasan hak akses *endpoint* menggunakan Middleware. 
  - 👨‍🏫 **Guru**: Memiliki akses penuh (Create, Update, Delete) untuk Mengelola Mata Pelajaran dan Nilai-nilai Siswa.
  - 👨‍🎓 **Siswa**: Hanya memiliki hak akses baca (Read-only) untuk melihat Mata Pelajaran dan Nilai.
- **CRUD Mata Pelajaran (Subjects)**: Fitur otomatisasi kepemilikan. `Teacher_id` disisipkan otomatis dari sesi Guru yang sedang Login.
- **CRUD Nilai Siswa (Scores)**: Manipulasi penilaian dengan pencarian efisien berbasis `student_id` & `subject_id` maupun Primary Key.
- **Dukungan Utilitas Skalabel**: Modul standarisasi JSON Validation Error dan JSON HTTP Response.


## 🛠 Teknologi yang Digunakan
- **Bahasa**: [Go (Golang)](https://go.dev/)
- **Framework Web**: [Gin Gonic](https://gin-gonic.com/)
- **Database**: MySQL dengan modul bawaan `database/sql`
- **Keamanan**: `golang.org/x/crypto/bcrypt` & `github.com/golang-jwt/jwt/v5`

---

## 💾 Skema Database
Sistem ini menggunakan 3 tabel utama:

1. **`users`**
   Menyimpan kredensial autentikasi dan peran. Kolom: `id`, `username`, `email`, `password`, `role` (enum: 'guru', 'siswa'), `created_at`.
2. **`subjects`**
   Menyimpan daftar mata pelajaran. Kolom: `id`, `name`, `teacher_id` (FK -> users.id).
3. **`scores`**
   Menyimpan nilai siswa per mapel. Kolom: `id`, `student_id` (FK -> users.id), `subject_id` (FK -> subjects.id), `score`, `created_at`.

---

## 🚀 Cara Menjalankan Aplikasi
1. Pastikan memiliki Go dan MySQL yang tersambung.
2. Atur konfigurasi koneksi MySQL di dalam `config/database_config.go` (atau file `.env`).
3. Jalankan server dengan perintah:
   ```bash
   go run ./cmd
   ```
4. Server akan otomatis menyala di **Port 8080** (http://localhost:8080).

---

## 📚 Panduan Struktur Folder (Clean Architecture)
Tujuan dari membagi-bagi file ke dalam beberapa folder adalah agar kode kita lebih rapi, gampang dites secara individual, dan lebih mudah dikembangkan jika aplikasinya semakin membesar.

*(Di bawah ini adalah penjelasan sederhana khusus untuk Pemula)*

### 📁 `cmd/`
- **Fungsi:** Titik utama di mana program ini mulai dijalankan `(go run ./cmd)`.
- **Analogi:** Seperti "Gerbang Pintu Masuk" dari sebuah pabrik.

### 📁 `config/`
- **Fungsi:** Berisi konfigurasi lingkungan global aplikasi (seperti Database & JWT Secret).
- **Analogi:** "Ruang Panel Mesin Listrik".

### 📁 `internal/` (Rahasia & Tertutup)
Ini adalah map tertutup khas Go. Inti dari semua rahasia logika bisnis aplikasi diletakkan di bawah payung folder ini dan tidak bisa diekspor oleh aplikasi luar.

#### 📂 `internal/handler/` (Controllers)
- **Fungsi:** Tempat bertemunya HTTP Request JSON (dari Browser/Postman).
- **Isi:** Membongkar JSON ke Struct (DTO) dan mencetak Response. Tidak boleh ada logika bisnis berat di sini.
- **Analogi:** "Pelayan Restoran" yang mencatat dan mengirim pesanan.

#### 📂 `internal/middleware/`
- **Fungsi:** Pasukan pencegat spesifik di tengah rute.
- **Isi:** Terdapat `jwt.go` (verifikasi Token otentikasi) dan `role.go` (pengecekan level otorisasi akses 'guru'/'siswa').
- **Analogi:** "Satpam Resepsionis" yang mengecek KTP pengunjung.

#### 📂 `internal/service/` (Business Logic)
- **Fungsi:** "Otak" utama aplikasi.
- **Isi:** Disinilah diputuskan aturan seperti enkripsi password, mengecek duplikasi sebelum Insert DB, dsb.
- **Analogi:** "Koki Spesialis Dapur" yang meracik masakan menggunakan resep rahasia restoran.

#### 📂 `internal/repository/` (Database Layer)
- **Fungsi:** Pengkoneksi aplikasi langsung ke MySQL.
- **Isi:** Seluruh perintah SQL mentah (SELECT, UPDATE, INSERT, DELETE) hidup di level ini. 
- **Analogi:** "Petugas Gudang Rak Buku".

#### 📂 `internal/dto/` (Data Transfer Object)
- **Fungsi:** Jembatan pembungkus struktur format body Request/Response. Mengandung validasi tag _binding_ bawaan Gin.

#### 📂 `internal/utils/` (Helper)
- **Fungsi:** Menaruh fungsi mungil serba-guna untuk merapikan pesan Error/Validasi/Response agar struktur JSON kembalian selalu seragam.
- **Analogi:** "Pisau Dapur Multifungsi".

---
*Happy Coding & Selamat Belajar Clean Architecture dengan Golang!* ✨

package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// Sesuaikan DSN (Data Source Name) dengan konfigurasi database kamu
	db, err := sql.Open("mysql", "root:Gridy20@tcp(localhost:3306)/golang_study")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 1. Insert Teachers (Guru) - Dibutuhkan oleh Classes dan Subjects
	_, err = db.Exec(`
		INSERT INTO teachers (id, name, email) VALUES
		(1, 'Budi Santoso', 'budi.santoso@sekolah.id'),
		(2, 'Siti Aminah', 'siti.aminah@sekolah.id')
	`)
	if err != nil {
		log.Fatal("Error seeding teachers: ", err)
	}

	// 2. Insert Classes (Kelas) - Dibutuhkan oleh Students
	_, err = db.Exec(`
		INSERT INTO classes (id, name, teacher_id) VALUES
		(1, 'Kelas 10-A', 1),
		(2, 'Kelas 10-B', 2)
	`)
	if err != nil {
		log.Fatal("Error seeding classes: ", err)
	}

	// 3. Insert Students (Siswa) - Dibutuhkan oleh Grades
	_, err = db.Exec(`
		INSERT INTO students (id, name, email, class_id) VALUES
		(1, 'Andi Wijaya', 'andi.wijaya@student.id', 1),
		(2, 'Bunga Citra', 'bunga.citra@student.id', 1),
		(3, 'Candra Gupta', 'candra.gupta@student.id', 2)
	`)
	if err != nil {
		log.Fatal("Error seeding students: ", err)
	}

	// 4. Insert Subjects (Mata Pelajaran) - Dibutuhkan oleh Grades
	_, err = db.Exec(`
		INSERT INTO subjects (id, name, teacher_id) VALUES
		(1, 'Matematika', 1),
		(2, 'Bahasa Inggris', 2),
		(3, 'Fisika', 1)
	`)
	if err != nil {
		log.Fatal("Error seeding subjects: ", err)
	}

	// 5. Insert Grades (Nilai)
	_, err = db.Exec(`
		INSERT INTO grades (student_id, subject_id, score) VALUES
		(1, 1, 85), -- Andi, Matematika
		(1, 2, 90), -- Andi, B. Inggris
		(2, 1, 78), -- Bunga, Matematika
		(3, 2, 88), -- Candra, B. Inggris
		(3, 3, 82)  -- Candra, Fisika
	`)
	if err != nil {
		log.Fatal("Error seeding grades: ", err)
	}

	log.Println("Seeding database berhasil dilakukan!")
}

package mysql

import (
	"database/sql"
	"golang-relational-db/internal/model"
)

// struct repo menyimpan koneksi DB
type StudentRepoImpl struct {
	DB *sql.DB
}

// constructor
func NewStudentRepo(db *sql.DB) *StudentRepoImpl {
	return &StudentRepoImpl{DB: db}
}

// Insert data student ke database
func (r *StudentRepoImpl) Create(student *model.Students) error {
	result, err := r.DB.Exec(
		"INSERT INTO students (name, email, class_id) VALUES (?, ?, ?)",
		student.Name,
		student.Email,
		student.ClassID,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	student.ID = int(id)
	return err // hanya return error
}

func (r *StudentRepoImpl) FindByID(id int) (*model.Students, error) {
	var s model.Students

	// ambil 1 data berdasarkan id
	err := r.DB.QueryRow(
		"SELECT id, name, email, class_id FROM students WHERE id = ?",
		id,
	).Scan(&s.ID, &s.Name, &s.Email, &s.ClassID)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // ⚠️ ini penting → tanda data tidak ada
		}
		return nil, err // error lain (DB error)
	}

	return &s, nil
}

func (r *StudentRepoImpl) FindAll(limit, offset int) ([]model.Students, error) {
	rows, err := r.DB.Query("SELECT id, name, email, class_id FROM students LIMIT ? OFFSET ?", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []model.Students

	// loop semua hasil query
	for rows.Next() {
		var s model.Students
		err := rows.Scan(&s.ID, &s.Name, &s.Email, &s.ClassID)
		if err != nil {
			return nil, err
		}
		students = append(students, s)
	}

	return students, nil
}
func (r *StudentRepoImpl) Count() (int, error) {
	var total int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM students").Scan(&total)
	return total, err
}

func (r *StudentRepoImpl) Update(id int, student model.Students) error {
	_, err := r.DB.Exec(
		"UPDATE students SET name=?, email=?, class_id=? WHERE id=?",
		student.Name,
		student.Email,
		student.ClassID,
		id,
	)
	return err
}

func (r *StudentRepoImpl) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM students WHERE id=?", id)
	return err
}

package mysql

import (
	"database/sql"
	"errors"
	"fmt"
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

func (r *StudentRepoImpl) IsStudentEmailExists(email string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) from students where email=?`
	err := r.DB.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, nil
	}
	return count > 0, nil
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
			return nil, errors.New("siswa tidak ditemukan") // ⚠️ ini penting → tanda data tidak ada
		}
		return nil, err // error lain (DB error)
	}

	return &s, nil
}

func (r *StudentRepoImpl) FindAll(limit, offset int, search string, classID int, sortBy, order string) ([]model.Students, error) {
	query := "SELECT id, name, email, class_id FROM students WHERE 1=1"
	args := []interface{}{}

	if search != "" {
		query += " AND (LOWER(name) LIKE LOWER(?) OR LOWER(email) LIKE LOWER(?))"
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	if classID != 0 {
		query += " AND class_id =?"
		args = append(args, classID)
	}

	query += " ORDER BY " + sortBy + " " + order

	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	students := []model.Students{}

	for rows.Next() {
		var student model.Students
		err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.ClassID)
		if err != nil {
			return nil, err
		}
		students = append(students, student)
	}
	fmt.Println("QUERY:", query)
	fmt.Println("ARGS:", args)
	return students, nil
}

func (r *StudentRepoImpl) Count(search string, classID int) (int, error) {
	query := "SELECT COUNT(*) FROM students WHERE 1=1"
	args := []interface{}{}

	if search != "" {
		query += " AND name LIKE ? OR email LIKE ?"
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	if classID != 0 {
		query += " AND class_id = ?"
		args = append(args, classID)
	}

	var total int
	err := r.DB.QueryRow(query, args...).Scan(&total)
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

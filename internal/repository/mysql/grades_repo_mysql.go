package mysql

import (
	"database/sql"
	"golang-relational-db/internal/model"
)

// struct repo menyimpan koneksi DB
type GradesRepoImpl struct {
	DB *sql.DB
}

// constructor
func NewGradesRepo(db *sql.DB) *GradesRepoImpl {
	return &GradesRepoImpl{DB: db}
}

// Insert data student ke database
func (r *GradesRepoImpl) Create(grade *model.Grades) error {
	result, err := r.DB.Exec(
		"INSERT INTO grades (student_id, subject_id, score) VALUES (?, ?, ?)",
		grade.StudentID,
		grade.SubjectID,
		grade.Score,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	grade.ID = int(id)
	return err // hanya return error
}

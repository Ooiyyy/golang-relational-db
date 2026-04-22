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

func (r *GradesRepoImpl) FindGradeByID(id int) (*model.GradesDetail, error) {
	var grade model.GradesDetail

	// ambil 1 data berdasarkan id
	err := r.DB.QueryRow(
		`SELECT g.id, 
		s.id, s.name AS nama_siswa,
		g.score AS nilai,
		sj.id, sj.name AS mapel
		FROM grades AS g
		JOIN students AS s ON s.id = g.student_id
		JOIN subjects AS sj ON sj.id = g.subject_id 
		WHERE g.id=?`, id).Scan(
		&grade.ID, &grade.StudentID, &grade.StudentName, &grade.Score, &grade.SubjectID, &grade.SubjectName,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // ⚠️ ini penting → tanda data tidak ada
		}
		return nil, err // error lain (DB error)
	}
	return &grade, nil
}

func (r *GradesRepoImpl) UpdateGrade(id int, grade model.Grades) error {
	_, err := r.DB.Exec(
		"UPDATE grades SET student_id=?, subject_id=?, score=? WHERE id=?",
		grade.StudentID,
		grade.SubjectID,
		grade.Score,
		id,
	)
	return err
}

func (r *GradesRepoImpl) DeleteGrade(id int) error {
	_, err := r.DB.Exec("DELETE FROM grades WHERE id=?", id)
	return err
}

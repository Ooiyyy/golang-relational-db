package mysql

import (
	"database/sql"
	"fmt"
	"golang-relational-db/internal/model"
)

// struct repo menyimpan koneksi DB
type SubjectsRepoImpl struct {
	DB *sql.DB
}

// constructor
func NewSubjectsRepo(db *sql.DB) *SubjectsRepoImpl {
	return &SubjectsRepoImpl{DB: db}
}

// Insert data student ke database
func (r *SubjectsRepoImpl) Create(subject *model.Subjects) error {
	result, err := r.DB.Exec(
		"INSERT INTO subjects (name, teacher_id) VALUES (?, ?)",
		subject.Name,
		subject.TeacherID,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	subject.Id = int(id)
	return err // hanya return error
}

func (r *SubjectsRepoImpl) FindAllSubjects(limit, offset int, search string, sortBy, order string) ([]model.Subjects, error) {
	query := "SELECT * FROM subjects WHERE 1=1"
	args := []interface{}{}

	if search != "" {
		query += " AND (LOWER(name) LIKE LOWER(?))"
		args = append(args, "%"+search+"%")
	}

	// if classID != 0 {
	// 	query += " AND class_id =?"
	// 	args = append(args, classID)
	// }

	query += " ORDER BY " + sortBy + " " + order

	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []model.Subjects

	for rows.Next() {
		var subject model.Subjects
		err := rows.Scan(&subject.Id, &subject.Name, &subject.TeacherID)
		if err != nil {
			return nil, err
		}
		subjects = append(subjects, subject)
	}
	fmt.Println("QUERY:", query)
	fmt.Println("ARGS:", args)
	return subjects, nil
}

func (r *SubjectsRepoImpl) AllSubjects(search string) (int, error) {
	query := "SELECT COUNT(*) FROM subjects WHERE 1=1"
	args := []interface{}{}

	if search != "" {
		query += " AND name LIKE ?"
		args = append(args, "%"+search+"%")
	}

	var total int
	err := r.DB.QueryRow(query, args...).Scan(&total)
	return total, err
}

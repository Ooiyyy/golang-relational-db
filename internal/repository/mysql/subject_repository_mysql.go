package mysql

import (
	"database/sql"
	"golang-sekolah/internal/dto"
	"golang-sekolah/internal/model"
)

type SubjectRepoImpl struct {
	DB *sql.DB
}

func NewSubjectRepo(db *sql.DB) *SubjectRepoImpl {
	return &SubjectRepoImpl{
		DB: db,
	}
}

func (s *SubjectRepoImpl) Create(req dto.SubjectReq) error {
	_, err := s.DB.Exec("INSERT INTO subjects (name, teacher_id) VALUES (?, ?)", req.Name, req.Teacher_id)
	return err
}

func (s *SubjectRepoImpl) FindAll() ([]model.Subjects, error) {
	rows, err := s.DB.Query("SELECT id, name, teacher_id FROM subjects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var subjects []model.Subjects
	for rows.Next() {
		var sub model.Subjects
		if err := rows.Scan(&sub.Id, &sub.Name, &sub.Teacher_id); err != nil {
			return nil, err
		}
		subjects = append(subjects, sub)
	}
	return subjects, nil
}

func (s *SubjectRepoImpl) FindByID(id int) (*model.Subjects, error) {
	var sub model.Subjects
	err := s.DB.QueryRow("SELECT id, name, teacher_id FROM subjects WHERE id = ?", id).Scan(&sub.Id, &sub.Name, &sub.Teacher_id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Data tidak ditemukan
		}
		return nil, err
	}
	return &sub, nil
}

func (s *SubjectRepoImpl) Update(id int, req dto.SubjectReq) error {
	_, err := s.DB.Exec("UPDATE subjects SET name = ?, teacher_id = ? WHERE id = ?", req.Name, req.Teacher_id, id)
	return err
}

func (s *SubjectRepoImpl) Delete(id int) error {
	_, err := s.DB.Exec("DELETE FROM subjects WHERE id = ?", id)
	return err
}

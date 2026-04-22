package service

import (
	"fmt"
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/repository"
)

// service pegang business logic
type SubjectsService struct {
	repo repository.SubjectsCepository
}

// constructor
func NewSubjectsService(repo repository.SubjectsCepository) *SubjectsService {
	return &SubjectsService{repo: repo}
}

// Create student baru
func (s *SubjectsService) CreateSubjects(subject *model.Subjects) error {
	// bisa tambah validasi manual di sini kalau mau
	return s.repo.Create(subject)
}

func (s *SubjectsService) GetAllSubjects(page, limit int, search, sortBy, order string) ([]model.Subjects, int, error) {
	if page < 1 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	allowedSort := map[string]bool{
		"id":   true,
		"name": true,
	}
	if !allowedSort[sortBy] {
		sortBy = "id"
	}

	if order != "asc" && order != "desc" {
		order = "asc"
	}

	offset := (page - 1) * limit

	data, err := s.repo.FindAllSubjects(limit, offset, search, sortBy, order)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repo.AllSubjects(search)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (s *SubjectsService) GetSubjectByID(id int) (*model.Subjects, error) {
	subject, err := s.repo.FindSubjectByID(id)
	if err != nil {
		return nil, err // error DB
	}

	if subject == nil {
		return nil, fmt.Errorf("subject tidak ditemukan") // handle not found
	}

	return subject, nil
}

func (s *SubjectsService) UpdateSubject(id int, subject model.Subjects) error {
	// cek dulu apakah data ada menggunakan fungsi dari repo yang mencari berdasarkan id
	_, err := s.GetSubjectByID(id)
	if err != nil {
		return err // kalau tidak ada → langsung stop
	}

	return s.repo.UpdateSubject(id, subject)
}

func (s *SubjectsService) DeleteSubject(id int) error {
	// cek dulu
	_, err := s.GetSubjectByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteSubject(id)
}

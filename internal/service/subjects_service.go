package service

import (
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

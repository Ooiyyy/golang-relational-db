package service

import (
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/repository"
)

// service pegang business logic
type GradesService struct {
	repo repository.GradesRepository
}

// constructor
func NewGradesService(repo repository.GradesRepository) *GradesService {
	return &GradesService{repo: repo}
}

// Create student baru
func (s *GradesService) CreateGrades(grades *model.Grades) error {
	// bisa tambah validasi manual di sini kalau mau
	return s.repo.Create(grades)
}

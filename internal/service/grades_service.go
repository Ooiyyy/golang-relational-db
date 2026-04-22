package service

import (
	"fmt"
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

func (s *GradesService) GetGradeByID(id int) (*model.GradesDetail, error) {
	class, err := s.repo.FindGradeByID(id)
	if err != nil {
		return nil, err // error DB
	}

	if class == nil {
		return nil, fmt.Errorf("nilai tidak ditemukan") // handle not found
	}

	return class, nil
}

func (s *GradesService) UpdateGrade(id int, grade model.Grades) error {
	// cek dulu apakah data ada menggunakan fungsi dari repo yang mencari berdasarkan id
	_, err := s.GetGradeByID(id)
	if err != nil {
		return err // kalau tidak ada → langsung stop
	}

	return s.repo.UpdateGrade(id, grade)
}

func (s *GradesService) DeleteGrade(id int) error {
	// cek dulu
	_, err := s.GetGradeByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteGrade(id)
}

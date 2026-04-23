package service

import (
	"errors"
	"fmt"
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/repository"
)

// service pegang business logic
type GradesService struct {
	repo         repository.GradesRepository
	studentRepo  repository.StudentRepository
	subjectsRepo repository.SubjectsCepository
}

// constructor
func NewGradesService(repo repository.GradesRepository, studentRepo repository.StudentRepository, subjectsRepo repository.SubjectsCepository) *GradesService {
	return &GradesService{repo: repo, studentRepo: studentRepo, subjectsRepo: subjectsRepo}
}

// Create student baru
func (s *GradesService) CreateGrades(grades *model.Grades) error {
	// bisa tambah validasi manual di sini kalau mau
	_, err := s.studentRepo.FindByID(grades.StudentID)
	if err != nil {
		return errors.New("siswa tidak ditemukan")
	}

	_, err = s.subjectsRepo.FindSubjectByID(grades.SubjectID)
	if err != nil {
		return errors.New("mapel tidak ditemukan")
	}
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
	_, err = s.studentRepo.FindByID(grade.StudentID)
	if err != nil {
		return errors.New("siswa tidak ditemukan")
	}

	_, err = s.subjectsRepo.FindSubjectByID(grade.SubjectID)
	if err != nil {
		return errors.New("mapel tidak ditemukan")
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

package service

import (
	"fmt"
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/repository"
)

// service pegang business logic
type StudentService struct {
	repo repository.StudentRepository
}

// constructor
func NewStudentService(repo repository.StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

// Create student baru
func (s *StudentService) CreateStudent(student *model.Students) error {
	// bisa tambah validasi manual di sini kalau mau
	return s.repo.Create(student)
}

func (s *StudentService) GetStudentByID(id int) (*model.Students, error) {
	student, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err // error DB
	}

	if student == nil {
		return nil, fmt.Errorf("student tidak ditemukan") // handle not found
	}

	return student, nil
}

func (s *StudentService) GetAllStudents(page, limit int) ([]model.Students, int, error) {
	if page < 1 {
		page = 1
	}
	if limit <= 0 || limit > 50 {
		limit = 5
	}
	offset := (page - 1) * limit

	data, err := s.repo.FindAll(limit, offset)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repo.Count()
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (s *StudentService) UpdateStudent(id int, student model.Students) error {
	// cek dulu apakah data ada
	_, err := s.GetStudentByID(id)
	if err != nil {
		return err // kalau tidak ada → langsung stop
	}

	return s.repo.Update(id, student)
}

func (s *StudentService) DeleteStudent(id int) error {
	// cek dulu
	_, err := s.GetStudentByID(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}

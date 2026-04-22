package service

import (
	"errors"
	"fmt"
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/repository"
)

// service pegang business logic
type TeacherService struct {
	repo repository.TeacherRepository
}

// constructor
func NewTeacherService(repo repository.TeacherRepository) *TeacherService {
	return &TeacherService{repo: repo}
}

// Create student baru
func (s *TeacherService) CreateTeacher(teacher *model.Teachers) error {
	// bisa tambah validasi manual di sini kalau mau
	email := teacher.Email
	exists, err := s.repo.IsTeacherEmailExists(email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email sudah digunakan")
	}
	return s.repo.Create(teacher)
}

func (s *TeacherService) GetTeacherByID(id int) (*model.Teachers, error) {
	teacher, err := s.repo.FindTeacherByID(id)
	if err != nil {
		return nil, err // error DB
	}

	if teacher == nil {
		return nil, fmt.Errorf("guru tidak ditemukan") // handle not found
	}

	return teacher, nil
}

func (s *TeacherService) GetAllTeachers(page, limit int, search, sortBy, order string) ([]model.Teachers, int, error) {
	if page < 1 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	var err error

	allowedSort := map[string]bool{
		"id":    true,
		"name":  true,
		"email": true,
	}
	if !allowedSort[sortBy] {
		sortBy = "id"
	}

	if order != "asc" && order != "desc" {
		order = "asc"
	}

	offset := (page - 1) * limit

	data, err := s.repo.FindAllTeachers(limit, offset, search, sortBy, order)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repo.AllTeachers(search)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (s *TeacherService) UpdateTeacher(id int, teacher model.Teachers) error {
	// cek dulu apakah data ada menggunakan fungsi dari repo yang mencari berdasarkan id
	_, err := s.GetTeacherByID(id)
	if err != nil {
		return err // kalau tidak ada → langsung stop
	}
	return s.repo.UpdateTeacher(id, teacher)
}

func (s *TeacherService) DeleteTeacher(id int) error {
	// cek dulu
	_, err := s.GetTeacherByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteTeacher(id)
}

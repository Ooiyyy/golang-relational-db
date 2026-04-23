package service

import (
	"errors"
	"fmt"
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/repository"
)

// service pegang business logic
type ClassService struct {
	repo        repository.ClassRepository
	teacherRepo repository.TeacherRepository
}

// constructor
func NewClassService(repo repository.ClassRepository, teacherRepo repository.TeacherRepository) *ClassService {
	return &ClassService{repo: repo, teacherRepo: teacherRepo}
}

// Create student baru
func (s *ClassService) CreateClass(class *model.Classes) error {
	// bisa tambah validasi manual di sini kalau mau
	_, err := s.teacherRepo.FindTeacherByID(class.TeacherID)
	if err != nil {
		return errors.New("guru tidak ditemukan") // kalau tidak ada → langsung stop
	}
	return s.repo.Create(class)
}

func (s *ClassService) GetAllClass(page, limit int, search, sortBy, order string) ([]model.DetailClass, int, error) {
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

	data, err := s.repo.FindAllClasses(limit, offset, search, sortBy, order)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repo.AllClasses(search)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (s *ClassService) GetClassByID(id int) (*model.DetailClass, error) {
	class, err := s.repo.FindClassByID(id)
	if err != nil {
		return nil, err // error DB
	}

	if class == nil {
		return nil, fmt.Errorf("subject tidak ditemukan") // handle not found
	}

	return class, nil
}

func (s *ClassService) UpdateClass(id int, class model.Classes) error {
	// cek dulu apakah data ada menggunakan fungsi dari repo yang mencari berdasarkan id
	_, err := s.GetClassByID(id)
	if err != nil {
		return err // kalau tidak ada → langsung stop
	}
	_, err = s.teacherRepo.FindTeacherByID(class.TeacherID)
	if err != nil {
		return errors.New("guru tidak ditemukan") // kalau tidak ada → langsung stop
	}

	return s.repo.UpdateClass(id, class)
}

func (s *ClassService) DeleteClass(id int) error {
	// cek dulu
	_, err := s.GetClassByID(id)
	if err != nil {
		return err
	}

	return s.repo.DeleteClass(id)
}

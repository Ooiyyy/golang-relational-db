package service

import (
	"fmt"
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/repository"
)

// service pegang business logic
type ClassService struct {
	repo repository.ClassRepository
}

// constructor
func NewClassService(repo repository.ClassRepository) *ClassService {
	return &ClassService{repo: repo}
}

// Create student baru
func (s *ClassService) CreateClass(class *model.Classes) error {
	// bisa tambah validasi manual di sini kalau mau
	return s.repo.Create(class)
}

func (s *ClassService) GetAllClass(page, limit int, search, sortBy, order string) ([]model.DetailClass, int, error) {
	if page < 1 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	// var classIDInt int
	// var err error
	// if classID != "" {
	// 	classIDInt, err = strconv.Atoi(classID)
	// 	if err != nil {
	// 		return nil, 0, fmt.Errorf("class_id harus angka")
	// 	}
	// 	return nil, classIDInt, err
	// }

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

package service

import (
	"fmt"
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/repository"
	"strconv"
)

// service pegang business logic
type StudentRelationalService struct {
	repo repository.StudentRelational
}

// constructor
func NewStudentRelationalService(repo repository.StudentRelational) *StudentRelationalService {
	return &StudentRelationalService{repo: repo}
}

func (s *StudentRelationalService) GetStudentsDetail(page, limit int, search, classID, sortBy, order string) ([]model.StudentDetail, int, error) {
	if page < 1 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	var classIDInt int
	var err error
	if classID != "" {
		classIDInt, err = strconv.Atoi(classID)
		if err != nil {
			return nil, 0, fmt.Errorf("class_id harus angka")
		}
		return nil, classIDInt, err
	}

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

	data, err := s.repo.StudentsDetail(limit, offset, search, classIDInt, sortBy, order)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repo.Pages(search, classIDInt)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (s *StudentRelationalService) GetStudentsGrade(page, limit int, search, classID, sortBy, order string) ([]model.StudentsGrade, int, error) {
	if page < 1 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	var classIDInt int
	var err error
	if classID != "" {
		classIDInt, err = strconv.Atoi(classID)
		if err != nil {
			return nil, 0, fmt.Errorf("class_id harus angka")
		}
		return nil, classIDInt, err
	}

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

	data, err := s.repo.StudentsGrade(limit, offset, search, classIDInt, sortBy, order)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repo.PageGrades(search, classIDInt)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (s *StudentRelationalService) GetStudensAvg(page, limit int, search, classID, sortBy, order string) ([]model.StudentsAvg, int, error) {
	if page < 1 {
		page = 1
	}
	if limit <= 0 || limit > 100 {
		limit = 10
	}

	var classIDInt int
	var err error
	if classID != "" {
		classIDInt, err = strconv.Atoi(classID)
		if err != nil {
			return nil, 0, fmt.Errorf("class_id harus angka")
		}
		return nil, classIDInt, err
	}

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

	data, err := s.repo.StudentsAvg(limit, offset, search, classIDInt, sortBy, order)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repo.PageAvg(search, classIDInt)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

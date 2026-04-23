package service

import (
	"fmt"
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/repository"
)

// service pegang business logic
type LogService struct {
	repo repository.ActivityLogRepository
}

// constructor
func NewLogService(repo repository.ActivityLogRepository) *LogService {
	return &LogService{repo: repo}
}

// // Create student baru
// func (s *LogService) CreateClass(class *model.Classes) error {
// 	// bisa tambah validasi manual di sini kalau mau
// 	return s.repo.Create(class)
// }

func (s *LogService) GetAllLog(page, limit int, search, sortBy, order string) ([]model.ActivityLog, int, error) {
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

	data, err := s.repo.FindAllLogs(limit, offset, search, sortBy, order)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repo.AllLogs(search)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (s *LogService) GetLogByID(id int) (*model.ActivityLog, error) {
	class, err := s.repo.FindLogByID(id)
	if err != nil {
		return nil, err // error DB
	}

	if class == nil {
		return nil, fmt.Errorf("log tidak ditemukan") // handle not found
	}

	return class, nil
}

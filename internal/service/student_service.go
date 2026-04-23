package service

import (
	"errors"
	"fmt"
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/repository"
	"strconv"
)

// service pegang business logic
type StudentService struct {
	repo      repository.StudentRepository
	classRepo repository.ClassRepository
}

// constructor
func NewStudentService(repo repository.StudentRepository, classRepo repository.ClassRepository) *StudentService {
	return &StudentService{repo: repo, classRepo: classRepo}
}

// Create student baru
func (s *StudentService) CreateStudent(student *model.Students) error {
	// bisa tambah validasi manual di sini kalau mau
	_, err := s.classRepo.FindClassByID(student.ClassID)
	if err != nil {
		return errors.New("kelas tidak ditemukan") // kalau tidak ada → langsung stop
	}
	email := student.Email
	exists, err := s.repo.IsStudentEmailExists(email)
	if err != nil {
		return err
	}
	if exists {
		return errors.New("email sudah digunakan")
	}
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

func (s *StudentService) GetAllStudents(page, limit int, search, classID, sortBy, order string) ([]model.Students, int, error) {
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

	data, err := s.repo.FindAll(limit, offset, search, classIDInt, sortBy, order)
	if err != nil {
		return nil, 0, err
	}
	total, err := s.repo.Count(search, classIDInt)
	if err != nil {
		return nil, 0, err
	}
	return data, total, nil
}

func (s *StudentService) UpdateStudent(id int, student model.Students) error {
	// cek dulu apakah data ada menggunakan fungsi dari repo yang mencari berdasarkan id
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

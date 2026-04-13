package service

import (
	"fmt"
	"golang-sekolah/internal/dto"
	"golang-sekolah/internal/model"
	"golang-sekolah/internal/repository"
)

type SubjectService struct {
	repo repository.SubjectRepository
}

func NewSubjectService(repo repository.SubjectRepository) *SubjectService {
	return &SubjectService{repo: repo}
}

func (s *SubjectService) CreateSubject(req dto.SubjectReq) error {
	// Pengecekan manual dihilangkan karena diatasi oleh DTO + tag binding
	subjectReq := dto.SubjectReq{
		Name:       req.Name,
		Teacher_id: req.Teacher_id,
	}
	return s.repo.Create(subjectReq)
}

func (s *SubjectService) GetAllSubjects() ([]model.Subjects, error) {
	return s.repo.FindAll()
}

func (s *SubjectService) GetSubjectByID(id int) (*model.Subjects, error) {
	subject, err := s.repo.FindByID(id)
	if err != nil {
		return nil,  err
	}
	if subject == nil {
		return nil, fmt.Errorf("mata pelajaran tidak ditemukan")
	}
	return subject, nil
}

func (s *SubjectService) UpdateSubject(id int, req dto.SubjectReq) error {
	// Verifikasi apakah mata pelajaran yang mau diedit ada
	_, err := s.GetSubjectByID(id)
	if err != nil {
		return err
	}
	
	return s.repo.Update(id, req)
}

func (s *SubjectService) DeleteSubject(id int) error {
	// Verifikasi apakah mata pelajaran yang mau dihapus ada
	_, err := s.GetSubjectByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

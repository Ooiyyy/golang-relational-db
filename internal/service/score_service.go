package service

import (
	"fmt"
	"golang-relational-db/internal/dto"
	"golang-relational-db/internal/model"
	"golang-relational-db/internal/repository"
)

type ScoreService struct {
	repo repository.ScoreRepository
}

func NewScoreService(repo repository.ScoreRepository) *ScoreService {
	return &ScoreService{repo: repo}
}

func (s *ScoreService) Create(req dto.ScoreReq) error {
	scoreReq := dto.ScoreReq{
		SubjectID: req.SubjectID,
		StudentID: req.StudentID,
		Score:     req.Score,
	}
	return s.repo.Create(scoreReq)
}

func (s *ScoreService) GetAllScores() ([]model.Scores, error) {
	return s.repo.FindAll()
}

func (s *ScoreService) GetScoreByID(id int) (*model.Scores, error) {
	score, err := s.repo.FindByID(id)
	if err != nil {
		return nil, err
	}
	if score == nil {
		return nil, fmt.Errorf("nilai tidak ditemukan")
	}
	return score, nil
}

func (s *ScoreService) UpdateScore(id int, req dto.ScoreReq) error {
	// Verifikasi apakah nilai yang mau diedit ada
	_, err := s.GetScoreByID(id)
	if err != nil {
		return err
	}

	return s.repo.Update(id, req)
}

func (s *ScoreService) DeleteScore(id int) error {
	// Verifikasi apakah nilai yang mau dihapus ada
	_, err := s.GetScoreByID(id)
	if err != nil {
		return err
	}
	return s.repo.Delete(id)
}

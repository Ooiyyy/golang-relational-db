package repository

import (
	"golang-sekolah/internal/dto"
	"golang-sekolah/internal/model"
)

type ScoreRepository interface {
	Create(req dto.ScoreReq) error
	FindAll() ([]model.Scores, error)
	FindByID(id int) (*model.Scores, error)
	Update(id int, req dto.ScoreReq) error
	Delete(id int) error
}

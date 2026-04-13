package repository

import (
	"golang-sekolah/internal/dto"
	"golang-sekolah/internal/model"
)

type SubjectRepository interface {
	Create(req dto.SubjectReq) error
	FindAll() ([]model.Subjects, error)
	FindByID(id int) (*model.Subjects, error)
	Update(id int, req dto.SubjectReq) error
	Delete(id int) error
}

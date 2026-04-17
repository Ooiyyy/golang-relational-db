package repository

import "golang-relational-db/internal/model"

// Repository hanya berurusan dengan data & DB
type StudentRepository interface {
	Create(student *model.Students) error
	Count(search string, classID int) (int, error)
	FindAll(limit, offset int, search string, classID int, sortBy, order string) ([]model.Students, error)
	FindByID(id int) (*model.Students, error)
	Update(id int, student model.Students) error
	Delete(id int) error
}

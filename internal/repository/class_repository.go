package repository

import "golang-relational-db/internal/model"

type ClassRepository interface {
	Create(class *model.Classes) error
	FindAllClasses(limit, offset int, search string, sortBy, order string) ([]model.DetailClass, error)
	AllClasses(search string) (int, error)
	// FindByID(id int) (*model.Students, error)
	// Update(id int, student model.Students) error
	// Delete(id int) error
}

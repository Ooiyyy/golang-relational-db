package repository

import "golang-relational-db/internal/model"

type ClassRepository interface {
	Create(class *model.Classes) error
	FindAllClasses(limit, offset int, search string, sortBy, order string) ([]model.DetailClass, error)
	AllClasses(search string) (int, error)
	FindClassByID(id int) (*model.DetailClass, error)
	UpdateClass(Classid int, student model.Classes) error
	DeleteClass(id int) error
}

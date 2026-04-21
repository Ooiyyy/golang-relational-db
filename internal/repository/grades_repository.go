package repository

import "golang-relational-db/internal/model"

type GradesRepository interface {
	Create(class *model.Grades) error
}

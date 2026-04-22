package repository

import "golang-relational-db/internal/model"

type GradesRepository interface {
	Create(grade *model.Grades) error
	// FindAllGrades(limit, offset int, search string, sortBy, order string) ([]model.GradesDetail, error)
	// AllGrades(search string) (int, error)
	FindGradeByID(id int) (*model.GradesDetail, error)
	UpdateGrade(id int, grade model.Grades) error
	DeleteGrade(id int) error
}

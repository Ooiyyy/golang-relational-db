package repository

import "golang-relational-db/internal/model"

type SubjectsCepository interface {
	Create(class *model.Subjects) error
	FindAllSubjects(limit, offset int, search string, sortBy, order string) ([]model.Subjects, error)
	AllSubjects(search string) (int, error)
	// FindByID(id int) (*model.Students, error)
	// Update(id int, student model.Students) error
	// Delete(id int) error
}

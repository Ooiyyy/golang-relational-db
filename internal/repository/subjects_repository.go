package repository

import "golang-relational-db/internal/model"

type SubjectsCepository interface {
	Create(class *model.Subjects) error
	FindAllSubjects(limit, offset int, search string, sortBy, order string) ([]model.Subjects, error)
	AllSubjects(search string) (int, error)
	FindSubjectByID(id int) (*model.Subjects, error)
	UpdateSubject(id int, student model.Subjects) error
	DeleteSubject(id int) error
}

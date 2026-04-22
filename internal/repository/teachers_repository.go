package repository

import "golang-relational-db/internal/model"

// Repository hanya berurusan dengan data & DB
type TeacherRepository interface {
	IsTeacherEmailExists(email string) (bool, error)
	Create(teacher *model.Teachers) error
	AllTeachers(search string) (int, error)
	FindAllTeachers(limit, offset int, search string, sortBy, order string) ([]model.Teachers, error)
	FindTeacherByID(id int) (*model.Teachers, error)
	UpdateTeacher(id int, teacher model.Teachers) error
	DeleteTeacher(id int) error
}

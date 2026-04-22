package repository

import "golang-relational-db/internal/model"

type StudentRelational interface {
	Pages(search string, classID int) (int, error)
	PageGrades(search string, classID int) (int, error)
	PageAvg(search string, classID int) (int, error)
	StudentsDetail(limit, offset int, search string, classID int, sortBy, order string) ([]model.StudentDetail, error)
	StudentsGrade(limit, offset int, search string, classID int, sortBy, order string) ([]model.GradesDetail, error)
	StudentsAvg(limit, offset int, search string, classID int, sortBy, order string) ([]model.StudentsAvg, error)
}

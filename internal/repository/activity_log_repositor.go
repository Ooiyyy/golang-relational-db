package repository

import "golang-relational-db/internal/model"

type ActivityLogRepository interface {
	Create(log model.ActivityLog) error
	FindAllLogs(limit, offset int, search string, sortBy, order string) ([]model.ActivityLog, error)
	AllLogs(search string) (int, error)
	FindLogByID(id int) (*model.ActivityLog, error)
	// UpdateClass(Classid int, student model.Classes) error
	// DeleteClass(id int) error
}

package mysql

import (
	"database/sql"
	"golang-relational-db/internal/model"
)

type ActivityLogRepo struct {
	DB *sql.DB
}

func NewActivityLogRepo(db *sql.DB) *ActivityLogRepo {
	return &ActivityLogRepo{DB: db}
}

func (r *ActivityLogRepo) Create(log model.ActivityLog) error {
	query := "INSERT INTO activity_logs (ip, aktivitas) VALUES (?, ?)"
	_, err := r.DB.Exec(query, log.IP, log.Aktivitas)
	return err
}

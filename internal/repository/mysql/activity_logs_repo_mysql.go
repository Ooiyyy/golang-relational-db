package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-relational-db/internal/model"
)

type ActivityLogRepoImpl struct {
	DB *sql.DB
}

func NewActivityLogRepo(db *sql.DB) *ActivityLogRepoImpl {
	return &ActivityLogRepoImpl{DB: db}
}

func (r *ActivityLogRepoImpl) Create(log model.ActivityLog) error {
	query := "INSERT INTO activity_logs (ip, aktivitas) VALUES (?, ?)"
	_, err := r.DB.Exec(query, log.IP, log.Aktivitas)
	return err
}

func (r *ActivityLogRepoImpl) FindAllLogs(limit, offset int, search string, sortBy, order string) ([]model.ActivityLog, error) {
	query := "SELECT * FROM activity_logs WHERE 1=1"
	args := []interface{}{}

	if search != "" {
		query += " AND (LOWER(aktivitas) LIKE LOWER(?))"
		args = append(args, "%"+search+"%")
	}

	query += " ORDER BY " + sortBy + " " + order

	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	logs := []model.ActivityLog{}

	for rows.Next() {
		var log model.ActivityLog
		err := rows.Scan(&log.ID, &log.IP, &log.Aktivitas, &log.CreatedAt)
		if err != nil {
			return nil, err
		}
		logs = append(logs, log)
	}
	fmt.Println("QUERY:", query)
	fmt.Println("ARGS:", args)
	return logs, nil
}

func (r *ActivityLogRepoImpl) AllLogs(search string) (int, error) {
	query := "SELECT COUNT(*) FROM activity_logs WHERE 1=1"
	args := []interface{}{}

	if search != "" {
		query += " AND aktivitas LIKE ?"
		args = append(args, "%"+search+"%")
	}

	var total int
	err := r.DB.QueryRow(query, args...).Scan(&total)
	return total, err
}

func (r *ActivityLogRepoImpl) FindLogByID(id int) (*model.ActivityLog, error) {
	var log model.ActivityLog

	// ambil 1 data berdasarkan id
	err := r.DB.QueryRow(
		`SELECT id, ip, aktivitas, created_at FROM activity_logs
		WHERE id= ?`, id).Scan(&log.ID, &log.IP, &log.Aktivitas, &log.CreatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("log tidak ditemukan") // ⚠️ ini penting → tanda data tidak ada
		}
		return nil, err // error lain (DB error)
	}

	return &log, nil
}

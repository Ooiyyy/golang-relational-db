package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-relational-db/internal/model"
	"strings"
)

// struct repo menyimpan koneksi DB
type TeacherRepoImpl struct {
	DB *sql.DB
}

// constructor
func NewTeacherRepo(db *sql.DB) *TeacherRepoImpl {
	return &TeacherRepoImpl{DB: db}
}

func (r *TeacherRepoImpl) IsTeacherEmailExists(email string) (bool, error) {
	var count int
	query := `SELECT COUNT(*) from teachers where email=?`
	err := r.DB.QueryRow(query, email).Scan(&count)
	if err != nil {
		return false, nil
	}
	return count > 0, nil
}

// Insert data teacher ke database
func (r *TeacherRepoImpl) Create(teacher *model.Teachers) error {
	result, err := r.DB.Exec(
		"INSERT INTO teachers (name, email) VALUES (?, ?)",
		teacher.Name,
		teacher.Email,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	teacher.ID = int(id)
	return err // hanya return error
}

func (r *TeacherRepoImpl) FindAllTeachers(limit, offset int, search string, sortBy, order string) ([]model.Teachers, error) {
	query := "SELECT id, name, email FROM teachers WHERE 1=1"
	args := []interface{}{}

	if search != "" {
		query += " AND (LOWER(name) LIKE LOWER(?) OR LOWER(email) LIKE LOWER(?))"
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	// ######### Mencegah sql injection ###################
	allowedSort := map[string]string{
		"id":    "id",
		"name":  "name",
		"email": "email",
	}

	sortColumn, ok := allowedSort[sortBy]
	if !ok {
		sortColumn = "id"
	}

	order = strings.ToUpper(order)
	if order != "ASC" && order != "DESC" {
		order = "ASC"
	}

	query += " ORDER BY " + sortColumn + " " + order

	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)
	fmt.Println("QUERY:", query)
	fmt.Println("ARGS:", args)
	// #############################################################

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	teachers := []model.Teachers{}

	for rows.Next() {
		var teacher model.Teachers
		err := rows.Scan(&teacher.ID, &teacher.Name, &teacher.Email)
		if err != nil {
			return nil, err
		}
		teachers = append(teachers, teacher)
	}
	return teachers, nil
}

func (r *TeacherRepoImpl) AllTeachers(search string) (int, error) {
	query := "SELECT COUNT(*) FROM teachers WHERE 1=1"
	args := []interface{}{}

	if search != "" {
		query += " AND name LIKE ? OR email LIKE ?"
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	var total int
	err := r.DB.QueryRow(query, args...).Scan(&total)
	return total, err
}

func (r *TeacherRepoImpl) FindTeacherByID(id int) (*model.Teachers, error) {
	var teacher model.Teachers

	// ambil 1 data berdasarkan id
	err := r.DB.QueryRow(
		"SELECT id, name, email FROM teachers WHERE id = ?",
		id,
	).Scan(&teacher.ID, &teacher.Name, &teacher.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("guru tidak ditemukan") // ⚠️ ini penting → tanda data tidak ada
		}
		return nil, err // error lain (DB error)
	}

	return &teacher, nil
}

func (r *TeacherRepoImpl) UpdateTeacher(id int, teacher model.Teachers) error {
	_, err := r.DB.Exec(
		"UPDATE teachers SET name=?, email=? WHERE id=?",
		teacher.Name,
		teacher.Email,
		id,
	)
	return err
}

func (r *TeacherRepoImpl) DeleteTeacher(id int) error {
	_, err := r.DB.Exec("DELETE FROM teachers WHERE id=?", id)
	return err
}

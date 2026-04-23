package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"golang-relational-db/internal/model"
)

// struct repo menyimpan koneksi DB
type ClassRepoImpl struct {
	DB *sql.DB
}

// constructor
func NewClassRepo(db *sql.DB) *ClassRepoImpl {
	return &ClassRepoImpl{DB: db}
}

// Insert data student ke database
func (r *ClassRepoImpl) Create(class *model.Classes) error {
	result, err := r.DB.Exec(
		"INSERT INTO classes (name, teacher_id) VALUES (?, ?)",
		class.Name,
		class.TeacherID,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	class.ID = int(id)
	return err // hanya return error
}

func (r *ClassRepoImpl) FindAllClasses(limit, offset int, search string, sortBy, order string) ([]model.DetailClass, error) {
	query := "SELECT c.id, c.name AS nama_kelas, c.teacher_id, t.name AS nama_guru FROM classes as c JOIN teachers AS t ON t.id = c.teacher_id WHERE 1=1"
	args := []interface{}{}

	if search != "" {
		query += " AND (LOWER(c.name) LIKE LOWER(?) OR LOWER(t.name) LIKE LOWER(?))"
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	query += " ORDER BY " + sortBy + " " + order

	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	classes := []model.DetailClass{}

	for rows.Next() {
		var class model.DetailClass
		err := rows.Scan(&class.ID, &class.Name, &class.TeacherID, &class.TeacherName)
		if err != nil {
			return nil, err
		}
		classes = append(classes, class)
	}
	fmt.Println("QUERY:", query)
	fmt.Println("ARGS:", args)
	return classes, nil
}

func (r *ClassRepoImpl) AllClasses(search string) (int, error) {
	query := "SELECT COUNT(*) FROM classes WHERE 1=1"
	args := []interface{}{}

	if search != "" {
		query += " AND c.name LIKE ? OR t.name LIKE ?"
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	var total int
	err := r.DB.QueryRow(query, args...).Scan(&total)
	return total, err
}

func (r *ClassRepoImpl) FindClassByID(id int) (*model.DetailClass, error) {
	var class model.DetailClass

	// ambil 1 data berdasarkan id
	err := r.DB.QueryRow(
		`SELECT c.id, c.name AS nama_kelas, c.teacher_id, t.name AS nama_guru
		FROM classes as c JOIN teachers AS t ON t.id = c.teacher_id
		WHERE c.id=?`, id).Scan(&class.ID, &class.Name, &class.TeacherID, &class.TeacherName)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("kelas tidak ditemukan") // ⚠️ ini penting → tanda data tidak ada
		}
		return nil, err // error lain (DB error)
	}
	return &class, nil
}

func (r *ClassRepoImpl) UpdateClass(id int, class model.Classes) error {
	_, err := r.DB.Exec(
		"UPDATE classes SET name=?, teacher_id=? WHERE id=?",
		class.Name,
		class.TeacherID,
		id,
	)
	return err
}

func (r *ClassRepoImpl) DeleteClass(id int) error {
	_, err := r.DB.Exec("DELETE FROM classes WHERE id=?", id)
	return err
}

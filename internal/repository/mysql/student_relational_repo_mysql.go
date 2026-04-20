package mysql

import (
	"database/sql"
	"fmt"
	"golang-relational-db/internal/model"
)

type StudentRelationalImpl struct {
	DB *sql.DB
}

func NewStudentRelationalRepo(db *sql.DB) *StudentRelationalImpl {
	return &StudentRelationalImpl{
		DB: db,
	}
}

func (r *StudentRelationalImpl) StudentsDetail(limit, offset int, search string, classID int, sortBy, order string) ([]model.StudentDetail, error) {
	query := `SELECT s.id, s.name, s.email,
				c.name AS kelas,
				t.name AS wali_kelas
			FROM students AS s
			JOIN classes AS c ON c.id = s.class_id
			JOIN teachers AS t ON t.id = c.teacher_id WHERE 1=1`
	args := []interface{}{}

	if search != "" {
		query += " AND (LOWER(s.name) LIKE LOWER(?) OR LOWER(s.email) LIKE LOWER(?))"
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	if classID != 0 {
		query += " AND s.class_id = ?"
		args = append(args, classID)
	}

	query += " ORDER BY " + sortBy + " " + order

	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var detailStudents []model.StudentDetail

	for rows.Next() {
		var student model.StudentDetail
		err := rows.Scan(&student.ID, &student.Name, &student.Email, &student.Kelas, &student.WaliKelas)
		if err != nil {
			return nil, err
		}
		detailStudents = append(detailStudents, student)
	}
	fmt.Println("QUERY:", query)
	fmt.Println("ARGS:", args)
	return detailStudents, nil
}

func (r *StudentRelationalImpl) Pages(search string, classID int) (int, error) {
	query := `SELECT COUNT(*) FROM students s 
				JOIN classes AS c ON c.id = s.class_id
				JOIN teachers AS t ON t.id = c.teacher_id
				WHERE 1=1`
	args := []interface{}{}

	if search != "" {
		query += " AND (s.name LIKE ? OR s.email LIKE ?)"
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	if classID != 0 {
		query += " AND class_id = ?"
		args = append(args, classID)
	}

	var total int
	err := r.DB.QueryRow(query, args...).Scan(&total)
	return total, err
}

func (r *StudentRelationalImpl) StudentsGrade(limit, offset int, search string, classID int, sortBy, order string) ([]model.StudentsGrade, error) {
	query := `SELECT s.id, s.name AS nama_siswa,
				g.score AS nilai,
				sj.name AS mapel
				FROM grades AS g
				JOIN students AS s ON s.id = g.student_id
				JOIN subjects AS sj ON sj.id = g.subject_id WHERE 1=1`
	args := []interface{}{}

	if search != "" {
		query += " AND (LOWER(s.name) LIKE LOWER(?) OR LOWER(s.email) LIKE LOWER(?))"
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	if classID != 0 {
		query += " AND s.class_id = ?"
		args = append(args, classID)
	}

	query += " ORDER BY " + sortBy + " " + order

	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studentsGrade []model.StudentsGrade

	for rows.Next() {
		var student model.StudentsGrade
		err := rows.Scan(&student.ID, &student.Name, &student.Nilai, &student.Mapel)
		if err != nil {
			return nil, err
		}
		studentsGrade = append(studentsGrade, student)
	}
	fmt.Println("QUERY:", query)
	fmt.Println("ARGS:", args)
	return studentsGrade, nil
}

func (r *StudentRelationalImpl) PageGrades(search string, classID int) (int, error) {
	query := `SELECT COUNT(*) FROM grades g 
				JOIN students AS s ON s.id = g.student_id
				JOIN subjects AS sj ON sj.id = g.subject_id
				WHERE 1=1`
	args := []interface{}{}

	if search != "" {
		query += " AND (s.name LIKE ? OR s.email LIKE ?)"
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	if classID != 0 {
		query += " AND s.class_id = ?"
		args = append(args, classID)
	}

	var total int
	err := r.DB.QueryRow(query, args...).Scan(&total)
	return total, err
}

func (r *StudentRelationalImpl) StudentsAvg(limit, offset int, search string, classID int, sortBy, order string) ([]model.StudentsAvg, error) {
	query := `SELECT s.id, s.name, AVG(g.score) AS rata_rata
				FROM grades AS g
				JOIN students AS s on s.id = g.student_id
				WHERE 1=1`
	args := []interface{}{}

	if search != "" {
		query += " AND (LOWER(s.name) LIKE LOWER(?) OR LOWER(s.email) LIKE LOWER(?))"
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	if classID != 0 {
		query += " AND s.class_id = ?"
		args = append(args, classID)
	}

	query += " GROUP BY s.name"

	query += " ORDER BY " + sortBy + " " + order

	query += " LIMIT ? OFFSET ?"
	args = append(args, limit, offset)

	rows, err := r.DB.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var studentsAvg []model.StudentsAvg

	for rows.Next() {
		var student model.StudentsAvg
		err := rows.Scan(&student.ID, &student.Name, &student.Avg)
		if err != nil {
			return nil, err
		}
		studentsAvg = append(studentsAvg, student)
	}
	fmt.Println("QUERY:", query)
	fmt.Println("ARGS:", args)
	return studentsAvg, nil
}

func (r *StudentRelationalImpl) PageAvg(search string, classID int) (int, error) {
	query := `SELECT COUNT(DISTINCT s.id) FROM grades g 
				JOIN students AS s on s.id = g.student_id
				WHERE 1=1`
	args := []interface{}{}

	if search != "" {
		query += " AND (s.name LIKE ? OR s.email LIKE ?)"
		args = append(args, "%"+search+"%", "%"+search+"%")
	}

	if classID != 0 {
		query += " AND s.class_id = ?"
		args = append(args, classID)
	}

	var total int
	err := r.DB.QueryRow(query, args...).Scan(&total)
	return total, err
}

package model

type Subjects struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	TeacherID   int    `json:"teacher_id"`
	TeacherName string `json:"nama_guru"`
}

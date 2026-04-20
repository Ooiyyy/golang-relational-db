package model

type Classes struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	TeacherID int    `json:"teacher_id"`
}

type DetailClass struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	TeacherName string `json:"nama_guru"`
}

package dto

type CreateSubjectsReq struct {
	Name      string `json:"name" binding:"required"`
	TeacherID int    `json:"teacher_id" binding:"required"`
}

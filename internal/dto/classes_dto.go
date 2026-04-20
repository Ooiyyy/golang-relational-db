package dto

type CreateClassReq struct {
	Name      string `json:"name" binding:"required"`
	TeacherID int    `json:"teacher_id" binding:"required"`
}

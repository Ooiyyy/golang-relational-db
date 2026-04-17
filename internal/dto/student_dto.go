package dto

type CreateStudentReq struct {
	Name    string `json:"name" binding:"required"`
	Email   string `json:"email" binding:"required,email"`
	ClassID int    `json:"class_id" binding:"required"`
}

type UpdateStudentReq struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	ClassID int    `json:"class_id"`
}

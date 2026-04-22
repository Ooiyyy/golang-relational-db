package dto

type CreateTeacherReq struct {
	Name  string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
}

type UpdateTeacherReq struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

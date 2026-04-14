package dto

type SubjectReq struct {
	Name      string `json:"name" binding:"required"`
	TeacherID int    `json:"-"` // Tidak akan dibaca dari body JSON, akan diisi manual dari Token
}

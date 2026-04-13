package dto

type SubjectReq struct {
	Name       string `json:"name" binding:"required"`
	Teacher_id int    `json:"-"` // Tidak akan dibaca dari body JSON, akan diisi manual dari Token
}

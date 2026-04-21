package dto

type CreateGradesReq struct {
	StudentID int `json:"student_id" binding:"required"`
	SubjectID int `json:"subject_id" biding:"required"`
	Score     int `json:"score" binding:"required"`
}

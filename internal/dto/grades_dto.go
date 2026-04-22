package dto

type CreateGradesReq struct {
	StudentID int `json:"student_id" binding:"required"`
	SubjectID int `json:"subject_id" biding:"required"`
	Score     int `json:"score" binding:"required"`
}

type UpdateGradesReq struct {
	StudentID int `json:"student_id"`
	SubjectID int `json:"subject_id"`
	Score     int `json:"score"`
}

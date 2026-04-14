package dto

type ScoreReq struct {
	SubjectID int `json:"subject_id" binding:"required"`
	StudentID int `json:"student_id" binding:"required"`
	Score     int `json:"score" binding:"required"`
}

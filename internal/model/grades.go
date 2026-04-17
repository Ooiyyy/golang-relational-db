package model

type Grades struct {
	ID        int `json:"id"`
	StudentID int `json:"student_id"`
	SubjectID int `json:"subject_id"`
	Score     int `json:"score"`
}

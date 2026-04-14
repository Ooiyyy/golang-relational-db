package model

import "time"

type Scores struct {
	ID        int       `json:"id"`
	StudentID int       `json:"student_id"`
	SubjectID int       `json:"subject_id"`
	Score     int       `json:"score"`
	CreatedAt time.Time `json:"created_at"`
}

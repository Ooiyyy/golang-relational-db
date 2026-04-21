package model

type Grades struct {
	ID        int `json:"id"`
	StudentID int `json:"student_id"`
	SubjectID int `json:"subject_id"`
	Score     int `json:"score"`
}

type GradesDetail struct {
	ID          int `json:"id"`
	StudentName int `json:"student_name"`
	SubjectName int `json:"subject_name"`
	Score       int `json:"score"`
}

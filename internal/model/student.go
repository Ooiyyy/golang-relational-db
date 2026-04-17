package model

type Students struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	ClassID int    `json:"class_id"`
}

package mysql

import (
	"database/sql"
	"golang-relational-db/internal/dto"
	"golang-relational-db/internal/model"
)

type ScoreRepoImpl struct {
	DB *sql.DB
}

func NewScoreRepo(db *sql.DB) *ScoreRepoImpl {
	return &ScoreRepoImpl{DB: db}
}

func (s *ScoreRepoImpl) Create(req dto.ScoreReq) error {
	_, err := s.DB.Exec("INSERT INTO scores (student_id, subject_id, score, created_at) VALUES (?, ?, ?, NOW())", req.StudentID, req.SubjectID, req.Score)
	return err
}

func (s *ScoreRepoImpl) FindAll() ([]model.Scores, error) {
	rows, err := s.DB.Query("SELECT id, student_id, subject_id, score FROM scores")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var scores []model.Scores
	for rows.Next() {
		var score model.Scores
		if err := rows.Scan(&score.ID, &score.SubjectID, &score.StudentID, &score.Score); err != nil {
			return nil, err
		}
		scores = append(scores, score)
	}
	return scores, nil
}

func (s *ScoreRepoImpl) FindByID(id int) (*model.Scores, error) {
	var score model.Scores
	err := s.DB.QueryRow("SELECT id, subject_id, student_id, score FROM scores WHERE id = ?", id).Scan(&score.ID, &score.SubjectID, &score.StudentID, &score.Score)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Data tidak ditemukan
		}
		return nil, err
	}
	return &score, nil
}

func (s *ScoreRepoImpl) Update(id int, req dto.ScoreReq) error {
	_, err := s.DB.Exec("UPDATE scores SET score = ? WHERE id = ?", req.Score, id)
	return err
}

func (s *ScoreRepoImpl) Delete(id int) error {
	_, err := s.DB.Exec("DELETE FROM scores WHERE id = ?", id)
	return err
}

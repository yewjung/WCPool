package scoreRepo

import (
	"WCPool/models"
	"database/sql"
	"fmt"
	"strings"
)

type ScoreRepo struct{}

func (s *ScoreRepo) GetScore(db *sql.DB, scoreID int) (models.Score, error) {
	row := db.QueryRow("select * from Score where ScoreID = $1", scoreID)
	score := models.Score{}
	err := row.Scan(&score.UserID, &score.Points)
	if err != nil {
		return score, err
	}
	return score, nil
}

func (s *ScoreRepo) GetScores(db *sql.DB, userIDs []string) ([]models.Score, error) {
	sql := fmt.Sprintf("select * from Score where UserID in (%s)", strings.Join(userIDs, ","))
	rows, err := db.Query(sql)
	if err != nil {
		return []models.Score{}, err
	}
	defer rows.Close()
	scores := []models.Score{}
	for rows.Next() {
		score := models.Score{}
		err = rows.Scan(&score.UserID, &score.Points)
		scores = append(scores, score)
	}
	if err != nil {
		return []models.Score{}, err
	}
	return scores, nil
}

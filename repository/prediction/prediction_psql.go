package predictionRepo

import (
	"WCPool/models"
	"database/sql"
)

type PredictionRepo struct{}

func (p *PredictionRepo) GetPredictionsByUserIDNMatchID(db *sql.DB, userID int, matchID int) (models.Prediction, error) {
	row := db.QueryRow("select * from Prediction where UserID = $1 and MatchID = $2", userID, matchID)
	prediction := models.Prediction{}
	err := row.Scan(&prediction.UserID, &prediction.MatchID, &prediction.GoalsA, &prediction.GoalsB)
	if err != nil {
		return prediction, err
	}
	return prediction, nil
}

func (p *PredictionRepo) GetPredictionsByUserID(db *sql.DB, userID int) ([]models.Prediction, error) {
	rows, err := db.Query("select * from Prediction where UserID = $1", userID)
	if err != nil {
		return []models.Prediction{}, err
	}
	defer rows.Close()
	predictions := []models.Prediction{}
	for rows.Next() {
		prediction := models.Prediction{}
		err = rows.Scan(&prediction.UserID, &prediction.MatchID, &prediction.GoalsA, &prediction.GoalsB)
		predictions = append(predictions, prediction)
	}
	if err != nil {
		return []models.Prediction{}, err
	}
	return predictions, nil
}

func (p *PredictionRepo) AddPrediction(db *sql.DB, prediction models.Prediction) (models.Prediction, error) {
	db.QueryRow("insert into Prediction (UserID, MatchID, GoalsA, GoalsB) values ($1, $2, $3, $4) returning PredictionID", prediction.UserID, prediction.MatchID, prediction.GoalsA, prediction.GoalsB)
	return prediction, nil
}

// add a bunch of predictions
func (p *PredictionRepo) AddPredictions(db *sql.DB, predictions []models.Prediction) ([]models.Prediction, error) {
	// bulk imports
	tx, err := db.Begin()
	if err != nil {
		return []models.Prediction{}, err
	}
	stmt, err := tx.Prepare("insert into Prediction (UserID, MatchID, GoalsA, GoalsB) values ($1, $2, $3, $4)")
	if err != nil {
		return []models.Prediction{}, err
	}
	for _, prediction := range predictions {
		_, err = stmt.Exec(prediction.UserID, prediction.MatchID, prediction.GoalsA, prediction.GoalsB)
		if err != nil {
			return []models.Prediction{}, err
		}
	}
	err = stmt.Close()
	if err != nil {
		return []models.Prediction{}, err
	}
	err = tx.Commit()
	if err != nil {
		return []models.Prediction{}, err
	}
	return predictions, nil
}

func (p *PredictionRepo) UpdatePrediction(db *sql.DB, prediction models.Prediction) (models.Prediction, error) {
	db.QueryRow("update Prediction set GoalsA = $1, GoalsB = $2 where UserID = $3 and MatchID = $4 returning PredictionID", prediction.GoalsA, prediction.GoalsB, prediction.UserID, prediction.MatchID)
	return prediction, nil
}

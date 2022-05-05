package controller

import (
	"WCPool/models"
	predictionRepo "WCPool/repository/prediction"
	"WCPool/utils"
	"database/sql"
	"net/http"
)

type PredictionController struct{}

// add prediction
func (c *PredictionController) AddPrediction(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// add prediction
		predictionRepo := predictionRepo.PredictionRepo{}
		addedPrediction, err := predictionRepo.AddPrediction(db, utils.GetReqBody(r, models.Prediction{}))
		utils.HandleResponse(w, r, err, addedPrediction)
	}
}

// add a bunch of predictions
func (c *PredictionController) AddPredictions(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// add prediction
		predictionRepo := predictionRepo.PredictionRepo{}
		addedPredictions, err := predictionRepo.AddPredictions(db, utils.GetReqBody(r, []models.Prediction{}))
		utils.HandleResponse(w, r, err, addedPredictions)
	}
}

// update prediction
func (c *PredictionController) UpdatePrediction(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// update prediction
		predictionRepo := predictionRepo.PredictionRepo{}
		updatedPrediction, err := predictionRepo.UpdatePrediction(db, utils.GetReqBody(r, models.Prediction{}))
		utils.HandleResponse(w, r, err, updatedPrediction)
	}
}

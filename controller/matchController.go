package controller

import (
	"WCPool/models"
	matchRepo "WCPool/repository/match"
	"WCPool/utils"
	"database/sql"
	"net/http"
)

type MatchController struct{}

// get matches by matchday
func (c *MatchController) GetMatchesByMatchday(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		matchRepo := matchRepo.MatchRepo{}
		matches, err := matchRepo.GetMatchesByMatchday(db, utils.GetIntVar(r, "matchday"))
		utils.HandleResponse(w, r, err, matches)
	}
}

// add matches
func (c *MatchController) AddMatches(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		matchRepo := matchRepo.MatchRepo{}
		addedMatches, err := matchRepo.AddMatches(db, utils.GetReqBody(r, []models.Match{}))
		utils.HandleResponse(w, r, err, addedMatches)
	}
}

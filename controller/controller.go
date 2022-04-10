package controller

import (
	"WCPool/models"
	leaderboardRepo "WCPool/repository/leaderboard"
	matchRepo "WCPool/repository/match"
	partyRepo "WCPool/repository/party"
	predictionRepo "WCPool/repository/prediction"
	userRepo "WCPool/repository/user"
	"WCPool/utils"
	"database/sql"
	"net/http"
)

type Controller struct{}

func (c *Controller) GetLeaderboard(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		leaderboardRepo := leaderboardRepo.LeaderboardRepo{}
		leaderboardRepo.GetLeaderboard(db)
		// TODO: handle responsewriter
	}
}

/********************** party **********************/

// create party
func (c *Controller) CreateParty(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		partyRepo := partyRepo.PartyRepo{}
		party, err := partyRepo.AddParty(db, utils.GetReqBody(r, models.Party{}))
		utils.HandleResponse(w, r, err, party)
	}
}

// get party by id
func (c *Controller) GetParty(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		partyRepo := partyRepo.PartyRepo{}
		party, err := partyRepo.GetParty(db, utils.GetIntVar(r, "id"))
		if utils.SendServerErrorIfErr(w, r, err) {
			return
		}
		partyDTO := models.PartyDTO{}
		partyDTO.PartyID = party.PartyID
		partyDTO.PartyName = party.PartyName

		// map each userId to a user, collect them in a slice, set them to the partyDTO.Members
		userRepo := userRepo.UserRepo{}
		partyDTO.Members, err = userRepo.GetUsers(db, party.UserIDs)
		utils.HandleResponse(w, r, err, partyDTO)
	}
}

/************************ User ************************/

// Add Member to Party
func (c *Controller) AddMemberToParty(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRepo := userRepo.UserRepo{}
		user := utils.GetReqBody(r, models.User{})
		addedUser, err := userRepo.AddUser(db, user)
		if utils.SendServerErrorIfErr(w, r, err) {
			return
		}
		// add user to party
		partyRepo := partyRepo.PartyRepo{}
		err = partyRepo.AddUserToParty(db, utils.GetIntVar(r, "partyID"), addedUser.UserID)
		utils.HandleResponse(w, r, err, addedUser)
	}
}

func (c *Controller) DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRepo := userRepo.UserRepo{}
		err := userRepo.DeleteUser(db, utils.GetIntVar(r, "id"))
		utils.HandleResponse(w, r, err, nil)
	}
}

// update user
func (c *Controller) UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRepo := userRepo.UserRepo{}
		user, err := userRepo.UpdateUser(db, utils.GetReqBody(r, models.User{}))
		utils.HandleResponse(w, r, err, user)
	}
}

// delete user from party
func (c *Controller) RemoveUserFromParty(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		partyID := utils.GetIntVar(r, "partyID")
		userID := utils.GetIntVar(r, "userID")
		partyRepo := partyRepo.PartyRepo{}
		party, err := partyRepo.GetParty(db, partyID)
		if utils.SendServerErrorIfErr(w, r, err) {
			return
		}
		// remove userID from party.UserIDs
		party.UserIDs = utils.RemoveFromArray(party.UserIDs, userID)
		// update party
		party, err = partyRepo.AddParty(db, party)
		utils.HandleResponse(w, r, err, party)
	}
}

/******************* Matches ***********************/

// get matches by matchday
func (c *Controller) GetMatchesByMatchday(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		matchRepo := matchRepo.MatchRepo{}
		matches, err := matchRepo.GetMatchesByMatchday(db, utils.GetIntVar(r, "matchday"))
		utils.HandleResponse(w, r, err, matches)
	}
}

/************** Predictions ***********************/

// add prediction
func (c *Controller) AddPrediction(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// add prediction
		predictionRepo := predictionRepo.PredictionRepo{}
		addedPrediction, err := predictionRepo.AddPrediction(db, utils.GetReqBody(r, models.Prediction{}))
		utils.HandleResponse(w, r, err, addedPrediction)
	}
}

// add a bunch of predictions
func (c *Controller) AddPredictions(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// add prediction
		predictionRepo := predictionRepo.PredictionRepo{}
		addedPredictions, err := predictionRepo.AddPredictions(db, utils.GetReqBody(r, []models.Prediction{}))
		utils.HandleResponse(w, r, err, addedPredictions)
	}
}

// update prediction
func (c *Controller) UpdatePrediction(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// update prediction
		predictionRepo := predictionRepo.PredictionRepo{}
		updatedPrediction, err := predictionRepo.UpdatePrediction(db, utils.GetReqBody(r, models.Prediction{}))
		utils.HandleResponse(w, r, err, updatedPrediction)
	}
}

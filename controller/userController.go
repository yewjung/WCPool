package controller

import (
	"WCPool/models"
	partyRepo "WCPool/repository/party"
	userRepo "WCPool/repository/user"
	"WCPool/utils"
	"database/sql"
	"net/http"
)

type UserController struct{}

// Add Member to Party
func (c *UserController) AddMemberToParty(db *sql.DB) http.HandlerFunc {
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

func (c *UserController) DeleteUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRepo := userRepo.UserRepo{}
		err := userRepo.DeleteUser(db, utils.GetIntVar(r, "id"))
		utils.HandleResponse(w, r, err, nil)
	}
}

// update user
func (c *UserController) UpdateUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userRepo := userRepo.UserRepo{}
		user, err := userRepo.UpdateUser(db, utils.GetReqBody(r, models.User{}))
		utils.HandleResponse(w, r, err, user)
	}
}

// delete user from party
func (c *UserController) RemoveUserFromParty(db *sql.DB) http.HandlerFunc {
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

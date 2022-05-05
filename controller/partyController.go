package controller

import (
	"WCPool/models"
	partyRepo "WCPool/repository/party"
	userRepo "WCPool/repository/user"
	"WCPool/utils"
	"database/sql"
	"net/http"
)

type PartyController struct{}

// create party
func (c *PartyController) CreateParty(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		partyRepo := partyRepo.PartyRepo{}
		party, err := partyRepo.AddParty(db, utils.GetReqBody(r, models.Party{}))
		utils.HandleResponse(w, r, err, party)
	}
}

// get party by id
func (c *PartyController) GetParty(db *sql.DB) http.HandlerFunc {
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

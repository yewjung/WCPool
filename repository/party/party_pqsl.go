package partyRepo

import (
	"WCPool/models"
	"database/sql"

	"github.com/lib/pq"
)

type PartyRepo struct{}

func (p PartyRepo) GetParty(db *sql.DB, partyID int) (models.Party, error) {
	row := db.QueryRow("select * from Party where PartyID = $1", partyID)
	party := models.Party{}
	err := row.Scan(&party.PartyID, &party.PartyName, &party.UserIDs)
	if err != nil {
		return party, err
	}
	return party, nil
}

// add party
func (p *PartyRepo) AddParty(db *sql.DB, party models.Party) (models.Party, error) {
	row := db.QueryRow("insert into Party (UserIDs, PartyName) values ($1, $2) returning PartyID", party.UserIDs, pq.Array(party.UserIDs))
	err := row.Scan(&party.PartyID)
	if err != nil {
		return party, err
	}
	return party, nil
}

// add userId to party members list
func (p *PartyRepo) AddUserToParty(db *sql.DB, partyID int, userID int) error {
	sql := "update Party set UserIDs = array_append(UserIDs, $1) where PartyID = $2"
	_, err := db.Exec(sql, userID, partyID)
	if err != nil {
		return err
	}
	return nil
}

// update party
func (p *PartyRepo) UpdateParty(db *sql.DB, party models.Party) (models.Party, error) {
	sql := "update Party set PartyName = $1, UserIDs = $2 where PartyID = $3"
	_, err := db.Exec(sql, party.PartyName, pq.Array(party.UserIDs), party.PartyID)
	if err != nil {
		return party, err
	}
	return party, nil
}

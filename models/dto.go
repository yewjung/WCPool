package models

type PartyDTO struct {
	PartyID   int    `json:"party_id"`
	PartyName string `json:"party_name"`
	Members   []User `json:"members"`
}

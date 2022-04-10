package models

import "time"

type Leaderboard struct{}

type Party struct {
	PartyID   int    `json:"party_id"` // PK
	PartyName string `json:"party_name"`
	UserIDs   []int  `json:"members"`
}

type User struct {
	UserID   int    `json:"user_id"` // PK
	Name     string `json:"name"`
	PartyIDs []int  `json:"parties"`
}

type Score struct {
	UserID int `json:"user_id"` // PK, FK
	Points int `json:"points"`
}

type Match struct {
	MatchID  int       `json:"match_id"` // PK
	TeamA    string    `json:"team_a"`
	TeamB    string    `json:"team_b"`
	GoalsA   int       `json:"goals_a"`
	GoalsB   int       `json:"goals_b"`
	Time     time.Time `json:"time"`
	Matchday int       `json:"matchday"`
}

type Prediction struct {
	UserID  int `json:"user_id"`  // PK, FK to User.UserID
	MatchID int `json:"match_id"` // PK
	GoalsA  int `json:"goals_a"`
	GoalsB  int `json:"goals_b"`
}

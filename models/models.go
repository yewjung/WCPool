package models

import (
	"database/sql"
	"time"

	"github.com/golang-jwt/jwt"
)

type Leaderboard struct{}

type Party struct {
	PartyID   int    `json:"partyid"` // PK
	PartyName string `json:"partyname"`
	UserIDs   []int  `json:"members"`
}

type User struct {
	UserID   int    `json:"userid"` // PK
	Name     string `json:"name"`
	PartyIDs []int  `json:"parties"`
}

type Score struct {
	UserID int `json:"userid"` // PK, FK
	Points int `json:"points"`
}

type Match struct {
	MatchID  int       `json:"matchid"` // PK
	TeamA    string    `json:"teama"`
	TeamB    string    `json:"teamb"`
	GoalsA   int       `json:"goalsa"`
	GoalsB   int       `json:"goalsb"`
	Time     time.Time `json:"time"`
	Matchday int       `json:"matchday"`
}

type Prediction struct {
	UserID  int `json:"userid"`  // PK, FK to User.UserID
	MatchID int `json:"matchid"` // PK
	GoalsA  int `json:"goalsa"`
	GoalsB  int `json:"goalsb"`
}

type SecUser struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthUser struct {
	Email        string        `json:"email"`
	PasswordHash string        `json:"passwordhash"`
	UserId       sql.NullInt64 `json:"userid"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

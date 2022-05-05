package matchRepo

import (
	"WCPool/models"
	"database/sql"
)

type MatchRepo struct{}

func (m *MatchRepo) GetMatchesByMatchday(db *sql.DB, matchday int) ([]models.Match, error) {
	rows, err := db.Query("select * from Match where Matchday = $1", matchday)
	if err != nil {
		return []models.Match{}, err
	}
	defer rows.Close()
	matches := []models.Match{}
	for rows.Next() {
		match := models.Match{}
		err = rows.Scan(&match.MatchID, &match.TeamA, &match.TeamB, &match.GoalsA, &match.GoalsB, &match.Time, &match.Matchday)
		matches = append(matches, match)
	}
	if err != nil {
		return []models.Match{}, err
	}
	return matches, nil
}

func (m *MatchRepo) GetMatch(db *sql.DB, matchID int) (models.Match, error) {
	row := db.QueryRow("select * from Match where MatchID = $1", matchID)
	match := models.Match{}
	err := row.Scan(&match.MatchID, &match.TeamA, &match.TeamB, &match.Time, &match.Matchday)
	if err != nil {
		return match, err
	}
	return match, nil
}

// add all matches
func (m *MatchRepo) AddMatches(db *sql.DB, matches []models.Match) ([]models.Match, error) {
	// bulk add
	tx, err := db.Begin()
	if err != nil {
		return []models.Match{}, err
	}
	stmt, err := tx.Prepare("insert into Match (TeamA, TeamB, GoalsA, GoalsB, Time, Matchday) values ($1, $2, $3, $4, $5, $6)")
	if err != nil {
		return []models.Match{}, err
	}
	for _, match := range matches {
		_, err = stmt.Exec(match.TeamA, match.TeamB, match.GoalsA, match.GoalsB, match.Time, match.Matchday)
		if err != nil {
			return []models.Match{}, err
		}
	}
	err = stmt.Close()
	if err != nil {
		return []models.Match{}, err
	}
	err = tx.Commit()
	if err != nil {
		return []models.Match{}, err
	}
	return matches, nil
}

package leaderboardRepo

import (
	"WCPool/models"
	"database/sql"
)

type LeaderboardRepo struct{}

func (l *LeaderboardRepo) GetLeaderboard(db *sql.DB) (models.Leaderboard, error) {
	return models.Leaderboard{}, nil
}

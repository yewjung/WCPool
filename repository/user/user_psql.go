package userRepo

import (
	"WCPool/models"
	"database/sql"
	"fmt"
	"strings"

	"github.com/lib/pq"
)

type UserRepo struct{}

func (u *UserRepo) GetUser(db *sql.DB, userID int) (models.User, error) {
	row := db.QueryRow("select * from User where UserID = $1", userID)
	user := models.User{}
	err := row.Scan(&user.UserID, &user.Name)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (u UserRepo) GetUsers(db *sql.DB, userIDs []int) ([]models.User, error) {
	// convert slice of int into a string delimited by commas
	userIDsString := fmt.Sprint(userIDs)
	userIDsString = strings.Replace(userIDsString, "[", "", -1)
	userIDsString = strings.Replace(userIDsString, "]", "", -1)
	sql := fmt.Sprintf("select * from User where UserID in (%s)", userIDsString)
	rows, err := db.Query(sql)
	if err != nil {
		return []models.User{}, err
	}
	defer rows.Close()
	users := []models.User{}
	for rows.Next() {
		user := models.User{}
		err = rows.Scan(&user.UserID, &user.Name)
		users = append(users, user)
	}
	if err != nil {
		return []models.User{}, err
	}
	return users, nil
}

// update user
func (u *UserRepo) UpdateUser(db *sql.DB, user models.User) (models.User, error) {
	sql := "update User set Name = $1, PartyIDs = $2 where UserID = $2"
	_, err := db.Exec(sql, user.Name, pq.Array(user.PartyIDs), user.UserID)
	if err != nil {
		return user, err
	}
	return user, nil
}

// add user
func (u *UserRepo) AddUser(db *sql.DB, user models.User) (models.User, error) {
	sql := "insert into User (Name, PartyIDs) values ($1, $2) returning UserID"
	row := db.QueryRow(sql, user.Name, pq.Array(user.PartyIDs))
	err := row.Scan(&user.UserID)
	if err != nil {
		return user, err
	}
	return user, nil
}

// delete user
func (u *UserRepo) DeleteUser(db *sql.DB, userID int) error {
	sql := "delete from User where UserID = $1"
	_, err := db.Exec(sql, userID)
	if err != nil {
		return err
	}
	return nil
}

package repository

import (
	"database/sql"
	"errors"
	"fmt"
)

type UserActionPostgres struct {
	db *sql.DB
}

func NewUserActionPostgres(db *sql.DB) *UserActionPostgres {
	return &UserActionPostgres{db: db}

}

func (c *UserActionPostgres) IsUserExists(username string) (bool, error) {

	const isExistsQuery = `SELECT 1 FROM users WHERE username=$1`
	var exists int
	err := c.db.QueryRow(isExistsQuery, username).Scan(&exists)

	if errors.Is(err, sql.ErrNoRows) {
		return false, fmt.Errorf("user-action-repo: IsUserExists() : user not exists: %s", username)
	}
	if err != nil {
		return false, fmt.Errorf("user-action-repo: IsUserExists() : cant't check is user exist: %v", err)
	}

	return true, nil
}
func (c *UserActionPostgres) TakeUserTgID(username string) (int64, error) {
	var telegramID int64
	const getTgIDQuery = `SELECT telegram_id FROM users WHERE username=$1`
	err := c.db.QueryRow(getTgIDQuery, username).Scan(&telegramID)
	if err != nil {
		return 0, fmt.Errorf("user-action-repo: TakeUserTgID() : cant' get user tgID: %v", err)

	}

	return telegramID, nil
}

func (c *UserActionPostgres) AddUser(username, userid string, tgID int64) error {

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	const addUserQuery = "INSERT INTO users (username,id,telegram_id) VALUES ($1,$2,$3)"
	_, err = tx.Exec(addUserQuery, username, userid, tgID)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("user-action-repo: AddUser() : cant't do rollback table users: %v", err)
		}
		return fmt.Errorf("user-action-repo: AddUser() : cant't add user: %s", err)
	}

	return tx.Commit()
}

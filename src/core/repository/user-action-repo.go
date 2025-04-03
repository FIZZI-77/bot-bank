package repository

import (
	"database/sql"
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
	result, err := c.db.Query(isExistsQuery, username)
	if err != nil {
		return false, fmt.Errorf("не удалось проверить существование пользователя: %v", err)
	}

	if result == nil {
		return false, fmt.Errorf("пользователь не существет")
	}
	return true, nil
}
func (c *UserActionPostgres) TakeUserTgID(username string) (int64, error) {
	var telegramID int64
	const getTgIDQuery = `SELECT telegram_id FROM users WHERE username=$1`
	err := c.db.QueryRow(getTgIDQuery, username).Scan(&telegramID)
	if err != nil {
		return 0, fmt.Errorf("\n cant' get user tgID: %v", err)

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
			return fmt.Errorf("не удалось сделать rollback таблицы users: %v", err)
		}
		return fmt.Errorf("не удалось добавить пользователя: %s", err)
	}

	return tx.Commit()
}

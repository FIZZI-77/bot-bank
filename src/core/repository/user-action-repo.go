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

	isExistsQuery := fmt.Sprintf("SELECT 1 FROM users WHERE username=$1")
	result, err := c.db.Query(isExistsQuery, username)
	if err != nil {
		return false, fmt.Errorf("не удалось проверить существование пользователя: %v", err)
	}

	if result == nil {
		return false, fmt.Errorf("пользователь не существет")
	}
	return true, nil
}

func (c *UserActionPostgres) AddUser(username string, tgID int) error {

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	addUserQuery := fmt.Sprintf("INSERT INTO users (username,telegram_id) VALUES ($1,$2)")
	_, err = tx.Exec(addUserQuery, username, tgID)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("не удалось сделать rollback таблицы users: %v", err)
		}
		return fmt.Errorf("не удалось добавить пользователя: %s", err)
	}

	return tx.Commit()
}

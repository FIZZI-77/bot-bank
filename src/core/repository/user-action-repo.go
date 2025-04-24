package repository

import (
	"context"
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

func (c *UserActionPostgres) UserExistsByUsername(username string) (bool, error) {

	const isExistsQuery = `SELECT EXISTS (SELECT 1 FROM users WHERE username=$1)`
	var exists bool
	err := c.db.QueryRow(isExistsQuery, username).Scan(&exists)

	if err != nil {
		return false, fmt.Errorf("user-action-repo: UserExistsByUsername() : cant't check is user exist: %v", err)
	}

	if !exists {
		return false, fmt.Errorf("user-action-repo: UserExistsByUsername() : user not exists: %s", username)
	}

	return true, nil
}
func (c *UserActionPostgres) GetUserTgIDByUsername(ctx context.Context, username string) (int64, error) {
	var telegramID int64
	const getTgIDQuery = `SELECT telegram_id FROM users WHERE username=$1`
	err := c.db.QueryRowContext(ctx, getTgIDQuery, username).Scan(&telegramID)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return 0, fmt.Errorf("getTgIDQuery timed out: %w", err)
		}
		return 0, fmt.Errorf("user-action-repo: GetUserTgIDByUsername() : cant' get user tgID: %v", err)

	}

	return telegramID, nil
}

func (c *UserActionPostgres) PersistUser(username, userid string, tgID int64) error {

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	const addUserQuery = "INSERT INTO users (username,id,telegram_id) VALUES ($1,$2,$3)"
	_, err = tx.Exec(addUserQuery, username, userid, tgID)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("user-action-repo: PersistUser() : cant't do rollback table users: %v", err)
		}
		return fmt.Errorf("user-action-repo: PersistUser() : cant't add user: %s", err)
	}

	return tx.Commit()
}

func (c *UserActionPostgres) GetIDByUsername(ctx context.Context, username string) (string, error) {
	var uuid string
	const getUUIDByUsernameQuery = `SELECT id FROM users WHERE username=$1`

	err := c.db.QueryRowContext(ctx, getUUIDByUsernameQuery, username).Scan(&uuid)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return "", fmt.Errorf("getUUIDByUsernameQuery timed out: %w", err)
		}
		return "", fmt.Errorf("user-action-repo: GetUserUUIDByUsername() : cant't get uuid: %v", err)
	}
	return uuid, nil
}

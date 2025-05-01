package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
)

type UserActionPostgres struct {
	db *sql.DB
}

func NewUserActionPostgres(db *sql.DB) *UserActionPostgres {
	return &UserActionPostgres{db: db}

}

func (c *UserActionPostgres) UserExistsByUsername(ctx context.Context, username string) (bool, error) {

	const isExistsQuery = `SELECT EXISTS (SELECT 1 FROM users WHERE username=$1)`
	var exists bool
	err := c.db.QueryRowContext(ctx, isExistsQuery, username).Scan(&exists)

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

func (c *UserActionPostgres) PersistUser(ctx context.Context, username string, userid uuid.UUID, tgID int64) error {

	tx, err := c.db.Begin()
	if err != nil {
		return err
	}
	const addUserQuery = "INSERT INTO users (username,id,telegram_id) VALUES ($1,$2,$3)"
	_, err = tx.ExecContext(ctx, addUserQuery, username, userid, tgID)

	if err != nil {
		if err := tx.Rollback(); err != nil {
			return fmt.Errorf("user-action-repo: PersistUser() : cant't do rollback table users: %v", err)
		}
		return fmt.Errorf("user-action-repo: PersistUser() : cant't add user: %s", err)
	}

	return tx.Commit()
}

func (c *UserActionPostgres) GetIDByUsername(ctx context.Context, username string) (uuid.UUID, error) {
	var id uuid.UUID
	const getUUIDByUsernameQuery = `SELECT id FROM users WHERE username=$1`

	err := c.db.QueryRowContext(ctx, getUUIDByUsernameQuery, username).Scan(&id)
	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return uuid.UUID{}, fmt.Errorf("getUUIDByUsernameQuery timed out: %w", err)
		}
		return uuid.UUID{}, fmt.Errorf("user-action-repo: GetUserUUIDByUsername() : cant't get uuid: %v", err)
	}
	return id, nil
}

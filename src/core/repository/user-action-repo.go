package repository

import "database/sql"

type UserActionPostgres struct {
	db *sql.DB
}

func NewUserActionPostgres(db *sql.DB) *UserActionPostgres {
	return &UserActionPostgres{db: db}
}

func (c *UserActionPostgres) IsUserExists(username string) (bool, error) {
	return false, nil
}

func (c *UserActionPostgres) AddUser(username string) error {
	return nil
}

package repository

import (
	"database/sql"
)

type TopUpPostgres struct {
	db *sql.DB
}

func NewTopUpPostgres(db *sql.DB) *TopUpPostgres {
	return &TopUpPostgres{db: db}
}

func (c *TopUpPostgres) TopUpMoney(username string, amount int) error {
	return nil
}

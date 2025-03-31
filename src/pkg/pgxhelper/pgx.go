package pgxhelper

import (
	"database/sql"
	"fmt"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	SSLMode  string
}

func NewPostgresDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DbName, cfg.Password, cfg.SSLMode))

	if err != nil {
		return nil, err
	}
	
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

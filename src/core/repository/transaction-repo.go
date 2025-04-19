package repository

import (
	"database/sql"
	"fmt"
	"github.com/sirupsen/logrus"
	"tg_transaction/src/core/models"
)

type TransactionPostgres struct {
	db *sql.DB
}

func NewTransactionPostgres(db *sql.DB) *TransactionPostgres {
	return &TransactionPostgres{db: db}
}

func (c *TransactionPostgres) PersistTransaction(userId string, amount float64, recipient int64, transactionId string, cur models.CurrencyEnum) error {
	var currency string
	currency, err := FromCurrencyEnum(cur)
	if err != nil {
		return err
	}

	const sendQuery = "INSERT INTO transactions (id,sender_uuid,amount,currency,recipient_tg_id) VALUES ($1,$2,$3,$4,$5)"

	_, err = c.db.Exec(sendQuery, transactionId, userId, amount, currency, recipient)
	if err != nil {

		return fmt.Errorf("run sql sendQuery: %w", err)
	}
	return nil
}

func (c *TransactionPostgres) GetTotalSentAmount(userID string) (models.Balance, error) {

	const getTotalSendQuery = `SELECT
    	currency,
		COALESCE(SUM(CASE WHEN sender_uuid = $1 THEN amount ELSE 0 END ),0)
	FROM transactions
	GROUP BY currency
`
	rows, err := c.db.Query(getTotalSendQuery, userID)
	if err != nil {
		return models.Balance{}, fmt.Errorf("run sql getTotalSendQuery: %w", err)
	}

	defer func() {
		if err = rows.Close(); err != nil {
			logrus.Error("Error rows.Close(): %v", err)
		}
	}()

	balance := make(models.Balance)
	for rows.Next() {
		var currency string
		var amount float64
		if err := rows.Scan(&currency, &amount); err != nil {
			return models.Balance{}, fmt.Errorf("error rows.Scan(): %v", err)
		}
		curEnum, err := ParseCurrencyToEnum(currency)
		if err != nil {
			return models.Balance{}, fmt.Errorf("error ParseCurrencyToEnum(currency): %w", err)
		}

		balance[curEnum] = amount
	}
	return balance, nil
}

func (c *TransactionPostgres) GetTotalReceivedAmount(userTgID int64) (models.Balance, error) {

	const getTotalReceivedQuery = `SELECT
    currency,
		COALESCE(SUM(CASE WHEN recipient_tg_id = $1 THEN amount ELSE 0 END ),0)
	FROM transactions
	GROUP BY currency
`
	rows, err := c.db.Query(getTotalReceivedQuery, userTgID)
	if err != nil {
		return models.Balance{}, fmt.Errorf("run sql getTotalSendQuery: %w", err)
	}

	defer func() {
		if err = rows.Close(); err != nil {
			logrus.Error("Error rows.Close(): %v", err)
		}
	}()

	balance := make(models.Balance)
	for rows.Next() {
		var currency string
		var amount float64
		if err := rows.Scan(&currency, &amount); err != nil {
			return models.Balance{}, fmt.Errorf("error rows.Scan(): %v", err)
		}
		curEnum, err := ParseCurrencyToEnum(currency)
		if err != nil {
			return models.Balance{}, fmt.Errorf("error ParseCurrencyToEnum(currency): %w", err)
		}

		balance[curEnum] = amount
	}
	return balance, nil
}

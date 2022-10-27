package models

import (
	"database/sql"
	"time"
	"wallet-server/helpers"
)

type Transaction struct {
	ID              string    `json:"id"`
	TransactionType uint      `json:"transaction_type"`
	Value           float32   `json:"value"`
	CreatedAt       time.Time `json:"created_at"`
	WalletID        string    `json:"wallet_id"`
}

type TransactionModel struct {
	DB *sql.DB
}

func (t *TransactionModel) Create(tp uint, value float32, walletID string) (int, error) {
	randomID, err := helpers.CreateRandomID(100)
	if err != nil {
		return 0, err
	}
	stmt := `INSERT INTO transactions(id,transaction_type,value,created_at,wallet_id)VALUES(?,?,?,UTC_TIMESTAMP(),?)`

	res, err := t.DB.Exec(stmt, randomID, tp, value, walletID)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

package models

import (
	"database/sql"
	"time"
	"wallet-server/helpers"
)

type Wallet struct {
	ID        string    `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WalletModel struct {
	DB *sql.DB
}

func (wm *WalletModel) Create(fn, ln string) (int, error) {
	randomID, err := helpers.CreateRandomID(100)
	if err != nil {
		return 0, err
	}

	stmt := `INSERT INTO wallets(id,first_name,last_name,created_at,updated_at)VALUES(?,?,?,UTC_TIMESTAMP(),UTC_TIMESTAMP())`

	res, err2 := wm.DB.Exec(stmt, randomID, fn, ln)
	if err2 != nil {
		return 0, err2
	}

	id, err := res.LastInsertId()
	return int(id), nil

}

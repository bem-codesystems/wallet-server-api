package models

import (
	"context"
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
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()
	randomID, err := helpers.CreateRandomID(100)
	if err != nil {
		return 0, err
	}

	stmt := `INSERT INTO wallets(
                    id,
                    first_name,
                    last_name,
                    created_at,
                    updated_at)VALUES(?,?,?,UTC_TIMESTAMP(),UTC_TIMESTAMP())`

	res, err2 := wm.DB.ExecContext(ctx, stmt, randomID, fn, ln)
	if err2 != nil {
		return 0, err2
	}

	id, err := res.LastInsertId()
	return int(id), nil

}

func (wm *WalletModel) GetSingle(id string) (*Wallet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()
	stmt := `SELECT * FROM wallets WHERE id = ?`

	walletRef := &Wallet{}

	row := wm.DB.QueryRowContext(ctx, stmt, id)

	err := row.Scan(&walletRef.ID, &walletRef.FirstName, &walletRef.LastName, &walletRef.CreatedAt, &walletRef.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return walletRef, nil
}

func (wm *WalletModel) GetList() ([]*Wallet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)

	defer cancel()
	stmt := `SELECT * FROM wallets ORDER BY first_name LIMIT 20`

	var walletList []*Wallet

	rows, err := wm.DB.QueryContext(ctx, stmt)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		wallet := &Wallet{}
		err := rows.Scan(
			&wallet.ID,
			&wallet.FirstName,
			&wallet.LastName,
			&wallet.CreatedAt,
			&wallet.UpdatedAt)
		if err != nil {
			return nil, err
		}
		walletList = append(walletList, wallet)
	}
	return walletList, nil
}

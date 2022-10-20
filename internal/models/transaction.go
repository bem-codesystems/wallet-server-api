package models

import "database/sql"

type Transaction struct {
	Id              string  `json:"id"`
	TransactionType uint    `json:"transaction_type"`
	Value           float32 `json:"value"`
	WalletId        string  `json:"wallet_id"`
}

type TransactionModel struct {
	DB *sql.DB
}

package models

type Transaction struct {
	Id              string  `json:"id"`
	TransactionType uint    `json:"transactionType"`
	Value           float32 `json:"value"`
	WalletId        string  `json:"walletId"`
}

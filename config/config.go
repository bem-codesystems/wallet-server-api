package config

import (
	"log"
	"wallet-server/internal/models"
)

type Application struct {
	ErrorLog    *log.Logger
	InfoLog     *log.Logger
	Wallet      *models.WalletModel
	Transaction *models.TransactionModel
}

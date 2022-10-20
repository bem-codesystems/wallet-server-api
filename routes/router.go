package routes

import (
	"net/http"
	"wallet-server/config"
)

func Router(app *config.Application) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", PingHandler(app))
	mux.HandleFunc("/wallet/new", CreateWalletHandler(app))
	return mux
}

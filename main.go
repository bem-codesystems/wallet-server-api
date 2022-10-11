package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
	"wallet-server/routes"
)

type Fetch struct {
	url    string
	client *http.Client
}

func New() (*Fetch, error) {
	return &Fetch{
		url: os.Getenv("API_BASE_URL"),
		client: &http.Client{
			Timeout: time.Second * 60,
		},
	}, nil
}

func initRouter(mux *http.ServeMux) error {
	mux.HandleFunc("/", routes.PingHandler)
	mux.HandleFunc("/wallet/new", routes.CreateWalletHandler)
	return nil
}

func main() {
	port := os.Getenv("PORT")
	exec := fmt.Sprintf(":%s", port)
	mux := http.NewServeMux()

	err := initRouter(mux)
	if err != nil {
		log.Fatalf("error while initiating router:%s", err)
	}

	err = http.ListenAndServe(exec, mux)
	if err != nil {
		log.Fatalf("error running server: %s", err)
	}
	msg := fmt.Sprintf("server listening on %s", port)
	fmt.Println(msg)
}

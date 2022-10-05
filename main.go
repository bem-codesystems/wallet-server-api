package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
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
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		payload := map[string]interface{}{
			"message":    "Server running.Pong.",
			"statusCode": http.StatusOK,
		}
		jsonData, err := json.Marshal(payload)
		if err != nil {
			log.Fatalf("error found while marshalling payload: %s", err)
			return
		}
		if _, err := writer.Write(jsonData); err != nil {
			log.Fatalf("error found while writing bytes: %s", err)
			return
		}
		return
	})
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

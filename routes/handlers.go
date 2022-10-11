package routes

import (
	"encoding/json"
	"log"
	"net/http"
	"wallet-server/helpers"
	"wallet-server/models"
)

func PingHandler(writer http.ResponseWriter, request *http.Request) {
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
}

func CreateWalletHandler(writer http.ResponseWriter, request *http.Request) {

	_, err := helpers.CheckRequestMethod(request, "POST")
	if err != nil {
		payload := map[string]interface{}{
			"message":    "invalid method provided",
			"statusCode": http.StatusMethodNotAllowed,
		}
		jsonData, err := json.Marshal(payload)
		if err != nil {
			log.Fatalf("error while marshalling data: %s", err)
			return
		}
		writer.WriteHeader(http.StatusMethodNotAllowed)
		if _, err := writer.Write(jsonData); err != nil {
			log.Fatalf("error while marshalling data: %s", err)
			return
		}
		log.Fatalf("invalid method provided")
		return
	}

	validContent := "application/json"
	_, err = helpers.CheckValidContent(request, validContent)
	if err != nil {
		payload := map[string]interface{}{
			"message":    "invalid request content",
			"statusCode": http.StatusForbidden,
		}
		jsonData, err := json.Marshal(payload)
		if err != nil {
			log.Fatalf("error while marshalling data: %s", err)
			return
		}
		writer.WriteHeader(http.StatusForbidden)
		if _, err := writer.Write(jsonData); err != nil {
			log.Fatalf("error while marshalling data: %s", err)
			return
		}
		log.Fatalf("invalid request content")
		return
	}

	body := request.Body
	decodedBody := json.NewDecoder(body)
	decodedBody.DisallowUnknownFields()
	wallet := models.NewWallet{}
	err = decodedBody.Decode(&wallet)
	if err != nil {
		log.Fatalf("error while decoding payload: %s", err)
		return
	}

}

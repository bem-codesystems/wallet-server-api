package routes

import (
	"encoding/json"
	"net/http"
	"wallet-server/config"
	"wallet-server/helpers"
	"wallet-server/internal/models"
)

func PingHandler(app *config.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.WriteHeader(http.StatusOK)
		payload := map[string]interface{}{
			"message":    "Server running.Pong.",
			"statusCode": http.StatusOK,
		}
		jsonData, err := json.Marshal(payload)
		if err != nil {
			app.ErrorLog.Fatalf("error found while marshalling payload: %s", err)
			return
		}
		if _, err := writer.Write(jsonData); err != nil {
			app.ErrorLog.Fatalf("error found while writing bytes: %s", err)
			return
		}
		return
	}
}

func CreateWalletHandler(app *config.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		_, err := helpers.CheckRequestMethod(request, "POST")
		if err != nil {
			payload := map[string]interface{}{
				"message":    "invalid method provided",
				"statusCode": http.StatusMethodNotAllowed,
			}
			jsonData, err := json.Marshal(payload)
			if err != nil {
				app.ErrorLog.Fatalf("error while marshalling data: %s", err)
				return
			}
			writer.WriteHeader(http.StatusMethodNotAllowed)
			if _, err := writer.Write(jsonData); err != nil {
				app.ErrorLog.Fatalf("error while marshalling data: %s", err)
				return
			}
			app.ErrorLog.Fatalf("invalid method provided")
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
				app.ErrorLog.Fatalf("error while marshalling data: %s", err)
				return
			}
			writer.WriteHeader(http.StatusForbidden)
			if _, err := writer.Write(jsonData); err != nil {
				app.ErrorLog.Fatalf("error while marshalling data: %s", err)
				return
			}
			app.ErrorLog.Fatalf("invalid request content")
			return
		}

		body := request.Body
		decodedBody := json.NewDecoder(body)
		decodedBody.DisallowUnknownFields()
		wallet := models.Wallet{}
		err = decodedBody.Decode(&wallet)
		if err != nil {
			app.ErrorLog.Fatalf("error while decoding payload: %s", err)
			return
		}
	}
}

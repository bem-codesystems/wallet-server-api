package helpers

import (
	"errors"
	"math/rand"
	"net/http"
	"time"
)

func CheckRequestMethod(request *http.Request, allowedMethod string) (bool, error) {
	validMethod := request.Method == allowedMethod
	if !validMethod {
		return false, errors.New("")
	}
	return true, nil
}

func CheckValidContent(request *http.Request, acceptedContent string) (bool, error) {
	validContent := len(acceptedContent) > 0 && request.Header.Get("Content-Type") == acceptedContent
	if !validContent {
		return false, errors.New("")
	}
	return true, nil
}

func CreateRandomID(length int) (string, error) {
	rand.Seed(time.Now().UnixNano())
	bytes := make([]byte, length)
	for i := 0; i < length; i++ {
		bytes[i] = byte(65 + rand.Intn(90-65))
	}
	return string(bytes), nil
}

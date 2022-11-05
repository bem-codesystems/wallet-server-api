package helpers

import (
	"errors"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"unicode/utf8"
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

func CheckMainStringVars(keyToCheck string, replacer string) error {
	if strings.TrimSpace(keyToCheck) == "" || utf8.RuneCountInString(keyToCheck) == 0 {
		keyToCheck = replacer
	}
	return errors.New("invalid key checked")

}

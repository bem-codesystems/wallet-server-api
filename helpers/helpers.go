package helpers

import (
	"errors"
	"net/http"
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

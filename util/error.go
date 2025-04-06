package util

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/GabriellaAmah/go-url-shortner/config"
)

type MakeError struct {
	IsError bool   `json:"error"`
	Message string `json:"message"`
	Code    int    `json:"status"`
	Error   any  `json:"errorDetails"`
}

func ConstructError(err any, message string, code int) MakeError {
	errorStruct := MakeError{
		Error:   err,
		Message: message,
		Code:    code,
		IsError: true,
	}
	return errorStruct
}

func HttpErrorResponse(w http.ResponseWriter, errPayload MakeError)  {
	payload := errPayload

	if config.EnvData.ENVIROMENT != "developement" {
		payload = MakeError{Message: errPayload.Message, Code: errPayload.Code, IsError: true}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errPayload.Code)
	if err := json.NewEncoder(w).Encode(payload); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
	}
}

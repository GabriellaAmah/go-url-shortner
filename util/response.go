package util

import (
	"encoding/json"
	"net/http"
)

type SuccessResponseType[T any] struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    T    `json:"data"`
}

func HttpSuccessResponse[T any](w http.ResponseWriter, status int, msg string, data T) {
	if status == 0 {
		status = http.StatusOK
	}
	resp := SuccessResponseType[T]{Success: true, Message: msg, Data: data}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	err := json.NewEncoder(w).Encode(resp)
	if err != nil{
		panic(err)
	}
}

package jopher

import (
	"encoding/json"
	"errors"
	"net/http"
)

// Write transforms the passed message to a json object,
// then sends response with the status code passed
func Write(w http.ResponseWriter, status int, msg interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(status)
	enc, err := json.Marshal(msg)
	if err != nil {
		InternalServerError(w, errors.New("Failed to encode response body to json"))
	}
	w.Write([]byte(enc))
}

// Success return a 200 response with the passed message as the body of the response
func Success(w http.ResponseWriter, msg interface{}) {
	Write(w, http.StatusOK, msg)
}

// Created return a 201 response with the passed message as the body of the response
func Created(w http.ResponseWriter, msg interface{}) {
	Write(w, http.StatusOK, msg)
}

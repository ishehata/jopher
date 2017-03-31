package jopher

import "net/http"

// Error sends a json error message
func Error(w http.ResponseWriter, status int, err error) {
	w.Header().Set("Content-Type", "application/json")
	Write(w, status, err)
}

// BadRequest return 400 response
func BadRequest(w http.ResponseWriter, err error) {
	Error(w, http.StatusBadRequest, err)
}

// Unauthorized return 401 response
func Unauthorized(w http.ResponseWriter, err error) {
	Error(w, http.StatusUnauthorized, err)
}

// NotFound return a 404 response
func NotFound(w http.ResponseWriter, err error) {
	Error(w, http.StatusNotFound, err)
}

// MethodNotAllowed return 405 response
func MethodNotAllowed(w http.ResponseWriter, err error) {
	Error(w, http.StatusMethodNotAllowed, err)
}

// InternalServerError return 505 response
func InternalServerError(w http.ResponseWriter, err error) {
	Error(w, http.StatusInternalServerError, err)
}

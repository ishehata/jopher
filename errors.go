package jopher

import "net/http"

// Error sends a json error message
func Error(w http.ResponseWriter, status int, err error) {
	m := M{"error": err.Error()}
	Write(w, status, m)
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

// HandleError checks if the passed err is not nil,
// if it is: it will write InternalServerError to ResponseWriter
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		InternalServerError(w, err)
	}
}

// HandleErrorWithStatus checks if the passed err is not nil,
// if it is: it will write the passed status code with error message to ResponseWriter
func HandleErrorWithStatus(w http.ResponseWriter, status int, err error) {
	if err != nil {
		Error(w, status, err)
		return
	}
}

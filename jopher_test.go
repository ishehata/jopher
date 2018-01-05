package jopher

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestSuccessResponse(t *testing.T) {
	handler := func(w http.ResponseWriter, r *http.Request) {
		Success(w, M{"Hello": "World"})
	}

	req := httptest.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()
	handler(w, req)

	resp := w.Result()
	body := make(map[string]string)
	if err := json.NewDecoder(resp.Body).Decode(&body); err != nil {
		t.Log(err.Error())
	}

	defer resp.Body.Close()

	// parse body as json

	if body["Hello"] != "World" {
		t.Log("Failed to get correct values from response body")
		t.Fail()
	}

	if resp.Status != "OK" {
		t.Log("Response status code != 200")
		t.Fail()
	}

	if !strings.Contains(resp.Header.Get("Content-Type"), "application/json") {
		t.Log("Response header Content-Type should be application/json")
		t.Fail()
	}
}

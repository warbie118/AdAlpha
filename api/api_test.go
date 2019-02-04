package api_test

import (
	"AdAlpha/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

var a api.Api

func executeRequest(r *http.Request) *httptest.ResponseRecorder {
	resp := httptest.NewRecorder()
	a.Router.ServeHTTP(resp, r)

	return resp
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. got %d\n", expected, actual)
	}
}

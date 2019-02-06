package api_test

import (
	"AdAlpha/api"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
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

func SetTestingEnvVariables() {

	err := os.Setenv("DB_NAME", "test")
	err = os.Setenv("DB_USERNAME", "test")
	err = os.Setenv("DB_PASSWORD", "test")
	err = os.Setenv("DB_PORT", "5001")

	if err != nil {
		fmt.Println("Issue setting test env variables")
	}
}

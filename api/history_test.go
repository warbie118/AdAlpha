package api_test

import (
	"net/http"
	"testing"
)

func TestGetInvestorHistoryWhenProvidedInvestorIdDoesNotExistReturnsHttpCode404(t *testing.T) {
	a.Initialise()
	req, _ := http.NewRequest("POST", "/history/investor/FAKE", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, resp.Code)
}

func TestNewBuyWhenInvestorIdIsValidReturnsHttpCode200(t *testing.T) {
	a.Initialise()
	req, _ := http.NewRequest("POST", "/history/investor/1", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)
}

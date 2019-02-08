package api_test

import (
	"bytes"
	"net/http"
	"testing"
)

func init() {
	a.Initialise()
}

func TestNewBuyWhenRequestBodyIsValidReturnsHttpCode200(t *testing.T) {
	req, _ := http.NewRequest("POST", "/instruction/buy", bytes.NewBuffer(buyRequestBody(`GB00BG0QP828`)))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)
}

func TestNewBuyWhenRequestBodyIsNilReturnsHttpCode400(t *testing.T) {
	req, _ := http.NewRequest("POST", "/instruction/buy", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewBuyWhenRequestBodyDoesNotMapToBuyStructReturnsHttpCode400(t *testing.T) {
	req, _ := http.NewRequest("POST", "/instruction/buy", bytes.NewBuffer(investRequestBody("blah")))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewBuyWhenCalculationFailsReturnsHttpCode400(t *testing.T) {
	req, _ := http.NewRequest("POST", "/instruction/buy", bytes.NewBuffer(buyRequestBody("sausage")))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewInvestWhenRequestBodyIsValidReturnsHttpCode200(t *testing.T) {
	req, _ := http.NewRequest("POST", "/instruction/invest", bytes.NewBuffer(investRequestBody("GB00BQ1YHQ70")))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)
}

func TestNewInvestWhenRequestBodyIsNilReturnsHttpCode400(t *testing.T) {
	req, _ := http.NewRequest("POST", "/instruction/invest", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewInvestWhenRequestBodyDoesNotMapToInvestStructReturnsHttpCode400(t *testing.T) {
	req, _ := http.NewRequest("POST", "/instruction/invest", bytes.NewBuffer(buyRequestBody("blah")))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewInvestWhenCalculationFailsReturnsHttpCode400(t *testing.T) {
	req, _ := http.NewRequest("POST", "/instruction/invest", bytes.NewBuffer(investRequestBody("chips")))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewSellWhenRequestBodyIsValidReturnsHttpCode200(t *testing.T) {
	req, _ := http.NewRequest("POST", "/instruction/sell", bytes.NewBuffer(sellRequestBody("GB00B3X7QG63")))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)
}

func TestNewSellWhenRequestBodyIsNilReturnsHttpCode400(t *testing.T) {
	req, _ := http.NewRequest("POST", "/instruction/sell", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewSellWhenRequestBodyDoesNotMapToInvestStructReturnsHttpCode400(t *testing.T) {
	req, _ := http.NewRequest("POST", "/instruction/sell", bytes.NewBuffer(raiseRequestBody("blah")))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewSellWhenCalculationFailsReturnsHttpCode400(t *testing.T) {
	req, _ := http.NewRequest("POST", "/instruction/sell", bytes.NewBuffer(sellRequestBody("eggs")))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewRaiseWhenRequestBodyIsValidReturnsHttpCode200(t *testing.T) {
	req, _ := http.NewRequest("POST", "/instruction/raise", bytes.NewBuffer(raiseRequestBody("IE00B1S74Q32")))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)
}

func TestNewRaiseWhenRequestBodyIsNilReturnsHttpCode400(t *testing.T) {
	req, _ := http.NewRequest("POST", "/instruction/raise", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewRaiseWhenRequestBodyDoesNotMapToInvestStructReturnsHttpCode400(t *testing.T) {
	req, _ := http.NewRequest("POST", "/instruction/raise", bytes.NewBuffer(sellRequestBody("blah")))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewRaiseWhenCalculationFailsReturnsHttpCode400(t *testing.T) {
	req, _ := http.NewRequest("POST", "/instruction/raise", bytes.NewBuffer(raiseRequestBody("beans")))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func buyRequestBody(isin string) []byte {
	return []byte(`{"investor_id":1,"isin":"` + isin + `","units":5}`)
}

func investRequestBody(isin string) []byte {
	return []byte(`{"investor_id":1, "isin":"` + isin + `", "currency_code":"GBP", "amount":5.55}`)
}

func sellRequestBody(isin string) []byte {
	return []byte(`{"investor_id":1, "isin":"` + isin + `", "units":5}`)
}

func raiseRequestBody(isin string) []byte {
	return []byte(`{"investor_id":1, "isin":"` + isin + `", "currency_code":"GBP", "amount":11.25}`)
}

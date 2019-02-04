package api_test

import (
	"bytes"
	"net/http"
	"testing"
)

func TestNewBuyWhenRequestBodyIsValidReturnsHttpCode200(t *testing.T) {
	a.Initialise()
	req, _ := http.NewRequest("POST", "/instruction/buy", bytes.NewBuffer(buyRequestBody()))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)
}

func TestNewBuyWhenRequestBodyIsNilReturnsHttpCode400(t *testing.T) {
	a.Initialise()
	req, _ := http.NewRequest("POST", "/instruction/buy", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewBuyWhenRequestBodyDoesNotMapToBuyStructReturnsHttpCode400(t *testing.T) {
	a.Initialise()
	req, _ := http.NewRequest("POST", "/instruction/buy", bytes.NewBuffer(investRequestBody()))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewInvestWhenRequestBodyIsValidReturnsHttpCode200(t *testing.T) {
	a.Initialise()
	req, _ := http.NewRequest("POST", "/instruction/invest", bytes.NewBuffer(investRequestBody()))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)
}

func TestNewInvestWhenRequestBodyIsNilReturnsHttpCode400(t *testing.T) {
	a.Initialise()
	req, _ := http.NewRequest("POST", "/instruction/invest", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewInvestWhenRequestBodyDoesNotMapToInvestStructReturnsHttpCode400(t *testing.T) {
	a.Initialise()
	req, _ := http.NewRequest("POST", "/instruction/invest", bytes.NewBuffer(buyRequestBody()))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewSellWhenRequestBodyIsValidReturnsHttpCode200(t *testing.T) {
	a.Initialise()
	req, _ := http.NewRequest("POST", "/instruction/sell", bytes.NewBuffer(sellRequestBody()))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)
}

func TestNewSellWhenRequestBodyIsNilReturnsHttpCode400(t *testing.T) {
	a.Initialise()
	req, _ := http.NewRequest("POST", "/instruction/sell", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewSellWhenRequestBodyDoesNotMapToInvestStructReturnsHttpCode400(t *testing.T) {
	a.Initialise()
	req, _ := http.NewRequest("POST", "/instruction/sell", bytes.NewBuffer(raiseRequestBody()))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewRaiseWhenRequestBodyIsValidReturnsHttpCode200(t *testing.T) {
	a.Initialise()
	req, _ := http.NewRequest("POST", "/instruction/raise", bytes.NewBuffer(raiseRequestBody()))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)
}

func TestNewRaiseWhenRequestBodyIsNilReturnsHttpCode400(t *testing.T) {
	a.Initialise()
	req, _ := http.NewRequest("POST", "/instruction/raise", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func TestNewRaiseWhenRequestBodyDoesNotMapToInvestStructReturnsHttpCode400(t *testing.T) {
	a.Initialise()
	req, _ := http.NewRequest("POST", "/instruction/sell", bytes.NewBuffer(sellRequestBody()))
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusBadRequest, resp.Code)
}

func buyRequestBody() []byte {
	return []byte(`{"investor_id":"1", "isin":"TESTISIN01", "units":"5"`)
}

func investRequestBody() []byte {
	return []byte(`{"investor_id":"1", "isin":"TESTISIN01", "currency_code":"GBP", "amount":"5"`)
}

func sellRequestBody() []byte {
	return []byte(`{"investor_id":"1", "isin":"TESTISIN01", "units":"5"`)
}

func raiseRequestBody() []byte {
	return []byte(`{"investor_id":"1", "isin":"TESTISIN01", "currency_code":"GBP", "amount":"5"`)
}

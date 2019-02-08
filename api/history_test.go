package api_test

import (
	"AdAlpha/model"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func init() {
	SetTestingEnvVariables()
	a.Initialise()
}

func TestGetInvestorHistoryWhenProvidedInvestorIdDoesNotExistReturnsHttpCode404(t *testing.T) {
	req, _ := http.NewRequest("GET", "/history/investor/3000", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, resp.Code)
}

func TestGetInvestorHistoryWhenInvestorIdIsNotProvidedReturnsHttpCode400(t *testing.T) {
	req, _ := http.NewRequest("GET", "/history/investor/", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, resp.Code)
}

func TestGetInvestorHistoryWhenInvestorIdIsValidReturnsHttpCode200(t *testing.T) {
	req, _ := http.NewRequest("GET", "/history/investor/3", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	var history []model.History

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&history); err != nil {
		fmt.Println(err.Error())
	}

	if len(history) != 2 {
		t.Errorf("Expected history length %v. got %v\n", 2, len(history))
	}

	h1 := history[0]
	if h1.Instruction != "BUY" && h1.Isin != "GB00BG0QP828" && h1.Amount != 10.99 &&
		h1.CurrencyCode != "GBP" && h1.Asset != "Legal & General Japan Index Trust C Class Accumulation" {
		t.Error("Does not match expected data")
	}

	h2 := history[1]
	if h2.Instruction != "BUY" && h2.Isin != "GB00BG0QP828" && h2.Amount != 10.99 &&
		h2.CurrencyCode != "GBP" && h2.Asset != "Legal & General Japan Index Trust C Class Accumulation" {
		t.Error("Does not match expected data")
	}
}

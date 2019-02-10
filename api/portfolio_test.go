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

func TestGetInvestorPortfolioWhenProvidedInvestorIdDoesNotExistReturnsHttpCode404(t *testing.T) {
	req, _ := http.NewRequest("GET", "/portfolio/investor/3000", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, resp.Code)
}

func TestGetInvestorPortfolioWhenInvestorIdIsNotProvidedReturnsHttpCode400(t *testing.T) {
	req, _ := http.NewRequest("GET", "/portfolio/investor/", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusNotFound, resp.Code)
}

func TestGetInvestorPortfolioWhenInvestorIdIsValidReturnsHttpCode200(t *testing.T) {
	req, _ := http.NewRequest("GET", "/portfolio/investor/7", nil)
	resp := executeRequest(req)
	checkResponseCode(t, http.StatusOK, resp.Code)

	var portfolio []model.Portfolio

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&portfolio); err != nil {
		fmt.Println(err.Error())
	}

	if len(portfolio) != 6 {
		t.Errorf("Expected history length %v. got %v\n", 6, len(portfolio))
	}

	p1 := portfolio[0]
	if p1.Isin != "IE00B52L4369" && p1.Asset != "BlackRock Institutional Cash Series Sterling Liquidity Agency Inc" && p1.Units != 50 {
		t.Error("Does not match expected data")
	}

	p2 := portfolio[1]
	if p2.Isin != "GB00BQ1YHQ70" && p2.Asset != "Threadneedle UK Property Authorised Investment Net GBP 1 Acc" && p2.Units != 50 {
		t.Error("Does not match expected data")
	}

	p3 := portfolio[2]
	if p3.Isin != "GB00B3X7QG63" && p3.Asset != "Vanguard FTSE U.K. All Share Index Unit Trust Accumulation" && p3.Units != 50 {
		t.Error("Does not match expected data")
	}

	p4 := portfolio[3]
	if p4.Isin != "GB00BG0QP828" && p4.Asset != "Legal & General Japan Index Trust C Class Accumulation" && p4.Units != 50 {
		t.Error("Does not match expected data")
	}

	p5 := portfolio[4]
	if p5.Isin != "GB00BPN5P238" && p5.Asset != "Vanguard US Equity Index Institutional Plus GBP Accumulation" && p5.Units != 50 {
		t.Error("Does not match expected data")
	}

	p6 := portfolio[5]
	if p6.Isin != "IE00B1S74Q32" && p6.Asset != "Vanguard U.K. Investment Grade Bond Index Fund GBP Accumulation" && p6.Units != 50 {
		t.Error("Does not match expected data")
	}
}

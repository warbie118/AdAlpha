package model_test

import (
	"AdAlpha/exchange_rate"
	"AdAlpha/model"
	"github.com/jarcoal/httpmock"
	"os"
	"testing"
)

func TestInvest_IsValidReturnsFalseWhenInvestorIdIs0(t *testing.T) {
	b := model.Invest{0, "test01", "GBP", 5.55}
	if b.IsValid() == true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestInvest_IsValidReturnsFalseWhenIsinIsEmpty(t *testing.T) {
	b := model.Invest{1, "", "GBP", 5.55}
	if b.IsValid() == true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestInvest_IsValidReturnsFalseWhenCurrencyCodeIsInvalid(t *testing.T) {
	b := model.Invest{1, "test01", "FAKE", 5.55}
	if b.IsValid() == true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestInvest_IsValidReturnsFalseWhenAmountIs0(t *testing.T) {
	b := model.Invest{1, "test01", "GBP", 0}
	if b.IsValid() == true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestInvest_IsValidReturnsTrueWhenAllFieldsValid(t *testing.T) {
	b := model.Invest{1, "test01", "GBP", 5.55}
	if b.IsValid() == false {
		t.Errorf("Expected response %t. got %t\n", true, false)
	}
}

func TestInvest_CalculateInvestIsSuccessfulWhenCurrencyCodesDiffer(t *testing.T) {
	SetTestingEnvVariables()
	i := model.Invest{3, "GB00BQ1YHQ70", "USD", 50}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", exchange_rate.GenerateUrl("USD", os.Getenv("BASE_CC")),
		httpmock.NewStringResponder(200, `{"base": "GBP", "date": "2019-02-07", "rates":{"USD":2}}`))
	err := model.CalculateInvest(5, i)
	if err != nil {
		t.Error("Expected no errors")
	}

	units, valid := correctNumberOfUnitsInPortfolioTable(3, "GB00BQ1YHQ70", 25, 5)
	if !valid {
		t.Errorf("Wrong number of units in investor portfolio, expected 25, got %v", units)
	}
}

func TestInvest_CalculateInvestIsSuccessfulWhenCurrencyCodesMatch(t *testing.T) {
	SetTestingEnvVariables()
	i := model.Invest{3, "GB00BQ1YHQ70", "GBP", 50}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", exchange_rate.GenerateUrl("USD", os.Getenv("BASE_CC")),
		httpmock.NewStringResponder(200, `{"base": "GBP", "date": "2019-02-07", "rates":{"USD":2}}`))
	err := model.CalculateInvest(5, i)
	if err != nil {
		t.Error("Expected no errors")
	}

	units, valid := correctNumberOfUnitsInPortfolioTable(3, "GB00BQ1YHQ70", 15, 5)

	if !valid {
		t.Errorf("Wrong number of units in investor portfolio, expected 15, got %v", units)
	}
}

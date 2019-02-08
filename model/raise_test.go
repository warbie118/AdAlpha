package model_test

import (
	"AdAlpha/exchange_rate"
	"AdAlpha/model"
	"github.com/jarcoal/httpmock"
	"os"
	"testing"
)

func TestRaise_IsValidReturnsFalseWhenRaiseorIdIs0(t *testing.T) {
	b := model.Raise{0, "test01", "GBP", 5.55}
	if b.IsValid() == true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestRaise_IsValidReturnsFalseWhenIsinIsEmpty(t *testing.T) {
	b := model.Raise{1, "", "GBP", 5.55}
	if b.IsValid() == true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestRaise_IsValidReturnsFalseWhenCurrencyCodeIsInvalid(t *testing.T) {
	b := model.Raise{1, "test01", "FAKE", 5.55}
	if b.IsValid() == true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestRaise_IsValidReturnsFalseWhenAmountIs0(t *testing.T) {
	b := model.Raise{1, "test01", "GBP", 0}
	if b.IsValid() == true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestRaise_IsValidReturnsTrueWhenAllFieldsValid(t *testing.T) {
	b := model.Raise{1, "test01", "GBP", 5.55}
	if b.IsValid() == false {
		t.Errorf("Expected response %t. got %t\n", true, false)
	}
}

func TestRaise_CalculateRaiseIsSuccessful(t *testing.T) {
	SetTestingEnvVariables()
	r := model.Raise{4, "IE00B52L4369", "USD", 50}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", exchange_rate.GenerateUrl("USD", os.Getenv("BASE_CC")),
		httpmock.NewStringResponder(200, `{"base": "GBP", "date": "2019-02-07", "rates":{"USD":2}}`))
	err := model.CalculateRaise(5, r)
	if err != nil {
		t.Error("Expected no errors")
	}

	units, valid := correctNumberOfUnitsInPortfolioTable(4, "IE00B52L4369", 30, 50)
	if !valid {
		t.Errorf("Wrong number of units in investor portfolio, expected 30, got %v", units)
	}
}

func TestRaise_CalculateRaiseIsSuccessfulWhenCurrencyCodesMatch(t *testing.T) {
	SetTestingEnvVariables()
	r := model.Raise{4, "GB00BQ1YHQ70", "GBP", 50}
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", exchange_rate.GenerateUrl("USD", os.Getenv("BASE_CC")),
		httpmock.NewStringResponder(200, `{"base": "GBP", "date": "2019-02-07", "rates":{"USD":2}}`))
	err := model.CalculateRaise(5, r)
	if err != nil {
		t.Error("Expected no errors")
	}

	units, valid := correctNumberOfUnitsInPortfolioTable(4, "GB00BQ1YHQ70", 40, 50)
	if !valid {
		t.Errorf("Wrong number of units in investor portfolio, expected 15, got %v", units)
	}
}

package model_test

import (
	"AdAlpha/model"
	"testing"
)

func TestSell_IsValidReturnsFalseWhenInvestorIdIs0(t *testing.T) {
	b := model.Sell{0, "test01", 5}
	if b.IsValid() == true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestSell_IsValidReturnsFalseWhenIsinIsEmpty(t *testing.T) {
	b := model.Sell{1, "", 5}
	if b.IsValid() == true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestSell_IsValidReturnsFalseWhenUnitsIs0(t *testing.T) {
	b := model.Sell{1, "test01", 0}
	if b.IsValid() == true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestSell_IsValidReturnsTrueWhenAllFieldsAreValid(t *testing.T) {
	b := model.Sell{1, "test01", 10}
	if b.IsValid() == false {
		t.Errorf("Expected response %t. got %t\n", true, false)
	}
}

func TestSell_CalculateSellIsSuccessful(t *testing.T) {
	SetTestingEnvVariables()
	s := model.Sell{3, "GB00BQ1YHQ70", 3}
	err := model.CalculateSell(5.50, s)
	if err != nil {
		t.Error("Expected no errors")
	}

	units, valid := correctNumberOfUnitsInPortfolioTable(3, "GB00BQ1YHQ70", 2, 5)

	if !valid {
		t.Errorf("Wrong number of units in investor portfolio, expected 2, got %v", units)
	}
}

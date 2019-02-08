package model_test

import (
	"AdAlpha/model"
	"testing"
)

func TestBuy_IsValidReturnsFalseWhenInvestorIdIs0(t *testing.T) {
	b := model.Buy{0, "test01", 5}
	if b.IsValid() == true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestBuy_IsValidReturnsFalseWhenIsinIsEmpty(t *testing.T) {
	b := model.Buy{1, "", 5}
	if b.IsValid() == true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestBuy_IsValidReturnsFalseWhenUnitsIs0(t *testing.T) {
	b := model.Buy{1, "test01", 0}
	if b.IsValid() == true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestBuy_IsValidReturnsTrueWhenAllFieldsAreValid(t *testing.T) {
	b := model.Buy{1, "test01", 10}
	if b.IsValid() == false {
		t.Errorf("Expected response %t. got %t\n", true, false)
	}
}

func TestBuy_CalculateBuyIsSuccessful(t *testing.T) {
	SetTestingEnvVariables()
	b := model.Buy{2, "GB00BQ1YHQ70", 2}
	err := model.CalculateBuy(5.50, b)
	if err != nil {
		t.Error("Expected no errors")
	}

	units, valid := correctNumberOfUnitsInPortfolioTable(2, "GB00BQ1YHQ70", 2, 0)
	if !valid {
		t.Errorf("Wrong number of units in investor portfolio, expected 2, got %v", units)
	}
}

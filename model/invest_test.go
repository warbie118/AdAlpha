package model_test

import (
	"AdAlpha/model"
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

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

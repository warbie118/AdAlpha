package model_test

import (
	"AdAlpha/model"
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

package model_test

import (
	"AdAlpha/model"
	"testing"
)

func TestValidCurrencyCodeWhenCodeIsGBP(t *testing.T) {
	valid := model.ValidCurrencyCode("GBP")
	if valid != true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestValidCurrencyCodeWhenCodeIsUSD(t *testing.T) {
	valid := model.ValidCurrencyCode("USD")
	if valid != true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestValidCurrencyCodeWhenCodeIsEUR(t *testing.T) {
	valid := model.ValidCurrencyCode("EUR")
	if valid != true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestValidCurrencyCodeWhenCodeIsCNY(t *testing.T) {
	valid := model.ValidCurrencyCode("CNY")
	if valid != true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestValidCurrencyCodeWhenCodeIsNotValid(t *testing.T) {
	valid := model.ValidCurrencyCode("FAKE")
	if valid != false {
		t.Errorf("Expected response %t. got %t\n", true, false)
	}
}

func TestValidCurrencyCodeWhenCodeIsEmpty(t *testing.T) {
	valid := model.ValidCurrencyCode("")
	if valid != false {
		t.Errorf("Expected response %t. got %t\n", true, false)
	}
}

func TestContainsWhenArrayDoesContainStringReturnsTrue(t *testing.T) {
	contains := model.Contains([]string{"1", "2"}, "1")
	if contains != true {
		t.Errorf("Expected response %t. got %t\n", false, true)
	}
}

func TestContainsWhenArrayDoesNotContainStringReturnsFalse(t *testing.T) {
	contains := model.Contains([]string{"1", "2"}, "3")
	if contains != false {
		t.Errorf("Expected response %t. got %t\n", true, false)
	}
}

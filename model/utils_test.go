package model_test

import (
	"AdAlpha/model"
	"database/sql"
	"fmt"
	"os"
	"testing"
)

func SetTestingEnvVariables() {

	err := os.Setenv("DB_NAME", "test")
	err = os.Setenv("DB_USERNAME", "test")
	err = os.Setenv("DB_PASSWORD", "test")
	err = os.Setenv("DB_PORT", "5001")
	err = os.Setenv("DB_HOST", "localhost")
	err = os.Setenv("BASE_CC", "GBP")

	if err != nil {
		fmt.Println("Issue setting test env variables")
	}
}

func correctNumberOfUnitsInPortfolioTable(id int, isin string, expected float64, resetVal float64) (float64, bool) {
	var units float64
	dbCon := model.GetDbConnection()
	err := dbCon.Pg.QueryRow("select units from portfolio where investor_id=$1 and isin=$2", id, isin).Scan(&units)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("error checking units in portfolio table")
	}

	_, err = dbCon.Pg.Exec("update portfolio set units =$3 where investor_id=$1 and isin=$2", id, isin, resetVal)
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("error resetting DB data")
	}
	return units, units == expected

}

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

func TestContainsReturnsTrueWhenArrayContainsVal(t *testing.T) {
	arr := []string{"cat", "dog"}
	contains := model.Contains(arr, "cat")
	if !contains {
		t.Error("Expected contains to be true")
	}
}

func TestContainsReturnsFalseWhenArrayDoesNotContainVal(t *testing.T) {
	arr := []string{"cat", "dog"}
	contains := model.Contains(arr, "frog")
	if contains {
		t.Error("Expected contains to be false")
	}
}

func TestCalculateUnitsReturnsCorrectValue(t *testing.T) {
	res := model.CalculateUnits(5, 2, 10)
	if res != 1 {
		t.Errorf("Calculation wrong, expected 1 and received %b\n", res)
	}
}

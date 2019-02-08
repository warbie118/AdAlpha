package model_test

import (
	"AdAlpha/model"
	"database/sql"
	"fmt"
	"testing"
)

func TestGetInvestorHistory(t *testing.T) {
	SetTestingEnvVariables()
	dbCon := model.GetDbConnection()
	history, err := model.GetInvestorHistory(dbCon.Pg, 6)

	if err != nil {
		t.Errorf("error thrown - " + err.Error())
	}
	if len(history) != 2 {
		t.Errorf("Wrong amount of history items, expected %d, got %d", 2, len(history))
	}
}

func TestAddInvestorHistory(t *testing.T) {
	SetTestingEnvVariables()
	dbCon := model.GetDbConnection()
	err := model.AddInvestorHistory(dbCon.Pg, 5, "BUY", "GB00BQ1YHQ70", 5, 5, "TST", 10)
	if err != nil {
		t.Errorf("error thrown - " + err.Error())
	}
	history, err := model.GetInvestorHistory(dbCon.Pg, 5)

	//test date for investor_id 5 has 0 history rows prior to test, expect 1
	if len(history) != 1 {
		t.Errorf("Wrong amount of history items, expected %d, got %d", 1, len(history))
	}

	_, err = dbCon.Pg.Exec("delete from instructions where currency_code='TST'")
	if err != nil && err != sql.ErrNoRows {
		fmt.Println("error resetting DB data")
	}
}

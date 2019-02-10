package model_test

import (
	"AdAlpha/model"
	"testing"
)

func TestGetInvestorPortfolio(t *testing.T) {
	SetTestingEnvVariables()
	dbCon := model.GetDbConnection()
	portfolio, err := model.GetInvestorPortfolio(dbCon.Pg, 4)

	if err != nil {
		t.Errorf("error thrown - " + err.Error())
	}
	if len(portfolio) != 6 {
		t.Errorf("Wrong amount of portfolio items, expected %d, got %d", 6, len(portfolio))
	}

	for _, p := range portfolio {

		if p.CurrentPrice == 0 {
			t.Error("Portfolio item does not contain current price of asset")
		}
	}
}

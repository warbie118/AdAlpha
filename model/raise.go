package model

import (
	"AdAlpha/exchange_rate"
	"os"
)

type Raise struct {
	InvestorId   int     `json:"investor_id"`
	Isin         string  `json:"isin"`
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
}

//checks if raise struct is valid
func (r *Raise) IsValid() bool {
	if r.InvestorId == 0 || len(r.Isin) == 0 || !ValidCurrencyCode(r.CurrencyCode) || r.Amount == 0 {
		return false
	}
	return true
}

//Initialises a new raise, gets the current price of the asset before calling CalculateRaise
func (r *Raise) New() error {
	err, price := getCurrentPrice(r.Isin)
	if err != nil {
		return err
	}

	err = CalculateRaise(price, *r)
	if err != nil {
		return err
	}
	return err
}

//Takes the asset price and raise information and calculates the raise.
//Gets the current exchange rate for given currency code to calculate how many units to sell in GBP (BASE_CC)
//Updates portfolio table with raise info and add row to Instructions table (history)
func CalculateRaise(assetPrice float64, r Raise) error {
	//default val is 1 as default currency code is GBP
	var excr float64 = 1

	dbCon := GetDbConnection()

	if r.CurrencyCode != os.Getenv("BASE_CC") {
		err, rate := exchange_rate.GetExchangeRate(r.CurrencyCode, os.Getenv("BASE_CC"))
		excr = rate
		if err != nil {
			return err
		}
	}

	units := CalculateUnits(r.Amount, excr, assetPrice)

	err := RemoveAssetsInPortfolio(dbCon.Pg, units, r.InvestorId, r.Isin)
	if err != nil {
		return err
	}

	err = AddInvestorHistory(dbCon.Pg, r.InvestorId, "RAISE", r.Isin, assetPrice, units, r.CurrencyCode, r.Amount)
	if err != nil {
		AddAssetsInPortfolio(dbCon.Pg, units, r.InvestorId, r.Isin)
		return err
	}

	return err
}

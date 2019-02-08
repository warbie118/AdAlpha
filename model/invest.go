package model

import (
	"AdAlpha/exchange_rate"
	"os"
)

type Invest struct {
	InvestorId   int     `json:"investor_id"`
	Isin         string  `json:"isin"`
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
}

//checks if invest struct is valid
func (i *Invest) IsValid() bool {
	if i.InvestorId == 0 || len(i.Isin) == 0 || !ValidCurrencyCode(i.CurrencyCode) || i.Amount == 0 {
		return false
	}
	return true
}

//Initialises a new invest, gets the current price of the asset before calling CalculateInvest
func (i *Invest) New() error {
	err, price := getCurrentPrice(i.Isin)
	if err != nil {
		return err
	}

	err = CalculateInvest(price, *i)
	if err != nil {
		return err
	}
	return err
}

//Takes the asset price and invest information and calculates the invest.
//Gets the current exchange rate for given currency code to calculate how many units to purchase in GBP (BASE_CC)
//Updates portfolio table with invest info and add row to Instructions table (history)
func CalculateInvest(assetPrice float64, i Invest) error {

	//default val is 1 as default currency code is GBP
	var excr float64 = 1

	dbCon := GetDbConnection()

	if i.CurrencyCode != os.Getenv("BASE_CC") {
		err, rate := exchange_rate.GetExchangeRate(i.CurrencyCode, os.Getenv("BASE_CC"))
		excr = rate
		if err != nil {
			return err
		}
	}

	units := CalculateUnits(i.Amount, excr, assetPrice)

	err := AddAssetsInPortfolio(dbCon.Pg, units, i.InvestorId, i.Isin)
	if err != nil {
		return err
	}

	err = AddInvestorHistory(dbCon.Pg, i.InvestorId, "INVEST", i.Isin, assetPrice, units, i.CurrencyCode, i.Amount)
	if err != nil {
		RemoveAssetsInPortfolio(dbCon.Pg, units, i.InvestorId, i.Isin)
		return err
	}

	return err

}

package model

import (
	"AdAlpha/exchange_rate"
	"AdAlpha/logger"
	"fmt"
	"os"
	"time"
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
	esLog.LogInfo(logger.CreateInfoLog("INFO",
		fmt.Sprintf("New Raise for investor_id: %d, isin: %s", r.InvestorId, r.Isin), time.Now()))
	err, price := getCurrentPrice(r.Isin)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Problem getting current asset price for isin: %s", r.Isin), err.Error(), logger.Trace(), time.Now()))
		return err
	}

	err = CalculateRaise(price, *r)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Problem calculating raise for investor_id: %d, isin: %s", r.InvestorId, r.Isin), err.Error(), logger.Trace(), time.Now()))
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
			esLog.LogError(logger.CreateLog("ERROR",
				fmt.Sprintf("Problem getting exchange rate, currency code: %s", r.CurrencyCode), err.Error(), logger.Trace(), time.Now()))
			return err
		}
	}

	units := CalculateUnits(r.Amount, excr, assetPrice)

	err := RemoveAssetsInPortfolio(dbCon.Pg, units, r.InvestorId, r.Isin)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Problem removing asssets from portfolio, investor_id: %d, isin: %s, units: %v",
				r.InvestorId, r.Isin, units), err.Error(), logger.Trace(), time.Now()))
		return err
	}

	err = AddInvestorHistory(dbCon.Pg, r.InvestorId, "RAISE", r.Isin, assetPrice, units, r.CurrencyCode, r.Amount)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Problem adding raise instruction to instructions, investor_id: %d, isin: %s, instruction: %s, asset price: %v, currency code: %s, amount: %v, units: %v",
				r.InvestorId, r.Isin, "RAISE", assetPrice, "GBP", assetPrice*units, units), err.Error(), logger.Trace(), time.Now()))
		err := AddAssetsInPortfolio(dbCon.Pg, units, r.InvestorId, r.Isin)
		if err != nil {
			esLog.LogError(logger.CreateLog("ERROR",
				fmt.Sprintf("Problem rolling back assets to portfolio after failure, investor_id: %d, isin: %s, units: %v",
					r.InvestorId, r.Isin, units), err.Error(), logger.Trace(), time.Now()))
		}
		return err
	}

	return err
}

package model

import (
	"AdAlpha/exchange_rate"
	"AdAlpha/logger"
	"fmt"
	"os"
	"time"
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

	esLog.LogInfo(logger.CreateInfoLog("INFO",
		fmt.Sprintf("New Invest for investor_id: %d, isin: %s", i.InvestorId, i.Isin), time.Now()))

	err, price := getCurrentPrice(i.Isin)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Problem getting current asset price for isin: %s", i.Isin), err.Error(), logger.Trace(), time.Now()))
		return err
	}

	err = CalculateInvest(price, *i)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Problem calculating invest for investor_id: %d, isin: %s", i.InvestorId, i.Isin), err.Error(), logger.Trace(), time.Now()))
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
			esLog.LogError(logger.CreateLog("ERROR",
				fmt.Sprintf("Problem getting exchange rate, currency code: %s", i.CurrencyCode), err.Error(), logger.Trace(), time.Now()))
			return err
		}
	}

	units := CalculateUnits(i.Amount, excr, assetPrice)

	err := AddAssetsInPortfolio(dbCon.Pg, units, i.InvestorId, i.Isin)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Problem adding asssets to portfolio, investor_id: %d, isin: %s, units: %v",
				i.InvestorId, i.Isin, units), err.Error(), logger.Trace(), time.Now()))
		return err
	}

	err = AddInvestorHistory(dbCon.Pg, i.InvestorId, "INVEST", i.Isin, assetPrice, units, i.CurrencyCode, i.Amount)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Problem adding invest instruction to instructions, investor_id: %d, isin: %s, instruction: %s, asset price: %v, currency code: %s, amount: %v, units: %v",
				i.InvestorId, i.Isin, "INVEST", assetPrice, "GBP", assetPrice*units, units), err.Error(), logger.Trace(), time.Now()))
		err := RemoveAssetsInPortfolio(dbCon.Pg, units, i.InvestorId, i.Isin)
		if err != nil {
			esLog.LogError(logger.CreateLog("ERROR",
				fmt.Sprintf("Problem rolling back assets to portfolio after failure, investor_id: %d, isin: %s, units: %v",
					i.InvestorId, i.Isin, units), err.Error(), logger.Trace(), time.Now()))
		}
		return err
	}

	return err

}

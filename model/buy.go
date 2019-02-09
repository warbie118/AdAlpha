package model

import (
	"AdAlpha/logger"
	"fmt"
	"time"
)

type Buy struct {
	InvestorId int     `json:"investor_id"`
	Isin       string  `json:"isin"`
	Units      float64 `json:"units"`
}

//Checks if buy struct is valid
func (b *Buy) IsValid() bool {
	if b.InvestorId == 0 || len(b.Isin) == 0 || b.Units == 0 {
		return false
	}
	return true
}

//Initialises a new buy, gets the current price of the asset before calling CalculateBuy
func (b *Buy) New() error {

	esLog.LogInfo(logger.CreateInfoLog("INFO",
		fmt.Sprintf("New Buy for investor_id: %d, isin: %s", b.InvestorId, b.Isin), time.Now()))
	err, price := getCurrentPrice(b.Isin)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Problem getting current asset price for isin: %s", b.Isin), err.Error(), logger.Trace(), time.Now()))
		return err
	}

	err = CalculateBuy(price, *b)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Problem calculating buy for investor_id: %d, isin: %s", b.InvestorId, b.Isin), err.Error(), logger.Trace(), time.Now()))
		return err
	}
	return err

}

//Takes the asset price and buy information and calculates the buy.
//Updates portfolio table with buy info and add row to Instructions table (history)
func CalculateBuy(assetPrice float64, b Buy) error {

	dbCon := GetDbConnection()

	err := AddAssetsInPortfolio(dbCon.Pg, b.Units, b.InvestorId, b.Isin)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Problem adding asssets to portfolio, investor_id: %d, isin: %s, units: %v",
				b.InvestorId, b.Isin, b.Units), err.Error(), logger.Trace(), time.Now()))
		return err
	}

	err = AddInvestorHistory(dbCon.Pg, b.InvestorId, "BUY", b.Isin, assetPrice, b.Units, "GBP", assetPrice*b.Units)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Problem adding buy instruction to instructions, investor_id: %d, isin: %s, instruction: %s, asset price: %v, currency code: %s, amount: %v, units: %v",
				b.InvestorId, b.Isin, "BUY", assetPrice, "GBP", assetPrice*b.Units, b.Units), err.Error(), logger.Trace(), time.Now()))
		err := RemoveAssetsInPortfolio(dbCon.Pg, b.Units, b.InvestorId, b.Isin)
		if err != nil {
			esLog.LogError(logger.CreateLog("ERROR",
				fmt.Sprintf("Problem rolling back assets to portfolio after failure, investor_id: %d, isin: %s, units: %v",
					b.InvestorId, b.Isin, b.Units), err.Error(), logger.Trace(), time.Now()))
		}
		return err
	}

	return err

}

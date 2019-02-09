package model

import (
	"AdAlpha/logger"
	"fmt"
	"time"
)

type Sell struct {
	InvestorId int     `json:"investor_id"`
	Isin       string  `json:"isin"`
	Units      float64 `json:"units"`
}

//checks if sell struct is valid
func (s *Sell) IsValid() bool {
	if s.InvestorId == 0 || len(s.Isin) == 0 || s.Units == 0 {
		return false
	}
	return true
}

//Initialises a new sell, gets the current price of the asset before calling CalculateSell
func (s *Sell) NewSell() error {
	esLog.LogInfo(logger.CreateInfoLog("INFO",
		fmt.Sprintf("New Sell for investor_id: %d, isin: %s", s.InvestorId, s.Isin), time.Now()))
	err, price := getCurrentPrice(s.Isin)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Problem getting current asset price for isin: %s", s.Isin), err.Error(), logger.Trace(), time.Now()))
		return err
	}

	err = CalculateSell(price, *s)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Problem calculating sell for investor_id: %d, isin: %s", s.InvestorId, s.Isin), err.Error(), logger.Trace(), time.Now()))
		return err
	}
	return err
}

//Takes the asset price and sell information and calculates the sell.
//Updates portfolio table with buy info and add row to Instructions table (history)
func CalculateSell(assetPrice float64, s Sell) error {

	dbCon := GetDbConnection()

	err := RemoveAssetsInPortfolio(dbCon.Pg, s.Units, s.InvestorId, s.Isin)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Problem selling asssets from portfolio, investor_id: %d, isin: %s, units: %v",
				s.InvestorId, s.Isin, s.Units), err.Error(), logger.Trace(), time.Now()))
		return err
	}

	err = AddInvestorHistory(dbCon.Pg, s.InvestorId, "SELL", s.Isin, assetPrice, s.Units, "GBP", assetPrice*s.Units)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Problem adding sell instruction to instructions, investor_id: %d, isin: %s, instruction: %s, asset price: %v, currency code: %s, amount: %v, units: %v",
				s.InvestorId, s.Isin, "SELL", assetPrice, "GBP", assetPrice*s.Units, s.Units), err.Error(), logger.Trace(), time.Now()))
		err := AddAssetsInPortfolio(dbCon.Pg, s.Units, s.InvestorId, s.Isin)
		if err != nil {
			esLog.LogError(logger.CreateLog("ERROR",
				fmt.Sprintf("Problem rolling back assets to portfolio after failure, investor_id: %d, isin: %s, units: %v",
					s.InvestorId, s.Isin, s.Units), err.Error(), logger.Trace(), time.Now()))
		}
		return err
	}

	return err
}

package model

import (
	"AdAlpha/db"
	"AdAlpha/logger"
	"AdAlpha/price_scrape"
	"database/sql"
	"errors"
	"fmt"
	"time"
)

//adds a logger instance for package to use
var esLog = logger.GetInstance()

//checks if valid currency code
func ValidCurrencyCode(cc string) bool {
	codes := []string{"GBP", "USD", "EUR", "CNY"}
	return Contains(codes, cc)
}

//checks if value in array
func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

//calculates the number of units based on exchange rate and given amount
func CalculateUnits(amount float64, exchangeRate float64, unitPrice float64) float64 {

	units := (amount * exchangeRate) / unitPrice
	return units
}

//gets a DB connection
func GetDbConnection() db.Db {
	d := db.Db{}
	d.Initialise()
	return d
}

// DB Update to add asset units to an investors portfolio
func AddAssetsInPortfolio(db *sql.DB, u float64, id int, isin string) error {

	_, err :=
		db.Exec("update portfolio SET units = units + $1 where investor_id=$2 AND isin=$3",
			u, id, isin)

	return err

}

// DB Update to remove asset units to an investors portfolio
func RemoveAssetsInPortfolio(db *sql.DB, u float64, id int, isin string) error {

	_, err :=
		db.Exec("update portfolio SET units = units - $1 where investor_id=$2 AND isin=$3",
			u, id, isin)

	if err != nil {
		if err != nil {
			esLog.LogError(logger.CreateLog("ERROR",
				fmt.Sprintf("Failed to remove assets from portfolio, investor_id: %d, isin: %s, units: %v",
					id, isin, u), err.Error(), logger.Trace(), time.Now()))
		}
		return err
	}
	return err
}

// Calls the scraper to get current price of given asset
func getCurrentPrice(isin string) (error, float64) {
	err, price := price_scrape.GetCurrentPrice(isin)
	if price == 0 {
		err = errors.New("Issue getting current asset price")
	}
	return err, price
}

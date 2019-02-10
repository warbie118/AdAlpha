package model

import (
	"AdAlpha/logger"
	"database/sql"
	"fmt"
	"time"
)

type Portfolio struct {
	Isin  string `json:"Isin"`
	Asset string `json:"Asset"`
	//CurrentPrice float64 `json:"Current_price"`
	Units float64 `json:"Units"`
}

//gets the investors portfolio and gets the current asset price
func GetInvestorPortfolio(db *sql.DB, id int) ([]Portfolio, error) {
	rows, err := db.Query(
		"SELECT p.isin, a.asset_name, p.units FROM portfolio p, assets a WHERE investor_id=$1 AND p.isin=a.isin", id)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR",
			fmt.Sprintf("Getting investor portfolio, investor_id: %d", id), err.Error(), logger.Trace(), time.Now()))
		return nil, err
	}

	defer rows.Close()

	var portfolio []Portfolio

	for rows.Next() {
		var p Portfolio
		if err := rows.Scan(&p.Isin, &p.Asset, &p.Units); err != nil {
			esLog.LogError(logger.CreateLog("ERROR",
				fmt.Sprintf("Getting investor instruction history error, investor_id: %d", id), err.Error(), logger.Trace(), time.Now()))
			return nil, err
		}
		portfolio = append(portfolio, p)
	}

	// Problem scraping - Too many requests?

	//for i, p := range portfolio {
	//	err, price := getCurrentPrice(p.Isin)
	//	if err != nil {
	//		esLog.LogError(logger.CreateLog("ERROR",
	//			fmt.Sprintf("Getting current price of asset, isin: %s", p.Isin), err.Error(), logger.Trace(), time.Now()))
	//	}
	//	portfolio[i].CurrentPrice = price
	//}

	return portfolio, nil
}

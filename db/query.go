package db

import (
	"AdAlpha/model"
	"database/sql"
	"log"
)

//Checks if investor_id exists in investors table
func InvestorIdExists(db *sql.DB, id int) (bool, error) {

	var exists bool
	err := db.QueryRow("select exists(select 1 from investors where investor_id=$1)", id).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		log.Fatalf("error checking if investor id exists '%d' %v", id, err)
	}
	return exists, err

}

//Get investor history and return as array of History model
func GetInvestorHistory(db *sql.DB, id int) ([]model.History, error) {
	rows, err := db.Query(
		"SELECT i.instruction, i.isin, a.asset_name, i.asset_price, i.units, i.currency_code, i.amount FROM instructions i, assets a WHERE investor_id=$1 AND i.isin=a.isin", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var history []model.History

	for rows.Next() {
		var h model.History
		if err := rows.Scan(&h.Instruction, &h.Isin, &h.Asset, &h.AssetPrice, &h.Units, &h.CurrencyCode, &h.Amount); err != nil {
			return nil, err
		}
		history = append(history, h)
	}

	return history, nil
}

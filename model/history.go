package model

import (
	"database/sql"
)

type History struct {
	Instruction  string  `json:"Instruction"`
	Isin         string  `json:"Isin"`
	Asset        string  `json:"Asset"`
	AssetPrice   float64 `json:"Asset_price"`
	Units        float64 `json:"Units"`
	CurrencyCode string  `json:"CurrencyCode"`
	Amount       float64 `json:"Amount"`
}

//Get investor history and return as array of History model
func GetInvestorHistory(db *sql.DB, id int) ([]History, error) {
	rows, err := db.Query(
		"SELECT i.instruction, i.isin, a.asset_name, i.asset_price, i.units, i.currency_code, i.amount FROM instructions i, assets a WHERE investor_id=$1 AND i.isin=a.isin", id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var history []History

	for rows.Next() {
		var h History
		if err := rows.Scan(&h.Instruction, &h.Isin, &h.Asset, &h.AssetPrice, &h.Units, &h.CurrencyCode, &h.Amount); err != nil {
			return nil, err
		}
		history = append(history, h)
	}

	return history, nil
}

//Adds a row to instructions table which holds investor history
func AddInvestorHistory(db *sql.DB, id int, instruction string, isin string, price float64, units float64, cc string, amount float64) error {

	var newId int
	err := db.QueryRow(
		"INSERT INTO instructions VALUES(nextval('instructions_instruction_id_seq'), $1, $2, $3, $4, $5, $6, $7) RETURNING instruction_id",
		id, isin, price, instruction, cc, amount, units).Scan(&newId)

	return err
}

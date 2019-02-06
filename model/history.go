package model

import "database/sql"

type History struct {
	Instruction  string          `json:"Instruction"`
	Isin         string          `json:"Isin"`
	Asset        string          `json:"Asset"`
	AssetPrice   sql.NullFloat64 `json:"Asset_price"`
	Units        sql.NullInt64   `json:"Units"`
	CurrencyCode sql.NullString  `json:"CurrencyCode"`
	Amount       sql.NullFloat64 `json:"Amount"`
}

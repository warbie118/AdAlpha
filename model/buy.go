package model

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
	err, price := getCurrentPrice(b.Isin)
	if err != nil {
		return err
	}

	err = CalculateBuy(price, *b)
	if err != nil {
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
		return err
	}

	err = AddInvestorHistory(dbCon.Pg, b.InvestorId, "BUY", b.Isin, assetPrice, b.Units, "GBP", assetPrice*b.Units)
	if err != nil {
		RemoveAssetsInPortfolio(dbCon.Pg, b.Units, b.InvestorId, b.Isin)
		return err
	}

	return err

}

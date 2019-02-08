package model

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
	err, price := getCurrentPrice(s.Isin)
	if err != nil {
		return err
	}

	err = CalculateSell(price, *s)
	if err != nil {
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
		return err
	}

	err = AddInvestorHistory(dbCon.Pg, s.InvestorId, "SELL", s.Isin, assetPrice, s.Units, "GBP", assetPrice*s.Units)
	if err != nil {
		AddAssetsInPortfolio(dbCon.Pg, s.Units, s.InvestorId, s.Isin)
		return err
	}

	return err
}

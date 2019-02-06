package model

import "fmt"

type Raise struct {
	InvestorId   int     `json:"investor_id"`
	Isin         string  `json:"isin"`
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
}

func (r *Raise) IsValid() bool {
	if r.InvestorId == 0 || len(r.Isin) == 0 || !ValidCurrencyCode(r.CurrencyCode) || r.Amount == 0 {
		return false
	}
	return true
}

func (*Raise) CalculateRaise() {
	fmt.Println("calculate raise")

}

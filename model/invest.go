package model

import "fmt"

type Invest struct {
	InvestorId   int     `json:"investor_id"`
	Isin         string  `json:"isin"`
	CurrencyCode string  `json:"currency_code"`
	Amount       float64 `json:"amount"`
}

func (i *Invest) IsValid() bool {
	if i.InvestorId == 0 || len(i.Isin) == 0 || !ValidCurrencyCode(i.CurrencyCode) || i.Amount == 0 {
		return false
	}
	return true
}

func (*Invest) CalculateInvest() {
	fmt.Println("calculate invest")

}

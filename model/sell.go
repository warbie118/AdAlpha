package model

import "fmt"

type Sell struct {
	InvestorId int    `json:"investor_id"`
	Isin       string `json:"isin"`
	Units      int    `json:"units"`
}

func (s *Sell) IsValid() bool {
	if s.InvestorId == 0 || len(s.Isin) == 0 || s.Units == 0 {
		return false
	}
	return true
}

func (*Sell) CalculateSell() {
	fmt.Println("calculate sell")

}

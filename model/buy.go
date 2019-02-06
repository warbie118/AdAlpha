package model

import "fmt"

type Buy struct {
	InvestorId int    `json:"investor_id"`
	Isin       string `json:"isin"`
	Units      int    `json:"units"`
}

func (b *Buy) IsValid() bool {
	if b.InvestorId == 0 || len(b.Isin) == 0 || b.Units == 0 {
		return false
	}
	return true
}

func (*Buy) CalculateBuy() {
	fmt.Println("calculate buy")

}

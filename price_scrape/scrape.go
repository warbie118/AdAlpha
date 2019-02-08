package price_scrape

import (
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
)

const (
	allowedDomains    = "markets.ft.com"
	url               = "https://markets.ft.com/data/funds/tearsheet/summary?s=%s"
	parent            = ".mod-tearsheet-overview__quote__bar"
	child             = ".mod-ui-data-list__value"
	assetPriceElement = ".mod-ui-data-list__value"
)

var collector = colly.NewCollector(colly.AllowedDomains(allowedDomains))

func GetCurrentPrice(shareCode string) (error, float64) {

	var p string
	var price float64

	// On every a element which has href attribute call callback
	collector.OnHTML(parent, func(e *colly.HTMLElement) {
		p = e.ChildText(child)
		e.ForEach(assetPriceElement, func(i int, element *colly.HTMLElement) {
			if i == 0 {
				p = element.Text
			}
		})

	})

	// Start scraping
	err := collector.Visit(fmt.Sprintf(url, shareCode))
	if err != nil {
		fmt.Printf("Error scraping price for asset: %s", shareCode)
		return err, price
	}

	price, err = strconv.ParseFloat(p, 64)
	if err != nil {
		fmt.Println("Error during string to float64 conversion")
		return err, price
	}

	return err, price
}

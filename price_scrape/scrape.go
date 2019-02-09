package price_scrape

import (
	"AdAlpha/logger"
	"fmt"
	"github.com/gocolly/colly"
	"strconv"
	"time"
)

const (
	allowedDomains    = "markets.ft.com"
	url               = "https://markets.ft.com/data/funds/tearsheet/summary?s=%s"
	parent            = ".mod-tearsheet-overview__quote__bar"
	child             = ".mod-ui-data-list__value"
	assetPriceElement = ".mod-ui-data-list__value"
)

var esLog = logger.GetInstance()
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
		esLog.LogError(logger.CreateLog("ERROR", fmt.Sprintf("Error scraping price for asset: %s", shareCode),
			err.Error(), logger.Trace(), time.Now()))
		return err, price
	}

	price, err = strconv.ParseFloat(p, 64)
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR", "Error during string to float64 conversion",
			err.Error(), logger.Trace(), time.Now()))
		return err, price
	}

	return err, price
}

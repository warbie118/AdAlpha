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
var lastUpdateTime = make(map[string]time.Time)
var prices = make(map[string]float64)

func GetCurrentPrice(shareCode string) (error, float64) {

	checkPricesMapInitialised()

	if time.Since(lastUpdateTime[shareCode]).Hours() > 1 {
		err, price := scrapePriceFromFT(shareCode)
		if err != nil {
			return err, price
		}
		prices[shareCode] = price
		lastUpdateTime[shareCode] = time.Now()

	}

	return nil, prices[shareCode]
}

func scrapePriceFromFT(shareCode string) (error, float64) {

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

	return nil, price
}

func checkPricesMapInitialised() {
	if len(lastUpdateTime) == 0 {
		isins := [...]string{"IE00B52L4369", "GB00BQ1YHQ70", "GB00B3X7QG63", "GB00BG0QP828", "GB00BPN5P238", "IE00B1S74Q32"}

		for _, isin := range isins {
			err, price := scrapePriceFromFT(isin)
			if err != nil {
				esLog.LogError(logger.CreateLog("ERROR",
					fmt.Sprintf("Getting current price of asset, isin: %s", isin), err.Error(), logger.Trace(), time.Now()))
			}
			prices[isin] = price
			lastUpdateTime[isin] = time.Now()
		}
	}
}

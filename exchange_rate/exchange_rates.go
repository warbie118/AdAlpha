package exchange_rate

import (
	"AdAlpha/logger"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

var esLog = logger.GetInstance()

const url = "https://api.exchangeratesapi.io/latest?base=%s&symbols=%s"

func GetExchangeRate(basecc string, exchangecc string) (error, float64) {

	resp, err := http.Get(GenerateUrl(basecc, exchangecc))
	if err == nil && resp.StatusCode != 200 {
		err = errors.New("GET request to exchangeratesapi.io failed")
	}
	if err != nil {
		esLog.LogError(logger.CreateLog("ERROR", "Request to get exchange rate failed",
			err.Error(), logger.Trace(), time.Now()))
		return err, 0
	}

	var result map[string]interface{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(&result); err != nil {
		fmt.Println(err.Error())
		esLog.LogError(logger.CreateLog("ERROR", "Issue decoding result into map[string]",
			err.Error(), logger.Trace(), time.Now()))
		return err, 0
	}

	rate := GetExchangeRateFromResponse(result, basecc)

	return err, rate
}

func GetExchangeRateFromResponse(resp map[string]interface{}, cc string) float64 {
	rate := resp["rates"].(map[string]interface{})
	ex := rate[cc].(float64)
	return ex
}

func GenerateUrl(basecc string, exchangecc string) string {
	return fmt.Sprintf(url, exchangecc, basecc)
}

package exchange_rate_test

import (
	"AdAlpha/exchange_rate"
	"fmt"
	"github.com/jarcoal/httpmock"
	"testing"
)

func TestGenerateUrl(t *testing.T) {
	url := exchange_rate.GenerateUrl("GBP", "USD")
	expectedUrl := "https://api.exchangeratesapi.io/latest?base=USD&symbols=GBP"

	if url != expectedUrl {
		t.Errorf("Expected url to be %v. got %v\n", expectedUrl, url)
	}

}

func TestGetExchangeRateFromResponse(t *testing.T) {
	var resp = map[string]interface{}{"base": "GBP", "date": "2019-02-7", "rates": map[string]interface{}{"GBP": 1.29}}
	rate := exchange_rate.GetExchangeRateFromResponse(resp, "GBP")
	if rate != 1.29 {
		fmt.Println("Error getting exchange rate from response")
	}
}

func TestGetExchangeRateIsSuccessful(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", exchange_rate.GenerateUrl("USD", "GBP"),
		httpmock.NewStringResponder(200, `{"base": "GBP", "date": "2019-02-07", "rates":{"USD":1.2}}`))

	err, res := exchange_rate.GetExchangeRate("USD", "GBP")

	if res != 1.2 || err != nil {
		t.Errorf("Expected response to be %v. got %v\n", 1.2, res)
	}
}

func TestGetExchangeRateJsonDecodeFailsAndReturnsError(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	httpmock.RegisterResponder("GET", exchange_rate.GenerateUrl("USD", "GBP"),
		httpmock.NewStringResponder(200, `{"base": "GBP", "date": 2019-02-07, "rates":{"USD":1.2}}`))

	err, res := exchange_rate.GetExchangeRate("USD", "GBP")

	if res != 0 && err == nil {
		t.Errorf("Expected an error")
	}
}

func TestGetExchangeRateApiGetCallFailsAndReturnsError(t *testing.T) {

	err, res := exchange_rate.GetExchangeRate("FAKE", "GBP")

	if res != 0 && err == nil {
		t.Errorf("Expected an error")
	}
}

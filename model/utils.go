package model

func ValidCurrencyCode(cc string) bool {
	codes := []string{"GBP", "USD", "EUR", "CNY"}
	return Contains(codes, cc)
}

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

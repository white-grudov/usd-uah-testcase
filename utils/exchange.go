package utils

import (
	"encoding/json"
	"net/http"
)

type ExchangeRateResponse struct {
	Rate float64 `json:"rate"`
}

func FetchExchangeRate() (float64, error) {
	resp, err := http.Get("https://api.exchangerate-api.com/v4/latest/USD")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}

	rates := result["rates"].(map[string]interface{})
	uahRate := rates["UAH"].(float64)
	return uahRate, nil
}

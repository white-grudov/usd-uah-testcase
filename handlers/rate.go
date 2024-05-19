package handlers

import (
	"encoding/json"
	"net/http"
	"usd-uah-testcase/utils"
)

func GetRate(w http.ResponseWriter, r *http.Request) {
	rate, err := utils.FetchExchangeRate()
	if err != nil {
		http.Error(w, "Error fetching exchange rate", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(rate)
}

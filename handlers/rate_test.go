package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"usd-uah-testcase/utils"
)

func TestGetRate(t *testing.T) {
	req, err := http.NewRequest("GET", "/rate", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(GetRate)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	rate, err := utils.FetchExchangeRate()
	if err != nil {
		t.Errorf("FetchExchangeRate returned an error: %v", err)
	}

	expected := fmt.Sprintf("%.2f", rate)
	result := strings.Trim(rr.Body.String(), "\n")

	if result != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", result, expected)
	}
}

package utils

import (
	"testing"
)

func TestFetchExchangeRate(t *testing.T) {
	// Call the FetchExchangeRate function
	rate, err := FetchExchangeRate()

	// Check if there was an error
	if err != nil {
		t.Errorf("FetchExchangeRate returned an error: %v", err)
	}

	// Check if the rate is greater than 0
	if rate <= 0 {
		t.Errorf("FetchExchangeRate returned an invalid rate: %f", rate)
	}
}

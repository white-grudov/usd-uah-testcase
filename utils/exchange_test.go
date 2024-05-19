package utils

import (
	"testing"
)

func TestFetchExchangeRate(t *testing.T) {
	rate, err := FetchExchangeRate()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if rate <= 0 {
		t.Fatalf("Expected positive rate, got %v", rate)
	}
}

package db

import (
	"os"
	"testing"
)

func TestInitDB(t *testing.T) {
	err := InitDB(os.Getenv("DATABASE_URL"))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

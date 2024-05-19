package utils

import (
	"testing"
)

type MockEmailProvider struct{}

func (m *MockEmailProvider) SendMail(to string, subject string, body string) error {
	return nil
}

func TestSendMail(t *testing.T) {
	emailProvider = &MockEmailProvider{}
	err := SendMail("test@example.com", "Test Subject", "Test Body")
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}
}

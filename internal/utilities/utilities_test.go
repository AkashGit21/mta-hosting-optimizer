package utilities

import (
	"os"
	"testing"
)

func TestIsEmptyString(t *testing.T) {
	tests := []struct {
		name        string
		value       string
		expectation bool
	}{
		{
			name:        "empty string",
			value:       "",
			expectation: true,
		},
		{
			name:        "not empty string",
			value:       "not empty",
			expectation: false,
		},
		{
			name:        "string containing only spaces",
			value:       "  ",
			expectation: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsEmptyString(tt.value)
			if got != tt.expectation {
				t.Errorf("got %v\n want %v", got, tt.expectation)
			}
		})
	}
}

func TestGetEnvValue(t *testing.T) {
	tests := []struct {
		name          string
		key           string
		defaultValue  string
		expectedValue string
	}{
		{
			name:          "get the env value of X",
			key:           "X",
			defaultValue:  "1",
			expectedValue: "1",
		},
		{
			name:          "get the env value of Y",
			key:           "Y",
			defaultValue:  "4",
			expectedValue: "2",
		},
	}
	os.Setenv("Y", "2")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := GetEnvValue(tt.key, "1")
			if got != tt.expectedValue {
				t.Errorf("got %v \nwant %v", got, tt.expectedValue)
			}
		})
	}
}

func TestWarnLog(t *testing.T) {
	WarnLog("this", "is a test logging", 1, "warn")
}

func TestErrorLog(t *testing.T) {
	ErrorLog("this", "is a test logging", 1, "error")
}

func TestDebugLog(t *testing.T) {
	DebugLog("this", "is a test logging", 1, "debug")
}

func TestInfoLog(t *testing.T) {
	InfoLog("this", "is a test logging", 1, "info")
}

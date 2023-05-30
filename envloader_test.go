package envloader_test

import (
	"os"
	"strings"
	"testing"

	"github.com/golangsugar/envloader"
)

func TestLoadFromFile(t *testing.T) {
	testCases := []struct {
		line,
		key,
		expected string
	}{
		// Invalid lines
		{line: "# Commented line", key: "Commented line", expected: ""}, // starts with #.
		{line: "#", key: "", expected: ""},                              // starts with #.
		{line: "", key: "", expected: ""},                               // Empty.
		{line: "\n", key: "", expected: ""},                             // Empty.
		{line: "INVALID_LINE", key: "INVALID_LINE", expected: ""},       // Invalid syntax, without =.
		{line: "_Invalid", key: "_Invalid", expected: ""},               // has to start with a letter.
		{line: "#XYZ", key: "XYZ", expected: ""},                        // commented.
		{line: "_LETTERS", key: "_LETTERS", expected: ""},               // has to start with a letter.
		{line: "X", key: "X", expected: ""},                             // key should contain 2 or more chars.
		// Valid lines
		{line: "KEY1=value1", key: "KEY1", expected: "value1"},
		{line: "KEY2=value 2", key: "KEY2", expected: "value 2"},
		{line: "KEY3=", key: "KEY3", expected: ""},
		{line: `ABC="42378462%&&3 178964@"`, key: "ABC", expected: `"42378462%&&3 178964@"`},
		{line: "mnoPQR=42378462%&&3 ###", key: "mnoPQR", expected: "42378462%&&3 ###"},
	}

	var sb strings.Builder
	{
		for _, tc := range testCases {
			sb.WriteString(tc.line)
			sb.WriteString("\n")
		}

		// Create a temporary file with test data
		tmpfile, err := os.CreateTemp("", "envloader_test")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(tmpfile.Name())

		err = os.WriteFile(tmpfile.Name(), []byte(sb.String()), 0644)
		if err != nil {
			t.Fatalf("Failed to write test data to temporary file: %v", err)
		}

		// Load environment variables from the temporary file
		err = envloader.LoadFromFile(tmpfile.Name(), false)
		if err != nil {
			t.Fatalf("Failed to load environment variables: %v", err)
		}
	}

	for _, tc := range testCases {
		if tc.key == "" {
			continue
		}

		actualValue := os.Getenv(tc.key)
		if actualValue != tc.expected {
			t.Errorf("Environment variable %s has incorrect value. Expected: %s, Got: %s", tc.key, tc.expected, actualValue)
		}
	}
}

func TestLoadFromFile_NonexistentFile(t *testing.T) {
	err := envloader.LoadFromFile("nonexistent.txt", true)
	if err == nil {
		t.Error("Expected an error when loading from a nonexistent file, but got nil")
	}
}

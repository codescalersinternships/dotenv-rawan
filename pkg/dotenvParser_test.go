package dotenvParser

import (
	"reflect"
	"testing"
)

func TestParseFile(t *testing.T) {
	tests := []struct {
		name     string
		filePath string
		expected map[string]string
		err      error
	}{
		{
			name:     "Valid .env file",
			filePath: "testdata/.env",
			expected: map[string]string{
				"APP_NAME":       "Linktree",
				"APP_ENV":        "development",
				"APP_PORT":       "8080",
				"DB_HOST":        "localhost",
				"DB_PORT":        "5432",
				"DB_USER":        "admin",
				"DB_PASSWORD":    "secretpassword",
				"DB_NAME":        "linktree_db",
				"JWT_SECRET":     "my_super_secret_key",
				"JWT_EXPIRATION": "3600",
				"API_KEY":        "123456789abcdef",
				"API_URL":        "https://api.example.com",
				"DEBUG":          "true",
				"CACHE_ENABLED":  "false",
			},
			err: nil,
		},
		{
			name:     "Unsupported file type",
			filePath: "testdata/env.txt",
			expected: nil,
			err:      ErrUnsupporteFileType,
		},
		{
			name:     "Non-existent file",
			filePath: "testdata/nonexistent.env",
			expected: nil,
			err:      ErrFileNotExist,
		},
		{
			name:     "Empty .env file",
			filePath: "testdata/empty.env",
			expected: nil,
			err:      ErrEmptyFile,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			got, err := Parse(testcase.filePath)

			if err != testcase.err {
				t.Errorf("expected error: %v, got: %v", testcase.err, err)
			}

			if !reflect.DeepEqual(got, testcase.expected) {
				t.Fatalf("mismatch:\nexpected: %#v\ngot: %#v", testcase.expected, got)
			}
		})
	}
}

func TestParseSyntaxErr(t *testing.T) {
	tests := []struct {
		name       string
		fileString string
		expected   map[string]string
		err        error
	}{
		{
			name:       "Missing Value",
			fileString: "DB_HOST=",
			expected:   nil,
			err:        ErrMissingValue,
		},
		{
			name:       "Missing Key",
			fileString: "=12345",
			expected:   nil,
			err:        ErrMissingKey,
		},
		{
			name:       "Space in Key",
			fileString: "DB USER=admin",
			expected:   nil,
			err:        ErrKeyName,
		},
		{
			name:       "No equal sign",
			fileString: "DB_USER",
			expected:   nil,
			err:        ErrNoEqual,
		},
	}

	for _, testcase := range tests {
		t.Run(testcase.name, func(t *testing.T) {
			got, err := parseString(testcase.fileString)

			if err != testcase.err {
				t.Errorf("expected error: %v, got: %v", testcase.err, err)
			}

			if !reflect.DeepEqual(got, testcase.expected) {
				t.Fatalf("mismatch:\nexpected: %#v\ngot: %#v", testcase.expected, got)
			}
		})
	}
}

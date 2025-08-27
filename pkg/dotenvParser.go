package dotenvParser

import (
	"errors"
	"os"
	"strings"
)

var ErrUnsupporteFileType = errors.New("file type is not supported")
var ErrFileNotExist = errors.New("file doesn't exist")
var ErrEmptyFile = errors.New("file is empty")

var ErrNoEqual = errors.New("syntax error: line must have an equal sign")
var ErrMissingKey = errors.New("syntax error: missing key")
var ErrMissingValue = errors.New("syntax error: missing Value")
var ErrKeyName = errors.New("syntax error: invalid key name")

func loadEnv(path string) (string, error) {
	if !strings.HasSuffix(path, ".env") {
		return "", ErrUnsupporteFileType
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return "", ErrFileNotExist
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func Parse(path string) (map[string]string, error) {
	data, err := loadEnv(path)

	if err != nil {
		return nil, err
	}

	got, err := parseString(data)

	return got, err
}

func parseString(data string) (map[string]string, error) {

	if len(data) == 0 {
		return nil, ErrEmptyFile
	}

	envVars := make(map[string]string)

	lines := strings.Split(data, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)

		if len(parts) < 2 {
			return nil, ErrNoEqual
		}

		if cut, ok := strings.CutPrefix(parts[0], "export"); ok {
			parts[0] = cut
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		if key == "" {
			return nil, ErrMissingKey
		}

		if value == "" {
			return nil, ErrMissingValue
		}

		if strings.Contains(key, " ") {
			return nil, ErrKeyName
		}

		envVars[key] = value
	}

	return envVars, nil
}

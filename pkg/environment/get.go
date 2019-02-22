package environment

import (
	"fmt"
	"os"
	"strings"
)

func Get(key string, defaultValue ...string) (string, error) {
	if len(strings.TrimSpace(key)) == 0 {
		return "", fmt.Errorf("empty key")
	}

	value := os.Getenv(key)

	if len(strings.TrimSpace(value)) == 0 {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return "", fmt.Errorf("environment variable '%s' not found", key)
	}

	return value, nil
}

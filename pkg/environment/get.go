package environment

import (
	"fmt"
	"os"
	"strings"
)

func Get(key string, defaultValue ...string) (string, error) {
	if strings.TrimSpace(key) == "" {
		return "", fmt.Errorf("empty key")
	}

	value := os.Getenv(key)

	if strings.TrimSpace(value) == "" {
		if len(defaultValue) > 0 {
			return defaultValue[0], nil
		}
		return "", fmt.Errorf("environment variable '%s' not found", key)
	}

	return value, nil
}

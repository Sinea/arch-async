package environment

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func Get(key string, defaultValue ...string) string {
	value := os.Getenv(key)

	if len(strings.TrimSpace(value)) == 0 {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		panic(errors.New(fmt.Sprintf("environment variable '%s' not found", key)))
	}

	return value
}

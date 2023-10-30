package config

import (
	"os"
)

func GetString(text string) string {
	return os.Getenv(text)
}
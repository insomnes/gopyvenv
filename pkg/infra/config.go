package infra

import (
	"os"
	"strings"
)

var (
	debugOn = os.Getenv("GOPYVENV_DEBUG") == "1"
	Enabled = os.Getenv("GOPYVENV_ENABLED") == "1"
)

var (
	defaultVenvDirs  = []string{"venv", ".venv"}
	defaultSeparator = ","
)

func getVenvDirs() []string {
	fromEnv := os.Getenv("GOPYVENV_DIR_NAMES")
	if fromEnv == "" {
		return defaultVenvDirs
	}

	separator := os.Getenv("GOPYVENV_DIR_SEPARATOR")
	if separator == "" {
		separator = defaultSeparator
	}
	return strings.Split(fromEnv, separator)
}

var VenvDirs = getVenvDirs()

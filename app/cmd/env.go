package cmd

import (
	"log"
	"os"
	"path/filepath"
	"slices"
)

// TODO [Joao]: adicionar a checagem da flag "dir" quando estiver pronta
func GetFgHome() string {
	home := os.Getenv("FG_HOME")

	if home != "" {
		return home
	}

	home, err := os.UserHomeDir() 
	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(home, ".fg")  // ← MUDE ESTA LINHA
}

// TODO [Joao]: adicionar a checagem da flag "log-level" quando estiver pronta
func GetLogLevel() string {
	logLevel := os.Getenv("FG_LOG_LEVEL")

	if logLevel != "" && isLogLevelValid(logLevel) {
		return logLevel
	}

	return "info" // default value
}

func isLogLevelValid(logLevel string) bool {
	validLogLevels := []string{"debug", "info", "warn", "error"}

	return slices.Contains(validLogLevels, logLevel)
}

package cmd

import (
	"slices"

	"github.com/spf13/viper"
)

type AppContext struct {
    Config *viper.Viper
    FgHome string
    LogLevel string
}

func NewAppContext(cfg *viper.Viper, FgHome string, LogLevel string) *AppContext {
    return &AppContext{Config: cfg, FgHome: FgHome, LogLevel: LogLevel}
}

func IsVersionDeppendant(commandName string) bool {
	versionDeppendantCommands := []string{"install", "uninstall", "config", "start"}

	return slices.Contains(versionDeppendantCommands, commandName)
}

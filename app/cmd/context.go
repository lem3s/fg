package cmd

import "github.com/spf13/viper"

type AppContext struct {
    Config *viper.Viper
}

func NewAppContext(cfg *viper.Viper) *AppContext {
    return &AppContext{Config: cfg}
}

package app

import (
    "github.com/spf13/viper"
    "sync"
)

var once sync.Once
var cfg *viper.Viper

func GetConfig() *viper.Viper {
    once.Do(func() {
        v := viper.New()
        v.SetConfigName("config")
        v.SetConfigType("yaml")
        v.AddConfigPath(".")
        v.AutomaticEnv()

        err := v.ReadInConfig()
        if err != nil {
            panic("Erro ao ler config: " + err.Error())
        }

        cfg = v
    })
    return cfg
}

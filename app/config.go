package app

import (
	"sync"

	"github.com/spf13/viper"
)

var once sync.Once
var cfg *viper.Viper

func GetConfig() *viper.Viper {
	once.Do(func() {
		v := viper.New()
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(".") //aqui trabalharemos com o dir do FG, sendo o default home/.fg
		v.AutomaticEnv()

		err := v.ReadInConfig()
		if err != nil {
			panic("Erro ao ler config: " + err.Error())
		}

		cfg = v
	})
	return cfg
}

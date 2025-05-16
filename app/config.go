package app

import (
	"errors"
	"log"
	"sync"

	"github.com/lem3s/fg/app/cmd"
	"github.com/spf13/viper"
)

var once sync.Once
var cfg *viper.Viper

func GetConfig() *viper.Viper {
	once.Do(func() {
		v := viper.New()
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(cmd.GetFgHome()) //aqui trabalharemos com o dir do FG, sendo o default home/.fg
		v.AutomaticEnv()

		err := v.ReadInConfig()
		if err != nil {
			var configFileNotFoundError viper.ConfigFileNotFoundError
			if errors.As(err, &configFileNotFoundError) {
				return
			}
			log.Fatalf("Error reading config: %s\n", err)
			return
		}

		cfg = v
	})

	return cfg
}

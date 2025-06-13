package app

import (
	"errors"
	"log"
	"sync"

	"github.com/lem3s/fg/app/cmd"
	"github.com/spf13/viper"
)

type ConfigCommand struct {
	Ctx *cmd.AppContext
}

var once sync.Once
var cfg *viper.Viper

func GetConfig() *viper.Viper {
	once.Do(func() {
		v := viper.New()
		v.SetConfigName("config")
		v.SetConfigType("yaml")
		v.AddConfigPath(cmd.GetFgHome())
		v.AutomaticEnv()

		err := v.ReadInConfig()
		if err != nil {
			var configFileNotFoundError viper.ConfigFileNotFoundError
			if errors.As(err, &configFileNotFoundError) {
				log.Println("Arquivo de configuração não encontrado.")
				cfg = v
				return
			}
			log.Fatalf("Erro ao ler arquivo de configuração: %s\n", err)
			return
		}

		cfg = v
	})

	DisplayConfig()
	return cfg
}

func DisplayConfig() {
	config := GetConfig()
	if config == nil {
		log.Println("Nenhuma configuração informada.")
		return
	}

	log.Println("Configurações do FG:")

	settings := config.AllSettings()

	for key, value := range settings {
		if value == nil {
			log.Printf("%s: <null>\n", key)
		} else {
			log.Printf("%s: %v\n", key, value)
		}
	}
}

func (h *ConfigCommand) Run(args []string) error {
	message := "Hello " + h.Ctx.Config.GetString("jar") + "!"
	h.Ctx.Interactor.Info(message)
	return nil
}

func init() {
	cmd.Register("config", func(ctx *cmd.AppContext) cmd.Command {
		return &ConfigCommand{Ctx: ctx}
	})
}

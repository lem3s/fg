package main

import (
	"os"

	"github.com/lem3s/fg/app"
	"github.com/lem3s/fg/app/cmd"
	_ "github.com/lem3s/fg/app/commands" // Importa todos os comandos para registro
	"github.com/lem3s/fg/app/handlers"
	"github.com/spf13/viper"
)

func main() {
	commandName := os.Args[1]
	args := os.Args[2:]

	var cfg *viper.Viper
	//carrega as configurações do arquivo config.yaml
	if cmd.IsVersionDeppendant(commandName) {
		cfg = app.GetConfig()
	}

	//cria o contexto de configuração da aplicação
	ctx := cmd.NewAppContext(cfg, cmd.GetFgHome(), cmd.GetLogLevel())
	handlers.HandleParams(args, ctx)

	command, _ := cmd.CreateCommand(commandName, ctx)

	err := command.Run(args)
	handlers.HandleCallback(err, ctx)
}

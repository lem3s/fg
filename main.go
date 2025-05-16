package main

import (
	"fmt"
	"os"

	"github.com/lem3s/fg/app"
	"github.com/lem3s/fg/app/cmd"
	_ "github.com/lem3s/fg/app/commands" // Importa todos os comandos para registro
	"github.com/spf13/viper"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: fg <comando> [args...]")
		os.Exit(1)
	}

	commandName := os.Args[1]
	args := os.Args[2:]

	var cfg *viper.Viper
	//carrega as configurações do arquivo config.yaml
	if (cmd.IsVersionDeppendant(commandName)) {
		cfg = app.GetConfig()
	}

	//cria o contexto de configuração da aplicação
	ctx := cmd.NewAppContext(cfg, cmd.GetFgHome(), cmd.GetLogLevel())

	//cria o comando a partir do nome e do contexto
	//se o comando não existir, retorna erro
	command, err := cmd.CreateCommand(commandName, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	//executa o comando com os argumentos
	command.Run(args)
}

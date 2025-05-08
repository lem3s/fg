package main

import (
	"fmt"
	"os"

	"github.com/lem3s/fg/app"
	"github.com/lem3s/fg/app/cmd"
	_ "github.com/lem3s/fg/app/commands" // Importa todos os comandos para registro
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Uso: fg <comando> [args...]")
		os.Exit(1)
	}

	commandName := os.Args[1]
	args := os.Args[2:]

	cfg := app.GetConfig()
	ctx := cmd.NewAppContext(cfg)

	command, err := cmd.CreateCommand(commandName, ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	command.Run(args)
}

// common/services/list.go
package commands

import (
	"fmt"
	//"os"
	"github.com/lem3s/fg/app/cmd"
)

type ListCmd struct {
    Ctx *cmd.AppContext
}

func(h *ListCmd) Run(args []string) error {
	//validacoes especificas para o comando
    path := h.Ctx.Config.GetString("jar")
    if path == "" && len(args) > 0 {
        path = args[0]
    }
    if path == "" {
        path = "teste"
    }

	//ListVersionsFromFile(nome)
    fmt.Printf("%s!\n", path)
	return nil
}

func init() {
	cmd.Register("version", func(ctx *cmd.AppContext) cmd.Command {
        return &ListCmd{Ctx: ctx}
    })
}

func ListVersionsFromFile(filePath string) error {
	// _, err := os.Stat(filePath)
	// if err != nil {
	// 	return fmt.Errorf("arquivo não encontrado: %s", filePath)
	// }
	
	// data, err := os.ReadFile(filePath)
	// if err != nil {
	// 	return fmt.Errorf("erro ao ler o arquivo: %v", err)
	// }
	
	// displayVersionInfo(versionSystem)
	return nil
}
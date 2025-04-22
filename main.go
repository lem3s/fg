// main.go
package main

import (
	"fmt"
	"os"

	"github.com/lem3s/fg/common/services"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "fg",
	Short: "Gerenciador de versões para aplicações",
	Long: `Ferramenta para gerenciar, instalar e listar 
versões da aplicação. Permite visualizar o histórico de 
versões instaladas e verificar se há atualizações disponíveis.`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Gerencia versões do aplicativo",
	Long: `O comando 'version' permite gerenciar as versões 
da aplicação, incluindo listar, instalar ou desinstalar versões.`,
}

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.AddCommand(services.ListCmd)
}

func main() {
	fmt.Println("Gerenciador de Versões - Iniciando...")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

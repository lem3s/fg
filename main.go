// main.go
package main

import (
	"fmt"
	"os"

	"github.com/lem3s/fg/common"
	"github.com/lem3s/fg/common/services"

	"embed"

	"github.com/spf13/cobra"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
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

var assets embed.FS

func init() {
	rootCmd.AddCommand(versionCmd)

	versionCmd.AddCommand(services.ListCmd)

	rootCmd.AddCommand(common.ConfigCmd)
}

func main() {
	fmt.Println("Gerenciador de Versões - Iniciando...")

	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "myproject",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

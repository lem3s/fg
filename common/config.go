package common

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gopkg.in/yaml.v3"
)

type Configuration struct {
	Host      string `yaml:"Host"`
	Port      int    `yaml:"Port"`
	LogLevel  string `yaml:"LogLevel"`
	Debug     bool   `yaml:"Debug"`
	DataDir   string `yaml:"DataDir"`
	MaxMemory int    `yaml:"MaxMemory"`
	MaxCPU    int    `yaml:"MaxCPU"`
	Workers   int    `yaml:"Workers"`
}

var (
	Config Configuration
)

func InitConfig(cfgFile string) error {
	viper.SetConfigFile(cfgFile)
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.Unmarshal(&Config)
}

func displayConfigInfo(data Configuration) {

	fmt.Printf(`Configuração atual da máquina:
			Servidor:
			- Host: %s
			- Porta: %d
			- LogLevel: %s
			- Debug: %t
			- DataDir: %s
			
			Recursos: 
			- Memória máxima: %dMB
			- CPU máxima: %d
			- Workers: %d`,

		data.Host, data.Port, data.LogLevel, data.Debug, data.DataDir, data.MaxMemory, data.MaxCPU, data.Workers)
}

var ConfigCmd = &cobra.Command{
	Use:   "config",
	Short: "Exibe as configurações disponibilizadas pela máquina para a aplicação",
	Long: `O comando 'config' exibe as configurações de servidor e recursos oferecidos 
			pela máquina para a execução da aplicação`,

	Run: func(cmd *cobra.Command, args []string) {
		data, err := os.ReadFile("config.yaml")
		if err != nil {
			fmt.Println("Erro ao ler o arquivo YAML:", err)
		} else {
			var config Configuration
			err = yaml.Unmarshal(data, &config)
			if err != nil {
				fmt.Println("Erro ao parsear o YAML:", err)
				return
			}

			displayConfigInfo(config)
		}
	},
}

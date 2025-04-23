package services

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
)
func GetFgHome() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "/tmp/fg"
	}
	return filepath.Join(homeDir, ".fg")
}

func UninstallVersion(version string) error {
	fgHome := GetFgHome()
	
	version = strings.TrimSpace(version)

	versionPath := filepath.Join(fgHome, version)
	
	_, err := os.Stat(versionPath)
	if os.IsNotExist(err) {
		entries, err := os.ReadDir(fgHome)
		if err != nil {
			return fmt.Errorf("erro ao ler diretório principal: %v", err)
		}
		
		found := false
		for _, entry := range entries {
			if entry.IsDir() && strings.Contains(entry.Name(), version) {
				versionPath = filepath.Join(fgHome, entry.Name())
				version = entry.Name()
				found = true
				break
			}
		}
		
		if !found {
			return fmt.Errorf("versão '%s' não encontrada", version)
		}
	} else if err != nil {
		return fmt.Errorf("erro ao verificar versão: %v", err)
	}
	
	fmt.Printf("Desinstalando versão: %s\n", version)
	fmt.Printf("Localização: %s\n", versionPath)

	jarFiles := []string{}
	err = filepath.Walk(versionPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".jar") {
			jarFiles = append(jarFiles, path)
		}
		return nil
	})
	
	if err != nil {
		return fmt.Errorf("erro ao procurar arquivos: %v", err)
	}
	
	if len(jarFiles) > 0 {
		fmt.Println("Arquivos JAR encontrados:")
		for _, jar := range jarFiles {
			fmt.Printf("  - %s\n", filepath.Base(jar))
		}
	}

	fmt.Printf("Tem certeza que deseja desinstalar a versão %s? (s/N): ", version)
	var response string
	fmt.Scanln(&response)
	response = strings.ToLower(strings.TrimSpace(response))
	
	if response != "s" && response != "sim" && response != "y" && response != "yes" {
		fmt.Println("Desinstalação cancelada.")
		return nil
	}
	
	err = os.RemoveAll(versionPath)
	if err != nil {
		return fmt.Errorf("erro ao remover diretório: %v", err)
	}
	
	fmt.Printf("Versão '%s' desinstalada com sucesso!\n", version)
	return nil
}
func ListInstalledVersions() ([]string, error) {
	fgHome := GetFgHome()
	
	_, err := os.Stat(fgHome)
	if os.IsNotExist(err) {
		return []string{}, nil
	} else if err != nil {
		return nil, fmt.Errorf("erro ao verificar diretório principal: %v", err)
	}
	
	entries, err := os.ReadDir(fgHome)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler diretório principal: %v", err)
	}
	
	versions := []string{}
	for _, entry := range entries {
		if entry.IsDir() {

			name := entry.Name()
			if strings.Contains(name, ".") && !strings.HasPrefix(name, ".") {
				versions = append(versions, name)
			}
		}
	}
	
	return versions, nil
}
var UninstallCmd = &cobra.Command{
	Use:   "uninstall [versão]",
	Short: "Desinstala uma versão",
	Long: `Desinstala uma versão específica da aplicação.
Isso remove o diretório completo da versão, incluindo todos os arquivos JAR e dados.

Exemplo:
  fg uninstall 1.0.0
  fg uninstall 2.1.0`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]

		fgHome := GetFgHome()
		if _, err := os.Stat(fgHome); os.IsNotExist(err) {
			fmt.Println("Nenhuma versão instalada. Diretório principal não existe.")
			return
		}

		force, _ := cmd.Flags().GetBool("force")
		if force {
			fmt.Println("Modo força ativado: não será solicitada confirmação")
		}
		
		err := UninstallVersion(version)
		if err != nil {
			fmt.Println("Erro durante a desinstalação:", err)
			return
		}
	},
}

var uninstallListCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista as versões instaladas",
	Long: `Lista todas as versões atualmente instaladas no sistema.

Exemplo:
  fg uninstall list`,
	Run: func(cmd *cobra.Command, args []string) {
		versions, err := ListInstalledVersions()
		if err != nil {
			fmt.Println("Erro ao listar versões:", err)
			return
		}
		
		if len(versions) == 0 {
			fmt.Println("Nenhuma versão instalada.")
			return
		}
		
		fmt.Println("Versões instaladas:")
		for i, version := range versions {
			fmt.Printf("%d. %s\n", i+1, version)

			versionPath := filepath.Join(GetFgHome(), version)
			jarCount := 0
			
			err := filepath.Walk(versionPath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return nil 
				}
				if !info.IsDir() && strings.HasSuffix(strings.ToLower(info.Name()), ".jar") {
					jarCount++
				}
				return nil
			})
			
			if err == nil && jarCount > 0 {
				fmt.Printf("   Contém %d arquivo(s) JAR\n", jarCount)
			}
		}
	},
}

func init() {
	UninstallCmd.Flags().BoolP("force", "f", false, "Não solicitar confirmação para desinstalar")
	
	UninstallCmd.AddCommand(uninstallListCmd)
}


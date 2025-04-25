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
		if tempDir, err := os.MkdirTemp("", "fg"); err == nil {
			return tempDir
		}
		return filepath.Join(os.TempDir(), "fg")
	}
	return filepath.Join(homeDir, ".fg")
}

func ValidateVersionPath(version string) (string, error) {
	fgHome := GetFgHome()

	if _, err := os.Stat(fgHome); os.IsNotExist(err) {
		return "", fmt.Errorf("nenhuma versão instalada. Diretório principal não existe")
	}
	
	version = strings.TrimSpace(version)
	versionPath := filepath.Join(fgHome, version)
	
	_, err := os.Stat(versionPath)
	if os.IsNotExist(err) {
		entries, err := os.ReadDir(fgHome)
		if err != nil {
			return "", fmt.Errorf("erro ao ler diretório principal: %v", err)
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
			return "", fmt.Errorf("versão '%s' não encontrada", version)
		}
	} else if err != nil {
		return "", fmt.Errorf("erro ao verificar versão: %v", err)
	}
	
	return versionPath, nil
}

func UninstallVersion(version string) error {
	versionPath, err := ValidateVersionPath(version)
	if err != nil {
		return err
	}
	
	fmt.Printf("Desinstalando versão: %s\n", filepath.Base(versionPath))
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

	fmt.Printf("Tem certeza que deseja desinstalar a versão %s? (s/N): ", filepath.Base(versionPath))
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
	
	fmt.Printf("Versão '%s' desinstalada com sucesso!\n", filepath.Base(versionPath))
	return nil
}

var UninstallCmd = &cobra.Command{
	Use:   "uninstall [versão]",
	Short: "Desinstala uma versão",
	Long: `Desinstala uma versão específica da aplicação.`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		version := args[0]
		
		err := UninstallVersion(version)
		if err != nil {
			fmt.Println("Erro durante a desinstalação:", err)
			return
		}
	},
}
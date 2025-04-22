package services

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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


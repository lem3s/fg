package install

import (
	"io"
	"net/http"
	"os"

	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	Version      string `yaml:"version"`
	Applications map[string]struct {
		JDKVersion int    `yaml:"jdk_version"`
		Source     string `yaml:"source"`
		Command    string `yaml:"command"`
	} `yaml:"applications"`
}

func ParseYAMLFromURL(url string) (*AppConfig, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	yamlData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var config AppConfig
	err = yaml.Unmarshal(yamlData, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

func downloadFile(appConfig *AppConfig) error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err 
	}

	var folderPath = homeDir + "/.fg/" + appConfig.Version
	err = os.MkdirAll(folderPath, os.ModePerm)
	if err != nil {
		return err
	}

	out, err := os.Create(folderPath + "/" + "kafka" + ".targz")
	if err != nil {
		return err
	}
	defer out.Close()

	resp, err := http.Get(appConfig.Applications["kafka"].Source)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
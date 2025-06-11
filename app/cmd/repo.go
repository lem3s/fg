package cmd

import (
	"io"
	"net/http"

	"gopkg.in/yaml.v3"
)

type VersionsData struct {
	Versions []Version `yaml:"versions"`
}

type Version struct {
	Version string `yaml:"version"`
	Date    string `yaml:"date"`
	JDK     JDK    `yaml:"jdk"`
	Apps    []App  `yaml:"apps"`
}

type JDK struct {
	Name                         string            `yaml:"name"`
	PlatformsCompressedSourceUrl map[string]string `yaml:"platforms"` // keys: win, mac, linux
}

type App struct {
	Name            string   `yaml:"name"`
	Description     string   `yaml:"description"`
	Version         string   `yaml:"version"`
	DependenciesUrl []string `yaml:"dependencies"`
	FilesUrl        []string `yaml:"files"`
	Command         string   `yaml:"command"`
}

func GetVersionsData() ([]Version, error) {
	url := "https://raw.githubusercontent.com/lem3s/fg-example-app/refs/heads/main/setup.yaml"

	resp, err := http.Get(url)
	if err != nil {
		return []Version{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []Version{}, err
	}

	var data VersionsData
	err = yaml.Unmarshal(body, &data)
	if err != nil {
		return []Version{}, err
	}

	return data.Versions, nil
}

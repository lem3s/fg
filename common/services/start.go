package services

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

type CacheConfig struct {
	Pid     int    `yaml:"pid"`
	Jarfile string `yaml:"jarfile"`
	Status  string `yaml:"status"`
	Port    int    `yaml:"port"`
	Host    string `yaml:"host"`
	Uptime  int    `yaml:"uptime"`
}

var StartCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the FHIR Guard application",
	Long:  `Example: fg start <jarfile>`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Please provide the JAR file to start.")
			return
		}
		jarfile := args[0]

		start(jarfile)
	},
}

func start(jarfile string) {
	var cacheConfig map[string]interface{}

	if _, err := os.Stat(jarfile); os.IsNotExist(err) {
		log.Fatalf("JAR file does not exist: %s", jarfile)
	}

	absPath, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current working directory:", err)
	}

	jarfile = absPath + "/" + jarfile
	log.Printf("Absolute path of JAR file: %s", jarfile)
	port := findAvailablePort(8080)
	portStr := fmt.Sprintf("--server.port=%d", port)

	cmd := exec.Command("java", "-jar", jarfile, portStr)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Fatal("Error starting JAR file:", err)
	}

	pid := cmd.Process.Pid
	log.Printf("Started Java process with PID: %d", pid)

	//processCacheData(jarfile, pid)
	yamlFile, err := os.ReadFile("cache.yaml")
	if err != nil {
		log.Fatal("Error reading cache.yml:", err)
	}

	yaml.Unmarshal(yamlFile, &cacheConfig)

	yamlData, err := yaml.Marshal(cacheConfig)

	if err != nil {
		log.Fatal("Error marshalling YAML:", err)
	}

	err = os.WriteFile("cache.yaml", yamlData, 0644)
	if err != nil {
		log.Fatal("Error writing to config.yaml:", err)
	}
}

func findAvailablePort(startPort int) int {
	for port := startPort; port < 65535; port++ {
		addr := fmt.Sprintf(":%d", port)
		listener, err := net.Listen("tcp", addr)
		if err != nil {
			continue
		}
		listener.Close()
		return port
	}
	log.Fatal("No available ports found")
	return -1
}

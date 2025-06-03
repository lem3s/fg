package commands

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"

	"github.com/lem3s/fg/app/cmd"
	"gopkg.in/yaml.v3"
)

type StartCmd struct {
	Ctx *cmd.AppContext
}

func (s *StartCmd) Run(args []string) error {
	if len(args) < 1 {
		return fmt.Errorf("missing JAR file argument")
	}
	jarfile := args[0]

	start(jarfile)
	return nil
}

func init() {
	cmd.Register("start", func(ctx *cmd.AppContext) cmd.Command {
		return &StartCmd{Ctx: ctx}
	})
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

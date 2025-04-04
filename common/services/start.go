package services

import (
	"log"
	"os"
	"os/exec"
)

func init() {
}

func Start(jarfile string) {
	if _, err := os.Stat(jarfile); os.IsNotExist(err) {
		log.Fatalf("JAR file does not exist: %s", jarfile)
	}

	absPath, err := os.Getwd()
	if err != nil {
		log.Fatal("Error getting current working directory:", err)
	}

	jarfile = absPath + "/" + jarfile
	log.Printf("Absolute path of JAR file: %s", jarfile)

	cmd := exec.Command("java", "-jar", jarfile)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Fatal("Error starting JAR file:", err)
	}
	pid := cmd.Process.Pid
	log.Printf("Started Java process with PID: %d", pid)

	if err != nil {
		log.Fatal("Error running JAR file:", err)
	}
}

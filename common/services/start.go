package services

import (
	"log"
	"module/common/watcher"
	"os"
	"os/exec"
)

func init() {
	// This is a placeholder for package initialization logic.
	// You can add any necessary setup code here.
	// For example, you might want to initialize logging or configuration settings.
	// fmt.Println("Services package initialized")
}

func Start() {
	cmd := exec.Command("java", "-jar", "./myproject-0.0.1-SNAPSHOT.jar")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Start(); err != nil {
		log.Fatal("Error starting JAR file:", err)
	}
	pid := cmd.Process.Pid
	log.Printf("Started Java process with PID: %d", pid)
	w := watcher.NewWatcher(pid)
	if err := w.Watch(); err != nil {
		log.Fatal("Error watching process:", err)
	}

	err := cmd.Wait()

	if err != nil {
		log.Fatal("Error running JAR file:", err)
	}
}

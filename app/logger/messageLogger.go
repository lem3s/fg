package logger

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type DefaultInteractionHandler struct{}

func (d *DefaultInteractionHandler) Prompt(message string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(message + ": ")
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func (d *DefaultInteractionHandler) Confirm(message string) bool {
	response := d.Prompt(message + " [y/N]")
	return strings.ToLower(response) == "y"
}

func (d *DefaultInteractionHandler) Info(message string) {
	fmt.Println("[INFO]", message)
}

func (d *DefaultInteractionHandler) Warn(message string) {
	fmt.Println("[WARN]", message)
}

func (d *DefaultInteractionHandler) Error(message string) {
	fmt.Println("[ERROR]", message)
}

func (d *DefaultInteractionHandler) Verify(message string) bool {
	response := d.Prompt(message + " [y/N]")
	return strings.ToLower(response) == "y"
}

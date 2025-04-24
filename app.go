package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
)

type App struct {
	ctx                     context.Context
	currentWorkingDirectory string
}

type CommandResult struct {
	Output           string `json:"output"`
	CurrentDirectory string `json:"currentDirectory"`
	Error            string `json:"error"`
	IsBuiltinCommand bool   `json:"isBuiltinCommand"`
}

func NewApp() *App {
	currentDir, _ := os.Getwd()
	return &App{
		currentWorkingDirectory: currentDir,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods teste
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func getDefaultShell() []string {
	switch runtime.GOOS {
	case "windows":
		if path, err := exec.LookPath("powershell.exe"); err == nil {
			return []string{path, "-Command"}
		}
		return []string{"cmd.exe", "/C"}

	case "darwin", "linux":
		shell := os.Getenv("SHELL")
		if shell != "" {
			return []string{shell, "-c"}
		}

		shells := []string{
			"/bin/bash",
			"/bin/zsh",
			"/bin/sh",
		}

		for _, sh := range shells {
			if _, err := os.Stat(sh); err == nil {
				return []string{sh, "-c"}
			}
		}

		return []string{"/bin/sh", "-c"}

	default:
		return []string{"/bin/sh", "-c"}
	}
}

func (a *App) ExecuteCommand(cmdStr string) CommandResult {
	cmdStr = strings.TrimSpace(cmdStr)

	if strings.HasPrefix(cmdStr, "cd") {
		parts := strings.Fields(cmdStr)
		var newDir string

		if len(parts) > 1 {
			newDir = parts[1]
			if !filepath.IsAbs(newDir) {
				newDir = filepath.Join(a.currentWorkingDirectory, newDir)
			}
		} else {
			newDir, _ = os.UserHomeDir()
		}

		if err := os.Chdir(newDir); err == nil {
			a.currentWorkingDirectory = newDir
			return CommandResult{
				Output:           "",
				CurrentDirectory: newDir,
				IsBuiltinCommand: true,
			}
		} else {
			return CommandResult{
				Output:           "cd: " + err.Error(),
				Error:            err.Error(),
				IsBuiltinCommand: true,
			}
		}
	}

	if cmdStr == "clear" {
		return CommandResult{
			Output:           "",
			IsBuiltinCommand: true,
		}
	}

	if cmdStr == "exit" {
		return CommandResult{
			Output:           "Não é possível fechar o terminal nesta interface",
			IsBuiltinCommand: true,
		}
	}

	if strings.Contains(cmdStr, "nano") ||
		strings.Contains(cmdStr, "vim") ||
		strings.Contains(cmdStr, "ssh") ||
		strings.HasPrefix(cmdStr, "sudo") {
		return CommandResult{
			Error:  "Comandos interativos não são suportados nesta interface",
			Output: "Erro: Este terminal não suporta edição interativa ou conexões SSH",
		}
	}

	shellCmd := getDefaultShell()

	cmd := exec.Command(shellCmd[0], append(shellCmd[1:], cmdStr)...)

	cmd.Dir = a.currentWorkingDirectory

	cmd.Env = os.Environ()

	output, err := cmd.CombinedOutput()

	result := CommandResult{
		Output:           string(output),
		CurrentDirectory: a.currentWorkingDirectory,
	}

	if err != nil {
		result.Error = err.Error()
	}

	return result
}

func (a *App) OpenTerminalHere(path string) CommandResult {
	if path == "." || path == "" {
		path = a.currentWorkingDirectory
	} else if !filepath.IsAbs(path) {
		path = filepath.Join(a.currentWorkingDirectory, path)
	}

	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd.exe", "/C", "start", "cmd.exe", "/K", fmt.Sprintf("cd /d %s", path))
	case "darwin":
		cmd = exec.Command("open", "-a", "Terminal", path)
	case "linux":
		terminals := []string{
			"gnome-terminal",
			"konsole",
			"xterm",
			"mate-terminal",
		}

		for _, terminal := range terminals {
			if _, err := exec.LookPath(terminal); err == nil {
				switch terminal {
				case "gnome-terminal":
					cmd = exec.Command(terminal, "--working-directory="+path)
				case "konsole":
					cmd = exec.Command(terminal, "--workdir", path)
				case "xterm":
					cmd = exec.Command(terminal, "-e", fmt.Sprintf("cd %s && $SHELL", path))
				case "mate-terminal":
					cmd = exec.Command(terminal, "--working-directory="+path)
				}
				break
			}
		}

		if cmd == nil {
			return CommandResult{
				Error:  "Nenhum terminal encontrado",
				Output: "Não foi possível encontrar um terminal para abrir",
			}
		}
	default:
		return CommandResult{
			Error:  "Sistema não suportado",
			Output: "Abertura de terminal não suportada neste sistema",
		}
	}

	err := cmd.Start()
	if err != nil {
		return CommandResult{
			Error:  "Erro ao abrir terminal",
			Output: fmt.Sprintf("Não foi possível abrir o terminal: %v", err),
		}
	}

	return CommandResult{
		Output:           fmt.Sprintf("Terminal aberto em: %s", path),
		CurrentDirectory: path,
	}
}

package app

import (
	"fmt"
	"os"
	"syscall"
	"time"
)

type ProcessWatcher struct {
	Pid int
}

func NewWatcher(pid int) *ProcessWatcher {
	return &ProcessWatcher{Pid: pid}
}

func (w *ProcessWatcher) Watch() error {
	process, err := os.FindProcess(w.Pid)
	if err != nil {
		return fmt.Errorf("process not found: %v", err)
	}

	for {
		err := process.Signal(syscall.Signal(0))
		if err != nil {
			return fmt.Errorf("process no longer exists: %v", err)
		}
		time.Sleep(time.Second)
	}
}

func (w *ProcessWatcher) IsRunning() bool {
	process, err := os.FindProcess(w.Pid)
	if err != nil {
		return false
	}
	err = process.Signal(syscall.Signal(0))
	return err == nil
}

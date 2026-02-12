package services

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/mitchellh/go-ps"
	"golang.org/x/sys/windows"
)

// GetExecutableDir returns the directory of the executable
func GetExecutableDir() (string, error) {
	exePath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Dir(exePath), nil
}

// FileExists checks if a file exists
func FileExists(path string) bool {
	_, err :=
		os.Stat(path)
	return !os.IsNotExist(err)
}

// terminateProcessTree terminates a process and all its child processes
func TerminateProcessTree(pid int) error {
	p, err := ps.FindProcess(pid)
	if err != nil {
		return err
	}
	if p == nil {
		return fmt.Errorf("process not found")
	}

	// Recursively terminate child processes
	children, err := ps.Processes()
	if err != nil {
		return err
	}
	for _, child := range children {
		if child.PPid() == pid {
			if err := TerminateProcessTree(child.Pid()); err != nil {
				return err
			}
		}
	}

	// Terminate the process itself
	processHandle, err := windows.OpenProcess(windows.PROCESS_TERMINATE, false, uint32(pid))
	if err != nil {
		return fmt.Errorf("failed to open process: %v", err)
	}
	defer windows.CloseHandle(processHandle)

	if err := windows.TerminateProcess(processHandle, 0); err != nil {
		return fmt.Errorf("failed to terminate process: %v", err)
	}

	return nil
}

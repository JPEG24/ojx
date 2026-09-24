package util

import (
	"fmt"
	"os"
	"os/exec"
)

func CompileProgram() error {
	config, err := LoadConfig()
	if err != nil {
		return fmt.Errorf("Error loading config: %v", err)
	}
	command := config.Compile
	if command == nil {
		return nil
	}

	cmd := exec.Command("sh", "-c", *command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Run()
	if err != nil {
		return fmt.Errorf("Compilation failed: %v", err)
	}
	return nil
}

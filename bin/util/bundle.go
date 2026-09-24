package util

import (
	"fmt"
	"os"
	"os/exec"
)

func Bundle(outputFile string, acl bool) ([]byte, error) {
	config, err := LoadConfig()
	if err != nil {
		return nil, fmt.Errorf("Error loading config: %v", err)
	}

	var command *string
	if acl {
		command = config.Program.BundleACL
	} else {
		command = config.Program.Bundle
	}

	if command == nil {
		return nil, fmt.Errorf("No bundle command configured")
	}

	cmd := exec.Command("sh", "-c", *command)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("Bundling failed: %v", err)
	}
	if outputFile != "" {
		if err := os.WriteFile(outputFile, output, 0644); err != nil {
			return output, fmt.Errorf("Error writing output file: %v", err)
		}
	}
	return output, nil
}

package commands

import (
	"fmt"
	"os"
	"os/exec"
)

func fail(err error) int {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
	return 1
}

func contains(args []string, value string) bool {
	for _, arg := range args {
		if arg == value {
			return true
		}
	}
	return false
}

func runCommand(name string, args ...string) int {
	cmd := exec.Command(name, args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return exitCode(err)
	}
	return 0
}

func exitCode(err error) int {
	if exitError, ok := err.(*exec.ExitError); ok {
		return exitError.ExitCode()
	}
	return 1
}

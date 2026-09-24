package commands

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"

	lib "atc/bin/util"
)

func Bundle(args []string) int {
	output, err := lib.Bundle("", contains(args, "-a"))
	if err != nil {
		return fail(err)
	}
	cmd := exec.Command("xclip", "-selection", "clipboard")
	cmd.Stdin = bytes.NewReader(output)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	status := 0
	if err := cmd.Run(); err != nil {
		status = lib.ExitCode(err)
	}
	fmt.Fprintln(os.Stderr, "Bundled source code copied to clipboard")
	return status
}

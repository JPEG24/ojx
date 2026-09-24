package commands

import (
	"fmt"
	"os"

	"atc/bin/util"
)

func Run(args []string) int {
	config, err := util.LoadConfig()
	if err != nil {
		return fail(err)
	}
	if !contains(args, "-c") {
		if status := Compile(); status != 0 {
			return status
		}
	}
	if config.Program.Run == "" {
		fmt.Fprintln(os.Stderr, "Run command is not configured")
		return 1
	}
	return util.RunShellCommand(config.Program.Run)
}

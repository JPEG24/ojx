package commands

import (
	"encoding/json"
	"fmt"
	"ojx/bin/util"
	"os"
)

func Mode(args []string) int {
	mode, err := util.GetMode()
	if err != nil {
		fmt.Fprintln(os.Stderr, "mode: error loading config:", err)
		return 1
	}

	if len(args) == 0 {
		fmt.Printf("current mode: %s\n", mode)
		return 0
	}

	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "mode: expected 1 argument, got", len(args))
		return 1
	}

	newMode := args[0]
	configPath := util.ResolveConfigPath("config.json")
	data, err := json.MarshalIndent(util.ModeConfig{Mode: newMode}, "", "  ")
	if err != nil {
		fmt.Fprintln(os.Stderr, "mode: error marshaling config:", err)
		return 1
	}

	err = os.WriteFile(configPath, data, 0644)
	if err != nil {
		fmt.Fprintln(os.Stderr, "mode: error writing config file:", err)
		return 1
	}
	return 0
}

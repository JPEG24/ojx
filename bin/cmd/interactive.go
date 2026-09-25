package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"ojx/bin/util"
)

func InteractiveInit() int {
	config, err := util.LoadConfig()
	if err != nil {
		return fail(err)
	}
	return copyInteractiveTemplate(config.Interactive.Judge)
}

func copyInteractiveTemplate(file string) int {
	src := util.ResolveConfigPath("modes", file)
	dst := filepath.Base(src)
	if _, err := os.Stat(src); err != nil {
		fmt.Fprintf(os.Stderr, "Template not found: %s\n", src)
		return 1
	}
	if _, err := os.Stat(dst); err == nil {
		fmt.Fprintf(os.Stderr, "Already exists: %s\n", dst)
		return 1
	}
	data, err := os.ReadFile(src)
	if err != nil {
		return fail(err)
	}
	if err := os.WriteFile(dst, data, 0755); err != nil {
		return fail(err)
	}
	fmt.Fprintf(os.Stderr, "Created %s\n", dst)
	return 0
}

func Interactive(skipCompile bool) int {
	config, err := util.LoadConfig()
	if err != nil {
		return fail(err)
	}
	if !skipCompile && config.Interactive.Compile != nil {
		if status := runConfiguredScript(*config.Interactive.Compile); status != 0 {
			return status
		}
	}
	return runConfiguredScript(config.Interactive.Run)
}

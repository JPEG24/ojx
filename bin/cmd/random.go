package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"ojx/bin/util"
)

func RandomTestInit() int {
	config, err := util.LoadConfig()
	if err != nil {
		return fail(err)
	}
	for _, file := range []string{config.RandomTest.Naive, config.RandomTest.Generator} {
		copyTemplate(file)
	}
	return 0
}

func RandomTest(skipCompile bool) int {
	config, err := util.LoadConfig()
	if err != nil {
		return fail(err)
	}
	if !skipCompile && config.RandomTest.Compile != nil {
		if status := runConfiguredScript(*config.RandomTest.Compile); status != 0 {
			return status
		}
	}
	return runConfiguredScript(config.RandomTest.Run)
}

func copyTemplate(file string) int {
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

func runConfiguredScript(file string) int {
	path := util.ResolveConfigPath("modes", file)
	if _, err := os.Stat(path); err != nil {
		fmt.Fprintf(os.Stderr, "Run script not found: %s\n", path)
		return 1
	}
	fmt.Fprintf(os.Stderr, "Running %s\n", path)
	return util.RunScript(path)
}

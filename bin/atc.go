package main

import (
	"fmt"
	"os"
	"path/filepath"

	commands "atc/bin/cmd"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Fprintln(os.Stderr, "No command has been specified")
		os.Exit(1)
	}
	status := 0
	switch args[0] {
	case "new", "n":
		status = commands.New(args[1:])
	case "random_test", "rt":
		status = commands.RandomTestInit()
	case "interactive", "in":
		status = commands.InteractiveInit()
	case "compile", "c":
		status = commands.Compile()
	case "run", "r":
		status = commands.Run(args[1:])
	case "test", "t":
		status = commands.Test(args[1:])
	case "bundle", "b":
		status = commands.Bundle(args[1:])
	case "submit", "s":
		status = commands.Submit(args[1:])
	case "config-dir":
		fmt.Println(filepath.Join(os.Getenv("HOME"), ".config", "atc"))
	default:
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n", args[0])
		status = 1
	}
	os.Exit(status)
}

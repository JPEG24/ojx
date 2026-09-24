package commands

import (
	lib "atc/bin/util"
	"fmt"
	"os"
	"strconv"
)

func Test(args []string) int {
	config, err := lib.LoadConfig()
	if err != nil {
		return fail(err)
	}
	skipCompile := contains(args, "-c")
	if contains(args, "-r") {
		return RandomTest(skipCompile)
	}
	if contains(args, "-i") {
		return Interactive(skipCompile)
	}
	if !skipCompile {
		if status := Compile(); status != 0 {
			return status
		}
	}

	extraArgs := make([]string, 0, len(args))
	for _, arg := range args {
		if arg != "-c" && arg != "-r" && arg != "-i" {
			extraArgs = append(extraArgs, arg)
		}
	}
	baseArgs := []string{"t", "--command", config.Program.Run, "--ignore-spaces-and-newlines"}
	task, taskErr := lib.FindCurrentTask()
	if taskErr != nil {
		return runCommand("oj", append(baseArgs, extraArgs...)...)
	}
	timeLimit := task.TimeLimit
	if timeLimit <= 0 {
		timeLimit = 2
	}
	memoryLimit := task.MemoryLimit
	if memoryLimit <= 0 {
		memoryLimit = 1024
	}
	fmt.Fprintf(os.Stderr, "Time Limit : %g sec\n", timeLimit)
	fmt.Fprintf(os.Stderr, "Memory Limit : %d MiB\n", memoryLimit)
	baseArgs = append(baseArgs, "--tle", strconv.FormatFloat(timeLimit, 'f', -1, 64), "--mle", strconv.Itoa(memoryLimit))
	return runCommand("oj", append(baseArgs, extraArgs...)...)
}

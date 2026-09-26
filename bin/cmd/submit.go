package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"ojx/bin/util"
)

func Submit(args []string) int {
	contest, err := util.FindContestJson()
	if err != nil {
		return fail(err)
	}
	data, err := os.ReadFile(*contest)
	if err != nil {
		return fail(err)
	}
	var contestJson util.Contest
	err = json.Unmarshal(data, &contestJson)
	if err != nil {
		return fail(err)
	}
	if contestJson.Platform != "atcoder" && contestJson.Platform != "yukicoder" {
		fmt.Fprintln(os.Stderr, "Only atcoder and yukicoder platforms are supported for submit command")
		return 1
	}

	// outputFile := ".ojx_submit.cpp"
	// _, err = util.Bundle(outputFile, contains(args, "-a"))
	// if err != nil {
	// 	return fail(err)
	// }

	config, err := util.LoadConfig()
	if err != nil {
		return fail(err)
	}
	task, err := util.FindCurrentTask()
	if err != nil {
		fmt.Fprintln(os.Stderr, "contest.ojx.json not found")
		return 1
	}
	extraArgs := make([]string, 0, len(args))
	for _, arg := range args {
		if arg != "-a" {
			extraArgs = append(extraArgs, arg)
		}
	}
	return runCommand("oj", append([]string{"submit", task.URL, config.SubmitFile, "--yes"}, extraArgs...)...)
}

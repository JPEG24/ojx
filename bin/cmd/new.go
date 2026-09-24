package commands

import (
	"fmt"
	contestlib "ojx/bin/new"
	"ojx/bin/util"
	"os"
)

func New(args []string) int {
	if len(args) != 2 {
		fmt.Fprintln(os.Stderr, "Usage: ojx new <platform> <contestID>")
		return 1
	}

	platform := args[0]
	contestID := args[1]

	var contest *util.Contest
	var err error
	switch platform {
	case "atcoder":
		contest, err = contestlib.NewAtCoder(contestID)
	case "codeforces":
		contest, err = contestlib.NewCodeforces(contestID)
	case "yukicoder":
		contest, err = contestlib.NewYukicoder(contestID)
	default:
		fmt.Fprintf(os.Stderr, "Unknown platform: %s\n", platform)
		return 1
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating contest: %v\n", err)
		return 1
	}
	err = contestlib.CreateContest(contest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error creating contest: %v\n", err)
		return 1
	}
	return 0
}

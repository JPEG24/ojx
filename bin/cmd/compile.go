package commands

import "atc/bin/util"

func Compile() int {
	if err := util.CompileProgram(); err != nil {
		return fail(err)
	}
	return 0
}

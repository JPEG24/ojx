package util

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

func FindCurrentTask() (*Task, error) {
	contestJson, err := FindContestJson()

	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(*contestJson)
	if err != nil {
		return nil, err
	}

	var contestFile Contest
	err = json.Unmarshal(data, &contestFile)
	if err != nil {
		return nil, err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	for i := range contestFile.Tasks {
		task := &contestFile.Tasks[i]
		taskDir := filepath.Join(filepath.Dir(*contestJson), task.TaskID)
		if cwd == taskDir {
			return task, nil
		}
	}

	return nil, fmt.Errorf("current directory is not a task directory")
}

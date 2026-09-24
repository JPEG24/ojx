package new

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"

	"ojx/bin/util"
)

func CreateContest(contest *util.Contest) error {
	if err := os.MkdirAll(contest.ContestID, 0755); err != nil {
		return fmt.Errorf("failed to create contest directory: %w", err)
	}
	if err := os.Chdir(contest.ContestID); err != nil {
		return fmt.Errorf("failed to change directory to contest directory: %w", err)
	}
	if err := WriteContestJSON(contest); err != nil {
		return fmt.Errorf("failed to write contest JSON: %w", err)
	}
	if err := SetupTasks(contest); err != nil {
		return fmt.Errorf("failed to create task directories: %w", err)
	}
	return nil
}

func WriteContestJSON(contest *util.Contest) error {
	data, err := json.MarshalIndent(contest, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal contest data: %w", err)
	}
	return os.WriteFile("contest.ojx.json", data, 0644)
}

func SetupTasks(contest *util.Contest) error {
	config, err := util.LoadConfig()
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}
	templatePath := util.ResolveConfigPath(config.Template)
	for _, task := range contest.Tasks {
		if err := os.MkdirAll(task.TaskID, 0755); err != nil {
			return fmt.Errorf("failed to create task directory: %w", err)
		}

		cmd := exec.Command("oj", "download", task.URL)
		cmd.Dir = task.TaskID
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Fprintf(os.Stderr, "failed to download testcases for task %s: %v\n", task.TaskID, err)
		}

		// Create template files for the task
		if err := CopyTemplate(templatePath, task.TaskID); err != nil {
			return fmt.Errorf("failed to copy template for task %s: %w", task.TaskID, err)
		}
	}
	return nil
}

func CopyTemplate(src string, dst string) error {
	entries, err := os.ReadDir(src)
	if err != nil {
		return fmt.Errorf("failed to read template directory: %w", err)
	}
	for _, entry := range entries {
		if entry.IsDir() && entry.Name() == "test" {
			return fmt.Errorf("Don't include the test directory in the template directory")
		}
		if !entry.IsDir() && entry.Name() == "template.json" {
			continue // Skip template.json
		}
		srcPath := filepath.Join(src, entry.Name())
		dstPath := filepath.Join(dst, entry.Name())
		if entry.IsDir() {
			if err := os.MkdirAll(dstPath, 0755); err != nil {
				return fmt.Errorf("failed to create directory %s: %w", dstPath, err)
			}
			if err := CopyTemplate(srcPath, dstPath); err != nil {
				return err
			}
		} else {
			data, err := os.ReadFile(srcPath)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %w", srcPath, err)
			}
			if err := os.WriteFile(dstPath, data, 0644); err != nil {
				return fmt.Errorf("failed to write file %s: %w", dstPath, err)
			}
		}
	}
	return nil
}

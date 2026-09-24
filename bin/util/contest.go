package util

import (
	"os"
	"path/filepath"
)

func FindContestJson() (*string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	for {
		candidate := filepath.Join(dir, "contest.ojx.json")
		if _, err := os.Stat(candidate); err == nil {
			return &candidate, nil
		}

		parent := filepath.Dir(dir)
		if parent == dir {
			return nil, os.ErrNotExist
		}
		dir = parent
	}
}

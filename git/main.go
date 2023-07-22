package git

import (
	"fmt"
	"os"
	"os/exec"
)

func Log(limit int) ([]byte, error) {
	cwd, err := os.Getwd()

	if err != nil {
		return nil, err
	}

	args := []string{"--no-pager", "log", "--name-only", "--format="}

	if limit > 0 {
		args = append(args, fmt.Sprintf("-%d", limit))
	}

	cmd := exec.Command("git", args...)
	cmd.Dir = cwd

	return cmd.Output()
}

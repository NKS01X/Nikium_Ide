package runner

import (
	"context"
	"fmt"
	"os/exec"
	"time"

	"github.com/google/uuid"
)

type CodeRunnerRequest struct {
	ID       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"userId"`
	Code     string    `json:"code"`
	Input    string    `json:"input"`
	Language string    `json:"language"`
}

type CodeRunnerResponse struct {
	ID       uuid.UUID `json:"id"`
	UserId   uuid.UUID `json:"userId"`
	Output   string    `json:"output"`
	Error    string    `json:"error"`
	ExitCode int       `json:"exitCode"`
	Duration string    `json:"duration"`
}

func RunInDocker(image string, args []string, timeoutSec int) (string, int, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSec)*time.Second)
	defer cancel()

	dockerArgs := []string{
		"run", "--rm",
		"--network=none",
		"--memory=128m",
		"--cpus=0.5",
		"--pids-limit=50",
		image,
	}
	dockerArgs = append(dockerArgs, args...)

	cmd := exec.CommandContext(ctx, "docker", dockerArgs...)
	output, err := cmd.CombinedOutput()

	if ctx.Err() == context.DeadlineExceeded {
		return "", -1, fmt.Errorf("execution timed out")
	}

	exitCode := 0
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			exitCode = exitError.ExitCode()
		} else {
			return "", -1, err
		}
	}

	return string(output), exitCode, nil
}

func RunCode(req CodeRunnerRequest) (string, int, error) {
	var image string
	var args []string

	switch req.Language {
	case "python":
		image = "python:3.11-alpine"
		args = []string{"python", "-c", req.Code}
	case "javascript", "node":
		image = "node:20-alpine"
		args = []string{"node", "-e", req.Code}
	case "go":
		image = "golang:1.22-alpine"
		args = []string{"sh", "-c", req.Code}
	default:
		image = "alpine"
		args = []string{"sh", "-c", req.Code}
	}

	return RunInDocker(image, args, 10)
}

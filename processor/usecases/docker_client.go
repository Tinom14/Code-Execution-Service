package usecases

import (
	"bytes"
	"context"
	"fmt"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type DockerClient struct {
	cli *client.Client
}

func NewDockerClient() (*DockerClient, error) {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, fmt.Errorf("creating Docker client: %w", err)
	}
	return &DockerClient{cli: cli}, nil
}
func (dc *DockerClient) RunCodeInContainer(language, code string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	tmpDir, err := os.MkdirTemp("", "code_exec")
	if err != nil {
		return "", fmt.Errorf("creating temp dir: %w", err)
	}
	defer os.RemoveAll(tmpDir)

	codeFileName := "user_code." + language
	codeFilePath := filepath.Join(tmpDir, codeFileName)
	if err := os.WriteFile(codeFilePath, []byte(code), 0644); err != nil {
		return "", fmt.Errorf("writing code to temp file: %w", err)
	}

	image, cmd := getImageAndCommand(language, codeFileName)
	if image == "" {
		return "", fmt.Errorf("unsupported language: %s", language)
	}

	resp, err := dc.cli.ContainerCreate(ctx, &container.Config{
		Image:      image,
		Cmd:        cmd,
		WorkingDir: "/app",
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: tmpDir,
				Target: "/app",
			},
		},
	}, nil, nil, "")
	if err != nil {
		return "", fmt.Errorf("creating Docker container: %w", err)
	}
	defer dc.cli.ContainerRemove(ctx, resp.ID, container.RemoveOptions{Force: true})

	if err := dc.cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		return "", fmt.Errorf("starting Docker container: %w", err)
	}

	statusCh, errCh := dc.cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			return "", fmt.Errorf("waiting for Docker container: %w", err)
		}
	case <-statusCh:
	}

	out, err := dc.cli.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		return "", fmt.Errorf("getting Docker container logs: %w", err)
	}
	defer out.Close()

	var stdoutBuf, stderrBuf bytes.Buffer
	_, err = stdcopy.StdCopy(&stdoutBuf, &stderrBuf, out)
	if err != nil {
		return "", fmt.Errorf("reading Docker container logs: %w", err)
	}

	if stderrBuf.Len() > 0 {
		return stderrBuf.String(), nil
	}

	return stdoutBuf.String(), nil
}

func getImageAndCommand(language, codeFileName string) (string, []string) {
	switch strings.ToLower(language) {
	case "python3":
		return "python:3.12-alpine", []string{"python3", "/app/" + codeFileName}
	case "c":
		return "gcc:latest", []string{"sh", "-c", fmt.Sprintf(
			"mkdir -p /tmp/bin && g++ /app/%s -o /tmp/bin/output && chmod +x /tmp/bin/output && /tmp/bin/output",
			codeFileName)}
	case "cpp":
		return "gcc:latest", []string{"sh", "-c", fmt.Sprintf(
			"mkdir -p /tmp/bin && g++ /app/%s -o /tmp/bin/output && chmod +x /tmp/bin/output && /tmp/bin/output",
			codeFileName)}
	default:
		return "", nil
	}
}

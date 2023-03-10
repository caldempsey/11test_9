package docker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/unweave/unweave/api/types"
	"github.com/unweave/unweave/builder"
)

type FsLogger struct{}

func (l *FsLogger) GetLogs(ctx context.Context, buildID string) ([]types.LogEntry, error) {
	path := filepath.Join(buildLogsDir, buildID+".json")

	f, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open build log file: %w", err)
	}
	contents, err := io.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read build log file: %w", err)
	}
	var data builder.BuildLogsV1
	if err := json.Unmarshal(contents, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal build logs: %w", err)
	}
	return data.Logs, nil
}

func (l *FsLogger) SaveLogs(ctx context.Context, buildID string, logs []types.LogEntry) error {
	if err := os.MkdirAll(buildLogsDir, 0755); err != nil {
		return fmt.Errorf("failed to create build logs directory: %w", err)
	}

	path := filepath.Join(buildLogsDir, buildID+".json")
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create build log file: %w", err)
	}

	data := builder.BuildLogsV1{
		Version: 1,
		Logs:    logs,
	}
	contents, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("failed to marshal build logs: %w", err)
	}
	if _, err := f.Write(contents); err != nil {
		return fmt.Errorf("failed to write build logs: %w", err)
	}
	return nil
}

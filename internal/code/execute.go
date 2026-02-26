package code

import (
	"bytes"
	"context"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// ExecuteTimeout is the maximum duration for code block execution.
const ExecuteTimeout = 30 * time.Second

// Execute runs the given code block and returns its output.
// Code is executed with a 30-second timeout.
func Execute(block Block) (string, error) {
	lang, ok := Languages[block.Language]
	if !ok {
		return "", fmt.Errorf("unsupported language: %s", block.Language)
	}

	// Strip hidden lines before execution
	code := StripHiddenLines(block.Code)

	// Write to temp file
	tmpDir, err := os.MkdirTemp("", "deck-exec-*")
	if err != nil {
		return "", fmt.Errorf("creating temp dir: %w", err)
	}
	defer func() { _ = os.RemoveAll(tmpDir) }()

	fileName := "main" + lang.Extension
	filePath := filepath.Join(tmpDir, fileName)
	if err := os.WriteFile(filePath, []byte(code), 0o600); err != nil {
		return "", fmt.Errorf("writing temp file: %w", err)
	}

	// Build command with placeholder replacements
	args := make([]string, len(lang.Command))
	for i, arg := range lang.Command {
		arg = strings.ReplaceAll(arg, "<file>", filePath)
		arg = strings.ReplaceAll(arg, "<name>", fileName)
		arg = strings.ReplaceAll(arg, "<path>", tmpDir)
		args[i] = arg
	}

	ctx, cancel := context.WithTimeout(context.Background(), ExecuteTimeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, args[0], args[1:]...)
	cmd.Dir = tmpDir

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return "", fmt.Errorf("execution timed out after %s", ExecuteTimeout)
		}
		if stderr.Len() > 0 {
			return "", fmt.Errorf("%s", stderr.String())
		}
		return "", err
	}

	output := stdout.String()
	if stderr.Len() > 0 {
		output += stderr.String()
	}

	return strings.TrimRight(output, "\n"), nil
}

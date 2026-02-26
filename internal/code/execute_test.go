package code

import (
	"strings"
	"testing"
)

func TestExecute(t *testing.T) {
	t.Run("runs bash code and gets output", func(t *testing.T) {
		block := Block{
			Language: "bash",
			Code:     "echo hello world",
		}

		got, err := Execute(block)
		if err != nil {
			t.Fatalf("Execute() error = %v", err)
		}
		if got != "hello world" {
			t.Errorf("Execute() = %q, want %q", got, "hello world")
		}
	})

	t.Run("unsupported language returns error", func(t *testing.T) {
		block := Block{
			Language: "cobol",
			Code:     "DISPLAY 'HELLO'",
		}

		_, err := Execute(block)
		if err == nil {
			t.Fatal("Execute() expected error for unsupported language, got nil")
		}
		if !strings.Contains(err.Error(), "unsupported language") {
			t.Errorf("Execute() error = %q, want it to contain %q", err.Error(), "unsupported language")
		}
	})

	t.Run("stderr on error", func(t *testing.T) {
		block := Block{
			Language: "bash",
			Code:     "echo failure >&2; exit 1",
		}

		_, err := Execute(block)
		if err == nil {
			t.Fatal("Execute() expected error, got nil")
		}
		if !strings.Contains(err.Error(), "failure") {
			t.Errorf("Execute() error = %q, want it to contain %q", err.Error(), "failure")
		}
	})

	t.Run("strips hidden lines before execution", func(t *testing.T) {
		block := Block{
			Language: "bash",
			Code:     "/// this line is hidden\necho visible",
		}

		got, err := Execute(block)
		if err != nil {
			t.Fatalf("Execute() error = %v", err)
		}
		if got != "visible" {
			t.Errorf("Execute() = %q, want %q", got, "visible")
		}
	})
}

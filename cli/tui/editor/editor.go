package editor

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

//go:embed template.yaml
var template []byte

//go:embed job.schema.json
var schema []byte

var editors = map[string][]string{
	"nano":          {},
	"nano.exe":      {},
	"vim":           {},
	"vim.exe":       {},
	"vi":            {},
	"vi.exe":        {},
	"notepad++.exe": {"-nosession", "-multiInst"},
	"notepad.exe":   {},
}

const templateFormat = "# yaml-language-server: $schema=%s\n%s"

func isTerminalDumb() bool {
	_, ok := os.LookupEnv("TERM")
	return !ok
}

func getEditor() (editor string, args []string) {
	var ok bool
	terminalIsDumb := isTerminalDumb()

	editor, ok = os.LookupEnv("VISUAL")
	if ok && !terminalIsDumb {
		if editor, err := exec.LookPath(editor); err == nil {
			return editor, []string{}
		}
	}

	editor, ok = os.LookupEnv("EDITOR")
	if ok {
		if editor, err := exec.LookPath(editor); err == nil {
			return editor, []string{}
		}
	}

	for editor, args := range editors {
		if editor, err := exec.LookPath(editor); err == nil {
			return editor, args
		}
	}

	return "vi", []string{}
}

func prepareFiles() (jobSchemaPath string, jobPath string, err error) {
	tempDir := os.TempDir()
	date := time.Now().Unix()
	jobSchemaPath = filepath.Join(tempDir, fmt.Sprintf("job.schema.%d.json", date))
	jobPath = filepath.Join(tempDir, fmt.Sprintf("job.%d.yaml", date))

	// Insert the yaml-language-server parameter
	template = []byte(
		fmt.Sprintf(templateFormat, jobSchemaPath, template),
	)

	if err := os.WriteFile(jobSchemaPath, schema, 0644); err != nil {
		return "", "", fmt.Errorf("fail to write %s: %w", jobSchemaPath, err)
	}

	if err := os.WriteFile(jobPath, template, 0644); err != nil {
		return "", "", fmt.Errorf("fail to write %s: %w", jobPath, err)
	}

	return jobSchemaPath, jobPath, nil
}

func openAndWaitEditor(ctx context.Context, jobPath string) error {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	var editorCommand string
	var editorArgs []string

	switch runtime.GOOS {
	case "windows":
		editorCommand, editorArgs = getEditor()
		editorArgs = append(editorArgs, jobPath)
	case "darwin":
		editorCommand = "open" // Use 'open' command on macOS
		editorArgs = []string{"-e", jobPath}
	case "linux":
		editorCommand, editorArgs = getEditor()
		editorArgs = append(editorArgs, jobPath)
	default:
		fmt.Println("Unsupported operating system")
		return nil
	}
	c := exec.CommandContext(ctx, editorCommand, editorArgs...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	return c.Run()
}

func Open(ctx context.Context) ([]byte, error) {
	jobSchemaPath, jobPath, err := prepareFiles()
	if err != nil {
		return []byte{}, err
	}
	defer func() {
		_ = os.Remove(jobSchemaPath)
		_ = os.Remove(jobPath)
	}()
	if err := openAndWaitEditor(ctx, jobPath); err != nil {
		return []byte{}, err
	}
	return os.ReadFile(jobPath)
}

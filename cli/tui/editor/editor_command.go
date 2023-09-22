// Copyright (C) 2023 DeepSquare
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package editor

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"os/exec"
	"runtime"
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

func Command(ctx context.Context, jobPath string) *exec.Cmd {
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
	return c
}

// Copyright (C) 2023 DeepSquare Asociation
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
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

//go:build license

package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"golang.org/x/tools/imports"
)

const licenseGPLFormat = `// Copyright (C) %d %s
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
`

const licenseLGPLFormat = `// Copyright (C) %d %s
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.
`

func licenseGPL() string {
	return fmt.Sprintf(licenseGPLFormat, time.Now().Year(), "DeepSquare Association")
}

func licenseLGPL() string {
	return fmt.Sprintf(licenseLGPLFormat, time.Now().Year(), "DeepSquare Association")
}

func processDirectory(dirPath string, filePaths chan<- string) error {
	out, err := os.ReadDir(dirPath)
	if err != nil {
		return err
	}

	for _, file := range out {
		filePath := filepath.Join(dirPath, file.Name())

		if file.IsDir() {
			// If it's a directory, recursively process it
			if err := processDirectory(filePath, filePaths); err != nil {
				return err
			}
		} else {
			if filepath.Ext(file.Name()) == ".go" {
				filePaths <- filePath
			}
		}
	}

	return nil
}

func main() {
	opt := imports.Options{
		Comments:   true,
		FormatOnly: false,
	}

	filePaths := make(chan string, 1)
	go func() {
		defer close(filePaths)
		if err := processDirectory(".", filePaths); err != nil {
			fmt.Println(err.Error())
		}
	}()

	for path := range filePaths {
		path = filepath.Clean(path)
		data, err := os.ReadFile(path)
		if err != nil {
			log.Fatal(err.Error())
		}
		if strings.Contains(string(data), "DO NOT EDIT") {
			continue
		}

		var license string
		switch {
		case strings.HasPrefix(path, "cmd") || strings.HasPrefix(path, "tui"):
			license = licenseGPL()
		default:
			license = licenseLGPL()
		}

		lines := strings.Split(string(data), "\n")
		if strings.Contains(lines[0], "Copyright") {
			// Delete the first block of lines that starts with "//"
			for i, line := range lines {
				if !strings.HasPrefix(line, "//") {
					data = []byte(strings.Join(lines[i:], "\n"))
					break
				}
			}

			// Replace license
			formatted, err := imports.Process(
				path,
				[]byte(license+"\n"+string(data)),
				&opt,
			)
			if err != nil {
				log.Fatal(err.Error())
			}

			if err := os.WriteFile(path, formatted, os.ModePerm); err != nil {
				log.Fatal(err.Error())
			}
			continue
		}
		if !strings.HasPrefix(string(data), license) {
			formatted, err := imports.Process(
				path,
				[]byte(license+"\n"+string(data)),
				&opt,
			)
			if err != nil {
				log.Fatal(err.Error())
			}
			if err := os.WriteFile(path, formatted, os.ModePerm); err != nil {
				log.Fatal(err.Error())
			}
		}
	}
}

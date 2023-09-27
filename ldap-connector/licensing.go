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
	"regexp"
	"strings"
	"time"

	"golang.org/x/tools/imports"
)

var CopyrightYearRegex = regexp.MustCompile(`Copyright(.*)\d{4}(.*)\n`)

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

func licenseGPL() string {
	return fmt.Sprintf(licenseGPLFormat, time.Now().Year(), "DeepSquare Asociation")
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
		line, _, _ := strings.Cut(string(data), "\n")
		if strings.Contains(line, "Copyright") {
			data = CopyrightYearRegex.ReplaceAll(
				data,
				[]byte(fmt.Sprintf("Copyright${1}%d${2}\n", time.Now().Year())),
			)
			if err := os.WriteFile(path, data, os.ModePerm); err != nil {
				log.Fatal(err.Error())
			}
			continue
		}
		license := licenseGPL()
		if !strings.HasPrefix(string(data), license) {
			opt := imports.Options{
				Comments:   true,
				FormatOnly: false,
			}
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

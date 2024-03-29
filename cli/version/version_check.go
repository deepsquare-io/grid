// Copyright (C) 2024 DeepSquare Association
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

// Package version contains functions to track the module version.
package version

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"golang.org/x/mod/semver"
)

const repo = "https://github.com/deepsquare-io/grid.git"

// CheckLatest checks the latest version from git.
func CheckLatest(ctx context.Context) (string, error) {
	opts := &git.CloneOptions{
		URL:   repo,
		Depth: 1,
		Tags:  git.AllTags,
	}
	path, err := os.UserCacheDir()
	if err != nil {
		path = os.TempDir()
	}
	path = filepath.Join(path, "supervisor")
	repo, err := git.PlainCloneContext(ctx, path, true, opts)
	if errors.Is(err, git.ErrRepositoryAlreadyExists) {
		repo, err = git.PlainOpen(path)
		if err != nil {
			return "", err
		}
		if err := repo.Fetch(&git.FetchOptions{
			Depth: 1,
			Tags:  git.AllTags,
		}); err != nil && !errors.Is(err, git.NoErrAlreadyUpToDate) {
			return "", err
		}
	} else if err != nil {
		return "", err
	}

	tags, err := repo.Tags()
	if err != nil {
		return "", err
	}
	defer tags.Close()

	var latest string
	if err = tags.ForEach(func(tag *plumbing.Reference) error {
		tagName := tag.Name().Short()
		if strings.HasPrefix(tagName, "cli/") {
			version := strings.Replace(tagName, "cli/", "", 1)
			if latest == "" || semver.Compare(version, latest) > 0 {
				latest = version
			}
		}
		return nil
	}); err != nil {
		return "", err
	}

	return latest, nil
}

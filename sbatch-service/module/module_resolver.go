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
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package module

import (
	"context"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"gopkg.in/yaml.v3"
)

const moduleFileName = "module.yaml"

var commitHashRegex = regexp.MustCompile(`^[a-f0-9]{7,40}$`)

func isCommitHash(str string) bool {
	return commitHashRegex.MatchString(str)
}

func urlParse(url string) (host, owner, repo, path string) {
	host, rest, _ := strings.Cut(url, "/")
	owner, rest, _ = strings.Cut(rest, "/")
	repo, path, _ = strings.Cut(rest, "/")
	return host, owner, repo, path
}

func Resolve(
	ctx context.Context,
	j *model.Job,
	s *model.Step,
	repository string,
	ref string,
) (*model.Module, error) {
	host, owner, repo, path := urlParse(repository)

	opts := &git.CloneOptions{
		URL: fmt.Sprintf("https://%s/%s/%s", host, owner, repo),
	}

	switch {
	case ref == "":
		opts.ReferenceName = plumbing.HEAD
		opts.Depth = 1

	case ref != "" && !isCommitHash(ref):
		if path != "" {
			ref = fmt.Sprintf("%s/%s", path, ref)
		}
		opts.ReferenceName = plumbing.NewTagReferenceName(ref)
		opts.Depth = 1
	}

	// Clone
	r, err := git.CloneContext(ctx, memory.NewStorage(), memfs.New(), opts)
	if err != nil {
		return nil, fmt.Errorf("failed to clone: %w", err)
	}

	wt, err := r.Worktree()
	if err != nil {
		return nil, fmt.Errorf("failed to fetch worktree: %w", err)
	}

	if ref != "" && isCommitHash(ref) {
		h, err := r.ResolveRevision(plumbing.Revision(ref))
		if err != nil {
			return nil, fmt.Errorf("failed to resolve revision: %w", err)
		}
		if err = wt.Checkout(&git.CheckoutOptions{
			Hash: *h,
		}); err != nil {
			return nil, fmt.Errorf("failed to checkout: %w", err)
		}
	}

	// Read the contents of the module file
	filePath := moduleFileName
	if path != "" {
		filePath = fmt.Sprintf("%s/%s", path, moduleFileName)
	}
	file, err := wt.Filesystem.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to find module.yaml: %w", err)
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("found module.yaml but failed to read module.yaml: %w", err)
	}

	rContents, err := Render(j, s, string(contents))
	if err != nil {
		return nil, err
	}

	module := &model.Module{}
	if err := yaml.Unmarshal([]byte(rContents), module); err != nil {
		return nil, err
	}

	return module, nil
}

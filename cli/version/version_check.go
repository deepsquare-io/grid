package version

import (
	"context"
	"strings"

	"github.com/go-git/go-billy/v5/memfs"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/go-git/go-git/v5/storage/memory"
	"golang.org/x/mod/semver"
)

const repo = "https://github.com/deepsquare-io/grid.git"

func CheckLatest(ctx context.Context) (string, error) {
	opts := &git.CloneOptions{
		URL:   repo,
		Depth: 1,
		Tags:  git.AllTags,
	}
	repo, err := git.CloneContext(ctx, memory.NewStorage(), memfs.New(), opts)
	if err != nil {
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

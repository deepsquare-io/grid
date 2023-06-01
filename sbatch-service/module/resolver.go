package module

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"regexp"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
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

func Resolve(ctx context.Context, j *model.Job, s *model.Step, repository string, ref string) (*model.Module, error) {
	opts := &git.CloneOptions{
		URL: fmt.Sprintf("https://%s", repository),
	}

	switch {
	case ref == "":
		opts.ReferenceName = plumbing.HEAD
		opts.Depth = 1

	case ref != "" && !isCommitHash(ref):
		opts.ReferenceName = plumbing.NewTagReferenceName(ref)
		opts.Depth = 1
	}

	// Clone
	r, err := git.CloneContext(ctx, memory.NewStorage(), memfs.New(), opts)
	if err != nil {
		return nil, err
	}

	wt, err := r.Worktree()
	if err != nil {
		return nil, err
	}

	if ref != "" && isCommitHash(ref) {
		h, err := r.ResolveRevision(plumbing.Revision(ref))
		if err != nil {
			return nil, err
		}
		if err = wt.Checkout(&git.CheckoutOptions{
			Hash: *h,
		}); err != nil {
			return nil, err
		}
	}

	// Read the contents of the module file
	file, err := wt.Filesystem.Open(moduleFileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	contents, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	rContents, err := render(j, s, string(contents))
	if err != nil {
		return nil, err
	}

	module := &model.Module{}
	if err := yaml.Unmarshal([]byte(rContents), module); err != nil {
		return nil, err
	}

	return module, nil
}

func render(j *model.Job, s *model.Step, template string) (string, error) {
	tmpl, err := engine().Parse(template)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		Job  *model.Job
		Step *model.Step
	}{
		Job:  j,
		Step: s,
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

func funcMap() template.FuncMap {
	f := sprig.TxtFuncMap()
	return f
}

func engine() *template.Template {
	return template.New("gotpl").Funcs(funcMap())
}

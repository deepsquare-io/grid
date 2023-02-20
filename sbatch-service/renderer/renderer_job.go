package renderer

import (
	"bytes"
	"strconv"
	"strings"

	_ "embed"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/logger"
	"github.com/deepsquare-io/the-grid/sbatch-service/validate"
)

//go:embed renderer_job.sh.tpl
var jobTpl string

type JobRenderer struct {
	loggerHost string
	loggerPort int
}

func NewJobRenderer(loggerEndpoint string) *JobRenderer {
	host, port, ok := strings.Cut(loggerEndpoint, ":")
	if !ok {
		logger.I.Panic("logger endpoint failed to parse")
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		logger.I.Panic("logger port failed to convert to int")
	}
	return &JobRenderer{
		loggerHost: host,
		loggerPort: portInt,
	}
}

func (r *JobRenderer) RenderJob(j *model.Job) (string, error) {
	if err := validate.I.Struct(j); err != nil {
		return "", err
	}

	tmpl, err := engine().Parse(jobTpl)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	if err = tmpl.Execute(&out, struct {
		Job    *model.Job
		Logger struct {
			Endpoint string
			Port     string
		}
	}{
		Job: j,
		Logger: struct {
			Endpoint string
			Port     string
		}{
			Endpoint: r.loggerHost,
			Port:     strconv.Itoa(r.loggerPort),
		},
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

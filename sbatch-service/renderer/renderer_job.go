package renderer

import (
	"bytes"
	"os"
	"strconv"
	"strings"

	_ "embed"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/logger"
	"github.com/deepsquare-io/the-grid/sbatch-service/validate"
	"go.uber.org/zap"
)

type JobRendererOption func(*JobRenderer)

func WithPostscript(postScriptPath string) JobRendererOption {
	return func(jr *JobRenderer) {
		if postScriptPath == "" {
			return
		}
		dat, err := os.ReadFile(postScriptPath)
		if err != nil {
			logger.I.Error("failed to read postscript path", zap.Error(err))
			return
		}
		jr.postScript = string(dat)
	}
}

func WithPrescript(preScriptPath string) JobRendererOption {
	return func(jr *JobRenderer) {
		if preScriptPath == "" {
			return
		}
		dat, err := os.ReadFile(preScriptPath)
		if err != nil {
			logger.I.Error("failed to read prescript path", zap.Error(err))
			return
		}
		jr.preScript = string(dat)
	}
}

//go:embed renderer_job.sh.tpl
var jobTpl string

type JobRenderer struct {
	loggerHost string
	loggerPort int
	loggerPath string
	preScript  string
	postScript string
}

func NewJobRenderer(
	loggerEndpoint string,
	loggerPath string,
	opts ...JobRendererOption,
) *JobRenderer {
	host, port, ok := strings.Cut(loggerEndpoint, ":")
	if !ok {
		logger.I.Panic("logger endpoint failed to parse")
	}
	portInt, err := strconv.Atoi(port)
	if err != nil {
		logger.I.Panic("logger port failed to convert to int")
	}
	jr := &JobRenderer{
		loggerHost: host,
		loggerPort: portInt,
		loggerPath: loggerPath,
	}
	for _, opt := range opts {
		opt(jr)
	}
	return jr
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
			Path     string
		}
		PostScript string
		PreScript  string
	}{
		Job: j,
		Logger: struct {
			Endpoint string
			Port     string
			Path     string
		}{
			Endpoint: r.loggerHost,
			Port:     strconv.Itoa(r.loggerPort),
			Path:     r.loggerPath,
		},
		PostScript: r.postScript,
		PreScript:  r.preScript,
	}); err != nil {
		return "", err
	}
	return out.String(), nil
}

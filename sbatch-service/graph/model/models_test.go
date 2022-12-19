package model_test

import (
	"testing"

	"github.com/deepsquare-io/the-grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/the-grid/sbatch-service/utils"
	"github.com/deepsquare-io/the-grid/sbatch-service/validate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var cleanHTTPData = model.HTTPData{
	URL: "https://url.com",
}

func TestValidateHTTPData(t *testing.T) {
	tests := []struct {
		input         model.HTTPData
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: cleanHTTPData,
			title: "Positive test",
		},
		{
			input: model.HTTPData{
				URL: "notaurl",
			},
			isError:       true,
			errorContains: []string{"url", "URL"},
			title:         "Negative test: URL is not an URL",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			err := validate.I.Struct(tt.input)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

var cleanS3Data = model.S3Data{
	Region:          "us‑east‑2",
	BucketURL:       "s3://test",
	Path:            "/test",
	AccessKeyID:     "AccessKeyID",
	SecretAccessKey: "SecretAccessKey",
	EndpointURL:     "https://s3.us‑east‑2.amazonaws.com",
}

func TestValidateS3Data(t *testing.T) {
	tests := []struct {
		input         model.S3Data
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: cleanS3Data,
			title: "Positive test",
		},
		{
			input: func() model.S3Data {
				m := cleanS3Data
				m.BucketURL = "test"
				return m
			}(),
			isError:       true,
			errorContains: []string{"BucketURL", "url"},
			title:         "Negative test: BucketURL is not valid url",
		},
		{
			input: func() model.S3Data {
				m := cleanS3Data
				m.BucketURL = "http://test"
				return m
			}(),
			isError:       true,
			errorContains: []string{"BucketURL", "startswith"},
			title:         "Negative test: BucketURL is not starting with s3://",
		},
		{
			input: func() model.S3Data {
				m := cleanS3Data
				m.BucketURL = "s3://test/"
				return m
			}(),
			isError:       true,
			errorContains: []string{"BucketURL", "endsnotwith"},
			title:         "Negative test: BucketURL is ending with /",
		},
		{
			input: func() model.S3Data {
				m := cleanS3Data
				m.Path = "test/"
				return m
			}(),
			isError:       true,
			errorContains: []string{"Path", "startswith"},
			title:         "Negative test: Path is not starting with /",
		},
		{
			input: func() model.S3Data {
				m := cleanS3Data
				m.EndpointURL = "test"
				return m
			}(),
			isError:       true,
			errorContains: []string{"EndpointURL", "url"},
			title:         "Negative test: EndpointURL is not a valid url",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			err := validate.I.Struct(tt.input)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func cleanTransportDataWithHTTP(h model.HTTPData) *model.TransportData {
	return &model.TransportData{
		HTTP: &h,
	}
}

func cleanTransportDataWithS3(s model.S3Data) *model.TransportData {
	return &model.TransportData{
		S3: &s,
	}
}

func TestValidateTransportData(t *testing.T) {
	tests := []struct {
		input         model.TransportData
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: *cleanTransportDataWithHTTP(cleanHTTPData),
			title: "Positive test",
		},
		{
			input: *cleanTransportDataWithS3(cleanS3Data),
			title: "Positive test",
		},
		{
			input: *cleanTransportDataWithHTTP(func() model.HTTPData {
				m := cleanHTTPData
				m.URL = "test"
				return m
			}()),
			isError:       true,
			errorContains: []string{"HTTP"},
			title:         "Negative test: HTTP validation error",
		},
		{
			input: *cleanTransportDataWithS3(func() model.S3Data {
				m := cleanS3Data
				m.BucketURL = "test"
				return m
			}()),
			isError:       true,
			errorContains: []string{"S3"},
			title:         "Negative test: S3 validation error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			err := validate.I.Struct(tt.input)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

var cleanEnvVar = model.EnvVar{
	Key:   "Key",
	Value: "Value",
}

func TestValidateEnvVar(t *testing.T) {
	tests := []struct {
		input         model.EnvVar
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: cleanEnvVar,
			title: "Positive test",
		},
		{
			input:         model.EnvVar{},
			isError:       true,
			errorContains: []string{"Key", "required"},
			title:         "Negative test: Key is required",
		},
		{
			input: model.EnvVar{
				Key: "'tata",
			},
			isError:       true,
			errorContains: []string{"Key", "valid_envvar_name"},
			title:         "Negative test: Key is not valid",
		},
		{
			input: model.EnvVar{
				Key: "PATH",
			},
			isError:       true,
			errorContains: []string{"Key", "ne"},
			title:         "Negative test: Setting PATH is forbidden",
		},
		{
			input: model.EnvVar{
				Key: "LD_LIBRARY_PATH",
			},
			isError:       true,
			errorContains: []string{"Key", "ne"},
			title:         "Negative test: Setting LD_LIBRARY_PATH is forbidden",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			err := validate.I.Struct(tt.input)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

var cleanForRange = model.ForRange{
	Begin:     1,
	End:       2,
	Increment: utils.Ptr(1),
}

func TestValidateForRange(t *testing.T) {
	tests := []struct {
		input         model.ForRange
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: cleanForRange,
			title: "Positive test",
		},
		{
			input: model.ForRange{
				Begin: 1,
				End:   2,
			},
			title: "Positive test: implicit increment",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			err := validate.I.Struct(tt.input)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

var cleanStepRunResources = model.StepRunResources{
	Tasks:       utils.Ptr(1),
	CpusPerTask: utils.Ptr(4),
	MemPerCPU:   utils.Ptr(4096),
	GpusPerTask: utils.Ptr(1),
}

func TestValidateStepRunResources(t *testing.T) {
	tests := []struct {
		input         model.StepRunResources
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: cleanStepRunResources,
			title: "Positive test",
		},
		{
			input: model.StepRunResources{},
			title: "Positive test: omitempty",
		},
		{
			input: func() model.StepRunResources {
				r := cleanStepRunResources
				r.CpusPerTask = utils.Ptr(0)
				return r
			}(),
			isError:       true,
			errorContains: []string{"CpusPerTask", "gt"},
			title:         "Negative test: invalid cpu count",
		},
		{
			input: func() model.StepRunResources {
				r := cleanStepRunResources
				r.GpusPerTask = utils.Ptr(-1)
				return r
			}(),
			isError:       true,
			errorContains: []string{"GpusPerTask", "gt"},
			title:         "Negative test: invalid gpu count",
		},
		{
			input: func() model.StepRunResources {
				r := cleanStepRunResources
				r.MemPerCPU = utils.Ptr(0)
				return r
			}(),
			isError:       true,
			errorContains: []string{"MemPerCPU", "gt"},
			title:         "Negative test: invalid mem count",
		},
		{
			input: func() model.StepRunResources {
				r := cleanStepRunResources
				r.Tasks = utils.Ptr(0)
				return r
			}(),
			isError:       true,
			errorContains: []string{"Tasks", "gt"},
			title:         "Negative test: invalid tasks count",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			err := validate.I.Struct(tt.input)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func cleanStepRunWith(
	res *model.StepRunResources,
	env *model.EnvVar,
) *model.StepRun {
	return &model.StepRun{
		Image:     utils.Ptr("docker.io#bash:latest"),
		Resources: res,
		X11:       utils.Ptr(true),
		Env:       []*model.EnvVar{env},
		Command:   "hostname",
		Shell:     utils.Ptr("/bin/bash"),
	}
}

func TestValidateStepRun(t *testing.T) {
	tests := []struct {
		input         model.StepRun
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: *cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar),
			title: "Positive test",
		},
		{
			input: model.StepRun{
				Resources: &cleanStepRunResources,
			},
			title: "Positive test: omitempty",
		},
		{
			input: func() model.StepRun {
				r := cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar)
				r.Image = utils.Ptr("a,ze")
				return *r
			}(),
			isError:       true,
			errorContains: []string{"Image", "valid_container_image_url"},
			title:         "Negative test: bad image",
		},
		{
			input: func() model.StepRun {
				r := cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar)
				r.Image = utils.Ptr("aze:0&az")
				return *r
			}(),
			isError:       true,
			errorContains: []string{"Image", "valid_container_image_url"},
			title:         "Negative test: bad image (bad tag)",
		},
		{
			input: func() model.StepRun {
				r := cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar)
				r.Shell = utils.Ptr("/bin/fish")
				return *r
			}(),
			isError:       true,
			errorContains: []string{"Shell", "oneof"},
			title:         "Negative test: invalid shell",
		},
		{
			input: func() model.StepRun {
				r := cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar)
				r.Env = []*model.EnvVar{{}}
				return *r
			}(),
			isError:       true,
			errorContains: []string{"Env"},
			title:         "Negative test: Env validation error",
		},
		{
			input: func() model.StepRun {
				r := cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar)
				r.Env = []*model.EnvVar{nil}
				return *r
			}(),
			isError:       true,
			errorContains: []string{"Env", "required"},
			title:         "Negative test: Env nil error",
		},
		{
			input: func() model.StepRun {
				r := cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar)
				r.Resources = &model.StepRunResources{
					CpusPerTask: utils.Ptr(-1),
				}
				return *r
			}(),
			isError:       true,
			errorContains: []string{"Resources"},
			title:         "Negative test: Resources validation error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			err := validate.I.Struct(tt.input)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func cleanStepWith(
	r *model.StepRun,
	f *model.StepFor,
) *model.Step {
	return &model.Step{
		Name: "test",
		Run:  r,
		For:  f,
	}
}

func TestValidateStep(t *testing.T) {
	tests := []struct {
		input         model.Step
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: *cleanStepWith(
				cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar),
				cleanStepForWith(&cleanForRange, cleanStepWith(
					cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar),
					nil,
				)),
			),
			title: "Positive test",
		},
		{
			input: *cleanStepWith(
				&model.StepRun{},
				nil,
			),
			isError:       true,
			errorContains: []string{"Run"},
			title:         "Negative test: StepRun validation error",
		},
		{
			input: *cleanStepWith(
				nil,
				cleanStepForWith(&cleanForRange, cleanStepWith(
					&model.StepRun{},
					nil,
				)),
			),
			isError:       true,
			errorContains: []string{"Run"},
			title:         "Negative test: StepFor validation error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			err := validate.I.Struct(tt.input)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func cleanStepForWith(
	r *model.ForRange,
	s *model.Step,
) *model.StepFor {
	return &model.StepFor{
		Parallel: true,
		Items:    []string{},
		Range:    r,
		Steps:    []*model.Step{s},
	}
}

func TestValidateStepFor(t *testing.T) {
	tests := []struct {
		input         model.StepFor
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: *cleanStepForWith(&cleanForRange, cleanStepWith(
				cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar),
				nil,
			)),
			title: "Positive test",
		},
		{
			input: func() model.StepFor {
				f := *cleanStepForWith(&cleanForRange, cleanStepWith(
					cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar),
					nil,
				))
				f.Steps = []*model.Step{nil}
				return f
			}(),
			isError:       true,
			errorContains: []string{"Steps", "required"},
			title:         "Negative test: Steps must not contains nil",
		},
		{
			input: *cleanStepForWith(&cleanForRange, cleanStepWith(
				&model.StepRun{},
				nil,
			)),
			isError:       true,
			errorContains: []string{"Steps"},
			title:         "Negative test: Steps validation errors",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			err := validate.I.Struct(tt.input)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

var cleanJobResources = model.JobResources{
	Tasks:       1,
	CpusPerTask: 4,
	MemPerCPU:   4096,
	GpusPerTask: 1,
}

func TestValidateJobResources(t *testing.T) {
	tests := []struct {
		input         model.JobResources
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: cleanJobResources,
			title: "Positive test",
		},
		{
			input: func() model.JobResources {
				r := cleanJobResources
				r.CpusPerTask = 0
				return r
			}(),
			isError:       true,
			errorContains: []string{"CpusPerTask", "gt"},
			title:         "Negative test: invalid cpu count",
		},
		{
			input: func() model.JobResources {
				r := cleanJobResources
				r.GpusPerTask = -1
				return r
			}(),
			isError:       true,
			errorContains: []string{"GpusPerTask", "gt"},
			title:         "Negative test: invalid gpu count",
		},
		{
			input: func() model.JobResources {
				r := cleanJobResources
				r.MemPerCPU = 0
				return r
			}(),
			isError:       true,
			errorContains: []string{"MemPerCPU", "gt"},
			title:         "Negative test: invalid mem count",
		},
		{
			input: func() model.JobResources {
				r := cleanJobResources
				r.Tasks = 0
				return r
			}(),
			isError:       true,
			errorContains: []string{"Tasks", "gt"},
			title:         "Negative test: invalid tasks count",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			err := validate.I.Struct(tt.input)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func cleanJobWith(
	res *model.JobResources,
	env *model.EnvVar,
	input *model.TransportData,
	step *model.Step,
	output *model.TransportData,
) *model.Job {
	return &model.Job{
		Resources:     res,
		Env:           []*model.EnvVar{env},
		EnableLogging: utils.Ptr(true),
		Input:         input,
		InputMode:     utils.Ptr(493),
		Steps:         []*model.Step{step},
		Output:        output,
	}
}

func TestValidateJob(t *testing.T) {
	tests := []struct {
		input         model.Job
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: *cleanJobWith(
				&cleanJobResources,
				&cleanEnvVar,
				cleanTransportDataWithHTTP(cleanHTTPData),
				cleanStepWith(cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar), nil),
				cleanTransportDataWithHTTP(cleanHTTPData),
			),
			title: "Positive test",
		},
		{
			input: model.Job{
				Resources: &cleanJobResources,
			},
			title: "Positive test: omitempty",
		},
		{
			input: *cleanJobWith(
				nil,
				&cleanEnvVar,
				cleanTransportDataWithHTTP(cleanHTTPData),
				cleanStepWith(cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar), nil),
				cleanTransportDataWithHTTP(cleanHTTPData),
			),
			isError:       true,
			errorContains: []string{"Resources", "required"},
			title:         "Negative test: required Resources",
		},
		{
			input: *cleanJobWith(
				&model.JobResources{
					Tasks: -1,
				},
				&cleanEnvVar,
				cleanTransportDataWithHTTP(cleanHTTPData),
				cleanStepWith(cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar), nil),
				cleanTransportDataWithHTTP(cleanHTTPData),
			),
			isError:       true,
			errorContains: []string{"Resources"},
			title:         "Negative test: Resources validation errors",
		},
		{
			input: *cleanJobWith(
				&cleanJobResources,
				nil,
				cleanTransportDataWithHTTP(cleanHTTPData),
				cleanStepWith(cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar), nil),
				cleanTransportDataWithHTTP(cleanHTTPData),
			),
			isError:       true,
			errorContains: []string{"Env", "required"},
			title:         "Negative test: required Env",
		},
		{
			input: *cleanJobWith(
				&cleanJobResources,
				&model.EnvVar{},
				cleanTransportDataWithHTTP(cleanHTTPData),
				cleanStepWith(cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar), nil),
				cleanTransportDataWithHTTP(cleanHTTPData),
			),
			isError:       true,
			errorContains: []string{"Env"},
			title:         "Negative test: Env validation errors",
		},
		{
			input: func() model.Job {
				j := cleanJobWith(
					&cleanJobResources,
					&cleanEnvVar,
					cleanTransportDataWithHTTP(cleanHTTPData),
					cleanStepWith(cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar), nil),
					cleanTransportDataWithHTTP(cleanHTTPData),
				)
				j.InputMode = utils.Ptr(512)
				return *j
			}(),
			isError:       true,
			errorContains: []string{"InputMode", "lt"},
			title:         "Negative test: invalid InputMode",
		},
		{
			input: *cleanJobWith(
				&cleanJobResources,
				&cleanEnvVar,
				&model.TransportData{HTTP: &model.HTTPData{URL: "not a url"}},
				cleanStepWith(cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar), nil),
				&model.TransportData{HTTP: &model.HTTPData{URL: "not a url"}},
			),
			isError:       true,
			errorContains: []string{"Input", "Output"},
			title:         "Negative test: input and output validation errors",
		},
		{
			input: *cleanJobWith(
				&cleanJobResources,
				&cleanEnvVar,
				&model.TransportData{HTTP: &model.HTTPData{URL: "not a url"}},
				nil,
				&model.TransportData{HTTP: &model.HTTPData{URL: "not a url"}},
			),
			isError:       true,
			errorContains: []string{"Steps", "required"},
			title:         "Negative test: step is required (not nil)",
		},
		{
			input: *cleanJobWith(
				&cleanJobResources,
				&cleanEnvVar,
				&model.TransportData{HTTP: &model.HTTPData{URL: "not a url"}},
				&model.Step{Run: &model.StepRun{}},
				&model.TransportData{HTTP: &model.HTTPData{URL: "not a url"}},
			),
			isError:       true,
			errorContains: []string{"Steps"},
			title:         "Negative test: step validation errors",
		},
	}

	for _, tt := range tests {
		t.Run(tt.title, func(t *testing.T) {
			// Act
			err := validate.I.Struct(tt.input)

			// Assert
			if tt.isError {
				assert.Error(t, err)
				for _, contain := range tt.errorContains {
					assert.ErrorContains(t, err, contain)
				}
			} else {
				require.NoError(t, err)
			}
		})
	}
}
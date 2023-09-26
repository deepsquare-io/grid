package model_test

import (
	"testing"

	"github.com/deepsquare-io/grid/sbatch-service/graph/model"
	"github.com/deepsquare-io/grid/sbatch-service/utils"
	"github.com/deepsquare-io/grid/sbatch-service/validate"
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

var cleanContainerRunWith = model.ContainerRun{
	Image:    "image:latest",
	Username: utils.Ptr("username"),
	Password: utils.Ptr("password"),
	Registry: utils.Ptr("registry"),
	X11:      utils.Ptr(true),
	Mounts: []*model.Mount{
		{
			HostDir:      "/host",
			ContainerDir: "/container",
			Options:      "ro",
		},
	},
}

func TestValidateContainerRun(t *testing.T) {
	tests := []struct {
		input         model.ContainerRun
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: cleanContainerRunWith,
			title: "Positive test",
		},
		{
			input: model.ContainerRun{
				Image: "image:latest",
			},
			title: "Positive test: omitempty",
		},
		{
			input: func() model.ContainerRun {
				r := cleanContainerRunWith
				r.Image = "a,ze"
				return r
			}(),
			isError:       true,
			errorContains: []string{"Image", "valid_container_image_url"},
			title:         "Negative test: bad image",
		},
		{
			input: func() model.ContainerRun {
				r := cleanContainerRunWith
				r.Image = "aze:0&az"
				return r
			}(),
			isError:       true,
			errorContains: []string{"Image", "valid_container_image_url"},
			title:         "Negative test: bad image (bad tag)",
		},
		{
			input: func() model.ContainerRun {
				r := cleanContainerRunWith
				r.Registry = utils.Ptr("aze:0&az")
				return r
			}(),
			isError:       true,
			errorContains: []string{"Registry", "hostname"},
			title:         "Negative test: bad hostname",
		},
		{
			input: func() model.ContainerRun {
				r := cleanContainerRunWith
				r.Mounts = append(r.Mounts, &model.Mount{
					HostDir:      "aze",
					ContainerDir: "/aze",
					Options:      "ro",
				})
				return r
			}(),
			isError:       true,
			errorContains: []string{"HostDir", "startswith"},
			title:         "Negative test: mounts HostDir bad path",
		},
		{
			input: func() model.ContainerRun {
				r := cleanContainerRunWith
				r.Mounts = append(r.Mounts, &model.Mount{
					HostDir:      "/aze",
					ContainerDir: "aze",
					Options:      "ro",
				})
				return r
			}(),
			isError:       true,
			errorContains: []string{"ContainerDir", "startswith"},
			title:         "Negative test: mounts ContainerDir bad path",
		},
		{
			input: func() model.ContainerRun {
				r := cleanContainerRunWith
				r.Mounts = append(r.Mounts, &model.Mount{
					HostDir:      "/aze",
					ContainerDir: "/aze",
					Options:      "aze",
				})
				return r
			}(),
			isError:       true,
			errorContains: []string{"Options", "oneof"},
			title:         "Negative test: mount bad options",
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
		Container:         &cleanContainerRunWith,
		DisableCPUBinding: utils.Ptr(true),
		Network:           utils.Ptr("slirp4netns"),
		DNS:               []string{"1.1.1.1"},
		MapRoot:           utils.Ptr(true),
		WorkDir:           utils.Ptr("/dir"),
		CustomNetworkInterfaces: []*model.NetworkInterface{
			{
				Wireguard: &cleanWireguard,
			},
		},
		Mpi:       utils.Ptr("pmix_v4"),
		Resources: res,
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
			input: model.StepRun{},
			title: "Positive test: omitempty",
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
		{
			input: func() model.StepRun {
				r := cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar)
				r.Network = utils.Ptr("other")
				return *r
			}(),
			isError:       true,
			errorContains: []string{"Network", "oneof"},
			title:         "Negative test: Network validation error",
		},
		{
			input: func() model.StepRun {
				r := cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar)
				r.DNS = []string{"notanip"}
				return *r
			}(),
			isError:       true,
			errorContains: []string{"DNS", "ip"},
			title:         "Negative test: DNS validation error",
		},
		{
			input: func() model.StepRun {
				r := cleanStepRunWith(&cleanStepRunResources, &cleanEnvVar)
				r.Mpi = utils.Ptr("error")
				return *r
			}(),
			isError:       true,
			errorContains: []string{"Mpi", "oneof"},
			title:         "Negative test: MPI validation error",
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
		Name:      utils.Ptr("test"),
		DependsOn: []string{"test_test"},
		Run:       r,
		For:       f,
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
			input: model.Step{},
			title: "Positive test: omitempty",
		},
		{
			input: model.Step{
				Name:      utils.Ptr("test"),
				DependsOn: []string{"test-test"},
			},
			isError:       true,
			errorContains: []string{"DependsOn", "alphanum_underscore"},
			title:         "Negative test: DependsOn validation error",
		},
		{
			input: *cleanStepWith(
				&model.StepRun{
					WorkDir: utils.Ptr("aaa"),
				},
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
					&model.StepRun{
						WorkDir: utils.Ptr("aaa"),
					},
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
				&model.StepRun{
					WorkDir: utils.Ptr("aaa"),
				},
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
				cleanTransportDataWithHTTP(cleanHTTPData),
				nil,
				cleanTransportDataWithHTTP(cleanHTTPData),
			),
			isError:       true,
			errorContains: []string{"Steps", "required"},
			title:         "Negative test: step is required (not nil)",
		},
		{
			input: *cleanJobWith(
				&cleanJobResources,
				&cleanEnvVar,
				cleanTransportDataWithHTTP(cleanHTTPData),
				&model.Step{Run: &model.StepRun{
					WorkDir: utils.Ptr("aaz"),
				}},
				cleanTransportDataWithHTTP(cleanHTTPData),
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

var cleanWireguardPeer = model.WireguardPeer{
	PublicKey:           "pub",
	AllowedIPs:          []string{"0.0.0.0/0", "172.10.0.0/32"},
	PreSharedKey:        utils.Ptr("sha"),
	Endpoint:            utils.Ptr("10.0.0.0:30"),
	PersistentKeepalive: utils.Ptr(20),
}

func TestValidateWireguardPeer(t *testing.T) {
	tests := []struct {
		input         model.WireguardPeer
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: cleanWireguardPeer,
			title: "Positive test",
		},
		{
			input: func() model.WireguardPeer {
				w := cleanWireguardPeer
				w.Endpoint = nil
				return w
			}(),
			title: "Positive test: Endpoint allow empty",
		},
		{
			input: func() model.WireguardPeer {
				w := cleanWireguardPeer
				w.AllowedIPs = []string{"a"}
				return w
			}(),
			title:         "Negative test: AllowedIPs is not a cidr",
			isError:       true,
			errorContains: []string{"AllowedIPs", "cidr"},
		},
		{
			input: func() model.WireguardPeer {
				w := cleanWireguardPeer
				w.Endpoint = utils.Ptr("a")
				return w
			}(),
			title:         "Negative test: Endpoint is not an hostname_port",
			isError:       true,
			errorContains: []string{"Endpoint", "hostname_port"},
		},
		{
			input: func() model.WireguardPeer {
				w := cleanWireguardPeer
				w.Endpoint = utils.Ptr("a")
				return w
			}(),
			title:         "Negative test: Endpoint is not an hostname_port",
			isError:       true,
			errorContains: []string{"Endpoint", "hostname_port"},
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

var cleanWireguard = model.Wireguard{
	Address:    []string{"10.0.0.1/32", "11.0.0.1/24"},
	PrivateKey: "abc",
	Peers: []*model.WireguardPeer{
		&cleanWireguardPeer,
	},
}

func TestValidateWireguard(t *testing.T) {
	tests := []struct {
		input         model.Wireguard
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: cleanWireguard,
			title: "Positive test",
		},
		{
			input: func() model.Wireguard {
				w := cleanWireguard
				w.Address = []string{"a"}
				return w
			}(),
			title:         "Negative test: Address is not an cidr",
			isError:       true,
			errorContains: []string{"Address", "cidr"},
		},
		{
			input: func() model.Wireguard {
				w := cleanWireguard
				p := cleanWireguardPeer
				p.Endpoint = utils.Ptr("err")
				w.Peers = []*model.WireguardPeer{
					&p,
				}
				return w
			}(),
			title:         "Negative test: Dive peers",
			isError:       true,
			errorContains: []string{"Peers"},
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

var cleanBore = model.Bore{
	Address:    "10.0.0.1",
	Port:       2200,
	TargetPort: 80,
}

func TestValidateBore(t *testing.T) {
	tests := []struct {
		input         model.Bore
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: cleanBore,
			title: "Positive test",
		},
		{
			input: func() model.Bore {
				w := cleanBore
				w.Address = "a.com"
				return w
			}(),
			title: "Positive test: fqdn",
		},
		{
			input: func() model.Bore {
				w := cleanBore
				w.Address = "a"
				return w
			}(),
			title:         "Negative test: Address is not an ip|fqdn",
			isError:       true,
			errorContains: []string{"Address", "ip|fqdn"},
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

var cleanStepUse = model.StepUse{
	Source:      "github.com/deepsquare-io/workflow-module-example",
	Args:        []*model.EnvVar{&cleanEnvVar},
	ExportEnvAs: utils.Ptr("EXAMPLE_MODULE"),
}

func TestValidateStepUse(t *testing.T) {
	tests := []struct {
		input         model.StepUse
		isError       bool
		errorContains []string
		title         string
	}{
		{
			input: cleanStepUse,
			title: "Positive test",
		},
		{
			input: model.StepUse{},
			title: "Positive test: Empty",
		},
		{
			input: model.StepUse{
				ExportEnvAs: utils.Ptr("'tata"),
			},
			isError:       true,
			errorContains: []string{"Key", "valid_envvar_name"},
			title:         "Negative test: Key is not valid",
		},
		{
			input: model.StepUse{
				ExportEnvAs: utils.Ptr("PATH"),
			},
			isError:       true,
			errorContains: []string{"Key", "ne"},
			title:         "Negative test: Setting PATH is forbidden",
		},
		{
			input: model.StepUse{
				ExportEnvAs: utils.Ptr("LD_LIBRARY_PATH"),
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

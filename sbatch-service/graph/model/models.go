package model

// S3Data describes the necessary variables to connect to a HTTP storage.
type HTTPData struct {
	URL string `json:"url"`
}

// S3Data describes the necessary variables to connect to a S3 storage.
type S3Data struct {
	// S3 region. Example: "us‑east‑2".
	Region string `json:"region"`
	// The S3 Bucket URL. Must not end with "/".
	//
	// Example: "s3://my-bucket".
	BucketURL string `json:"bucketUrl" validate:"startswith=s3://,endsnotwith=/"`
	// An absolute path of the bucket. Must start with "/".
	Path string `json:"path" validate:"startswith=/"`
	// An access key ID for the S3 endpoint.
	AccessKeyID string `json:"accessKeyId"`
	// A secret access key for the S3 endpoint.
	SecretAccessKey string `json:"secretAccessKey"`
	// A S3 Endpoint URL used for authentication. Example: https://s3.us‑east‑2.amazonaws.com
	EndpointURL string `json:"endpointUrl" validate:"url"`
}

type TransportData struct {
	// Use http to download a file or archive, which will be autoextracted.
	HTTP *HTTPData `json:"http"`
	// Use s3 to sync a file or directory.
	S3 *S3Data `json:"s3"`
}

// An environment variable.
type EnvVar struct {
	Key   string `json:"key" validate:"valid_envvar_name,ne=PATH"`
	Value string `json:"value"`
}

// ForRange describes the parameter for a range loop.
type ForRange struct {
	Begin     int `json:"Begin"`
	End       int `json:"End"`
	Increment int `json:"Increment"`
}

// A Job is a finite sequence of instructions.
type Job struct {
	// Environment variables accessible for the entire job.
	Env []*EnvVar `json:"env"  validate:"dive"`
	// EnableLogging enables the DeepSquare GRID Logger.
	EnableLogging *bool `json:"enableLogging"`
	// Pull data at the start of the job.
	Input *TransportData `json:"input"`
	Steps []*Step        `json:"steps" validate:"dive"`
	// Push data at the end of the job.
	//
	// Continuous sync/push can be enabled using the `continuousOutputSync` flag.
	Output *TransportData `json:"output"`
	// ContinuousOutputSync will push data during the whole job.
	//
	// This is useful when it is not desired to lose data when the job is suddenly stopped.
	ContinuousOutputSync *bool `json:"continuousOutputSync"`
}

// Resources are the allocated resources for a command.
type Resources struct {
	// Number of tasks which are run in parallel.
	//
	// Can be greater or equal to 1.
	Tasks int `json:"tasks" validate:"gte=1"`
	// Allocated CPUs per task.
	//
	// Can be greater or equal to 1.
	CpusPerTask int `json:"cpusPerTask" validate:"gte=1"`
	// Allocated memory (MB) per task.
	//
	// Can be greater or equal to 1.
	MemPerCPU int `json:"memPerCpu" validate:"gte=1"`
	// Allocated GPUs per task.
	//
	// Can be greater or equal to 0.
	GpusPerTask int `json:"gpusPerTask" validate:"gte=0"`
}

// Step is one instruction.
type Step struct {
	// Name of the instruction.
	Name string `json:"name"`
	// Run a command if not null.
	//
	// Is exclusive with "for".
	Run *StepRun `json:"run"`
	// Run a for loop if not null.
	//
	// Is exclusive with "run".
	For *StepFor `json:"for"`
}

// StepFor describes a for loop.
type StepFor struct {
	// Do a parallel for loop. Each iteration is run in parallel.
	Parallel bool `json:"parallel"`
	// Item accessible via the "$item" variable.
	//
	// Exclusive with "range".
	Items []string `json:"items" validate:"dive"`
	// Index accessible via the "$index" variable.
	//
	// Exclusive with "items".
	Range *ForRange `json:"range"`
	// Steps are run sequentially in one iteration.
	Steps []*Step `json:"steps" validate:"dive"`
}

// StepRun is one script executed with the shell.
//
// Shared storage is accessible through the $STORAGE_PATH environment variable.
//
// echo "KEY=value" >> "$STORAGE_PATH/env" can be used to share environment variables.
type StepRun struct {
	// Allocated resources for the command.
	Resources *Resources `json:"resources" validate:"dive"`
	// Environment variables accessible over the command.
	Env []*EnvVar `json:"env" validate:"dive"`
	// EnableLogging enables the DeepSquare GRID Logger.
	X11 *bool `json:"x11"`
	// Run the command inside a container.
	//
	// Format [user:password]@<host>#<image>:<tag>.
	// If null or empty, run on the host.
	Image *string `json:"image" validate:"omitempty,valid_container_image_url"`
	// Command specifies a shell script.
	Command string `json:"command"`
}

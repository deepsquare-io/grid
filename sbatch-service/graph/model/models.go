package model

// S3Data describes the necessary variables to connect to a HTTP storage.
type HTTPData struct {
	URL string `json:"url" yaml:"url"`
}

// S3Data describes the necessary variables to connect to a S3 storage.
type S3Data struct {
	// S3 region. Example: "us‑east‑2".
	Region string `json:"region" yaml:"region"`
	// The S3 Bucket URL. Must not end with "/".
	//
	// Example: "s3://my-bucket".
	BucketURL string `json:"bucketUrl" yaml:"bucketUrl" validate:"startswith=s3://,endsnotwith=/"`
	// An absolute path of the bucket. Must start with "/".
	Path string `json:"path" yaml:"path" validate:"startswith=/"`
	// An access key ID for the S3 endpoint.
	AccessKeyID string `json:"accessKeyId" yaml:"accessKeyId"`
	// A secret access key for the S3 endpoint.
	SecretAccessKey string `json:"secretAccessKey" yaml:"secretAccessKey"`
	// A S3 Endpoint URL used for authentication. Example: https://s3.us‑east‑2.amazonaws.com
	EndpointURL string `json:"endpointUrl" yaml:"endpointUrl" validate:"url"`
}

type TransportData struct {
	// Use http to download a file or archive, which will be autoextracted.
	HTTP *HTTPData `json:"http" yaml:"http"`
	// Use s3 to sync a file or directory.
	S3 *S3Data `json:"s3" yaml:"s3"`
}

// An environment variable.
type EnvVar struct {
	Key   string `json:"key" yaml:"key" validate:"valid_envvar_name,ne=PATH,ne=LD_LIBRARY_PATH"`
	Value string `json:"value" yaml:"value"`
}

// ForRange describes the parameter for a range loop.
type ForRange struct {
	// Begin is inclusive.
	Begin int `json:"begin" yaml:"begin"`
	// End is inclusive.
	End int `json:"end" yaml:"end"`
	// Increment counter by x count. If null, defaults to 1.
	Increment *int `json:"increment" yaml:"increment"`
}

// A Job is a finite sequence of instructions.
type Job struct {
	// Allocated resources for the job.
	//
	// Each resources is available as environment variables:
	// - $NTASKS: number of allowed parallel tasks
	// - $CPUS_PER_TASK: number of CPUs per task
	// - $MEM_PER_CPU: MB of memory per CPU
	// - $GPUS_PER_TASK: number of GPUs per task
	// - $GPUS: total number of GPUS
	// - $CPUS: total number of CPUS
	// - $MEM: total number of memory in MB
	Resources *Resources `json:"resources" yaml:"resources" validate:"required"`
	// Environment variables accessible for the entire job.
	Env []*EnvVar `json:"env" yaml:"env" validate:"dive"`
	// EnableLogging enables the DeepSquare GRID Logger.
	EnableLogging *bool `json:"enableLogging" yaml:"enableLogging"`
	// Pull data at the start of the job.
	//
	// It is recommended to set the mode of the data by filling the `inputMode` field.
	Input *TransportData `json:"input" yaml:"input"`
	// InputMode takes an integer that will be used to change the mode recursively (chmod -R) of the input data.
	//
	// The number shouldn't be in octal but in decimal. A mode over 512 is not accepted.
	//
	// Common modes:
	//   - 511 (user:rwx group:rwx world:rwx)
	//   - 493 (user:rwx group:r-x world:r-x)
	//   - 448 (user:rwx group:--- world:---)
	//
	// If null, the mode won't change and will default to the source.
	InputMode *int    `json:"inputMode" yaml:"inputMode" validate:"omitempty,lt=512"`
	Steps     []*Step `json:"steps" yaml:"steps" validate:"dive"`
	// Push data at the end of the job.
	//
	// Continuous sync/push can be enabled using the `continuousOutputSync` flag.
	Output *TransportData `json:"output" yaml:"output"`
	// ContinuousOutputSync will push data during the whole job.
	//
	// This is useful when it is not desired to lose data when the job is suddenly stopped.
	ContinuousOutputSync *bool `json:"continuousOutputSync" yaml:"continuousOutputSync"`
}

// Resources are the allocated resources for a command in a job, or a job in a cluster.
type Resources struct {
	// Number of tasks which are run in parallel.
	//
	// Can be greater or equal to 1.
	Tasks int `json:"tasks" yaml:"tasks" validate:"gte=1"`
	// Allocated CPUs per task.
	//
	// Can be greater or equal to 1.
	CpusPerTask int `json:"cpusPerTask" yaml:"cpusPerTask" validate:"gte=1"`
	// Allocated memory (MB) per task.
	//
	// Can be greater or equal to 1.
	MemPerCPU int `json:"memPerCpu" yaml:"memPerCpu" validate:"gte=1"`
	// Allocated GPUs per task.
	//
	// Can be greater or equal to 0.
	GpusPerTask int `json:"gpusPerTask" yaml:"gpusPerTask" validate:"gte=0"`
}

// Step is one instruction.
type Step struct {
	// Name of the instruction.
	Name string `json:"name" yaml:"name"`
	// Run a command if not null.
	//
	// Is exclusive with "for".
	Run *StepRun `json:"run" yaml:"run"`
	// Run a for loop if not null.
	//
	// Is exclusive with "run".
	For *StepFor `json:"for" yaml:"for"`
}

// StepFor describes a for loop.
type StepFor struct {
	// Do a parallel for loop. Each iteration is run in parallel.
	Parallel bool `json:"parallel" yaml:"parallel"`
	// Item accessible via the "$item" variable.
	//
	// Exclusive with "range".
	Items []string `json:"items" yaml:"items" validate:"dive"`
	// Index accessible via the "$index" variable.
	//
	// Exclusive with "items".
	Range *ForRange `json:"range" yaml:"range"`
	// Steps are run sequentially in one iteration.
	Steps []*Step `json:"steps" yaml:"steps" validate:"dive"`
}

// StepRun is one script executed with the shell.
//
// Shared storage is accessible through the $STORAGE_PATH environment variable.
//
// echo "KEY=value" >> "$STORAGE_PATH/env" can be used to share environment variables.
type StepRun struct {
	// Allocated resources for the command.
	Resources *Resources `json:"resources" yaml:"resources" validate:"required"`
	// Run the command inside a container.
	//
	// Format [<user>@][<registry>#]<image>[:<tag>].
	// reg_user="[[:alnum:]_.!~*\'()%\;:\&=+$,-@]+"
	// reg_registry="[^#]+"
	// reg_image="[[:lower:][:digit:]/._-]+"
	// reg_tag="[[:alnum:]._:-]+"
	// reg_url="^docker://((${reg_user})@)?((${reg_registry})#)?(${reg_image})(:(${reg_tag}))?$"
	//
	// It is also possible to load a squashfs file by specifying an absolute path.
	//
	// If null or empty, run on the host.
	Image *string `json:"image" yaml:"image" validate:"omitempty,valid_container_image_url"`
	// X11 mounts /tmp/.X11-unix in the container.
	//
	// If image is not defined, there is no need to define x11.
	X11 *bool `json:"x11" yaml:"x11"`
	// Environment variables accessible over the command.
	Env []*EnvVar `json:"env" yaml:"env" validate:"dive"`
	// Command specifies a shell script.
	Command string `json:"command" yaml:"command"`
	// Shell to use.
	//
	// Accepted: /bin/bash, /bin/ash, /bin/sh
	// Default: /bin/sh
	Shell *string `json:"shell" yaml:"shell" validate:"omitempty,oneof=/bin/bash /bin/ash /bin/sh"`
}

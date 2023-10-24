// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

// jkuri/bore tunnel Transport for StepRun.
//
// Bore is a proxy to expose TCP sockets.
type Bore struct {
	// Bore server IP/Address.
	//
	// Go name: "Address".
	Address string `json:"address" yaml:"address" validate:"ip|fqdn"`
	// The bore server port.
	//
	// Go name: "Port".
	Port int `json:"port" yaml:"port"`
	// Target port.
	//
	// Go name: "TargetPort".
	TargetPort int `json:"targetPort" yaml:"targetPort"`
}

type ContainerRun struct {
	// Run the command inside a container with Enroot.
	//
	// Format: image:tag. Registry and authentication is not allowed on this field.
	//
	// If the default container runtime is used:
	//
	//   - Use an absolute path to load a squashfs file. By default, it will search inside $STORAGE_PATH. /input will be equivalent to $DEEPSQUARE_INPUT, /output is $DEEPSQUARE_OUTPUT
	//
	// If apptainer=true:
	//
	//   - Use an absolute path to load a sif file or a squashfs file. By default, it will search inside $STORAGE_PATH. /input will be equivalent to $DEEPSQUARE_INPUT, /output is $DEEPSQUARE_OUTPUT
	//
	// Examples:
	//
	//   - library/ubuntu:latest
	//   - /my.squashfs
	//
	// Go name: "Image".
	Image string `json:"image" yaml:"image" validate:"valid_container_image_url"`
	// [DEPRECATED] Mounts decribes a Bind Mount.
	//
	// Please use predefined mounts like $STORAGE_PATH, $DEEPSQUARE_TMP, ...
	//
	// Go name: "Mounts".
	Mounts []*Mount `json:"mounts,omitempty" yaml:"mounts,omitempty" validate:"omitempty,dive,required"`
	// Username of a basic authentication.
	//
	// Go name: "Username".
	Username *string `json:"username,omitempty" yaml:"username,omitempty"`
	// Password of a basic authentication.
	//
	// Go name: "Password".
	Password *string `json:"password,omitempty" yaml:"password,omitempty"`
	// Container registry host.
	//
	// Defaults to registry-1.docker.io.
	//
	// Go name: "Registry".
	Registry *string `json:"registry,omitempty" yaml:"registry,omitempty" validate:"omitempty,hostname"`
	// Run with Apptainer as Container runtime instead of Enroot.
	//
	// By running with apptainer, you get access Deepsquare-hosted images.
	//
	// When running Apptainer, the container file system is read-only.
	//
	// Defaults to false.
	//
	// Go name: "Apptainer".
	Apptainer *bool `json:"apptainer,omitempty" yaml:"apptainer,omitempty"`
	// Use DeepSquare-hosted images.
	//
	// By setting to true, apptainer will be set to true.
	//
	// Go name: "DeepsquareHosted".
	DeepsquareHosted *bool `json:"deepsquareHosted,omitempty" yaml:"deepsquareHosted,omitempty"`
	// X11 mounts /tmp/.X11-unix in the container.
	//
	// Go name: "X11".
	X11 *bool `json:"x11,omitempty" yaml:"x11,omitempty"`
	// Mount the home directories.
	//
	// Go name: "MountHome".
	MountHome *bool `json:"mountHome,omitempty" yaml:"mountHome,omitempty"`
}

// An environment variable.
//
// Accessible via: "$key". "Key" name must follows the POSIX specifications (alphanumeric with underscore).
type EnvVar struct {
	// Key of the environment variable.
	//
	// Go name: "Key".
	Key string `json:"key" yaml:"key" validate:"required,valid_envvar_name,ne=PATH,ne=LD_LIBRARY_PATH"`
	// Value of the environment variable.
	//
	// Go name: "Value".
	Value string `json:"value" yaml:"value"`
}

// ForRange describes the parameter for a range loop.
type ForRange struct {
	// Begin is inclusive.
	//
	// Go name: "Begin".
	Begin int `json:"begin" yaml:"begin"`
	// End is inclusive.
	//
	// Go name: "End".
	End int `json:"end" yaml:"end"`
	// Increment counter by x count. If null, defaults to 1.
	//
	// Go name: "Increment".
	Increment *int `json:"increment,omitempty" yaml:"increment,omitempty"`
}

// HTTPData describes the necessary variables to connect to a HTTP storage.
type HTTPData struct {
	// HTTP or HTTPS URL to a file.
	//
	// Go name: "URL".
	URL string `json:"url" yaml:"url" validate:"url"`
}

// A Job is a finite sequence of instructions.
type Job struct {
	// Allocated resources for the job.
	//
	// Each resource is available as environment variables:
	// - $NTASKS: number of allowed parallel tasks
	// - $CPUS_PER_TASK: number of CPUs per task
	// - $MEM_PER_CPU: MB of memory per CPU
	// - $GPUS_PER_TASK: number of GPUs per task
	// - $GPUS: total number of GPUS
	// - $CPUS: total number of CPUS
	// - $MEM: total number of memory in MB
	//
	// Go name: "Resources".
	Resources *JobResources `json:"resources" yaml:"resources" validate:"required"`
	// Environment variables accessible for the entire job.
	//
	// Go name: "Env".
	Env []*EnvVar `json:"env,omitempty" yaml:"env,omitempty" validate:"omitempty,dive,required"`
	// EnableLogging enables the DeepSquare Grid Logger.
	//
	// Go name: "EnableLogging".
	EnableLogging *bool `json:"enableLogging,omitempty" yaml:"enableLogging,omitempty"`
	// Pull data at the start of the job.
	//
	// It is recommended to set the mode of the data by filling the `inputMode` field.
	//
	// Go name: "Input".
	Input *TransportData `json:"input,omitempty" yaml:"input,omitempty"`
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
	//
	// Go name: "InputMode".
	InputMode *int `json:"inputMode,omitempty" yaml:"inputMode,omitempty" validate:"omitempty,lt=512"`
	// Group of steps that will be run sequentially.
	//
	// Go name: "Steps".
	Steps []*Step `json:"steps" yaml:"steps" validate:"dive,required"`
	// Push data at the end of the job.
	//
	// Continuous sync/push can be enabled using the `continuousOutputSync` flag.
	//
	// Go name: "Output".
	Output *TransportData `json:"output,omitempty" yaml:"output,omitempty"`
	// ContinuousOutputSync will push data during the whole job.
	//
	// This is useful when it is not desired to lose data when the job is suddenly stopped.
	//
	// ContinousOutputSync is not available with HTTP.
	//
	// Go name: "ContinuousOutputSync".
	ContinuousOutputSync *bool `json:"continuousOutputSync,omitempty" yaml:"continuousOutputSync,omitempty"`
}

// JobResources are the allocated resources for a job in a cluster.
type JobResources struct {
	// Number of tasks which are run in parallel.
	//
	// Can be greater or equal to 1.
	//
	// Go name: "Tasks".
	Tasks int `json:"tasks" yaml:"tasks" validate:"gte=1"`
	// Allocated CPUs per task.
	//
	// Can be greater or equal to 1.
	//
	// Go name: "CpusPerTask".
	CpusPerTask int `json:"cpusPerTask" yaml:"cpusPerTask" validate:"gte=1"`
	// Allocated memory (MB) per task.
	//
	// Can be greater or equal to 1.
	//
	// Go name: "MemPerCPU".
	MemPerCPU int `json:"memPerCpu" yaml:"memPerCpu" validate:"gte=1"`
	// Allocated GPUs per task.
	//
	// Can be greater or equal to 0.
	//
	// Go name: "GpusPerTask".
	GpusPerTask int `json:"gpusPerTask" yaml:"gpusPerTask" validate:"gte=0"`
}

// A module is basically a group of steps.
//
// The module.yaml file goes through a templating engine first before getting parsed. So some variables are available:
//
// - `{{ .Job }}` and its childs, which represent the Job object using the module. Can be useful if you want to dynamically set an value based on the job.
// - `{{ .Step }}` and its childs, which represent the Step object using the module. Can be useful if you want the step name.
//
// If you want your user to pass custom steps, you can use `{{- .Step.Use.Steps | toYaml | nindent <n> }}` which is the group of steps.
//
// Example:
//
// ```yaml
// # module.yaml
// steps:
//   - name: my step
//     {{- .Step.Use.Steps | toYaml | nindent 2 }}
//   - name: my other step
//
// ```
//
// ```yaml
// # job.yaml
// steps:
//   - name: module
//     use:
//     source: git/my-module
//     steps:
//   - name: step by user
//   - name: another step by user
//
// ```
//
// Will render:
//
// ```yaml
// # module.yaml
// steps:
//   - name: my step
//   - name: step by user
//   - name: another step by user
//   - name: my other step
//
// ```
//
// Notice that the templating follows the Go format. You can also apply [sprig](http://masterminds.github.io/sprig/) templating functions.
//
// To outputs environment variables, just append KEY=value to the "${DEEPSQUARE_ENV}" file, like this:
//
// ```
// echo "KEY=value" >> "${DEEPSQUARE_ENV}"
// ```
type Module struct {
	// Name of the module.
	//
	// Go name: "Name".
	Name string `json:"name" yaml:"name"`
	// Description of the module.
	//
	// Go name: "Description".
	Description string `json:"description" yaml:"description"`
	// Minimum job resources.
	//
	// Go name: "MinimumResources".
	MinimumResources *JobResources `json:"minimumResources" yaml:"minimumResources"`
	// List of allowed arguments.
	//
	// Go name: "Inputs".
	Inputs []*ModuleInput `json:"inputs,omitempty" yaml:"inputs,omitempty" validate:"omitempty,dive,required"`
	// List of exported environment variables.
	//
	// Go name: "Outputs".
	Outputs []*ModuleOutput `json:"outputs,omitempty" yaml:"outputs,omitempty" validate:"omitempty,dive,required"`
	// Steps of the module.
	//
	// Go name: "Steps".
	Steps []*Step `json:"steps" yaml:"steps" validate:"dive,required"`
}

type ModuleInput struct {
	// Name of the input.
	//
	// Go name: "Key".
	Key string `json:"key" yaml:"key" validate:"valid_envvar_name,ne=PATH,ne=LD_LIBRARY_PATH"`
	// Description of the input.
	//
	// Go name: "Description".
	Description string `json:"description" yaml:"description"`
	// Default value.
	//
	// If not set, will default to empty string.
	//
	// Go name: "Default".
	Default *string `json:"default,omitempty" yaml:"default,omitempty"`
}

type ModuleOutput struct {
	// Name of the output.
	//
	// Go name: "Key".
	Key string `json:"key" yaml:"key" validate:"valid_envvar_name,ne=PATH,ne=LD_LIBRARY_PATH"`
	// Description of the output.
	//
	// Go name: "Description".
	Description string `json:"description" yaml:"description"`
}

// DEPRECATED: Mount decribes a Bind Mount.
//
// Mount is now deprecated. Please use predefined mounts like $STORAGE_PATH, $DEEPSQUARE_TMP, ...
type Mount struct {
	// Directory on the host to be mounted inside the container.
	//
	// Go name: "HostDir".
	HostDir string `json:"hostDir" yaml:"hostDir" validate:"startswith=/"`
	// Target directory inside the container.
	//
	// Go name: "ContainerDir".
	ContainerDir string `json:"containerDir" yaml:"containerDir" validate:"startswith=/"`
	// Options modifies the mount options.
	//
	// Accepted: ro, rw
	//
	// Go name: "Options".
	Options string `json:"options" yaml:"options" validate:"omitempty,oneof=rw ro"`
}

// Connect a network interface on a StepRun.
//
// The network interface is connected via slirp4netns.
type NetworkInterface struct {
	// Use the wireguard transport.
	//
	// Go name: "Wireguard".
	Wireguard *Wireguard `json:"wireguard,omitempty" yaml:"wireguard,omitempty"`
	// Use the bore transport.
	//
	// Go name: "Bore".
	Bore *Bore `json:"bore,omitempty" yaml:"bore,omitempty"`
}

// S3Data describes the necessary variables to connect to a S3 storage.
type S3Data struct {
	// S3 region. Example: "us‑east‑2".
	//
	// Go name: "Region".
	Region string `json:"region" yaml:"region"`
	// The S3 Bucket URL. Must not end with "/".
	//
	// Example: "s3://my-bucket".
	//
	// Go name: "BucketURL".
	BucketURL string `json:"bucketUrl" yaml:"bucketUrl" validate:"url,startswith=s3://,endsnotwith=/"`
	// The absolute path to a directory/file inside the bucket. Must start with "/".
	//
	// Go name: "Path".
	Path string `json:"path" yaml:"path" validate:"startswith=/"`
	// An access key ID for the S3 endpoint.
	//
	// Go name: "AccessKeyID".
	AccessKeyID string `json:"accessKeyId" yaml:"accessKeyId"`
	// A secret access key for the S3 endpoint.
	//
	// Go name: "SecretAccessKey".
	SecretAccessKey string `json:"secretAccessKey" yaml:"secretAccessKey"`
	// A S3 Endpoint URL used for authentication. Example: https://s3.us‑east‑2.amazonaws.com
	//
	// Go name: "EndpointURL".
	EndpointURL string `json:"endpointUrl" yaml:"endpointUrl" validate:"url"`
	// DeleteSync removes destination files that doesn't correspond to the source.
	//
	// This applies to any type of source to any type of destination (s3 or filesystem).
	//
	// See: s5cmd sync --delete.
	//
	// If null, defaults to false.
	//
	// Go name: "DeleteSync".
	DeleteSync *bool `json:"deleteSync,omitempty" yaml:"deleteSync,omitempty"`
}

// Step is one instruction.
type Step struct {
	// Name of the instruction.
	//
	// Is used for debugging.
	//
	// Go name: "Name".
	Name *string `json:"name,omitempty" yaml:"name,omitempty"`
	// Depends on wait for async tasks to end before launching this step.
	//
	// DependsOn uses the `handleName` property of a `StepAsyncLaunch`.
	//
	// Only steps at the same level can be awaited.
	//
	// BE WARNED: Uncontrolled `dependsOn` may results in dead locks.
	//
	// Go name: "DependsOn".
	DependsOn []string `json:"dependsOn,omitempty" yaml:"dependsOn,omitempty" validate:"omitempty,dive,alphanum_underscore"`
	// "If" is a boolean test that skips the step if the test is false.
	//
	// The test format is bash and variables such as $PATH or $(pwd) can be expanded.
	//
	// Note that "If" will be run after the "DependsOn".
	//
	// Example: '3 -eq 3 && "${TEST}" = "test"'.
	//
	// Go name: "If".
	If *string `json:"if,omitempty" yaml:"if,omitempty"`
	// Group of steps that will be run sequentially.
	//
	// Is exclusive with "for", "launch", "use", "run".
	//
	// Go name: "Steps".
	Steps []*Step `json:"steps,omitempty" yaml:"steps,omitempty" validate:"omitempty,dive,required"`
	// Run a command if not null.
	//
	// Is exclusive with "for", "launch", "use", "steps".
	//
	// Go name: "Run".
	Run *StepRun `json:"run,omitempty" yaml:"run,omitempty"`
	// Run a for loop if not null.
	//
	// Is exclusive with "run", "launch", "use", "steps".
	//
	// Go name: "For".
	For *StepFor `json:"for,omitempty" yaml:"for,omitempty"`
	// Launch a background process to run a group of commands if not null.
	//
	// Is exclusive with "run", "for", "use", "steps".
	//
	// Go name: "Launch".
	Launch *StepAsyncLaunch `json:"launch,omitempty" yaml:"launch,omitempty"`
	// Use a third-party group of steps.
	//
	// Is exclusive with "run", "for", "launch", "steps".
	//
	// Go name: "Use".
	Use *StepUse `json:"use,omitempty" yaml:"use,omitempty"`
	// Group of steps that will be run sequentially on error.
	//
	// Go name: "Catch".
	Catch []*Step `json:"catch,omitempty" yaml:"catch,omitempty" validate:"omitempty,dive,required"`
	// Group of steps that will be run sequentially after the group of steps or command finishes.
	//
	// Go name: "Finally".
	Finally []*Step `json:"finally,omitempty" yaml:"finally,omitempty" validate:"omitempty,dive,required"`
}

// StepAsyncLaunch describes launching a background process.
//
// StepAsyncLaunch will be awaited at the end of the job.
type StepAsyncLaunch struct {
	// HandleName is the name used to await (dependsOn field of the Step).
	//
	// Naming style is snake_case. Case is insensitive. No symbol allowed.
	//
	// Go name: "HandleName".
	HandleName *string `json:"handleName,omitempty" yaml:"handleName,omitempty" validate:"omitempty,alphanum_underscore"`
	// SignalOnParentStepExit sends a signal to the step and sub-steps when the parent step ends.
	//
	// This function can be used as a cleanup function to avoid a zombie process.
	//
	// Zombie processes will continue to run after the main process dies and therefore will not stop the job.
	//
	// If null, SIGTERM will be sent. If 0, no signal will be sent.
	//
	// Current signal :
	//
	// 1 SIGHUP Hang-up detected on the control terminal or death of the control process.
	// 2 SIGINT Abort from keyboard
	// 3 SIGQUIT Quit the keyboard
	// 9 SIGKILL If a process receives this signal, it must quit immediately and will not perform any cleaning operations.
	// 15 SIGTERM Software stop signal
	//
	// It is STRONGLY RECOMMENDED to use SIGTERM to gracefully exit a process. SIGKILL is the most abrupt and will certainly work.
	//
	// If no signal is sent, the asynchronous step will be considered a fire and forget asynchronous step and will have to terminate itself to stop the job.
	//
	// WARNING: the "no signal sent" option is subject to removal to avoid undefined behavior. Please refrain from using it.
	//
	// Go name: "SignalOnParentStepExit".
	SignalOnParentStepExit *int `json:"signalOnParentStepExit,omitempty" yaml:"signalOnParentStepExit,omitempty"`
	// Steps are run sequentially.
	//
	// Go name: "Steps".
	Steps []*Step `json:"steps" yaml:"steps" validate:"dive,required"`
}

// StepFor describes a for loop.
type StepFor struct {
	// Do a parallel for loop. Each iteration is run in parallel.
	//
	// Go name: "Parallel".
	Parallel bool `json:"parallel" yaml:"parallel"`
	// Item accessible via the {{ .Item }} variable. Index accessible via the $item variable.
	//
	// Exclusive with "range".
	//
	// Go name: "Items".
	Items []string `json:"items,omitempty" yaml:"items,omitempty"`
	// Index accessible via the $index variable.
	//
	// Exclusive with "items".
	//
	// Go name: "Range".
	Range *ForRange `json:"range,omitempty" yaml:"range,omitempty"`
	// Steps are run sequentially in one iteration.
	//
	// Go name: "Steps".
	Steps []*Step `json:"steps" yaml:"steps" validate:"dive,required"`
}

// StepRun is one script executed with the shell.
//
// A temporary shared storage is accessible through the $STORAGE_PATH environment variable.
//
// Availables caches can be used by invoking one of the following environment variable:
//
// | Environment variables                   | Lifecycle                        |
// | --------------------------------------- | -------------------------------- |
// | STORAGE_PATH                            | job duration                     |
// | DEEPSQUARE_TMP or DEEPSQUARE_SHARED_TMP | provider's policy                |
// | DEEPSQUARE_SHARED_WORLD_TMP             | provider's policy                |
// | DEEPSQUARE_DISK_TMP                     | node reboot or provider's policy |
// | DEEPSQUARE_DISK_WORLD_TMP               | node reboot or provider's policy |
//
// echo "KEY=value" >> "$DEEPSQUARE_ENV" can be used to share environment variables between steps.
//
// $DEEPSQUARE_INPUT is the path that contains imported files.
//
// $DEEPSQUARE_OUTPUT is the staging directory for uploading files.
type StepRun struct {
	// Command specifies a shell script.
	//
	// If container is used, command automatically overwrite the ENTRYPOINT and CMD. If you want to execute the entrypoint, it MUST be re-specified.
	//
	// You can install and use skopeo to inspect an image without having to pull it.
	//
	// Example: skopeo inspect --config docker://curlimages/curl:latest will gives "/entrypoint.sh" as ENTRYPOINT and "curl" as CMD. Therefore command="/entrypoint.sh curl".
	//
	// Go name: "Command".
	Command string `json:"command" yaml:"command"`
	// Shell to use.
	//
	// Accepted: /bin/bash, /bin/ash, /bin/sh
	// Default: /bin/sh
	//
	// Go name: "Shell".
	Shell *string `json:"shell,omitempty" yaml:"shell,omitempty" validate:"omitempty,oneof=/bin/bash /bin/ash /bin/sh"`
	// Allocated resources for the command.
	//
	// Go name: "Resources".
	Resources *StepRunResources `json:"resources,omitempty" yaml:"resources,omitempty"`
	// Container definition.
	//
	// If null, run on the host.
	//
	// Go name: "Container".
	Container *ContainerRun `json:"container,omitempty" yaml:"container,omitempty"`
	// Type of core networking functionality.
	//
	// Either: "host" (default) or "slirp4netns" (rootless network namespace).
	//
	// Using "slirp4netns" will automatically enables mapRoot.
	//
	// Go name: "Network".
	Network *string `json:"network,omitempty" yaml:"network,omitempty" validate:"omitempty,oneof=host slirp4netns"`
	// Configuration for the DNS in "slirp4netns" mode.
	//
	// ONLY enabled if network is "slirp4netns".
	//
	// A comma-separated list of DNS IP.
	//
	// Go name: "DNS".
	DNS []string `json:"dns,omitempty" yaml:"dns,omitempty" validate:"omitempty,dive,ip"`
	// Add custom network interfaces.
	//
	// ONLY enabled if network is "slirp4netns".
	//
	// Due to the nature of slirp4netns, the user is automatically mapped as root in order to create network namespaces and add new network interfaces.
	//
	// The tunnel interfaces will be named net0, net1, ... netX.
	//
	// The default network interface is tap0, which is a TAP interface connecting the host and the network namespace.
	//
	// Go name: "CustomNetworkInterfaces".
	CustomNetworkInterfaces []*NetworkInterface `json:"customNetworkInterfaces,omitempty" yaml:"customNetworkInterfaces,omitempty" validate:"omitempty,dive,required"`
	// Environment variables accessible over the command.
	//
	// Go name: "Env".
	Env []*EnvVar `json:"env,omitempty" yaml:"env,omitempty" validate:"omitempty,dive,required"`
	// Remap UID to root. Does not grant elevated system permissions, despite appearances.
	//
	// If the "default" (Enroot) container runtime is used, it will use the `--container-remap-root` flags.
	//
	// If the "apptainer" container runtime is used, the `--fakeroot` flag will be passed.
	//
	// If no container runtime is used, `unshare --user --map-root-user --mount` will be used and a user namespace will be created.
	//
	// It is not recommended to use mapRoot with network=slirp4netns, as it will create 2 user namespaces (and therefore will be useless).
	//
	// If null, default to false.
	//
	// Go name: "MapRoot".
	MapRoot *bool `json:"mapRoot,omitempty" yaml:"mapRoot,omitempty"`
	// Working directory.
	//
	// If the "default" (Enroot) container runtime is used, it will use the `--container-workdir` flag.
	//
	// If the "apptainer" container runtime is used, the `--pwd` flag will be passed.
	//
	// If no container runtime is used, `cd` will be executed first.
	//
	// If null, default to use $STORAGE_PATH as working directory.
	//
	// Go name: "WorkDir".
	WorkDir *string `json:"workDir,omitempty" yaml:"workDir,omitempty" validate:"omitempty,startswith=/"`
	// DisableCPUBinding disables process affinity binding to tasks.
	//
	// Can be useful when running MPI jobs.
	//
	// If null, defaults to false.
	//
	// Go name: "DisableCPUBinding".
	DisableCPUBinding *bool `json:"disableCpuBinding,omitempty" yaml:"disableCpuBinding,omitempty"`
	// MPI selection.
	//
	// Must be one of: none, pmix_v4, pmi2.
	//
	// If null, will default to infrastructure provider settings (which may not be what you want).
	//
	// Go name: "Mpi".
	Mpi *string `json:"mpi,omitempty" yaml:"mpi,omitempty" validate:"omitempty,oneof=none pmix_v4 pmi2"`
}

// StepRunResources are the allocated resources for a command in a job.
type StepRunResources struct {
	// Number of tasks which are run in parallel.
	//
	// Can be greater or equal to 1.
	//
	// If null, default to 1.
	//
	// Go name: "Tasks".
	Tasks *int `json:"tasks,omitempty" yaml:"tasks,omitempty" validate:"omitempty,gte=1"`
	// Allocated CPUs per task.
	//
	// Can be greater or equal to 1.
	//
	// If null, defaults to the job resources.
	//
	// Go name: "CpusPerTask".
	CpusPerTask *int `json:"cpusPerTask,omitempty" yaml:"cpusPerTask,omitempty" validate:"omitempty,gte=1"`
	// Allocated memory (MB) per task.
	//
	// Can be greater or equal to 1.
	//
	// If null, defaults to the job resources.
	//
	// Go name: "MemPerCPU".
	MemPerCPU *int `json:"memPerCpu,omitempty" yaml:"memPerCpu,omitempty" validate:"omitempty,gte=1"`
	// Allocated GPUs per task.
	//
	// Can be greater or equal to 0.
	//
	// If null, defaults to the job resources.
	//
	// Go name: "GpusPerTask".
	GpusPerTask *int `json:"gpusPerTask,omitempty" yaml:"gpusPerTask,omitempty" validate:"omitempty,gte=0"`
}

type StepUse struct {
	// Source of the group of steps.
	//
	// Syntax: <url>@<tag/hash>
	//
	// Example: github.com/example/my-module@v1
	// Example: github.com/example/module-monorepo/my-module@v1
	//
	// The host must be a git repository accessible via HTTPS.
	// The path must indicates a directory. For example, `/my-module` indicates the root directory of the repository `my-module`.
	// `module-monorepo/my-module` indicates the subdirectory `my-module` of the repository `module-monorepo`.
	//
	// Go name: "Source".
	Source string `json:"source" yaml:"source"`
	// Arguments to be passed as inputs to the group of steps.
	//
	// Go name: "Args".
	Args []*EnvVar `json:"args,omitempty" yaml:"args,omitempty"`
	// Environment variables exported with be prefixed with the value of this field.
	//
	// Exemple: If exportEnvAs=MY_MODULE, and KEY is exported. Then you can invoke ${MY_MODULE_KEY} environment variable.
	//
	// Go name: "ExportEnvAs".
	ExportEnvAs *string `json:"exportEnvAs,omitempty" yaml:"exportEnvAs,omitempty" validate:"omitempty,valid_envvar_name,ne=PATH,ne=LD_LIBRARY_PATH"`
	// Additional children steps to the module.
	//
	// If the module allow children steps, these steps will be passed to the module to replace {{ .Step.Run.Steps }}.
	//
	// Go name: "Steps".
	Steps []*Step `json:"steps,omitempty" yaml:"steps,omitempty" validate:"omitempty,dive,required"`
}

type TransportData struct {
	// Use http to download a file or archive, which will be autoextracted.
	//
	// Go name: "HTTP".
	HTTP *HTTPData `json:"http,omitempty" yaml:"http,omitempty"`
	// Use s3 to sync a file or directory.
	//
	// Go name: "S3".
	S3 *S3Data `json:"s3,omitempty" yaml:"s3,omitempty"`
}

// Wireguard VPN Transport for StepRun.
//
// The Wireguard VPN can be used as a gateway for the steps. All that is needed is a Wireguard server outside the cluster that acts as a public gateway.
//
// Wireguard transport uses UDP hole punching to connect to the VPN Server.
//
// Disabled settings: PreUp, PostUp, PreDown, PostDown, ListenPort, Table, MTU, SaveConfig.
//
// If these features are necessary, please do contact DeepSquare developpers!
type Wireguard struct {
	// The IP addresses of the wireguard interface.
	//
	// Format is a CIDRv4 (X.X.X.X/X) or CIDRv6.
	//
	// Recommendation is to take one IP from the 10.0.0.0/24 range (example: 10.0.0.2/24).
	//
	// Go name: "Address".
	Address []string `json:"address,omitempty" yaml:"address,omitempty" validate:"omitempty,dive,cidr"`
	// The client private key.
	//
	// Go name: "PrivateKey".
	PrivateKey string `json:"privateKey" yaml:"privateKey"`
	// The peers connected to the wireguard interface.
	//
	// Go name: "Peers".
	Peers []*WireguardPeer `json:"peers,omitempty" yaml:"peers,omitempty" validate:"omitempty,dive,required"`
}

// A Wireguard Peer.
type WireguardPeer struct {
	// The peer private key.
	//
	// Go name: "PublicKey".
	PublicKey string `json:"publicKey" yaml:"publicKey"`
	// The peer pre-shared key.
	//
	// Go name: "PreSharedKey".
	PreSharedKey *string `json:"preSharedKey,omitempty" yaml:"preSharedKey,omitempty"`
	// Configuration of wireguard routes.
	//
	// Format is a CIDRv4 (X.X.X.X/X) or CIDRv6.
	//
	// 0.0.0.0/0 (or ::/0) would forward all packets to the tunnel. If you plan to use the Wireguard VPN as a gateway, you MUST set this IP range.
	//
	// <server internal IP>/32 (not the server's public IP) would forward all packets to the tunnel with the server IP as the destination. MUST be set.
	//
	// <VPN IP range> would forward all packets to the tunnel with the local network as the destination. Useful if you want peers to communicate with each other and want the gateway to act as a router.
	//
	// Go name: "AllowedIPs".
	AllowedIPs []string `json:"allowedIPs,omitempty" yaml:"allowedIPs,omitempty" validate:"omitempty,dive,cidr"`
	// The peer endpoint.
	//
	// Format is IP:port.
	//
	// This would be the Wireguard server.
	//
	// Go name: "Endpoint".
	Endpoint *string `json:"endpoint,omitempty" yaml:"endpoint,omitempty" validate:"omitempty,hostname_port"`
	// Initiate the handshake and re-initiate regularly.
	//
	// Takes seconds as parameter. 25 seconds is recommended.
	//
	// You MUST set the persistent keepalive to enables UDP hole-punching.
	//
	// Go name: "PersistentKeepalive".
	PersistentKeepalive *int `json:"persistentKeepalive,omitempty" yaml:"persistentKeepalive,omitempty"`
}

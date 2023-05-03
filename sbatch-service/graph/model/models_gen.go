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
	Mounts []*Mount `json:"mounts,omitempty" yaml:"mounts,omitempty" validate:"dive,required"`
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
	Env []*EnvVar `json:"env,omitempty" yaml:"env,omitempty" validate:"dive,required"`
	// EnableLogging enables the DeepSquare GRID Logger.
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
// - {{ .Job }} and its childs, which represent the Job object using the module. Can be useful if you want to dynamically set an value based on the job.
// - {{ .Step }} and its childs, which represent the Step object using the module. Can be useful if you want the step name.
//
// Notice that the templating follows the Go format. You can also apply sprig templating functions.
//
// To outputs environment variables, just append KEY=value to the "${DEEPSQUARE_ENV}" file.
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
	Inputs []*ModuleInput `json:"inputs,omitempty" yaml:"inputs,omitempty" validate:"dive,required"`
	// List of exported environment variables.
	//
	// Go name: "Outputs".
	Outputs []*ModuleOutput `json:"outputs,omitempty" yaml:"outputs,omitempty" validate:"dive,required"`
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
	DependsOn []string `json:"dependsOn,omitempty" yaml:"dependsOn,omitempty" validate:"dive,alphanum_underscore"`
	// Run a command if not null.
	//
	// Is exclusive with "for", "launch", "use".
	//
	// Go name: "Run".
	Run *StepRun `json:"run,omitempty" yaml:"run,omitempty"`
	// Run a for loop if not null.
	//
	// Is exclusive with "run", "launch", "use".
	//
	// Go name: "For".
	For *StepFor `json:"for,omitempty" yaml:"for,omitempty"`
	// Launch a background process to run a group of commands if not null.
	//
	// Is exclusive with "run", "for", "use".
	//
	// Go name: "Launch".
	Launch *StepAsyncLaunch `json:"launch,omitempty" yaml:"launch,omitempty"`
	// Use a third-party group of steps.
	//
	// Is exclusive with "run", "for", "launch".
	//
	// Go name: "Use".
	Use *StepUse `json:"use,omitempty" yaml:"use,omitempty"`
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
	Steps []*Step `json:"steps" yaml:"steps"`
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
	ExportEnvAs *string `json:"exportEnvAs,omitempty" yaml:"exportEnvAs,omitempty" validate:"valid_envvar_name,ne=PATH,ne=LD_LIBRARY_PATH"`
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
	Address []string `json:"address,omitempty" yaml:"address,omitempty" validate:"dive,cidr"`
	// The client private key.
	//
	// Go name: "PrivateKey".
	PrivateKey string `json:"privateKey" yaml:"privateKey"`
	// The peers connected to the wireguard interface.
	//
	// Go name: "Peers".
	Peers []*WireguardPeer `json:"peers,omitempty" yaml:"peers,omitempty" validate:"dive,required"`
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
	AllowedIPs []string `json:"allowedIPs,omitempty" yaml:"allowedIPs,omitempty" validate:"dive,cidr"`
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

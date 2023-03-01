package model

// S3Data describes the necessary variables to connect to a HTTP storage.
type HTTPData struct {
	URL string `json:"url" yaml:"url" validate:"url"`
}

// S3Data describes the necessary variables to connect to a S3 storage.
type S3Data struct {
	// S3 region. Example: "us‑east‑2".
	Region string `json:"region" yaml:"region"`
	// The S3 Bucket URL. Must not end with "/".
	//
	// Example: "s3://my-bucket".
	BucketURL string `json:"bucketUrl" yaml:"bucketUrl" validate:"url,startswith=s3://,endsnotwith=/"`
	// An absolute path of the bucket. Must start with "/".
	Path string `json:"path" yaml:"path" validate:"startswith=/"`
	// An access key ID for the S3 endpoint.
	AccessKeyID string `json:"accessKeyId" yaml:"accessKeyId"`
	// A secret access key for the S3 endpoint.
	SecretAccessKey string `json:"secretAccessKey" yaml:"secretAccessKey"`
	// A S3 Endpoint URL used for authentication. Example: https://s3.us‑east‑2.amazonaws.com
	EndpointURL string `json:"endpointUrl" yaml:"endpointUrl" validate:"url"`
	// DeleteSync removes destination files that doesn't correspond to the source.
	//
	// This applies to any type of source to any type of destination (s3 or filesystem).
	//
	// See: s5cmd sync --delete.
	//
	// If null, defaults to false.
	DeleteSync *bool `json:"deleteSync" yaml:"deleteSync"`
}

type TransportData struct {
	// Use http to download a file or archive, which will be autoextracted.
	HTTP *HTTPData `json:"http" yaml:"http"`
	// Use s3 to sync a file or directory.
	S3 *S3Data `json:"s3" yaml:"s3"`
}

// An environment variable.
type EnvVar struct {
	Key   string `json:"key" yaml:"key" validate:"required,valid_envvar_name,ne=PATH,ne=LD_LIBRARY_PATH"`
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
	Resources *JobResources `json:"resources" yaml:"resources" validate:"required"`
	// Environment variables accessible for the entire job.
	Env []*EnvVar `json:"env" yaml:"env" validate:"dive,required"`
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
	Steps     []*Step `json:"steps" yaml:"steps" validate:"dive,required"`
	// Push data at the end of the job.
	//
	// Continuous sync/push can be enabled using the `continuousOutputSync` flag.
	Output *TransportData `json:"output" yaml:"output"`
	// ContinuousOutputSync will push data during the whole job.
	//
	// This is useful when it is not desired to lose data when the job is suddenly stopped.
	//
	// ContinousOutputSync is not available with HTTP.
	ContinuousOutputSync *bool `json:"continuousOutputSync" yaml:"continuousOutputSync"`
}

// JobResources are the allocated resources for a command in a job, or a job in a cluster.
type JobResources struct {
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
	Items []string `json:"items" yaml:"items"`
	// Index accessible via the "$index" variable.
	//
	// Exclusive with "items".
	Range *ForRange `json:"range" yaml:"range"`
	// Steps are run sequentially in one iteration.
	Steps []*Step `json:"steps" yaml:"steps" validate:"dive,required"`
}

// StepRunResources are the allocated resources for a command in a job.
type StepRunResources struct {
	// Number of tasks which are run in parallel.
	//
	// Can be greater or equal to 1.
	//
	// If null, default to 1.
	Tasks *int `json:"tasks" yaml:"tasks" validate:"omitempty,gte=1"`
	// Allocated CPUs per task.
	//
	// Can be greater or equal to 1.
	//
	// If null, defaults to the job resources.
	CpusPerTask *int `json:"cpusPerTask" yaml:"cpusPerTask" validate:"omitempty,gte=1"`
	// Allocated memory (MB) per task.
	//
	// Can be greater or equal to 1.
	//
	// If null, defaults to the job resources.
	MemPerCPU *int `json:"memPerCpu" yaml:"memPerCpu" validate:"omitempty,gte=1"`
	// Allocated GPUs per task.
	//
	// Can be greater or equal to 0.
	//
	// If null, defaults to the job resources.
	GpusPerTask *int `json:"gpusPerTask" yaml:"gpusPerTask" validate:"omitempty,gte=0"`
}

// Mount decribes a Bind Mount.
type Mount struct {
	HostDir      string `json:"hostDir" yaml:"hostDir" validate:"startswith=/"`
	ContainerDir string `json:"containerDir" yaml:"containerDir" validate:"startswith=/"`
	// Options modifies the mount options.
	//
	// Accepted: ro, rw
	Options string `json:"options" yaml:"options" validate:"omitempty,oneof=rw ro"`
}

type ContainerRun struct {
	// Run the command inside a container with Pyxis.
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
	Image string `json:"image" yaml:"image" validate:"valid_container_image_url"`
	// Mount decribes a Bind Mount.
	Mounts []*Mount `json:"mounts" yaml:"mounts" validate:"dive"`
	// Username of a basic authentication.
	Username *string `json:"username" yaml:"username"`
	// Password of a basic authentication.
	Password *string `json:"password" yaml:"password"`
	// Container registry host.
	//
	// Defaults to registry-1.docker.io
	Registry *string `json:"registry" yaml:"registry" validate:"omitempty,hostname"`
	// Run with Apptainer as Container runtime instead of Pyxis.
	//
	// By running with apptainer, you get access Deepsquare-hosted images.
	//
	// Defaults to false.
	Apptainer *bool `json:"apptainer" yaml:"apptainer"`
	// Use DeepSquare-hosted images.
	//
	// By setting to true, apptainer will be set to true.
	DeepsquareHosted *bool `json:"deepsquareHosted" yaml:"deepsquareHosted"`
	// X11 mounts /tmp/.X11-unix in the container.
	X11 *bool `json:"x11" yaml:"x11"`
}

// StepRun is one script executed with the shell.
//
// Shared storage is accessible through the $STORAGE_PATH environment variable.
//
// echo "KEY=value" >> "$DEEPSQUARE_ENV" can be used to share environment variables between steps.
//
// $DEEPSQUARE_INPUT is the path that contains imported files.
//
// $DEEPSQUARE_OUTPUT is the staging directory for uploading files.
type StepRun struct {
	// Allocated resources for the command.
	Resources *StepRunResources `json:"resources" yaml:"resources" validate:"required"`
	// Container definition.
	//
	// If null, run on the host.
	Container *ContainerRun `json:"container" yaml:"container"`
	// DisableCPUBinding disables process affinity binding to tasks.
	//
	// Can be useful when running MPI jobs.
	//
	// If null, defaults to false.
	DisableCPUBinding *bool `json:"disableCpuBinding" yaml:"disableCpuBinding"`
	// Environment variables accessible over the command.
	Env []*EnvVar `json:"env" yaml:"env" validate:"dive,required"`
	// Command specifies a shell script.
	Command string `json:"command" yaml:"command"`
	// Shell to use.
	//
	// Accepted: /bin/bash, /bin/ash, /bin/sh
	// Default: /bin/sh
	Shell *string `json:"shell" yaml:"shell" validate:"omitempty,oneof=/bin/bash /bin/ash /bin/sh"`
	// Type of core networking functionality.
	//
	// Either: "host" (default) or "slirp4netns" (rootless network namespace).
	Network *string `json:"network" yaml:"network" validate:"omitempty,oneof=host slirp4netns"`
	// Configuration for the DNS in "slirp4netns" mode.
	//
	// ONLY enabled if network is "slirp4netns".
	//
	// A comma-separated list of DNS IP.
	DNS []string `json:"dns" yaml:"dns" validate:"dive,ip"`
	// Remap UID to root.
	//
	// If the "default" (enroot/pyxis) container runtime is used, it will use the `--root` (--container-remap-root for Pyxis) flags.
	//
	// If the "apptainer" container runtime is used, the `--fakeroot` flag will be passed.
	//
	// If no container runtime is used, `unshared --user --map-root-user --mount` will be used and a user namespace will be created.
	//
	// If null, default to false.
	MapRoot *bool `json:"mapRoot" yaml:"mapRoot"`
	// Add custom network interfaces.
	//
	// ONLY enabled if network is "slirp4netns".
	//
	// Due to the nature of slirp4netns, the user is automatically mapped as root in order to create network namespaces and add new network interfaces.
	//
	// The tunnel interfaces will be named net0, net1, ... netX.
	//
	// The default network interface is tap0, which is a TAP interface connecting the host and the network namespace.
	CustomNetworkInterfaces []*NetworkInterface `json:"customNetworkInterfaces" yaml:"customNetworkInterfaces" validate:"dive"`
}

// Connect a network interface on a StepRun.
//
// The network interface is connected via slirp4netns.
type NetworkInterface struct {
	// Use the wireguard transport.
	Wireguard *Wireguard `json:"wireguard" yaml:"wireguard"`
}

// Wireguard VPN Transport for StepRun.
//
// The Wireguard VPN can be used as a gateway for the steps. All that is needed is a Wireguard server outside the cluster that acts as a public gateway.
//
// Wireguard transport uses UDP hole punching to connect to the VPN Server.
//
// Disabled settings: PreUp, PostUp, PreDown, PostDown, Table, MTU, SaveConfig.
//
// If these features are necessary, please do contact DeepSquare developpers!
type Wireguard struct {
	// The IP addresses of the wireguard interface.
	//
	// Format is a CIDRv4 (X.X.X.X/X) or CIDRv6.
	//
	// Recommendation is to take one IP from the 10.0.0.0/24 range.
	Address []string `json:"address" yaml:"address" validate:"dive,cidr"`
	// The client private key.
	PrivateKey string `json:"privateKey" yaml:"privateKey"`
	// The peers connected to the wireguard interface.
	Peers []*WireguardPeer `json:"peers" yaml:"peers" validate:"dive"`
}

// A Wireguard Peer.
type WireguardPeer struct {
	// The peer private key.
	PublicKey string `json:"publicKey" yaml:"publicKey"`
	// The peer shared key.
	PreSharedKey *string `json:"preSharedKey" yaml:"preSharedKey"`
	// Configuration of wireguard routes.
	//
	// Format is a CIDRv4 (X.X.X.X/X) or CIDRv6.
	//
	// 0.0.0.0/0 (or ::/0) would forward all packets to the tunnel. If you plan to use the Wireguard VPN as a gateway, you MUST set this IP range.
	//
	// <server internal IP>/32 (not the server's public IP) would forward all packets to the tunnel with the server IP as the destination. MUST be set.
	//
	// <VPN IP range> would forward all packets to the tunnel with the local network as the destination. Useful if you want peers to communicate with each other.
	AllowedIPs []string `json:"allowedIPs" yaml:"allowedIPs" validate:"dive,cidr"`
	// The peer endpoint.
	//
	// Format is IP:port.
	//
	// This would be the Wireguard server IP.
	Endpoint *string `json:"endpoint" yaml:"endpoint" validate:"omitempty,hostname_port"`
	// Initiate the handshake and re-initiate regularly.
	//
	// Takes seconds as parameter. 25 seconds is recommended.
	//
	// You MUST to set the persistent keepalive to enables UDP hole-punching.
	PersistentKeepalive *int `json:"persistentKeepalive" yaml:"persistentKeepalive"`
}

package model

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

// DEPRECATED: Mount decribes a Bind Mount.
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
	DNS []string `json:"dns,omitempty" yaml:"dns,omitempty" validate:"dive,ip"`
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
	CustomNetworkInterfaces []*NetworkInterface `json:"customNetworkInterfaces,omitempty" yaml:"customNetworkInterfaces,omitempty" validate:"dive,required"`
	// Environment variables accessible over the command.
	//
	// Go name: "Env".
	Env []*EnvVar `json:"env,omitempty" yaml:"env,omitempty" validate:"dive,required"`
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

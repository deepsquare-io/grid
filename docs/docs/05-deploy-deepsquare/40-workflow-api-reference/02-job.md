---
toc_max_heading_level: 6
---

# Job Reference

Jobs submitted to DeepSquare use the JSON format. The top level object is the [Job](#job-top-level-object).

### Job (top level object)

A Job is a finite sequence of instructions.

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>resources</strong></td>
<td valign="top"><a href="#resources-jobresources">JobResources</a>!</td>
<td>

Allocated resources for the job.

Each resource is available as environment variables:

- $NTASKS: number of allowed parallel tasks
- $CPUS_PER_TASK: number of CPUs per task
- $MEM_PER_CPU: MB of memory per CPU
- $GPUS_PER_TASK: number of GPUs per task
- $GPUS: total number of GPUS
- $CPUS: total number of CPUS
- $MEM: total number of memory in MB

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>env</strong></td>
<td valign="top">[<a href="#env-envvar">EnvVar</a>!]</td>
<td>

Environment variables accessible for the entire job.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>enableLogging</strong></td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td>

EnableLogging enables the DeepSquare GRID Logger.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>input</strong></td>
<td valign="top"><a href="#input-and-output-transportdata">TransportData</a></td>
<td>

Pull data at the start of the job.

It is recommended to set the mode of the data by filling the `inputMode` field.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>inputMode</strong></td>
<td valign="top"><a href="#int">Int</a></td>
<td>

InputMode takes an integer that will be used to change the mode recursively (chmod -R) of the input data.

The number shouldn't be in octal but in decimal. A mode over 512 is not accepted.

Common modes:

- 511 (user:rwx group:rwx world:rwx)
- 493 (user:rwx group:r-x world:r-x)
- 448 (user:rwx group:--- world:---)

If null, the mode won't change and will default to the source.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>steps</strong></td>
<td valign="top">[<a href="#steps-step">Step</a>!]!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>output</strong></td>
<td valign="top"><a href="#input-and-output-transportdata">TransportData</a></td>
<td>

Push data at the end of the job.

Continuous sync/push can be enabled using the `continuousOutputSync` flag.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>continuousOutputSync</strong></td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td>

ContinuousOutputSync will push data during the whole job.

This is useful when it is not desired to lose data when the job is suddenly stopped.

ContinousOutputSync is not available with HTTP.

</td>
</tr>
</tbody>
</table>

<details>
  <summary>Examples</summary>

```json title="Job (minimal)"
{
  "resources": {
    "tasks": 1,
    "gpusPerTask": 0,
    "cpusPerTask": 1,
    "memPerCpu": 1024
  },
  "enableLogging": true,
  "steps": [
    {
      "name": "hello world",
      "run": {
        "command": "echo \"Hello World\""
      }
    }
  ]
}
```

```json title="Job (full)"
{
  "resources": {
    "tasks": 1,
    "gpusPerTask": 0,
    "cpusPerTask": 1,
    "memPerCpu": 1024
  },
  "enableLogging": true,
  "input": {
    "http": {
      "url": "https://my-server/my-file"
    }
  },
  "inputMode": 493,
  "output": {
    "http": {
      "url": "https://transfer.sh"
    }
  },
  "env": [
    {
      "key": "MY_ENV",
      "value": "my_value"
    }
  ],
  "continuousOutputSync": true,
  "steps": [
    {
      "command": "ping 10.0.0.1",
      "resources": {
        "tasks": 1,
        "cpusPerTask": 1,
        "memPerCpu": 500,
        "gpusPerTask": 0
      },
      "network": "slirp4netns",
      "dns": "1.1.1.1",
      "container": {
        "image": "library/ubuntu:latest",
        "registry": "registry-1.docker.io"
      },
      "customNetworkInterfaces": [
        {
          "wireguard": {
            "address": ["10.0.0.2/24"],
            "privateKey": "<TO FILL: Client PK>",
            "peers": [
              {
                "publicKey": "<TO FILL: Serv Pub>",
                "preSharedKey": "<TO FILL: SharedKey>",
                "allowedIPs": ["10.0.0.1/32"],
                "persistentKeepalive": 10,
                "endpoint": "192.168.0.0:51820"
              }
            ]
          }
        }
      ],
      "env": [
        {
          "key": "MY_VAR",
          "value": "myvalue"
        }
      ],
      "mapRoot": false,
      "workDir": "/app",
      "disableCpuBinding": false
    }
  ]
}
```

</details>

### `.resources` _JobResources_

JobResources are the allocated resources for a job in a cluster.

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>tasks</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

Number of tasks which are run in parallel.

Can be greater or equal to 1.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>cpusPerTask</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

Allocated CPUs per task.

Can be greater or equal to 1.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>memPerCpu</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

Allocated memory (MB) per task.

Can be greater or equal to 1.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>gpusPerTask</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

Allocated GPUs per task.

Can be greater or equal to 0.

</td>
</tr>
</tbody>
</table>

<details>
  <summary>Example</summary>

```json title="JobResources"
{
  "tasks": 1,
  "cpusPerTask": 1,
  "memPerCpu": 1000,
  "gpusPerTask": 2
}
```

</details>

### `.env[]` _EnvVar_

An environment variable.

Accessible via: "$key". "Key" name must follows the POSIX specifications (alphanumeric with underscore).

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>key</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>value</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
</tbody>
</table>

<details>
  <summary>Example</summary>

```json title="EnvVar"
{
  "key": "MY_ENV",
  "value": "my value"
}
```

</details>

### `.input` and `.output` _TransportData_

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>http</strong></td>
<td valign="top"><a href="#inputhttp-and-outputhttp-httpdata">HTTPData</a></td>
<td>

Use http to download a file or archive, which will be autoextracted.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>s3</strong></td>
<td valign="top"><a href="#inputs3-and-outputs3-s3data">S3Data</a></td>
<td>

Use s3 to sync a file or directory.

</td>
</tr>
</tbody>
</table>

<details>
  <summary>Examples</summary>

```json title="TransportData (http)"
{
  "http": {
    "url": "https://my-server/my-file"
  }
}
```

```json title="TransportData (s3)"
{
  "s3": {
    "region": "us‑east‑2",
    "bucketUrl": "s3://my-bucket",
    "path": "/",
    "accessKeyId": "accessKeyId",
    "secretAccessKey": "secretAccessKey",
    "endpointUrl": "https://s3.us‑east‑2.amazonaws.com",
    "deleteSync": true
  }
}
```

</details>

### `.input.http` and `.output.http` _HTTPData_

HTTPData describes the necessary variables to connect to a HTTP storage.

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>url</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
</tbody>
</table>

<details>
  <summary>Example</summary>

```json title="HTTPData"
{
  "url": "https://my-server/my-file"
}
```

</details>

### `.input.s3` and `.output.s3` _S3Data_

S3Data describes the necessary variables to connect to a S3 storage.

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>region</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

S3 region. Example: "us‑east‑2".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>bucketUrl</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

The S3 Bucket URL. Must not end with "/".

Example: "s3://my-bucket".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>path</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

The absolute path to a directory/file inside the bucket. Must start with "/".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>accessKeyId</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

An access key ID for the S3 endpoint.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>secretAccessKey</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

A secret access key for the S3 endpoint.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>endpointUrl</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

A S3 Endpoint URL used for authentication. Example: https://s3.us‑east‑2.amazonaws.com

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>deleteSync</strong></td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td>

DeleteSync removes destination files that doesn't correspond to the source.

This applies to any type of source to any type of destination (s3 or filesystem).

See: s5cmd sync --delete.

If null, defaults to false.

</td>
</tr>
</tbody>
</table>

<details>
  <summary>Examples</summary>

```json title="S3Data"
{
  "region": "us‑east‑2",
  "bucketUrl": "s3://my-bucket",
  "path": "/",
  "accessKeyId": "accessKeyId",
  "secretAccessKey": "secretAccessKey",
  "endpointUrl": "https://s3.us‑east‑2.amazonaws.com"
}
```

```json title="S3Data (remove non original data)"
{
  "region": "us‑east‑2",
  "bucketUrl": "s3://my-bucket",
  "path": "/",
  "accessKeyId": "accessKeyId",
  "secretAccessKey": "secretAccessKey",
  "endpointUrl": "https://s3.us‑east‑2.amazonaws.com",
  "deleteSync": true
}
```

</details>

### `.steps[]` _Step_

Step is one instruction.

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>name</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

Name of the instruction.

Is used for debugging.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>dependsOn</strong></td>
<td valign="top">[<a href="#string">String</a>!]</td>
<td>

Depends on wait for async tasks to end before launching this step.

DependsOn uses the `handleName` property of a `StepAsyncLaunch`.

Only steps at the same level can be awaited.

BE WARNED: Uncontrolled `dependsOn` may results in dead locks.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>run</strong></td>
<td valign="top"><a href="#steprun">StepRun</a></td>
<td>

Run a command if not null.

Is exclusive with "for", "launch".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>for</strong></td>
<td valign="top"><a href="#stepfor">StepFor</a></td>
<td>

Run a for loop if not null.

Is exclusive with "run", "launch".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>launch</strong></td>
<td valign="top"><a href="#stepslaunch-stepasynclaunch">StepAsyncLaunch</a></td>
<td>

Launch a background process to run a group of commands if not null.

Is exclusive with "run", "for".

</td>
</tr>
</tbody>
</table>

<details>
  <summary>Examples</summary>

```json title="Step (run)"
{
  "name": "print hello world",
  "run": {
    "command": "echo 'hello world'"
  }
}
```

```json title="Step (for)"
{
  "name": "print 1 2 3",
  "for": {
    "range": {
      "begin": 1,
      "end": 3
    },
    "steps": [
      {
        "name": "print variable",
        "run": {
          "command": "echo \"$index\""
        }
      }
    ]
  }
}
```

</details>

### `.steps[].run` _StepRun_

StepRun is one script executed with the shell.

Shared storage is accessible through the $STORAGE_PATH environment variable.

echo "KEY=value" >> "$DEEPSQUARE_ENV" can be used to share environment variables between steps.

$DEEPSQUARE_INPUT is the path that contains imported files.

$DEEPSQUARE_OUTPUT is the staging directory for uploading files.

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>command</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

Command specifies a shell script.

If container is used, command automatically overwrite the ENTRYPOINT and CMD. If you want to execute the entrypoint, it MUST be re-specified.

You can install and use skopeo to inspect an image without having to pull it.

Example: skopeo inspect --config docker://curlimages/curl:latest will gives "/entrypoint.sh" as ENTRYPOINT and "curl" as CMD. Therefore command="/entrypoint.sh curl".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>shell</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

Shell to use.

Accepted: /bin/bash, /bin/ash, /bin/sh
Default: /bin/sh

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>resources</strong></td>
<td valign="top"><a href="#stepsrunresources-steprunresources">StepRunResources</a></td>
<td>

Allocated resources for the command.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>container</strong></td>
<td valign="top"><a href="#stepsruncontainer-containerrun">ContainerRun</a></td>
<td>

Container definition.

If null, run on the host.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>network</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

Type of core networking functionality.

Either: "host" (default) or "slirp4netns" (rootless network namespace).

Using "slirp4netns" will automatically enables mapRoot.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>dns</strong></td>
<td valign="top">[<a href="#string">String</a>!]</td>
<td>

Configuration for the DNS in "slirp4netns" mode.

ONLY enabled if network is "slirp4netns".

A comma-separated list of DNS IP.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>customNetworkInterfaces</strong></td>
<td valign="top">[<a href="#stepsruncustomnetworkinterfacesnetworkinterface">NetworkInterface</a>!]</td>
<td>

Add custom network interfaces.

ONLY enabled if network is "slirp4netns".

Due to the nature of slirp4netns, the user is automatically mapped as root in order to create network namespaces and add new network interfaces.

The tunnel interfaces will be named net0, net1, ... netX.

The default network interface is tap0, which is a TAP interface connecting the host and the network namespace.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>env</strong></td>
<td valign="top">[<a href="#env-envvar">EnvVar</a>!]</td>
<td>

Environment variables accessible over the command.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>mapRoot</strong></td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td>

Remap UID to root. Does not grant elevated system permissions, despite appearances.

If the "default" (Enroot) container runtime is used, it will use the `--container-remap-root` flags.

If the "apptainer" container runtime is used, the `--fakeroot` flag will be passed.

If no container runtime is used, `unshare --user --map-root-user --mount` will be used and a user namespace will be created.

It is not recommended to use mapRoot with network=slirp4netns, as it will create 2 user namespaces (and therefore will be useless).

If null, default to false.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>workDir</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

Working directory.

If the "default" (Enroot) container runtime is used, it will use the `--container-workdir` flag.

If the "apptainer" container runtime is used, the `--pwd` flag will be passed.

If no container runtime is used, `cd` will be executed first.

If null, default to use $STORAGE_PATH as working directory.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>disableCpuBinding</strong></td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td>

DisableCPUBinding disables process affinity binding to tasks.

Can be useful when running MPI jobs.

If null, defaults to false.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>mpi</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

MPI selection.

Must be one of: none, pmix_v4, pmi2.

If null, will default to infrastructure provider settings (which may not be what you want).

</td>
</tr>
</tbody>
</table>

<details>
  <summary>Examples</summary>

```json title="StepRun (minimal)"
{
  "command": "echo 'hello world'"
}
```

```json title="StepRun (with resource limitation)"
{
  "command": "echo 'hello world'",
  "resources": {
    "tasks": 1,
    "cpusPerTask": 1,
    "memPerCpu": 500,
    "gpusPerTask": 0
  }
}
```

```json title="StepRun (with container)"
{
  "command": "echo 'hello world'",
  "container": {
    "image": "library/ubuntu:latest",
    "registry": "registry-1.docker.io"
  }
}
```

```json title="StepRun (with rootless network namespace)"
{
  "command": "echo 'hello world'",
  "network": "slirp4netns"
}
```

```json title="StepRun (with a Wireguard tunnel)"
{
  "command": "ping 10.0.0.1",
  "network": "slirp4netns",
  "dns": "1.1.1.1",
  "customNetworkInterfaces": [
    {
      "wireguard": {
        "address": ["10.0.0.2/24"],
        "privateKey": "<TO FILL: Client PK>",
        "peers": [
          {
            "publicKey": "<TO FILL: Serv Pub>",
            "preSharedKey": "<TO FILL: SharedKey>",
            "allowedIPs": ["10.0.0.1/32"],
            "persistentKeepalive": 10,
            "endpoint": "192.168.0.0:51820"
          }
        ]
      }
    }
  ]
}
```

```json title="StepRun (full)"
{
  "command": "ping 10.0.0.1",
  "resources": {
    "tasks": 1,
    "cpusPerTask": 1,
    "memPerCpu": 500,
    "gpusPerTask": 0
  },
  "network": "slirp4netns",
  "dns": "1.1.1.1",
  "container": {
    "image": "library/ubuntu:latest",
    "registry": "registry-1.docker.io"
  },
  "customNetworkInterfaces": [
    {
      "wireguard": {
        "address": ["10.0.0.2/24"],
        "privateKey": "<TO FILL: Client PK>",
        "peers": [
          {
            "publicKey": "<TO FILL: Serv Pub>",
            "preSharedKey": "<TO FILL: SharedKey>",
            "allowedIPs": ["10.0.0.1/32"],
            "persistentKeepalive": 10,
            "endpoint": "192.168.0.0:51820"
          }
        ]
      }
    }
  ],
  "env": [
    {
      "key": "MY_VAR",
      "value": "myvalue"
    }
  ],
  "mapRoot": false,
  "workDir": "/app",
  "disableCpuBinding": false
}
```

</details>

### `.steps[].run.resources` _StepRunResources_

StepRunResources are the allocated resources for a command in a job.

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>tasks</strong></td>
<td valign="top"><a href="#int">Int</a></td>
<td>

Number of tasks which are run in parallel.

Can be greater or equal to 1.

If null, default to 1.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>cpusPerTask</strong></td>
<td valign="top"><a href="#int">Int</a></td>
<td>

Allocated CPUs per task.

Can be greater or equal to 1.

If null, defaults to the job resources.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>memPerCpu</strong></td>
<td valign="top"><a href="#int">Int</a></td>
<td>

Allocated memory (MB) per task.

Can be greater or equal to 1.

If null, defaults to the job resources.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>gpusPerTask</strong></td>
<td valign="top"><a href="#int">Int</a></td>
<td>

Allocated GPUs per task.

Can be greater or equal to 0.

If null, defaults to the job resources.

</td>
</tr>
</tbody>
</table>

<details>
  <summary>Examples</summary>

```json title="StepRunResources (partial)"
{
  "tasks": 1,
  "gpusPerTask": 0
}
```

```json title="StepRunResources (full)"
{
  "tasks": 1,
  "cpusPerTask": 1,
  "memPerCpu": 1000,
  "gpusPerTask": 0
}
```

</details>

### `.steps[].run.container` _ContainerRun_

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>image</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

Run the command inside a container with Enroot.

Format: image:tag. Registry and authentication is not allowed on this field.

If the default container runtime is used:

- Use an absolute path to load a squashfs file. By default, it will search inside $STORAGE_PATH. /input will be equivalent to $DEEPSQUARE_INPUT, /output is $DEEPSQUARE_OUTPUT

If apptainer=true:

- Use an absolute path to load a sif file or a squashfs file. By default, it will search inside $STORAGE_PATH. /input will be equivalent to $DEEPSQUARE_INPUT, /output is $DEEPSQUARE_OUTPUT

Examples:

- library/ubuntu:latest
- /my.squashfs

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>mounts</strong></td>
<td valign="top">[<a href="#stepsruncontainermountsmount">Mount</a>!]</td>
<td>

Mount decribes a Bind Mount.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>username</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

Username of a basic authentication.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>password</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

Password of a basic authentication.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>registry</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

Container registry host.

Defaults to registry-1.docker.io

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>apptainer</strong></td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td>

Run with Apptainer as Container runtime instead of Enroot.

By running with apptainer, you get access Deepsquare-hosted images.

When running Apptainer, the container file system is read-only.

Defaults to false.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>deepsquareHosted</strong></td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td>

Use DeepSquare-hosted images.

By setting to true, apptainer will be set to true.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>x11</strong></td>
<td valign="top"><a href="#boolean">Boolean</a></td>
<td>

X11 mounts /tmp/.X11-unix in the container.

</td>
</tr>
</tbody>
</table>

<details>
  <summary>Example</summary>

```json title="ContainerRun (public registry)"
{
  "image": "library/ubuntu:latest",
  "registry": "registry-1.docker.io"
}
```

```json title="ContainerRun (private registry)"
{
  "image": "library/ubuntu:latest",
  "registry": "registry-1.docker.io",
  "username": "john",
  "password": "password"
}
```

```json title="ContainerRun (Apptainer runtime, public registry)"
{
  "image": "library/ubuntu:latest",
  "registry": "registry-1.docker.io",
  "apptainer": true
}
```

```json title="ContainerRun (Apptainer runtime, Deepsquare-Hosted images)"
{
  "image": "library/stable-diffusion:latest",
  "registry": "registry-1.deepsquare.run",
  "apptainer": true,
  "deepsquareHosted": true
}
```

```json title="ContainerRun (x11 mount)"
{
  "image": "library/stable-diffusion:latest",
  "registry": "registry-1.deepsquare.run",
  "x11": true
}
```

```json title="ContainerRun (full example)"
{
  "image": "ubuntu:latest",
  "registry": "registry-1.docker.io",
  "mounts": [
    {
      "hostDir": "/host",
      "containerDir": "/container",
      "options": "ro"
    }
  ],
  "username": "john",
  "password": "password",
  "apptainer": true,
  "deepsquareHosted": false,
  "x11": false
}
```

</details>

### `.steps[].run.container.mounts[]`_Mount_

Mount decribes a Bind Mount.

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>hostDir</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>containerDir</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td></td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>options</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

Options modifies the mount options.

Accepted: ro, rw

</td>
</tr>
</tbody>
</table>

<details>
  <summary>Example</summary>

```json title="Mount"
{
  "hostDir": "/host",
  "containerDir": "/container",
  "options": "ro"
}
```

```json title="Mount-bind tmp (can be used to pass data between steps)"
{
  "hostDir": "/tmp",
  "containerDir": "/tmp",
  "options": "rw"
}
```

</details>

### `.steps[].run.customNetworkInterfaces[]`_NetworkInterface_

Connect a network interface on a StepRun.

The network interface is connected via slirp4netns.

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>wireguard</strong></td>
<td valign="top"><a href="#stepsruncustomnetworkinterfaceswireguard-wireguard">Wireguard</a></td>
<td>

Use the wireguard transport.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>bore</strong></td>
<td valign="top"><a href="#stepsruncustomnetworkinterfacesbore-bore">Bore</a></td>
<td>

Use the bore transport.

</td>
</tr>
</tbody>
</table>

<details>
  <summary>Example</summary>

```json title="NetworkInterface (wireguard)"
{
  "wireguard": {
    "address": ["10.0.0.2/24"],
    "privateKey": "<TO FILL: Client PK>",
    "peers": [
      {
        "publicKey": "<TO FILL: Serv Pub>",
        "preSharedKey": "<TO FILL: SharedKey>",
        "allowedIPs": ["10.0.0.1/32"],
        "persistentKeepalive": 10,
        "endpoint": "192.168.0.0:51820"
      }
    ]
  }
}
```

You can generate a keypair and a shared key with:

```shell title="/bin/sh"
wg genkey | tee privatekey | wg pubkey > pubkey
wg genkey > sharedkey
```

</details>

### `.steps[].run.customNetworkInterfaces[].wireguard` _Wireguard_

Wireguard VPN Transport for StepRun.

The Wireguard VPN can be used as a gateway for the steps. All that is needed is a Wireguard server outside the cluster that acts as a public gateway.

Wireguard transport uses UDP hole punching to connect to the VPN Server.

Disabled settings: PreUp, PostUp, PreDown, PostDown, ListenPort, Table, MTU, SaveConfig.

If these features are necessary, please do contact DeepSquare developpers!

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>address</strong></td>
<td valign="top">[<a href="#string">String</a>!]</td>
<td>

The IP addresses of the wireguard interface.

Format is a CIDRv4 (X.X.X.X/X) or CIDRv6.

Recommendation is to take one IP from the 10.0.0.0/24 range (example: 10.0.0.2/24).

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>privateKey</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

The client private key.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>peers</strong></td>
<td valign="top">[<a href="#stepsruncustomnetworkinterfaceswireguardpeers-wireguardpeer">WireguardPeer</a>!]</td>
<td>

The peers connected to the wireguard interface.

</td>
</tr>
</tbody>
</table>

<details>
  <summary>Example</summary>

```json title="Wireguard"
{
  "address": ["10.0.0.2/24"],
  "privateKey": "<TO FILL: Client PK>",
  "peers": [
    {
      "publicKey": "<TO FILL: Serv Pub>",
      "preSharedKey": "<TO FILL: SharedKey>",
      "allowedIPs": ["10.0.0.1/32"],
      "persistentKeepalive": 10,
      "endpoint": "192.168.0.0:51820"
    }
  ]
}
```

You can generate a keypair and a shared key with:

```shell title="/bin/sh"
wg genkey | tee privatekey | wg pubkey > pubkey
wg genkey > sharedkey
```

</details>

### `.steps[].run.customNetworkInterfaces[].wireguard.peers[]` _WireguardPeer_

A Wireguard Peer.

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>publicKey</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

The peer private key.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>preSharedKey</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The peer pre-shared key.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>allowedIPs</strong></td>
<td valign="top">[<a href="#string">String</a>!]</td>
<td>

Configuration of wireguard routes.

Format is a CIDRv4 (X.X.X.X/X) or CIDRv6.

0.0.0.0/0 (or ::/0) would forward all packets to the tunnel. If you plan to use the Wireguard VPN as a gateway, you MUST set this IP range.

&lt;server internal IP&gt;/32 (not the server's public IP) would forward all packets to the tunnel with the server IP as the destination. MUST be set.

&lt;VPN IP range&gt; would forward all packets to the tunnel with the local network as the destination. Useful if you want peers to communicate with each other and want the gateway to act as a router.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>endpoint</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

The peer endpoint.

Format is IP:port.

This would be the Wireguard server.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>persistentKeepalive</strong></td>
<td valign="top"><a href="#int">Int</a></td>
<td>

Initiate the handshake and re-initiate regularly.

Takes seconds as parameter. 25 seconds is recommended.

You MUST set the persistent keepalive to enables UDP hole-punching.

</td>
</tr>
</tbody>
</table>

<details>
  <summary>Example</summary>

```json title="WireguardPeer"
{
  "publicKey": "<TO FILL: Serv Pub>",
  "preSharedKey": "<TO FILL: SharedKey>",
  "allowedIPs": ["10.0.0.1/32"],
  "persistentKeepalive": 10,
  "endpoint": "192.168.0.0:51820"
}
```

You can generate a keypair and a shared key with:

```shell title="/bin/sh"
wg genkey | tee privatekey | wg pubkey > pubkey
wg genkey > sharedkey
```

</details>

### `.steps[].run.customNetworkInterfaces[].bore` _Bore_

jkuri/bore tunnel Transport for StepRun.

Bore is a proxy to expose TCP sockets.

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>address</strong></td>
<td valign="top"><a href="#string">String</a>!</td>
<td>

Bore server IP/Address.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>port</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

The bore server port.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>targetPort</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

Target port.

</td>
</tr>
</tbody>
</table>

<details>
  <summary>Example</summary>

```json title="Bore"
{
  "address": "bore.deepsquare.run",
  "port": 2200,
  "targetPort": 8080
}
```

Will results in 3 steps.

</details>

### `.steps[].for` _StepFor_

StepFor describes a for loop.

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>parallel</strong></td>
<td valign="top"><a href="#boolean">Boolean</a>!</td>
<td>

Do a parallel for loop. Each iteration is run in parallel.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>items</strong></td>
<td valign="top">[<a href="#string">String</a>!]</td>
<td>

Item accessible via the {{ .Item }} variable. Index accessible via the $item variable.

Exclusive with "range".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>range</strong></td>
<td valign="top"><a href="#stepsforrange-forrange">ForRange</a></td>
<td>

Index accessible via the $index variable.

Exclusive with "items".

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>steps</strong></td>
<td valign="top">[<a href="#steps-step">Step</a>!]!</td>
<td>

Steps are run sequentially in one iteration.

</td>
</tr>
</tbody>
</table>

<details>
  <summary>Example</summary>

```json title="StepFor (items)"
{
  "parallel": true,
  "items": ["a", "b", "c"],
  "steps": [
    {
      "name": "print variable",
      "run": {
        "command": "echo \"$item\""
      }
    }
  ]
}
```

Will prints "a", "b", "c".

```json title="StepFor (range)"
{
  "parallel": true,
  "range": {
    "begin": 1,
    "end": 3
  },
  "steps": [
    {
      "name": "print variable",
      "run": {
        "command": "echo \"$index\""
      }
    }
  ]
}
```

Will prints "1", "2", "3".

</details>

### `.steps[].for.range` _ForRange_

ForRange describes the parameter for a range loop.

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>begin</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

Begin is inclusive.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>end</strong></td>
<td valign="top"><a href="#int">Int</a>!</td>
<td>

End is inclusive.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>increment</strong></td>
<td valign="top"><a href="#int">Int</a></td>
<td>

Increment counter by x count. If null, defaults to 1.

</td>
</tr>
</tbody>
</table>

<details>
  <summary>Example</summary>

```json title="ForRange"
{
  "begin": 1,
  "end": 3
}
```

Will results in 3 steps.

</details>

### `.steps[].launch` _StepAsyncLaunch_

StepAsyncLaunch describes launching a background process.

StepAsyncLaunch will be awaited at the end of the job.

<table>
<thead>
<tr>
<th colspan="2" align="left">Field</th>
<th align="left">Type</th>
<th align="left">Description</th>
</tr>
</thead>
<tbody>
<tr>
<td colspan="2" valign="top"><strong>handleName</strong></td>
<td valign="top"><a href="#string">String</a></td>
<td>

HandleName is the name used to await (dependsOn field of the Step).

Naming style is snake_case. Case is insensitive. No symbol allowed.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>signalOnParentStepExit</strong></td>
<td valign="top"><a href="#int">Int</a></td>
<td>

SignalOnParentStepExit sends a signal to the step and sub-steps when the parent step ends.

This function can be used as a cleanup function to avoid a zombie process.

Zombie processes will continue to run after the main process dies and therefore will not stop the job.

If null, SIGTERM will be sent. If 0, no signal will be sent.

Current signal :

1 SIGHUP Hang-up detected on the control terminal or death of the control process.
2 SIGINT Abort from keyboard
3 SIGQUIT Quit the keyboard
9 SIGKILL If a process receives this signal, it must quit immediately and will not perform any cleaning operations.
15 SIGTERM Software stop signal

It is STRONGLY RECOMMENDED to use SIGTERM to gracefully exit a process. SIGKILL is the most abrupt and will certainly work.

If no signal is sent, the asynchronous step will be considered a fire and forget asynchronous step and will have to terminate itself to stop the job.

WARNING: the "no signal sent" option is subject to removal to avoid undefined behavior. Please refrain from using it.

</td>
</tr>
<tr>
<td colspan="2" valign="top"><strong>steps</strong></td>
<td valign="top">[<a href="#step">Step</a>!]!</td>
<td>

Steps are run sequentially.

</td>
</tr>
</tbody>
</table>

<details>
  <summary>Examples</summary>

```json title="Async-await"
{
  "resources": {
    "tasks": 2,
    "cpusPerTask": 4,
    "memPerCpu": 4096,
    "gpusPerTask": 0
  },
  "steps": [
    {
      "launch": {
        "handleName": "task1",
        "steps": [
          {
            "name": "work",
            "run": {
              "command": "echo \"Working\"; sleep 15; echo \"Working done\""
            }
          }
        ]
      }
    },
    {
      "name": "do-stuff-in-foreground",
      "run": {
        "command": "echo I work fast"
      }
    },
    {
      "name": "wait for task1",
      "dependsOn": ["task1"],
      "run": {
        "command": "echo Task 1 is done!"
      }
    }
  ]
}
```

```json title="Fire-and-forget (explicit safe)"
{
  "resources": {
    "tasks": 2,
    "cpusPerTask": 4,
    "memPerCpu": 4096,
    "gpusPerTask": 0
  },
  "steps": [
    {
      "launch": {
        "signalOnParentStepExit": 15,
        "steps": [
          {
            "name": "work",
            "run": {
              "command": "echo \"Working\"; sleep 15; echo \"Working done\""
            }
          }
        ]
      }
    }
  ]
}
```

```json title="Fire-and-forget (explicit unsafe)"
{
  "resources": {
    "tasks": 2,
    "cpusPerTask": 4,
    "memPerCpu": 4096,
    "gpusPerTask": 0
  },
  "steps": [
    {
      "launch": {
        "signalOnParentStepExit": 0,
        "steps": [
          {
            "name": "work",
            "run": {
              "command": "echo \"Working\"; sleep 15; echo \"Working done\""
            }
          }
        ]
      }
    }
  ]
}
```

</details>

## Scalars

### Boolean

The `Boolean` scalar type represents `true` or `false`.

### Int

The `Int` scalar type represents non-fractional signed whole numeric values. Int can represent values between -(2^31) and 2^31 - 1.

### String

The `String` scalar type represents textual data, represented as UTF-8 character sequences. The String type is most often used to represent free-form human-readable text.

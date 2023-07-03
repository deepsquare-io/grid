# Mounting volumes from the host to the container

:::caution

The bind mounts feature is deprecated and will be removed in the future.

Please use predefined mounts by using the available [environment variables](environment-variables).

:::

## Usage

If you are using a container, you can mount some directories using the `mounts` parameter. Like so:

```json title="Workflow"
{
  "resources": {
    "tasks": 1,
    "gpusPerTask": 0,
    "cpusPerTask": 1,
    "memPerCpu": 1024
  },
  "steps": [
    {
      "name": "map-root",
      "run": {
        "command": "ping 10.0.0.1",
        "container": {
          "image": "ubuntu:latest",
          "registry": "registry-1.docker.io",
          "mounts": [
            {
              "hostDir": "/host",
              "containerDir": "/container",
              "options": "rw"
            }
          ]
        }
      }
    }
  ]
}
```

This feature allows you to mount `/sys/fs/cgroups` as read-only to monitor CPU and memory usage.

## Predefined mounts

DeepSquare automatically mounts a shared file system on all nodes during a job. This shared filesystem only lives during the job, so no worries about cleaning it.

See [environment variables](environment-variables) for the path to the shared storage.

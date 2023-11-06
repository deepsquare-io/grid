# Running the container as root

## MapRoot

Fundamentally, containers uses user namespaces, a feature of the Linux kernel that provide a way to isolate and control user and group permissions within a container or other process. This isolation can help improve security and prevent container breakouts or privilege escalation attacks.

When a container is launched, the user and group IDs used on the host system can be associated with different IDs inside the container.

The Enroot and Apptainer container runtimes can only associate the user ID with the root user, while Podman and Docker can associate the user ID with any ID in the container.

It is possible to enable remap root using the `step.mapRoot` parameter:

```yaml title="Workflow"
resources:
  tasks: 1
  gpusPerTask: 0
  cpusPerTask: 1
  memPerCpu: 1024

steps:
  - name: map-root
    run:
      command: ping 10.0.0.1
      container:
        image: library/ubuntu:latest
        registry: registry-1.docker.io
      mapRoot: true
```

When using Enroot, since the container is not read-only, you can remap the root, which should allow you to write to the container rootfs.

However, it is good practice to:

- Have a read-only container file system.
- Mount the writable volumes if you want data persistence (`$STORAGE_PATH` is already mounted)
- Do not run the container as root

Use **`mapRoot`** with caution.

## Run as another user

Due to the nature of HPC and decentralization, it is not possible to map one user to an other UID.

In the HPC domain, when using Enroot or Apptainer, we are already limited in solutions for matching one user to another UID. One solution is to use a user namespace with a UID/GID mapping, which [only Apptainer supports](https://apptainer.org/docs/admin/latest/user_namespace.html).

However, as we're using LDAP and an LDAP connector, users are created dynamically on the servers. Because of this "dynamic" creation of users, we can't define a UID/GID matching range.

One solution is to statically allocate the entire range, as suggested by [Rootless Containers](https://rootlesscontaine.rs/getting-started/common/subuid/). This solution is obviously not scalable, which is why R&D on this feature has been paused.

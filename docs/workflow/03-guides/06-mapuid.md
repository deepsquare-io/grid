# Mapping the UID of the container

Fundamentally, containers uses user namespaces, a feature of the Linux kernel that provide a way to isolate and control user and group permissions within a container or other process. This isolation can help improve security and prevent container breakouts or privilege escalation attacks.

When a container is launched, the user and group IDs used on the host system can be associated with different IDs inside the container.

We use `unshare` to map from UID to another. Compared to Podman or Docker, we are unable to map a range of UID.

It is possible to enable remap the UID using the `step.mapUid` and `step.mapGid` parameter:

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
      mapUid: 0
      mapGid: 0
```

Note that there is some limitation with the Apptainer container runtime. You cannot run Apptainer with an unknown user, and therefore, you cannot `mapUid` to an unknown user (which only leaves you to map to `root`).

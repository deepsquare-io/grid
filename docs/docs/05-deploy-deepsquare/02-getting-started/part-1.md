---
title: 'Part 1: Containerizing an application'
---

# Containerizing an application

The container act as a runtime environment for the application.

## The application

The application that we are using is a simple OpenMPI "hello world":

```c title="main.c"
/*
 * Copyright (c) 2004-2006 The Trustees of Indiana University and Indiana
 *                         University Research and Technology
 *                         Corporation.  All rights reserved.
 * Copyright (c) 2006      Cisco Systems, Inc.  All rights reserved.
 * Copyright (c) 2023      DeepSquare Association.  All rights reserved.
 *
 * Simple ring test program in C.
 */

#include "mpi.h"
#include <stdio.h>

int main(int argc, char *argv[]) {
  int rank, size, next, prev, message, tag = 201;

  /* Start up MPI */

  MPI_Init(&argc, &argv);
  MPI_Comm_rank(MPI_COMM_WORLD, &rank);
  MPI_Comm_size(MPI_COMM_WORLD, &size);

  /* Calculate the rank of the next process in the ring.  Use the
     modulus operator so that the last process "wraps around" to
     rank zero. */

  next = (rank + 1) % size;
  prev = (rank + size - 1) % size;

  /* If we are the "master" process (i.e., MPI_COMM_WORLD rank 0),
     put the number of times to go around the ring in the
     message. */

  if (0 == rank) {
    message = 10;

    printf("rank 0 sending %d to %d, tag %d (%d processes in ring)\n", message,
           next, tag, size);
    MPI_Send(&message, 1, MPI_INT, next, tag, MPI_COMM_WORLD);
    printf("rank 0 sent to %d\n", next);
  }

  /* Pass the message around the ring.  The exit mechanism works as
     follows: the message (a positive integer) is passed around the
     ring.  Each time it passes rank 0, it is decremented.  When
     each processes receives a message containing a 0 value, it
     passes the message on to the next process and then quits.  By
     passing the 0 message first, every process gets the 0 message
     and can quit normally. */

  while (1) {
    MPI_Recv(&message, 1, MPI_INT, prev, tag, MPI_COMM_WORLD,
             MPI_STATUS_IGNORE);
    printf("rank %d: received the message %d and is passing to the next: %d\n",
           rank, message, next);

    if (0 == rank) {
      --message;
      printf("rank 0 decremented value: %d\n", message);
    }

    MPI_Send(&message, 1, MPI_INT, next, tag, MPI_COMM_WORLD);
    if (0 == message) {
      printf("rank %d exiting\n", rank);
      break;
    }
  }

  /* The last process does one extra send to process 0, which needs
     to be received before the program can exit */

  if (0 == rank) {
    MPI_Recv(&message, 1, MPI_INT, prev, tag, MPI_COMM_WORLD,
             MPI_STATUS_IGNORE);
  }

  /* All done */

  MPI_Finalize();
  return 0;
}

```

We use OpenMPI which is an open-source implementation of the Message Passing Interface (MPI) standard, a way for multiple processes running on separate computers to communicate and coordinate with each other in parallel computing.

This example, when executed, will do the following:

- Rank 0 (or task 0 for slurm, or simply process 0): will send an initial message "10" to the next process.
- The processes pass the message in a circle. If rank 0 receives the message, it will decrease the message by 1.
- If the message is 0, the processes will stop.

## Building and packaging the application using Docker, Podman or Buildah

Container images is built using `Dockerfiles` or `Containerfiles`, by using an OCI container image builder like Buildah, Podman or Docker. We personally recommend Podman since it is rootless and OCI-first.

The container image is built on top of overlay filesystems.

<div style={{textAlign: 'center'}}>

![overlayfs.drawio](./part-1.assets/overlayfs.drawio.svg#invert-on-dark)

</div>

The image has a _base_ image indicated by a `FROM` instruction. Layers can be added using Docker instructions defined in the [Dockerfile reference](https://docs.docker.com/engine/reference/builder/#format).

Instead of starting from `scratch`, we will use the `ghcr.io/deepsquare-io/openmpi:devel` as our base image.

```dockerfile title="Dockerfile"
FROM ghcr.io/deepsquare-io/openmpi:devel
COPY main.c .
RUN mpicc -Wall -Wextra main.c -o main
```

Build the image using:

```shell title="user@~/example/"
# podman build -t <tag> <context>
podman build -t localhost/mpi-circle:latest .
```

## Testing locally

You can test the MPI application locally using:

```shell title="user@~/example/"
# podman run --rm(delete container when done) <tag> <args>
# Arguments: mpirun -np <number of processes> <executable>
podman run --rm localhost/mpi-circle:latest mpirun -np 4 main
```

Which should give something like this:

```log
rank 0 sending 10 to 1, tag 201 (4 processes in ring)
rank 0 sent to 1
rank 1: received the message 10 and is passing to the next: 2
rank 2: received the message 10 and is passing to the next: 3
rank 0: received the message 10 and is passing to the next: 1
rank 0 decremented value: 9
rank 3: received the message 10 and is passing to the next: 0
rank 2: received the message 9 and is passing to the next: 3
rank 1: received the message 9 and is passing to the next: 2
rank 1: received the message 8 and is passing to the next: 2
rank 3: received the message 9 and is passing to the next: 0
rank 3: received the message 8 and is passing to the next: 0
rank 3: received the message 7 and is passing to the next: 0
rank 3: received the message 6 and is passing to the next: 0
rank 1: received the message 7 and is passing to the next: 2
rank 1: received the message 6 and is passing to the next: 2
rank 2: received the message 8 and is passing to the next: 3
rank 2: received the message 7 and is passing to the next: 3
rank 2: received the message 6 and is passing to the next: 3
rank 2: received the message 5 and is passing to the next: 3
rank 0: received the message 9 and is passing to the next: 1
rank 0 decremented value: 8
rank 0: received the message 8 and is passing to the next: 1
rank 0 decremented value: 7
rank 0: received the message 7 and is passing to the next: 1
rank 0 decremented value: 6
rank 0: received the message 6 and is passing to the next: 1
rank 0 decremented value: 5
rank 3: received the message 5 and is passing to the next: 0
rank 3: received the message 4 and is passing to the next: 0
rank 3: received the message 3 and is passing to the next: 0
rank 1: received the message 5 and is passing to the next: 2
rank 1: received the message 4 and is passing to the next: 2
rank 1: received the message 3 and is passing to the next: 2
rank 1: received the message 2 and is passing to the next: 2
rank 1: received the message 1 and is passing to the next: 2
rank 2: received the message 4 and is passing to the next: 3
rank 2: received the message 3 and is passing to the next: 3
rank 2: received the message 2 and is passing to the next: 3
rank 2: received the message 1 and is passing to the next: 3
rank 0: received the message 5 and is passing to the next: 1
rank 0 decremented value: 4
rank 0: received the message 4 and is passing to the next: 1
rank 0 decremented value: 3
rank 0: received the message 3 and is passing to the next: 1
rank 0 decremented value: 2
rank 0: received the message 2 and is passing to the next: 1
rank 0 decremented value: 1
rank 0: received the message 1 and is passing to the next: 1
rank 0 decremented value: 0
rank 0 exiting
rank 2: received the message 0 and is passing to the next: 3
rank 2 exiting
rank 3: received the message 2 and is passing to the next: 0
rank 3: received the message 1 and is passing to the next: 0
rank 3: received the message 0 and is passing to the next: 0
rank 3 exiting
rank 1: received the message 0 and is passing to the next: 2
rank 1 exiting
```

Due to the parallel aspect, the logs are in the wrong order.

## Exporting to a registry

DeepSquare needs a remote registry to fetch your image. You can use the [GitHub Container Registry (ghcr.io)](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry#pushing-container-images), or [Docker Hub (registry-1.docker.io)](https://docs.docker.com/docker-hub/repos/#pushing-a-docker-container-image-to-docker-hub) or self-host your own registry using [Harbor](https://goharbor.io).

To avoid pushing the container image, you can already use our own image: `ghcr.io/deepsquare-io/mpi-example:latest`

You can inspect the image without pulling the image by using [Skopeo](https://github.com/containers/skopeo):

```shell title="user@~/example/"
skopeo inspect --config docker://ghcr.io/deepsquare-io/mpi-example:latest
```

```shell title="skopeo inspect --config docker://ghcr.io/deepsquare-io/mpi-example:latest"
{
    "created": "2023-03-03T13:29:50.795015312Z",
    "architecture": "amd64",
    "os": "linux",
    "config": {
        "Env": [
            "PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin",
            "CFLAGS=-O3 -flto",
            "CXXFLAGS=-O3 -flto"
        ],
        "Cmd": [
            "/bin/bash"
        ],
        "Labels": {
            "io.buildah.version": "1.29.0"
        }
    },
    "rootfs": {
        "type": "layers",
        "diff_ids": [
            "sha256:1d90ccbaa27de4c4a317d5cde657a653b72678d8fc815b669514a6f26a081ffd",
            # ...
            "sha256:5c6c568eca03691fc6c308b83c39eec940a05017e3eed1ef2e7e9a64eca88bff",
            "sha256:7db1f20d29abcf3760b043afe36af4f1efe3925100b9db8539428868a01f56a3"
        ]
    },
    "history": [
        {
            "created": "2023-02-20T19:21:14.683611105Z",
            "created_by": "/bin/sh -c #(nop) ADD file:d511fb4604410f5ccb4804cc7d40512ecd13b8e3e9494c873da424cabbb80020 in / "
        },
        # ...
        {
            "created": "2023-03-03T13:29:50.43538223Z",
            "created_by": "/bin/sh -c #(nop) COPY file:1869eee0b16c5f9b95d9a84e7f5e40ce5010c9c9395a700151f8217f42f6590c in . ",
            "comment": "FROM 20252f3746b4"
        },
        {
            "created": "2023-03-03T13:29:50.795587473Z",
            "created_by": "/bin/sh -c mpicc -Wall -Wextra main.c -o main",
            "comment": "FROM 87f9bb706013"
        }
    ]
}
```

## Next steps

You learned how to containerise your application into a container image and publish it.

Next, you're going to run it on the DeepSquare GRID, the decentralized HPC cloud.

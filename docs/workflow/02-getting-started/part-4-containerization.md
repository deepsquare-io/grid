---
title: 'Part 4: Containerization'
---

Welcome to the fourth part of our series. In this section, we'll delve into containerization, an efficient method for packaging applications alongside dependencies. This ensures your application runs seamlessly across different computational environments.

# Containerizing an application

Containerization is a method of software deployment where an application and its dependencies are packaged together as a container. It ensures that the application works uniformly despite differences in infrastructures. The container act as a runtime environment for the application.

## Our Application: An OpenMPI "hello world"

We'll use the same OpenMPI program from [Part 2](part-2-openmpi) which illustrates basic message passing between multiple processes:

<details>

<summary>OpenMPI Code</summary>

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

int main(int argc, char \*argv[]) {
int rank, size, next, prev, message, tag = 201;

/_ Start up MPI _/

MPI_Init(&argc, &argv);
MPI_Comm_rank(MPI_COMM_WORLD, &rank);
MPI_Comm_size(MPI_COMM_WORLD, &size);

/_ Calculate the rank of the next process in the ring. Use the
modulus operator so that the last process "wraps around" to
rank zero. _/

next = (rank + 1) % size;
prev = (rank + size - 1) % size;

/_ If we are the "master" process (i.e., MPI_COMM_WORLD rank 0),
put the number of times to go around the ring in the
message. _/

if (0 == rank) {
message = 10;

    printf("rank 0 sending %d to %d, tag %d (%d processes in ring)\n", message,
           next, tag, size);
    MPI_Send(&message, 1, MPI_INT, next, tag, MPI_COMM_WORLD);
    printf("rank 0 sent to %d\n", next);

}

/_ Pass the message around the ring. The exit mechanism works as
follows: the message (a positive integer) is passed around the
ring. Each time it passes rank 0, it is decremented. When
each processes receives a message containing a 0 value, it
passes the message on to the next process and then quits. By
passing the 0 message first, every process gets the 0 message
and can quit normally. _/

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

/_ The last process does one extra send to process 0, which needs
to be received before the program can exit _/

if (0 == rank) {
MPI_Recv(&message, 1, MPI_INT, prev, tag, MPI_COMM_WORLD,
MPI_STATUS_IGNORE);
}

/_ All done _/

MPI_Finalize();
return 0;
}

```

</details>

Copy the code in your workspace and name it `main.c`.

## Building The Application: Docker, Podman or Buildah

To containerize our application, we use tools such as _Docker_, _Podman_ or _Buildah_ to construct a container image. For this tutorial, we'll use Podman due to its rootless and OCI-first nature.

**Docker:** Docker is a popular open-source platform that allows developers to automate the deployment, scaling, and management of applications. It does this by creating lightweight, portable, self-sufficient containers from any application.

**Podman:** Podman (short for Pod Manager) is a container engine developed by Red Hat for developing, managing, and running containers and pods. It's designed to be a drop-in replacement for Docker while providing some additional features. Unlike Docker, Podman doesn't require a daemon to function and is rootless by default, making it more secure and easy to use.

**Buildah:** Buildah is a tool designed to help developers to create, build, and update container images. It supports Dockerfile but also provides more flexibility in building images, allowing you to build images without a Dockerfile.

Container images are built using `Dockerfiles` or `Containerfiles`.

The container image is built on top of overlay filesystems.

<div style={{textAlign: 'center'}}>

![overlayfs.drawio](./part-1.assets/overlayfs.drawio.svg#invert-on-dark)

</div>

The image has a _base_ image indicated by a `FROM` instruction. Layers can be added using Docker instructions defined in the [Dockerfile reference](https://docs.docker.com/engine/reference/builder/#format).

Instead of starting from `scratch`, we will use the `ghcr.io/deepsquare-io/openmpi:devel` as our base image. Create a `Dockerfile` in your workspace and add this:

```dockerfile title="Dockerfile"
FROM ghcr.io/deepsquare-io/openmpi:devel
COPY main.c .
RUN mpicc -Wall -Wextra main.c -o main
```

This Dockerfile copies the `main.c` in the container image and compiles it. `mpicc` is used to compile the program which is provided by the base image.

To build the image, follow these steps:

- Open a terminal or command prompt.
- Navigate to the directory containing the Dockerfile and main.c files.
- Run the following command to build the image:

```shell title="user@~/example/"
# podman build -t <tag> <context>
podman build -t localhost/mpi-circle:latest .
```

- The `-t` option is used to specify the tag or name for the image.
- `localhost/mpi-circle:latest` is the tag we chose for the image.

Once the command completes, Podman will build the container image based on the instructions in the Dockerfile and the contents of the current directory. The resulting image will be tagged as `localhost/mpi-circle:latest`.

## Testing The Application Locally

You can easily test the MPI application locally using Podman:

```shell title="user@~/example/"
# podman run --rm(delete container when done) <tag> <args>
# Arguments: mpirun -np <number of processes> <executable>
podman run --rm localhost/mpi-circle:latest mpirun -np 4 main
```

The output from this test should resemble:

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

## Publishing Your Application

Next, we'll publish our containerized application. You have various options for hosting your container image, such as [GitHub Container Registry (ghcr.io)](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry#pushing-container-images), [Docker Hub (registry-1.docker.io)](https://docs.docker.com/docker-hub/repos/#pushing-a-docker-container-image-to-docker-hub), or you can even host your own registry using [Harbor](https://goharbor.io). To inspect the image without pulling it, we'll use Skopeo, a handy tool for working with remote container images.

To avoid pushing the container image, you can already use our own image: `ghcr.io/deepsquare-io/mpi-example:latest`

If you want to publish your own image, you must tag your image correctly, login to the registry and push it:

```shell title="user@~/example/"
podman tag localhost/mpi-circle:latest ghcr.io/my-user/my-image:latest
```

This command tags the local image (`localhost/mpi-circle:latest`) with your desired tag (`ghcr.io/my-user/my-image:latest`).
Replace `my-user` with your GitHub username or organization, and `my-image` with the desired name for your image.

```shell title="user@~/example/"
podman login ghcr.io
```

This command logs you in to the GitHub Container Registry (`ghcr.io`) using Podman.
You need to provide your GitHub username and a personal access token for authentication.

Run the following command to push your image to the registry:

```shell title="user@~/example/"
podman push ghcr.io/my-user/my-image:latest
```

This command pushes the tagged image (`ghcr.io/my-user/my-image:latest`) to the GitHub Container Registry.
Replace `my-user` with your GitHub username or organization, and `my-image` with the desired name for your image.

Once the command completes, your container image will be published and available in the specified registry for others to use.

## Next steps

Great job on successfully containerizing your application! You've successfully containerized the program, which can now be used on DeepSquare.

Our next step is to optimize our HPC workloads for multi-node distributed training. In Part 5, we'll harness the power of supercomputers in the DeepSquare ecosystem, adapting our machine learning example with Horovod, and revising our resource allocation. Stay tuned!

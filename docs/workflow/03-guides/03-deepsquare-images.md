# Using pre-cached DeepSquare-hosted images

The images hosted by DeepSquare are unpacked container images distributed on a read-only file system called CernVM-FS.

CernVS-FS is a distributed read-only file system designed to provide software to compute nodes. One of the main advantages of CVMFS is its ability to globally distribute software and data, making it ideal for DeepSquare where many users need access to the same resources. It also has built-in caching and content distribution capabilities, which significantly reduces network traffic and improves download speeds.

Thanks to the nature of CVMFS, unpacked container image files are loaded dynamically, which significantly reduces network traffic. DeepSquare engineers can unpacked the container files by using CVMFS DUCC. This is extremely useful when container images are large, such as CUDA images that weigh 10GB or more.

## Availables images

| Registry                  | Image                           |
| ------------------------- | ------------------------------- |
| registry-1.deepsquare.run | library/tdp:latest              |
| registry-1.deepsquare.run | library/upscaling:latest        |
| registry-1.deepsquare.run | library/stable-diffusion:latest |
| registry-1.deepsquare.run | library/openfoam:v2212          |
| registry-1.deepsquare.run | library/pytorch:23.01-py3       |
| registry-1.docker.io      | linuxserver/blender:3.4.1       |
| registry-1.docker.io      | linuxserver/blender:3.3.1       |
| registry-1.docker.io      | linuxserver/blender:3.2.2       |
| registry-1.docker.io      | linuxserver/blender:3.1.2       |
| registry-1.docker.io      | horovod/horovod:latest          |
| registry-1.docker.io      | mathworks/matlab:r2022b         |

## Usage

Apptainer is the only container runtime that can run a unpacked container image, so be sure to enable `container.apptainer=true`.

```yaml title="Workflow"
resources:
  tasks: 1
  gpusPerTask: 0
  cpusPerTask: 1
  memPerCpu: 1024

steps:
  - name: hello world
    run:
      command: echo "Hello World"
      container:
        image: library/tdp:latest
        registry: registry-1.deepsquare.run
        deepsquareHosted: true
        apptainer: true
```

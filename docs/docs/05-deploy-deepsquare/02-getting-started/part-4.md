---
title: 'Part 4: Deploying multi nodes workload'
---

# Multi tasks workload

## MPI and the job scheduler

Since DeepSquare is using Slurm as the job scheduler on the infrastructure provider side, we are able to send MPI jobs across multiple nodes

## Adapting the MPI example for multi tasking

```json
{
  "resources": {
    "tasks": 32,
    "gpusPerTask": 0,
    "cpusPerTask": 1,
    "memPerCpu": 1024
  },
  "enableLogging": true,
  "steps": [
    {
      "name": "run the circle program",
      "run": {
        "command": "./main",
        "workDir": "/app",
        "resources": {
          "tasks": 32
        },
        "mpi": "pmix_v4",
        "container": {
          "image": "deepsquare-io/mpi-example:latest",
          "registry": "ghcr.io",
          "apptainer": true
        }
      }
    }
  ]
}
```

## Adapting the ML example for multi tasking

```json title="Workflow"
{
  "resources": {
    "tasks": 2,
    "gpusPerTask": 2,
    "cpusPerTask": 8,
    "memPerCpu": 2048
  },
  "enableLogging": true,
  "input": {
    "http": {
      "url": "https://www.cs.toronto.edu/~kriz/cifar-10-python.tar.gz"
    }
  },
  "output": {
    "s3": {
      "region": "us‑east‑2",
      "bucketUrl": "s3://my-bucket",
      "path": "/",
      "accessKeyId": "accessKeyId",
      "secretAccessKey": "secretAccessKey",
      "endpointUrl": "https://s3.us‑east‑2.amazonaws.com",
      "deleteSync": true
    }
  },
  "continuousOutputSync": true,
  "steps": [
    {
      "command": "python train.py",
      "resources": {
        "tasks": 2
      },
      "mpi": "pmix_v4",
      "container": {
        "image": "deepsquare-io/cifar-10:latest",
        "registry": "ghcr.io"
      },
      "workDir": "/app"
    }
  ]
}
```

---
title: 'Part 2: Writing a workflow file'
---

# Writing a workflow file

You containerized the MPI example application. In this part, you will now run it on the DeepSquare GRID.

A workflow file is a JSON file that describes the resources allocation and the suite of instructions.

## Understanding the workflow format

We recommend to use the [DeepSquare sandbox page](https://app.deepsquare.run/sandbox) to write workflow file and the [reference API](/docs/deploy-deepsquare/workflow-api-reference/job).

The sandbox should be filled with:

```json
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

## Writing the workflow file

We will use 32 tasks (= 32 processes in parallel spread on multiple nodes), 1 cpu per task, 1024 MB of RAM per cpu and no gpu.

We can map easily the "podman run" arguments to the workflow file. We have:

```json
{
  "resources": {
    "tasks": 32,
    "gpusPerTask": 0,
    "cpusPerTask": 1,
    "memPerCpu": 1024
  },
  "enableLogging": true,
  "env": [
    {
      "key": "OMPI_MCA_btl_vader_single_copy_mechanism",
      "value": "none"
    },
    {
      "key": "OMPI_MCA_pml",
      "value": "ob1"
    },
    {
      "key": "OMPI_MCA_btl",
      "value": "vader,tcp,self"
    },
    {
      "key": "OMPI_MCA_btl_tcp_if_include",
      "value": "ib0"
    }
  ],
  "steps": [
    {
      "name": "run the circle program",
      "run": {
        "workDir": "/project",
        "resources": {
          "tasks": 32
        },
        "command": "./main",
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

We specify the resource allocation using the `resources` block and use these resources during the steps. The `run the circle program` step uses 32 tasks and implicitly 1 cpu per task and 1024 MB of memory. Not specifying `run.resources.tasks` would only start a single process.

By enabling `enableLogging`, you authorize the application to send logs to the DeepSquare logging system, which you can read on the [DeepSquare GRID Portal](https://app.deepsquare.run).

Because MPI is already integrated with SLURM, it is not necessary to run `mpirun`.

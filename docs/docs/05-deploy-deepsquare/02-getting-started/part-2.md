---
title: 'Part 2: Writing a workflow file'
---

# Writing a workflow file

You containerized the MPI example application. In this part, you will now run it on the DeepSquare GRID.

A workflow file is a JSON file that describes the resources allocation and the suite of instructions necessary to run your application.

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
  "job": {
    "resources": {
      "tasks": 1,
      "gpusPerTask": 0,
      "cpusPerTask": 16,
      "memPerCpu": 1024
    },
    "enableLogging": true,
    "steps": [
      {
        "name": "run the circle program",
        "run": {
          "command": "mpirun -np 16 ./main",
          "workDir": "/app",
          "container": {
            "image": "deepsquare-io/mpi-example:latest",
            "registry": "ghcr.io",
            "apptainer": true
          }
        }
      }
    ]
  }
}
```

We specify the resource allocation using the `resources` block and use these resources during the steps. The `run the circle program` step uses 32 tasks and implicitly 1 cpu per task and 1024 MB of memory. Not specifying `run.resources.tasks` would only start a single process.

By enabling `enableLogging`, you authorize the application to send logs to the DeepSquare logging system, which you can read on the [DeepSquare GRID Portal](https://app.deepsquare.run).

## Next steps

You've learned how to launch workloads on the DeepSquare grid! We've done a simple fire-and-forget, but you may want to persist some data.

Data scientists often train their machine learning neural network models on HPCs and need to input training data and output model control points.

So you will learn how to send a dataset and get results.

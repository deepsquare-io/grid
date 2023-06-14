# Creating and Running Custom Workflows

With the [Dev](https://app.deepsquare.run/sandbox) application in DeepSquare, you can create and run custom workflows to harness the full potential of high-performance computing. This guide will walk you through the process of writing, creating, and running workflow files.

## Prerequisites

Before getting started, ensure you have access to the DeepSquare portal and the [Dev](https://app.deepsquare.run/sandbox) application.

## Writing a Workflow File

A workflow file is a JSON-formatted file that defines resource allocation and execution instructions for your tasks. Here's a basic structure of a workflow file:

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

To learn more about Workflow files, read [Part 2: Writing a workflow file](/workflow/getting-started/part-2).

## Running a Workflow in the Dev Application

1. Log in to the DeepSquare portal and navigate to the Dev application.
2. In the **Workflow Editor**, write or paste your workflow JSON.
3. If your workflow is valid, click **Submit** to run the job.

## Downloading Results and Review Logs

1. In the **Job Status** page, locate your job based on its name.
2. Monitor the progress of your job by checking its status.
3. Click on the **Logs** button in order to read the realtime logs.
4. Once your job completes, you can view the results, logs, and other output files in the designated output storage.

With DeepSquare's Dev application, you can efficiently create, manage, and run custom workflows to cater to a wide range of high-performance computing demands.

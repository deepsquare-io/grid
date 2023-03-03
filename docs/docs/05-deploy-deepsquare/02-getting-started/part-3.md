---
title: 'Part 3: Persisting Data'
---

# Persisting Data

Let's do some basic Machine Learning! We'll send a small dataset and receive checkpoints.

## S3 and HTTP transport

DeepSquare does NOT store data. Only you can manage your data as you wish. However, DeepSquare will need access to your data source to pull and push data.

DeepSquare supports S3 and HTTP data sources. We'll use HTTP to pull a public dataset, and use S3 to store our data.

The dataset is CIFAR-10 which is downloadable at this [address](https://www.cs.toronto.edu/~kriz/cifar-10-python.tar.gz). It is composed of 60000 32x32 colored images in 10 classes, with 6000 images per class. There are 5 training batches (50000 images) and 1 test batch (10000 images).

## The ML model

We'll train a [Deep Layer Aggregation Model](https://arxiv.org/abs/1707.06484)

TODO

```

```

## Setup an S3 quickly

TODO

## Writing the workflow file

```json title="Workflow"
{
  "resources": {
    "tasks": 1,
    "gpusPerTask": 2,
    "cpusPerTask": 2,
    "memPerCpu": 1024
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
      "container": {
        "image": "deepsquare-io/cifar-10:latest",
        "registry": "ghcr.io"
      },
      "workDir": "/app"
    }
  ]
}
```

TODO

## Next steps

You've learned how to run an ML workload on DeepSquare! What if we did it much faster?

You've run the job on a single task (process), but to take advantage of HPC and DeepSquare, you need to run it on multiple processes and multiple nodes. So, next time, you'll learn how to send MPI jobs to the cluster.

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

## Setup an S3

Using S3 is the standard for data storage, especially for control files. S3 is an [object storage](https://aws.amazon.com/what-is/object-storage/), which is a technology for storing data in an unstructured format, and therefore allows the storage to scale up by storing data on multiple devices.

If you wish to self-host your own S3 server, we recommend [Garage](https://garagehq.deuxfleurs.fr) (small to medium scale) or [MinIO](https://min.io) (medium to large scale).

If you don't want to bother setting up your own storage, you can use [Exoscale Simple Object Storage](https://www.exoscale.com/object-storage/) (~0.0198 per GB/month) or [Google Cloud Storage](https://cloud.google.com/storage) (~0.023 per GB/month).

If you successfully deployed your S3 server and created your bucket, you should fetch the API access keys to the S3 storage.

| Key               | Description                                           | Example                                                      |
| ----------------- | ----------------------------------------------------- | ------------------------------------------------------------ |
| Region            | The region of the S3 object storage.                  | us‑east‑2                                                    |
| Endpoint URL      | The URL to the S3 API, which should start with https. | [https://s3.us‑east‑2.amazonaws.com](https://s3.us‑east‑2.amazonaws.com) |
| Bucket URL        | The S3 bucket URL used to fetch data.                 | s3://my-bucket                                               |
| Path              | The path relative to the bucket root.                 | /my-directory                                                |
| Access Key ID     | The access key ID of the API access.                  | Varies with the host: starts with AKIA (amazon) or EXO (exoscale) |
| Secret Access Key | The password of the API access.                       | ***                                                          |



## Writing the workflow file

```json title="Workflow"
{
  "resources": {
    "tasks": 1,
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

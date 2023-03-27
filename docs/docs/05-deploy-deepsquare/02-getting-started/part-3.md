---
title: 'Part 3: Persisting Data'
---

# Persisting Data

Let's do some basic Machine Learning! We'll use a dataset and receive checkpoints.

## S3 and HTTP transport

DeepSquare does NOT store data. Only you can manage your data as you wish. However, DeepSquare will need access to your data source to pull and push data.

DeepSquare supports both S3 and HTTP data sources. We will use HTTP to fetch a public dataset and use S3 to push our data.

The dataset is CIFAR-10 and can be downloaded at this [address](https://www.cs.toronto.edu/~kriz/cifar-10-python.tar.gz). It consists of 60000 32×32 colored images in 10 classes, with 6000 images per class. There are 5 training batches (50000 images) and 1 test batch (10000 images).

## The ML model

We'll train a [Deep Layer Aggregation Model](https://arxiv.org/abs/1707.06484). The source code is located [here](https://github.com/deepsquare-io/cifar-10-example) and the Docker image is available at `ghcr.io/deepsquare-io/cifar-10-example:latest`.

<div style={{textAlign: 'center'}}>

![image-20230308175425478](./part-3.assets/image-20230308175425478.png#invert-on-dark)

</div>

<details>
  <summary>Source code</summary>

```python title="model.py (PyTorch Model)"
import torch
import torch.nn.functional as F
from torch import nn


class BasicBlock(nn.Module):
    expansion = 1

    def __init__(self, in_planes, planes, stride=1):
        super(BasicBlock, self).__init__()
        self.conv1 = nn.Conv2d(
            in_planes, planes, kernel_size=3, stride=stride, padding=1, bias=False
        )
        self.bn1 = nn.BatchNorm2d(planes)
        self.conv2 = nn.Conv2d(
            planes, planes, kernel_size=3, stride=1, padding=1, bias=False
        )
        self.bn2 = nn.BatchNorm2d(planes)

        self.shortcut = nn.Sequential()
        if stride != 1 or in_planes != self.expansion * planes:
            self.shortcut = nn.Sequential(
                nn.Conv2d(
                    in_planes,
                    self.expansion * planes,
                    kernel_size=1,
                    stride=stride,
                    bias=False,
                ),
                nn.BatchNorm2d(self.expansion * planes),
            )

    def forward(self, x):
        out = F.relu(self.bn1(self.conv1(x)))
        out = self.bn2(self.conv2(out))
        out += self.shortcut(x)
        out = F.relu(out)
        return out


class Root(nn.Module):
    def __init__(self, in_channels, out_channels, kernel_size=1):
        super(Root, self).__init__()
        self.conv = nn.Conv2d(
            in_channels,
            out_channels,
            kernel_size,
            stride=1,
            padding=(kernel_size - 1) // 2,
            bias=False,
        )
        self.bn = nn.BatchNorm2d(out_channels)

    def forward(self, xs: list[torch.Tensor]):
        x = torch.cat(xs, 1)
        out = F.relu(self.bn(self.conv(x)))
        return out


class Tree(nn.Module):
    def __init__(self, block, in_channels, out_channels, level=1, stride=1):
        super(Tree, self).__init__()
        self.root = Root(2 * out_channels, out_channels)
        if level == 1:
            self.left_tree = block(in_channels, out_channels, stride=stride)
            self.right_tree = block(out_channels, out_channels, stride=1)
        else:
            self.left_tree = Tree(
                block, in_channels, out_channels, level=level - 1, stride=stride
            )
            self.right_tree = Tree(
                block, out_channels, out_channels, level=level - 1, stride=1
            )

    def forward(self, x):
        out1 = self.left_tree(x)
        out2 = self.right_tree(out1)
        out = self.root([out1, out2])
        return out


class SimpleDLA(nn.Module):
    def __init__(self, block=BasicBlock, num_classes=10):
        super(SimpleDLA, self).__init__()
        self.base = nn.Sequential(
            nn.Conv2d(3, 16, kernel_size=3, stride=1, padding=1, bias=False),
            nn.BatchNorm2d(16),
            nn.ReLU(True),
        )

        self.layer1 = nn.Sequential(
            nn.Conv2d(16, 16, kernel_size=3, stride=1, padding=1, bias=False),
            nn.BatchNorm2d(16),
            nn.ReLU(True),
        )

        self.layer2 = nn.Sequential(
            nn.Conv2d(16, 32, kernel_size=3, stride=1, padding=1, bias=False),
            nn.BatchNorm2d(32),
            nn.ReLU(True),
        )

        self.layer3 = Tree(block, 32, 64, level=1, stride=1)
        self.layer4 = Tree(block, 64, 128, level=2, stride=2)
        self.layer5 = Tree(block, 128, 256, level=2, stride=2)
        self.layer6 = Tree(block, 256, 512, level=1, stride=2)
        self.linear = nn.Linear(512, num_classes)

    def forward(self, x):
        out = self.base(x)
        out = self.layer1(out)
        out = self.layer2(out)
        out = self.layer3(out)
        out = self.layer4(out)
        out = self.layer5(out)
        out = self.layer6(out)
        out = F.avg_pool2d(out, 4)
        out = out.view(out.size(0), -1)
        out = self.linear(out)
        return out

```

</details>

<details>
<summary>Visualization</summary>

![model.pt](./part-3.assets/model.pt.svg#invert-on-dark)

</details>

## Setup an S3

Using S3 is the standard for data storage, especially for control files. S3 is an [object storage](https://aws.amazon.com/what-is/object-storage/), which is a technology for storing data in an unstructured format, and therefore allows the storage to scale up by storing data on multiple devices.

If you wish to self-host your own S3 server, we recommend [Garage](https://garagehq.deuxfleurs.fr) (small to medium scale) or [MinIO](https://min.io) (medium to large scale).

If you don't want to bother setting up your own storage, you can use [Exoscale Simple Object Storage](https://www.exoscale.com/object-storage/) (~0.0198 per GB/month) or [Google Cloud Storage](https://cloud.google.com/storage) (~0.023 per GB/month).

If you successfully deployed your S3 server and created your bucket, you should fetch the API access keys to the S3 storage.

| Key               | Description                                           | Example                                                                  |
| ----------------- | ----------------------------------------------------- | ------------------------------------------------------------------------ |
| Region            | The region of the S3 object storage.                  | us‑east‑2                                                                |
| Endpoint URL      | The URL to the S3 API, which should start with https. | [https://s3.us‑east‑2.amazonaws.com](https://s3.us‑east‑2.amazonaws.com) |
| Bucket URL        | The S3 bucket URL used to fetch data.                 | s3://my-bucket                                                           |
| Path              | The path relative to the bucket root.                 | /my-directory                                                            |
| Access Key ID     | The access key ID of the API access.                  | Varies with the host: starts with AKIA (amazon) or EXO (exoscale)        |
| Secret Access Key | The password of the API access.                       | \*\*\*                                                                   |

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
      "region": "at-vie-1",
      "bucketUrl": "s3://cifar10",
      "path": "/",
      "accessKeyId": "EXO***",
      "secretAccessKey": "***",
      "endpointUrl": "https://sos-at-vie-1.exo.io"
    }
  },
  "continuousOutputSync": true,
  "steps": [
    {
      "name": "train",
      "run": {
        "command": "/.venv/bin/python3 main.py --checkpoint_out=$DEEPSQUARE_OUTPUT/ckpt.pth --dataset=$DEEPSQUARE_INPUT/",
        "container": {
          "image": "deepsquare-io/cifar-10-example:latest",
          "registry": "ghcr.io"
        },
        "workDir": "/app"
      }
    }
  ]
}
```

The `input` option will automatically unpack the archive.

The `main.py` script will produce a `ckpt.pth` checkpoint file every time. Using `continuousOutputSync`, the checkpoint will be uploaded each time the file is updated. You won't lose progress and can resume it later using `--checkpoint_in=$DEEPSQUARE_INPUT/ckpt.pth`, but you need to set up an S3 as input. One solution would be to use an `init` container:

```json title="Workflow with resume checkpoint"
{
  "resources": {
    "tasks": 1,
    "gpusPerTask": 2,
    "cpusPerTask": 8,
    "memPerCpu": 2048
  },
  "enableLogging": true,
  "input": {
    "s3": {
      "region": "at-vie-1",
      "bucketUrl": "s3://cifar10",
      "path": "/",
      "accessKeyId": "EXO284cde16bdbe4195b8fc4763",
      "secretAccessKey": "KYReUpY-8ipfAvO5wlYpd7Uq-IkadN9ac535H-C1mbI",
      "endpointUrl": "https://sos-at-vie-1.exo.io"
    }
  },
  "output": {
    "s3": {
      "region": "at-vie-1",
      "bucketUrl": "s3://cifar10",
      "path": "/",
      "accessKeyId": "EXO284cde16bdbe4195b8fc4763",
      "secretAccessKey": "KYReUpY-8ipfAvO5wlYpd7Uq-IkadN9ac535H-C1mbI",
      "endpointUrl": "https://sos-at-vie-1.exo.io"
    }
  },
  "continuousOutputSync": true,
  "steps": [
    {
      "name": "download dataset",
      "run": {
        "command": "curl -fsSL https://www.cs.toronto.edu/~kriz/cifar-10-python.tar.gz -o $STORAGE_PATH/cifar-10-python.tar.gz; tar -C $STORAGE_PATH -xvzf $STORAGE_PATH/cifar-10-python.tar.gz; ls -lah $STORAGE_PATH",
        "container": {
          "image": "curlimages/curl:latest",
          "registry": "registry-1.docker.io"
        }
      }
    },
    {
      "name": "train",
      "run": {
        "command": "/.venv/bin/python3 main.py --checkpoint_in=$DEEPSQUARE_INPUT/ckpt.pth --checkpoint_out=$DEEPSQUARE_OUTPUT/ckpt.pth --dataset=$STORAGE_PATH/",
        "container": {
          "image": "deepsquare-io/cifar-10-example:latest",
          "registry": "ghcr.io"
        },
        "workDir": "/app"
      }
    }
  ]
}
```

_(By the way, since we use PyTorch DataLoader, we could have downloaded the dataset directly from the code)._

## Next steps

You've learned how to run an ML workload on DeepSquare! What if we did it much faster?

You've run the job on a single task (process), but to take advantage of HPC and DeepSquare, you need to run it on multiple processes and multiple nodes. So, next, you'll learn how to send MPI jobs to the cluster.

# DeepSquare Overview

DeepSquare is a decentralized high performance computing (HPC) cloud. DeepSquare allows you to run your application on supercomputers around the world without having to worry about the complex underlying infrastructure. With DeepSquare, you can run compute-intensive workloads on the DeepSquare Grid, the interoperability layer made of decentralized heterogeneous supercomputers. By leveraging DeepSquare, you can develop, build, and scale applications with minimal restrictions and permissions due to vendor-locked infrastructure.

## Running on DeepSquare

DeepSquare provides the ability to run a workload on any supercomputer connected to the DeepSquare Grid using a combination of job scheduling, containers and Web3 technologies. Using job scheduling, jobs are intelligently distributed across infrastructure providers. Containers solve compatibility issues with the heterogeneity of different clusters. And Web3 provides true transparency, availability and scale for a global job scheduler.

## What can I use DeepSquare for?

**Fast and easy deployment of HPC workloads**

DeepSquare accelerates the deployment lifecycle by always providing a cluster HPC-ready.

**Scaling and clever scheduling**

DeepSquare's job scheduler uses several criteria to meet the developer's demand. Whether testing on a small server or running a full HPC infrastructure, there is always a choice.

The DeepSquare Grid's nature also facilitates to dynamically manage workloads, scaling up or down applications as business needs dictate, in near real time.

**Intelligent and transparent billing**

TODO: about billing

## The Job Scheduling Architecture

DeepSquare uses something similar to a Web-Queue-Worker style of architecture.

Using Web3, DeepSquare's smart contracts combine the "Web-Queue" by providing:

- An identity provider (the user's wallet)
- A persistent database
- A job queue
- A consistent billing system
- A common API

The infrastructure provider uses a worker, which we call the **Supervisor** of jobs that listen to the smart contract queue. The infrastructure provider itself uses another job scheduler called **Slurm** to manage multi-node workloads.

To avoid storing private data on the blockchain, the sbatch API is the intermediary between the Supervisor and the Client. The sbatch API converts a workflow defined by strict specifications into a slurm batch (SBatch) script and stores it temporarily, waiting for the supervisor to claim a job and pull the sbatch script.

![Deepsquare Architecture.drawio](./01-overview.assets/deepsquare-architecture.drawio.svg#invert-on-dark)

If you want to try an already existing client, there is the [DeepSquare Grid Portal](https://app.deepsquare.run). You can also develop your own client using the [DeepSquare SDK](TODO: reference).

## The underlying technology

DeepSquare is composed of a large number of technologies to achieve such a solution:

- The [Solidity programming language](https://docs.soliditylang.org/en/latest/) to write the smart-contracts.
- The [Go](https://go.dev) and [Python](https://www.python.org) programming languages for developing the running services.
- The [Apptainer](https://apptainer.org), and [Enroot](https://github.com/NVIDIA/enroot) container runtime and several Linux kernel features.
- The [Slurm job scheduler](https://slurm.schedmd.com/documentation.html) for resource allocation and isolation.
- And much more!

Each workload is allocated resources by Slurm using [`cgroups`](https://docs.kernel.org/admin-guide/cgroup-v2.html) and is executed in containers using Apptainer or Enroot. Apptainer and Enroot acts as the software compatibility layer.

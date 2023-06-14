---
title: Architecture
---

## DeepSquare's Job Scheduling Architecture

DeepSquare's architecture resembles the Web-Queue-Worker model, effectively combining elements through Web3:

- Identity provider (represented by the user's wallet address)
- A persistent database
- A job queue
- A consistent billing system
- A common API

Within this architecture, the infrastructure provider hosts a **Supervisor** worker, which is tasked with monitoring the smart contract queue. In addition, the provider employs the **Slurm** job scheduler, which is designed to handle multi-node workloads.

To bridge the Supervisor and the Client without storing sensitive data on the blockchain, we use the sbatch API. This API transcribes workflows into a Slurm batch (SBatch) script. This script is then stored and held until the supervisor is ready to claim and retrieve the job.

![Deepsquare Architecture.drawio](.assets-deploy-deepsquare/image-226926804-7515f9e2-9f5f-43fc-b3f7-edc134be683f.png)

1. Customer Web2/Web3 app usage
1. Job definition creation
1. SDK integration in apps
1. Job submission through SDK
1. On-chain event detection
1. Meta-scheduler matching workload with suitable cluster
1. Compute Provider claiming the job
1. Compute Provider retrieving Job Definition
1. Job computation

## Underlying Technologies of DeepSquare

DeepSquare brings together a range of technologies to provide a seamless HPC experience:

- [Solidity](https://docs.soliditylang.org/en/latest/) for smart contracts
- [Go](https://go.dev) and [Python](https://www.python.org) for services development
- [Apptainer](https://apptainer.org) and [Enroot](https://github.com/NVIDIA/enroot) for container runtime and several Linux kernel features
- [Slurm job scheduler](https://slurm.schedmd.com/documentation.html) for resource allocation and isolation
- And much more!

These technologies enable each workload to be assigned resources via Slurm and executed in containers.

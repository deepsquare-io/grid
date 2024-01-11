# Scheduling Architecture

## The course of the job after submission

DeepSquare's architecture resembles the Web-Queue-Worker model, effectively combining elements through Web3:

- Identity provider (represented by the user's wallet address)
- A persistent database
- A job queue
- A consistent billing system
- A common API

Within this architecture, the infrastructure provider hosts a **Supervisor** worker, which is tasked with monitoring the smart contract queue. In addition, the provider employs the **Slurm** job scheduler, which is designed to handle multi-node workloads.

To bridge the Supervisor and the Client without storing sensitive data on the blockchain, we use the sbatch API. This API transcribes workflows into a Slurm batch (SBatch) script. This script is then stored and held until the supervisor is ready to claim and retrieve the job.

![Deepsquare Architecture.drawio](.assets-deploy-deepsquare/image-226926804-7515f9e2-9f5f-43fc-b3f7-edc134be683f.png)

1. The **User** submits a workflow to a **Web2 or Web3 application**.
2. The **application** sends the job definition is sent to the **Job Definition service**, which is used to sanitize, validate and pre-process the job into a job script.
3. The **application** interacts with the **DeepSquare library** and transmit the job definition.
4. The **DeepSquare library** sends the job definition to the **Meta-scheduler smart-contract**.
5. The **Meta-scheduler smart-contract** will notify the **Meta-scheduler service** and trigger the meta-scheduling algorithm.
6. The **Meta-scheduler** match the resource allocation with the best cluster, taking consideration resources allocation, resources occupation and fairness.
7. The **supervisor** of the compute provider is notified about the allocation and claim the job from the **Provider Manager queue**.
8. The **supervisor** of the compute provider fetches the complete job definition from the **Job Definition service**.
9. The **supervisor** submits the complete job definition to SLURM, the internal job scheduler.

## The return trip of the job

A return trip is needed since we want the providers and users to have feedbacks and rewards.

1. **SLURM** reports the job duration to the **supervisor**.
2. The **supervisor** notify the state and the duration of the job to the meta-scheduler **smart-contracts**.
3. The **compute provider** is rewarded and the **User** retrieve the left-over credits.

The **supervisor** constantly notify the meta-schedulers about the state of the jobs.

## Underlying Technologies of DeepSquare

DeepSquare brings together a range of technologies to provide a seamless HPC experience:

- [Solidity](https://docs.soliditylang.org/en/latest/) for smart contracts
- [Go](https://go.dev) and [Python](https://www.python.org) for services development
- [Apptainer](https://apptainer.org) and [Enroot](https://github.com/NVIDIA/enroot) for container runtime and several Linux kernel features
- [Slurm job scheduler](https://slurm.schedmd.com/documentation.html) for resource allocation and isolation
- And much more!

These technologies enable each workload to be assigned resources via Slurm and executed in containers.

## Next steps

Now that you known about the course of the job, you should learn about the [core concepts of DeepSquare and its motivation](/workflow/learn/core-concepts).

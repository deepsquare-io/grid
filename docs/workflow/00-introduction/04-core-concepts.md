---
title: Core concepts
---

# Understanding the Essentials

As you step into the world of DeepSquare and distributed high performance computing (HPC), there are a few fundamental concepts that will help you grasp our operations more effectively. Here, we present you an overview of these concepts.

## Supercomputing

Supercomputers epitomize raw computational power. Comprising parallel processors, they excel in complex simulations, large data analysis, and intricate problem-solving at extreme speeds. Used in high-stakes computations like weather forecasting, climate research, quantum mechanics, and cryptography, their utilization has traditionally been restricted to certain scientific and governmental entities due to costs and complexity.

## High Performance Computing (HPC)

High Performance Computing (HPC) broadens supercomputing's reach by employing advanced technologies, including supercomputers, to perform complex, computationally intensive tasks. Unlike supercomputing, which focuses on the most powerful machines, HPC aims to maximize performance across diverse systems, from server clusters to massive interconnected supercomputer grids.

HPC is vital in various sectors requiring rapid, precise processing of large data volumes. The emergence of cloud-based HPC has democratized access to high-performance computing resources, making them on-demand, reducing the need for heavy hardware investment, and data center management.

In summary, while all supercomputers participate in HPC, HPC encompasses more than supercomputers, integrating various technologies and strategies to achieve peak computational performance.

## DeepSquare Grid

The DeepSquare Grid is a global network of computational resources. It represents a significant evolution in the field of distributed computing, allowing users to tap into a vast array of supercomputers and HPC resources around the world. With its decentralized architecture and meta-scheduling techniques, the DeepSquare Grid enables efficient distribution and execution of computational tasks, fostering a new age of collaborative computing.

## Meta-scheduling

Meta-scheduling is a strategic technique that involves managing and enhancing the scheduling of jobs or tasks across numerous distributed computing resources. It employs a superior-level scheduler to supervise lower-level schedulers, which are responsible for resource allocation and job scheduling on individual machines or clusters.

## Job Status

In the DeepSquare ecosystem, each job is assigned a status that represents its current state. The following are the different types of job statuses:

```c
enum JobStatus {
  PENDING,        // Job has been requested and is awaiting meta-scheduling
  META_SCHEDULED, // Cluster has been assigned for the job
  SCHEDULED,      // Job has been queued by the assigned cluster
  RUNNING,        // Job execution has started on the cluster
  CANCELLED,      // Job has been cancelled by the owner
  FINISHED,       // Job has successfully completed execution
  FAILED,         // Job execution has failed
  OUT_OF_CREDITS  // Job owner has insufficient credits to run the job
}
```

## Infrastructure Providers

Infrastructure providers are the pillars of the DeepSquare Grid. They are entities that provide computing resources to the users on the network. To maintain the reliability and efficiency of the network, these providers must fulfill certain prerequisites such as maintaining specific hardware or software configurations, regularly updating their systems, and enforcing strong security measures.

Participation as infrastructure providers offers visibility to potential customers, revenue opportunities, and access to a vast network of computing resources and expertise.

## Economy of Compute

The "Economy of Compute" concept focuses on optimizing the use of computing resources to achieve maximum computational output per unit of energy, time, or cost. Key strategies to enhance the economy of compute include parallelization, optimization, efficient resource allocation, and cloud computing. By concentrating on the economy of compute, organizations can not only reduce their operational costs but also boost workload efficiency and reduce environmental impact.

## Credit Allocation

Allocating credits to a job is a crucial part of running tasks on the DeepSquare platform. When initiating a job using the SDK, you will be prompted to specify the number of credit tokens you wish to allocate for your task. These tokens represent the computational resources you are reserving for your job. Given the dynamic pricing of resources by providers on the Grid, it's advised to allocate a higher number of tokens than the bare minimum to prevent premature job termination due to insufficient credits. The tokens are locked for the duration of the job, and any remaining tokens after the job completion are returned to your account.

## Allowance System

The crypto allowance is a limit set on the amount of cryptocurrency DeepSquare can utilize on your behalf. This aids in streamlining the payment process and minimizes transaction confirmations during service use.

## Top-Up Running Job

Should you underestimate your credit requirement, the SDK enables you to add more tokens during job execution. This ensures your job runs smoothly without credit hitches.

## Auto Top-Up Mode

Activating Auto Top-Up mode lets DeepSquare automatically add credits to your job until it reaches your set allowance. Remember to specify an allowance that matches your wallet balance to prevent job failure.

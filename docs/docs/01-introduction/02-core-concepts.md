# Core concepts

## Meta-scheduling

Meta-scheduling is a technique for coordinating and optimizing the scheduling of jobs or tasks across multiple distributed computing resources. It involves using a higher-level scheduler to manage lower-level schedulers, which allocate resources and schedule jobs on individual machines or clusters.

## Job Status

```plaintext
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

## Infrastructure providers

Infrastructure providers are businesses or organizations that supply computing resources. They are essential in sustaining the DeepSquare Grid, offering their resources to other network participants.

As members of the DeepSquare ecosystem, these infrastructure providers may need to comply with certain prerequisites or recommendations to guarantee the security, dependability, and efficiency of the network. For instance, they might have to conform to particular hardware or software setups, perform routine updates and upkeep, and implement rigorous security measures to defend against data leaks or other hazards.

By participating in the DeepSquare grid as infrastructure providers, organizations can enjoy greater visibility to potential clients and revenue opportunities, as well as connect to an extensive network of computing resources and know-how. Simultaneously, they fulfill a vital function in promoting the continuous progress and expansion of the DeepSquare ecosystem, contributing to innovation and the advancement of distributed computing.

## High Performance Computing

High Performance Computing (HPC) employs cutting-edge hardware and software technologies to handle intricate and computationally demanding tasks. HPC systems typically consist of clusters of interconnected servers or supercomputers that collaborate to process large-scale projects.

HPC is commonly utilized in fields such as scientific research, engineering, finance, and others where massive amounts of data must be processed rapidly and accurately. Applications of HPC include Artificial intelligence, weather prediction, drug development, aerospace engineering, and financial risk assessment, etc.

HPC systems generally demand specialized hardware components like high-speed interconnects, parallel processors, and substantial amounts of memory and storage. Additionally, they depend on advanced software tools and frameworks for managing and coordinating the distribution of computational workloads throughout the cluster.

In recent times, the adoption of cloud-based HPC services has grown significantly, enabling organizations to access high-performance computing resources on-demand without investing in costly hardware or managing their own data centers. These cloud-based services often offer flexible pricing models based on usage, simplifying the process for organizations to scale their HPC capabilities as needed.

In summary, HPC is a vital and rapidly progressing field that allows researchers, scientists, and businesses to address some of the world's most complex and challenging computational issues.

## Economy of Compute

"Economy of compute" involves maximizing the use of computing resources to achieve the greatest computational output per unit of energy, time, or expense. It aims to extract the highest value from available resources.

Factors influencing the economy of compute include hardware and software efficiency, algorithm and data structure complexity, and workload distribution.

Key strategies for improving the economy of compute are:

- Parallelization: Dividing tasks into smaller segments for concurrent execution across multiple processors or nodes.
- Optimization: Enhancing algorithms, data structures, or code for speed or efficiency.
- Resource allocation: Distributing resources to optimize utilization and minimize idle time or wasted energy.
- Cloud computing: Using on-demand cloud services to access resources, enabling cost-effective scaling.

By optimizing the economy of compute, organizations can reduce costs, improve workload efficiency, and decrease environmental impact through lower energy consumption.

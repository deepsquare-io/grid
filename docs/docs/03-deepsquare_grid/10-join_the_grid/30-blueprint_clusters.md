# Blueprint clusters

DeepSquare offers turnkey solutions for infrastructure operators, i.e. tailored clusters pre-loaded with ClusterFactory. 
These configurations have been optimized to be as powerful as energy efficient and to facilitate heat-reuse. 
To achieve this goal, these solutions combine a typical supercomputing architecture with state of the art IT cooling technologies.  
They are made of 3 planes, interconnected with high performance networking using RDMA:

- **The compute plane:**  
Set of stateless compute nodes running SquareOS, our tailored HPC compute image, as a RAM disk (in memory)
- **The storage plane:**  
Set of storage nodes running BeeGFS, a parallel file system and CVMFS, the ideal filesystem to distribute applications.
- **The control/management plane:**  
Redundant high-available master nodes running the ClusterFactory control plane managing all the necessary services, such as the job scheduler, controller, provisioning system, monitoring, etc… 

Solutions using **immersion cooling** and **Direct Liquid Cooling (DLC)** are available, with both 19'' and 21'' (OCP) servers options.
Servers can be configured with latest Intel and AMD CPUs, DDR5 memory and a wide choice of GPU models.

To know more, please contact blueprints@deepsquare.io

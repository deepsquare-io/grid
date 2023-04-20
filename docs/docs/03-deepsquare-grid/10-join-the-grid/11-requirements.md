# Requirements

In order to be eligible for DeepSquare, a cluster needs to meet some requirements:

- Servers must run GNU/Linux
- Production grade [Slurm](https://slurm.schedmd.com/overview.html) cluster with the [pyxis](https://github.com/NVIDIA/pyxis) SPANK plugin
- [Enroot](https://github.com/NVIDIA/enroot) and [Apptainer](https://apptainer.org/) must be installed on the compute nodes
- Only Nvidia GPUs are supported for now
- [BeeGFS](https://www.beegfs.io/) or another shared file stem (RDMA is recommended)
- [CVMFS](https://cernvm.cern.ch/fs/) stratum 1 to synchronise the DeepSquare software catalog
- A DeepSquare supervisor with an approved address

The easiest way to deploy the above software stack and join the grid is by using [ClusterFactory](clusterfactory), our open source cluster manager available on [github](https://github.com/SquareFactory/ClusterFactory).
[ClusterFactory](clusterfactory) makes the process of deploying a full fledged HPC cluster and joining the grid fast and easy.

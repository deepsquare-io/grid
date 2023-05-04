# About ClusterFactory

[**ClusterFactory**](https://clusterfactory.io) is a modern and very powerful cluster manager.

It brings together best-in-class solutions from the HPC, Cloud, and DevOps industries to deploy and manage compute clusters in a declarative way in combination with the GitOps practice.

**ClusterFactory** is the easiest method to make your infrastructure ready to join the [DeepSquare Grid](/docs/deepsquare-grid/overview).

**ClusterFactory** provides:

- A production-ready vanilla upstream Kubernetes
- A solution easily deploy, scale, backup, restore, and update a Kubernetes cluster with cfctl
- A GitOps-enabled Continuous Deployement with ArgoCD and Sealed Secrets
- HPC/Batch workloads scheduling with Slurm
- Bare-metal provisioning with Grendel
- TLS/SSL certificates management with cert-manager
- DeepSquare software library mirror (CVMFS Stratum 1)
- Monitoring stack (Grafana, Prometheus with ready-to-use exporters)

To learn more about ClusterFactory, please have a look to its [dedicated documentation](https://docs.clusterfactory.io)

**ClusterFactory** allows DeepSquare to maintain a common software stack to facilitate **provisioning**, **maintenance** and **writing infrastructure as code** for bare-metal and cloud clusters.

---
toc_min_heading_level: 2
toc_max_heading_level: 2
---

# Overview

## About ClusterFactory

[**ClusterFactory**](https://clusterfactory.io) is a modern and very powerful cluster manager.

It brings together best-in-class solutions from the HPC, Cloud, and DevOps industries to deploy and manage compute clusters in a declarative way in combination with the GitOps practice.

**ClusterFactory** is the easiest method to make your infrastructure ready to join the [DeepSquare Grid](/docs/deepsquare-grid/overview).

To learn more about ClusterFactory, please have a look to its [dedicated documentation](https://docs.clusterfactory.io)

**ClusterFactory** allows DeepSquare to maintain a common software stack to facilitate **provisioning**, **maintenance** and **writing infrastructure as code** for bare-metal and cloud clusters.

## What is deployed with ClusterFactory

Do remember that ClusterFactory acts as the **control plane** of the whole cluster. It plays a crucial role in managing the entire cluster. It hosts various components necessary for the proper functioning of DeepSquare's compute plane.

ClusterFactory will be primarily used to deploy the following stacks:

1. **The vanilla Kubernetes stack**: This stack provides the foundation for running Kubernetes services within the cluster.

2. **The network stack**:

   - [Traefik](https://traefik.io/traefik/): It acts as the main entry point for Kubernetes services, functioning as a Layer 7 router and load balancer.
   - [MetalLB](https://metallb.universe.tf/): This component announces IP addresses for Kubernetes LoadBalancer services.
   - [Multus](https://github.com/k8snetworkplumbingwg/multus-cni): It enables multiple network interfaces with CNI plugins and serves as a secondary entry point for Kubernetes services that require the L2 network layer of the local network.
   - [CoreDNS](https://coredns.io): This is the primary domain name server for the entire cluster.

3. **The GitOps stack**:

   - [ArgoCD](https://argo-cd.readthedocs.io/en/stable/): It facilitates continuous deployment based on Git as the source of truth.
   - [cert-manager](https://cert-manager.io/): A powerful tool for generating and managing TLS certificates.
   - [sealed-secrets](https://github.com/bitnami-labs/sealed-secrets): This component allows encrypted secrets to be stored on Git.

4. **The provisioning stack**:

   - [Grendel](https://github.com/SquareFactory/grendel): It is an all-in-one bare-metal provisioner that incorporates, a DHCP server, a PXE server, an HTTP server, a TFTP server, and an IPMI controller.

5. **The software stack**:
   - [NFS CSI driver](https://github.com/kubernetes-csi/csi-driver-nfs) and [local-path provisioner](https://github.com/rancher/local-path-provisioner): These components enable storage provisioning for Kubernetes.
   - [MariaDB](https://mariadb.org): The database used by SLURM, a batch job scheduler for High-Performance Computing (HPC).
   - [389ds](https://directory.fedoraproject.org): The LDAP server utilized by SLURM and the compute nodes.
   - [Provider LDAP connector:](https://github.com/deepsquare-io/the-grid/tree/main/ldap-connector) This DeepSquare solution automatically registers DeepSquare users with the LDAP server.
   - [SLURM](https://slurm.schedmd.com/documentation.html): The batch job scheduler responsible for managing HPC workloads.
     - A login container: It serves as the main entry point for submitting SLURM batch scripts.
     - The controller container: It manages the SLURM system.
     - The database container: This connects MariaDB with the SLURM controller.
   - [The provider Supervisor](https://github.com/deepsquare-io/the-grid/tree/main/supervisor): This DeepSquare solution bridges DeepSquare with SLURM.
   - [CVMFS Stratum 1](https://cvmfs.readthedocs.io/en/stable/cpt-replica.html): A CVMFS server that replicates software exported on the compute nodes for the DeepSquare Grid.

**Take time to learn about these softwares as they will be used during the deployment process.**

## Architecture of a deployed cluster with ClusterFactory

The goal is to deploy the stacks, but it can be too complex to understand how everything is connected, so let's start with the basics.

### At the beginning, there is a bare-metal Kubernetes Cluster

When setting up a Kubernetes Cluster, there are two types of nodes that can be deployed: the Kubernetes controllers and the Kubernetes workers. The controllers are like the brain of the Kubernetes system and are responsible for managing different aspects of the cluster such as the deployment, scaling, and rollout of applications. On the other hand, the workers are the ones responsible for running the actual workloads and executing the containers.

<div style={{textAlign: 'center'}}>

![k0s Controller processes](./10-overview.assets/k0s_controller_processes.png#invert-on-dark)

![k0s worker processes](./10-overview.assets/k0s_worker_processes.png#invert-on-dark)

</div>

Usually, it's recommended to have an odd number of controllers for better fault tolerance. However, for simplicity, in this guide, we will combine the controllers and workers. If you want to separate the controllers and workers, it's recommended to follow the [K0s control plane high availability guide](https://docs.k0sproject.io/latest/high-availability/).

After setting up the controllers and workers, the Kubernetes API will be available, which you can connect to by fetching the kubeconfig. You can then use tools like [Lens](https://k8slens.dev) and kubectl to manage the cluster.

The Kubernetes workers will have pods (group of containers) running on it which are controlled by either a ReplicatSet/Deployment, a StatefulSet or a DaemonSet.

If you don't know what these resources are, these are the different types of Kubernetes workloads! To learn more, check out the [Kubernetes documentation](https://kubernetes.io/docs/concepts/workloads/).

### Then, we add the network stack for inter-pod and external-to-service communications

On Kubernetes, the smallest unit of workload that can be deployed is a Pod, which itself is a group of containers. Each pod is assigned a static IP for pod-to-external communication. To facilitate inter-pod and external-to-pod communications, it is necessary to utilize a [Kubernetes Service](https://kubernetes.io/docs/concepts/services-networking/service/) object.

Why is this necessary? In a Kubernetes cluster, pods are ephemeral and their IP addresses can change frequently. To ensure stable network communication, Kubernetes introduces the concept of a [Service](https://kubernetes.io/docs/concepts/services-networking/service/). A Service assigns a consistent IP address and DNS name to a group of pods, ensuring continuous connectivity even as pods change. This abstraction eliminates the need for clients to connect directly to individual pods, providing a reliable and scalable solution for intra-cluster communication.

**For inter-pod communication,** [ClusterIP](https://kubernetes.io/docs/concepts/services-networking/service/#publishing-services-service-types) Services are commonly used, which is the default type of Service. To further facilitate communication between pods, Kubernetes automatically creates DNS records for services and pods.

Therefore, we deploy [**CoreDNS**](https://coredns.io) first. CoreDNS serves as the domain name server within the Kubernetes cluster, providing names to Kubernetes pods and services. It can also act as the main DNS for the compute plane.

We also intend to expose [**Traefik**](https://traefik.io/traefik/), the **primary L7 router and entrypoint**. So instead of doing a simple ClusterIP Service, we need to use a Kubernetes [LoadBalancer](https://kubernetes.io/docs/concepts/services-networking/service/#loadbalancer) Service that allows us to attach an external IP to the Service.

To achieve this, we utilize **[MetalLB](https://metallb.universe.tf/)**, which enables the **advertisement of the IP to the router**, enabling routing of packets destined for Traefik.

**Lastly, some pods require direct access to a local network.** For example, the Grendel provisioning system uses DHCP, an L2 protocol that requires the DHCP server to be on the same network as the client broadcasting a DHCP discovery message. If the DHCP server is not in the same network as the client (e.g., Grendel is inside the Kubernetes network), a technical limitation prevents a DHCP discovery message from traversing a router.

To overcome this, we utilize [**Multus**](https://github.com/k8snetworkplumbingwg/multus-cni) in combination with a CNI IPVLAN plugin. This setup allows the Grendel pod to be directly connected to the local network, permitting DHCP communication.

**The network stack is necessary to expose the different workloads of the Kubernetes Cluster, from pod-to-pod, external-to-service and external-to-pod.**

<div style={{textAlign: 'center'}}>

![architecture-cf-de-Page-2.drawio](./10-overview.assets/architecture-cf-de-Page-2.drawio-1683645431656-6.svg#invert-on-dark)

</div>

### We add GitOps to ease the continuous deployment and infrastructure automation

The foundation of GitOps starts with [**ArgoCD**](https://argo-cd.readthedocs.io/en/stable/). ArgoCD automates the deployment and lifecycle management of applications and configurations in Kubernetes clusters, ensuring they are always in the desired state by leveraging GitOps principles. It synchronizes the desired state defined in Git repositories with the actual state of the target environment, providing a reliable and scalable solution for managing complex deployments, promoting collaboration, and ensuring consistency across multiple clusters and environments.

To always follow the GitOps principles, we must also manage our secrets and not store them in plaintext in the Git repository. Therefore, there is two solutions:

- [**cert-manager**](https://cert-manager.io): A solution to generate and manage TLS certificates using Kubernetes annotations and objects
- [**sealed-secrets**](https://github.com/bitnami-labs/sealed-secrets): A solution to encrypt secrets, which permits the storage of secrets in the Git repository

With these three solutions, your configuration is:

- Declarative: the entire system has to be described
- Versioned and immutable: the configurations are version-controlled and promotes infrastructure-as-code practices by allowing rollbacks simple as checking out a commit.
- Collaboration: instead of separating the operators from the developers, both of them can use the collaboration tools around git like pull requests to fosters collaboration among teams.
- Pulled automatically: Any approved changes to the git repository will be applied to the system
- Continuously reconciled: ArgoCD can see the differences between expected configuration and the reality

**GitOps is an approach that combines version control, automation, and collaboration to manage configurations and deployments in a declarative and auditable manner, ensuring consistency and scalability in modern software delivery pipelines. **

<div style={{textAlign: 'center'}}>

![architecture-cf-de-Page-3.drawio](./10-overview.assets/architecture-cf-de-Page-3.drawio.svg#invert-on-dark)

</div>

### We add Grendel to network boot the Compute Nodes

To network boot the compute nodes, we use [**Grendel**](https://github.com/SquareFactory/grendel) which is a combination of PXE (Preboot Execution Environment), TFTP (Trivial File Transfer Protocol), and DHCP (Dynamic Host Configuration Protocol).

The procedure of a network boot using Grendel is the following:

1. **Client Machine Initialization**: The network boot process begins when a client machine, configured for network booting, powers on or restarts. Instead of booting from its local storage, the client sends a DHCP request to the network to obtain an IP address and other necessary network configuration parameters.
2. **DHCP Server Response:** Upon receiving the DHCP request, the DHCP server, which is configured with the necessary options for network booting, responds to the client with an IP address, the IP address of the TFTP server, and the location of the initial firmware file to be fetched. In this case, the initial firmware file is the iPXE firmware.
3. **iPXE Firmware Download:** The client, armed with the IP address of the TFTP server and the filename of the iPXE firmware, initiates a TFTP (Trivial File Transfer Protocol) request to download the iPXE firmware from the TFTP server. The TFTP server responds by sending the iPXE firmware file to the client.
4. **iPXE Execution:** Once the iPXE firmware is successfully downloaded, the client executes the iPXE firmware, which takes control of the network stack of the client machine. Within the iPXE firmware, an iPXE script is executed to fetch the initramfs (initial RAM filesystem) and kernel files required for booting. The script defines the location of the initramfs and kernel files, typically specified as URLs or network paths, and instructs iPXE to download these files.
5. **Initramfs Boot:** Once the initramfs and kernel files are successfully downloaded, iPXE transfers control to the initramfs. The initramfs is a compressed file system that contains essential tools and drivers necessary for booting the system.
6. **Dracut Configuration:** The initramfs, upon initialization, reads the Dracut configuration. Dracut is a modular initramfs infrastructure used in many Linux distributions. The configuration defines the steps required to boot the system, including the detection and loading of hardware drivers, mounting of the root file system, and execution of necessary scripts or hooks.
7. **Live Image Boot:** Based on the Dracut configuration, the initramfs mounts the root file system, which is can be a live image of an operating system. The live image allows the client machine to boot into an operating system environment without modifying the local hard drive. Then, the live image boots up, initializing the operating system and presenting the user with a fully functional environment. The client machine can now be used.

Because, Grendel uses DHCP, we use **Multus and the IPVLAN CNI plugin** to connect the Grendel pod to the local network.

IPVLAN is a Container Network Interface (CNI) plugin that leverages the IPVLAN technology provided by the Linux kernel. IPVLAN allows the creation of a virtual network interface with its own IP address. In this case, the IPVLAN CNI plugin is used in conjunction with Multus to connect the Grendel pod to the local network.

<div style={{textAlign: 'center'}}>

![architecture-cf-de-Page-4.drawio](./10-overview.assets/architecture-cf-de-Page-4.drawio.svg#invert-on-dark)

</div>

### Lastly, we deploy the software stack which permits DeepSquare to work

This is the diagram that interconnect every components:

![architecture-cf-de-Page-5.drawio](./10-overview.assets/architecture-cf-de-Page-5.drawio.svg#invert-on-dark)

From left to right:

1. The **LDAP Connector** creates DeepSquare users in the 389ds LDAP server.
2. The **Supervisor** retrieves jobs from the smart contract and forwards them to a **SLURM login**.
3. The **SLURM login** submits the batch job to the **SLURM controller** and starts accounting with the **SLURM DB**, which is connected to a **MariaDB**.
4. The **SLURM controller** transmits the batch job to a compute node via the **SLURM daemon** and starts running the job.
5. Job statuses are reported to the **Supervisor** via the **SLURM SPANK plugin** on the compute node and the **SLURM completion plugin** on the SLURM Controller.

Everything is authenticated via **389ds** and **SSSD** (SSSD is also run inside the SLURM connection and controller containers).

The **CVMFS server** is exposed so that the compute node can mount the DeepSquare software on the compute nodes.

Finally, the **ArgoCD dashboard** is exposed so that Kubernetes engineers can access the argocd dashboard to view the status of the infrastructure.

Using all of these **standard tools**, this permits your infrastructure to be maintainable and scalable in the long term.

## What's next

Now you know what we are going to deploy. This will be the approach to deployment:

- Install the prerequisites.
- Deploy the ClusterFactory core.
- Deploy most of the control plane services.
- Deploy the control plane with Grendel.
- Learn how to maintain ClusterFactory.

Follow the [next chapter](prerequisites) for the prerequisites.

# The Supervisor

The Supervisor is the **link between DeepSquare and the Infrastructure Provider**. This guide will help you understand how the Supervisor works and how you can edit its behavior.

## The Dependencies

The Supervisor depends on three components:

- The meta-scheduler smart-contract
- The sbatch service (Job Definition Service)
- The SLURM Login

To **link the supervisor to the meta-scheduler smart contract**, you should:

1. Provider a wallet private key used for transactions and rewards. Right now, you should have **one private key per cluster**.
2. Set these parameters:

   - RPC Endpoint: **`https://testnet.deepsquare.run/rpc`**
   - WS Endpoint: **`https://testnet.deepsquare.run/ws`**
   - Smart-Contract Address: **see [GitHub Releases](https://github.com/deepsquare-io/grid/releases?q=smart-contracts&expanded=true).**

   If you are unfamiliar with Web3, the endpoints serve to interact with the smart-contracts, and the address is the identifier of an immutable software stored on the Blockchain.

To **link the supervisor to the sbatch service**, you should set the endpoint to: **`sbatch.deepsquare.run:443`**. The sbatch service uses encrypted gRPC traffic to exchange Job Definition. It is also secured by a Web3 Key Exchange, meaning, it is impossible to read the job of someone if you are not authorized.

To **link the supervisor to the SLURM login**, you should:

1. Set the address to the SLURM login.
2. Set the `adminUser` to `slurm`.
3. Set a generated SSH private key (`ssh-keygen -t ed25519 -C supervisor`).
4. Select a [SLURM Partition](https://slurm.schedmd.com/slurm.conf.html#SECTION_PARTITION-CONFIGURATION). Check the [examples](https://slurm.schedmd.com/slurm.conf.html#SECTION_PARTITION-CONFIGURATION:~:text=PartitionName%3DDEFAULT%20MaxTime%3D30%20MaxNodes%3D10%20State%3DUP%0APartitionName%3Ddebug%20Nodes%3Ddev%5B0%2D8%2C18%2D25%5D%20Default%3DYES%0APartitionName%3Dbatch%20Nodes%3Ddev%5B9%2D17%5D%20%20MinNodes%3D4%0APartitionName%3Dlong%20Nodes%3Ddev%5B9%2D17%5D%20MaxTime%3D120%20AllowGroups%3Dadmin).

The supervisor will expose a public key over gRPC. This public key can be fetched by using the [`provider-ssh-authorized-keys`](https://github.com/deepsquare-io/grid/releases?q=provider-ssh-authorized-keys&expanded=true) utility ([source code](https://github.com/deepsquare-io/grid/tree/main/provider-ssh-authorized-keys)).

## The Behavior

The normal behavior:

1. Upon starting the supervisor, the supervisor will health check SLURM. 
2. After that, the supervisor will launch benchmarks.
3. The results are sent to the meta-scheduler. The cluster is now waiting for approval. This process is done manually. Please contact DeepSquare's administrators to join the DeepSquare Grid.
4. If the cluster is approved, the supervisor will pull jobs from the Meta-Scheduler. No need to open any TCP ports.
5. If a job is received, the job definition is pulled and sent to SLURM.
6. The Provider SLURM SPANK Plugin will report the job state to the supervisor, and the supervisor will pass it to the Smart-Contract.
7. When the job finishes, the Provider SLURM Job Completion Plugin will report the job duration, and the supervisor will send the duration to the Smart-Contract.

The supervisor also has a garbage collector. The supervisor will periodically check if there are zombie jobs, and clean them. Zombie jobs will reimburse the user, and therefore, the provider won't be rewarded.

The supervisor has a crash recovery strategy, so it might be able to detect zombie jobs and treat them. The most important part is **to avoid losing the Job Completion Plugin callback**, otherwise, the job will be reimbursed.

:::note

In the future, we will make sure that the callback can be fetched even if the supervisor is down.

:::

## Configuration

### Benchmark

The benchmark can be configured if your infrastructure contains specificities:

- Enable the Single Node mode if your infrastructure does not possesses a good network (RDMA).
- Enable UCX if your infrastructure can use RDMA (Infiniband or RoCE, RDMA over Converged Ethernet).
  - Use affinities to select the network interfaces.
  - Use transport to select the UCX transport. If you don't know, don't set anything.
- Enable the Unresponsive flags if you want to submit a job for every node defined in the SLURM partition.

:::note

The SLURM partition is defined in the `slurm.conf` file. If you want to filter nodes, change the `slurm.conf`.

:::

### Labeling

Your cluster is [automatically labeled with standard labels](https://docs.deepsquare.run/workflow/learn/providers-labels). Labels are generated from the benchmark.

You can also add your own labels to advertise your cluster. Try to use standard labels. You can ask the DeepSquare's administrators to add your label to the standard.

For example, you can add labels about:

- Your location.
- The installed software.
- The accelerators installed on the nodes.
- The name of the cluster.
- ...

### Pricing

You can set your price to be competitive against other clusters. Right now, we have:

```yaml
## CPU physical threads per min.
cpuPerMin: '950000000000000000' # 9.5 credits/CPU.min
## GPU per min.
gpuPerMin: '8500000000000000000' # 8.5 credits/GPU.min
## RAM (MB) per min.
memPerMin: '80000000000000' # 0.00008 creds/MB.min or 0.08 creds/GB.min
```

For:

- CPU: AMD Epyc 7302 16 Cores Processor
- GPU: NVIDIA RTX 3090
- Memory: 16GB 3200 MT/s DDR4 CL22

:::note

Since we are in test net, these values are useless.

:::

## Next steps

Now that you've learn about how the supervisor works. You can:

- Ask to [Join the Grid](https://docs.deepsquare.run/deepsquare-grid/join/requirements).
- Try to [deploy the software stack with ClusterFactory](https://docs.deepsquare.run/deepsquare-grid/clusterfactory/overview).

We will be happy to welcome you to the world of decentralized HPC, and build the future of HPC together!

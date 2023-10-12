# Provider Labels and Job Affinities/Uses

## Labels

All providers are tagged with labels, either defined by the provider, or generated through a benchmark. Labels are key-value elements.

Standard labels are:

| Labels                                    | Description                                                                                                  | Example                           |
| ----------------------------------------- | ------------------------------------------------------------------------------------------------------------ | --------------------------------- |
| `os`                                      | Operating system of the cluster.                                                                             | `linux`                           |
| `arch`                                    | CPU architecture.                                                                                            | `amd64`                           |
| `cpu`                                     | CPU reported by `lscpu`.                                                                                     | `amd-epyc-7302-16-core-processor` |
| `cpu.microarch`                           | CPU micro-architecture.                                                                                      | `zen2`                            |
| `gpu`                                     | GPU of the cluster.                                                                                          | `nvidia-geforce-rtx-3090`         |
| `compute.gflops`                          | Giga Floating Points Operations Per Second (GFLOPS) of the whole cluster, based on the benchmark HPL-MXP-AI. | `96280.00` (96.280 TFLOPS)        |
| `network.download.bw.mbps`                | Download bandwidth with the Internet in Mbps.                                                                | `1119.06`                         |
| `network.upload.bw.mbps`                  | Upload bandwidth with the Internet in Mbps.                                                                  | `1170.66`                         |
| `network.p2p.bw.mbps`                     | Point-to-Point bandwidth between two nodes in Mbps.                                                          | `40241.29`                        |
| `network.p2p.latency.us`                  | Point-to-Point latency between two nodes in µs.                                                              | `202.53`                          |
| `network.all-to-all.latency.us`           | Bidirectional broadcast latency in µs.                                                                       | `143.83`                          |
| `storage.scratch.read.bw.mibps`           | `$STORAGE_PATH` read bandwidth in MiB/s.                                                                     | `8927.18`                         |
| `storage.scratch.read.iops`               | `$STORAGE_PATH` read I/O per seconds.                                                                        | `4493.06`                         |
| `storage.scratch.write.bw.mibps`          | `$STORAGE_PATH` write bandwidth in MiB/s.                                                                    | `3862.60`                         |
| `storage.scratch.write.iops`              | `$STORAGE_PATH` write I/O per seconds.                                                                       | `1967.37`                         |
| `storage.shared-world-tmp.read.bw.mibps`  | `$DEEPSQUARE_SHARED_WORLD_TMP` read bandwidth in MiB/s.                                                      | `8050.95`                         |
| `storage.shared-world-tmp.read.iops`      | `$DEEPSQUARE_SHARED_WORLD_TMP` read I/O per seconds.                                                         | `4047.36`                         |
| `storage.shared-world-tmp.write.bw.mibps` | `$DEEPSQUARE_SHARED_WORLD_TMP` write bandwidth in MiB/s.                                                     | `4460.33`                         |
| `storage.shared-world-tmp.write.iops`     | `$DEEPSQUARE_SHARED_WORLD_TMP` write I/O per seconds.                                                        | `2243.52`                         |
| `storage.shared-tmp.read.bw.mibps`        | `$DEEPSQUARE_SHARED_TMP` read bandwidth in MiB/s.                                                            | `8619.40`                         |
| `storage.shared-tmp.read.iops`            | `$DEEPSQUARE_SHARED_TMP` read I/O per seconds.                                                               | `4335.14`                         |
| `storage.shared-tmp.write.bw.mibps`       | `$DEEPSQUARE_SHARED_TMP` write bandwidth in MiB/s.                                                           | `3920.92`                         |
| `storage.shared-tmp.write.iops`           | `$DEEPSQUARE_SHARED_TMP` write I/O per seconds.                                                              | `1971.63`                         |
| `storage.disk-world-tmp.read.bw.mibps`    | `$DEEPSQUARE_DISK_WORLD_TMP` read bandwidth in MiB/s.                                                        | `108934.02`                       |
| `storage.disk-world-tmp.read.iops`        | `$DEEPSQUARE_DISK_WORLD_TMP` read I/O per seconds.                                                           | `54486.40`                        |
| `storage.disk-world-tmp.write.bw.mibps`   | `$DEEPSQUARE_DISK_WORLD_TMP` write bandwidth in MiB/s.                                                       | `940.15`                          |
| `storage.disk-world-tmp.write.iops`       | `$DEEPSQUARE_DISK_WORLD_TMP` write I/O per seconds.                                                          | `470.09`                          |
| `storage.disk-tmp.read.bw.mibps`          | `$DEEPSQUARE_DISK_TMP` read bandwidth in MiB/s.                                                              | `104976.84`                       |
| `storage.disk-tmp.read.iops`              | `$DEEPSQUARE_DISK_TMP` read I/O per seconds.                                                                 | `52503.92`                        |
| `storage.disk-tmp.write.bw.mibps`         | `$DEEPSQUARE_DISK_TMP` write bandwidth in MiB/s.                                                             | `786.46`                          |
| `storage.disk-tmp.write.iops`             | `$DEEPSQUARE_DISK_TMP` write I/O per seconds.                                                                | `393.24`                          |

## Affinities and Use flags when submitting jobs

A job can filter the clusters by using affinities. An affinity is basically a rule. Affinities are a set of rules which limits the selection of clusters. A Use flag is a simplified notation of the affinity using the `=` operator.

An affinity is represented by a key, a value and an operator.

The key is used to select a label of a provider, the value is used for the comparison, and the operator is the type of rule.

There are 6 supported operators: `<` (less than), `<=` (less than or equal), `=` or `==` (equal), `>` (greater than), `>=` (greater than or equal), `in` (includes).

On the CLI, the `in` operator is wrapped with `:`. Example: `gpu:in:nvidia`.

For example, there were two clusters:

- Cluster A, labelled:
  - `arch=amd64`
  - `compute.gflops=1000`
- Cluster B, labelled:
  - `arch=arm64`
  - `compute.gflops=10000`

When submitting a job, it is possible to filter clusters:

- "Only if the CPU architecture is arm64", rule is `arch=arm64`, which is also equivalent to the **Use** flag `arch=arm64`.
- "Only if the GFLOPS is greater than 1000", rule is `compute.gflops>1000`.
- "Only if the GPU is NVIDIA", rule is `gpu:in:nvidia`.

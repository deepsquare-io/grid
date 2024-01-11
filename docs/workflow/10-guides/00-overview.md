# Overview

The guides lead you through the advanced features of DeepSquare.

## Modules development

Want to develop modules for the community? Then, you are welcome to continue to the ["Module Development" section](/workflow/guides/module-development/about).

## Container specific features

Containers running on DeepSquare can be configured! You might want to see:

- [About using different container runtimes](container-runtimes)
- [Using the super-fast DeepSquare-hosted unpacked images](deepsquare-images) or [adding new images to the DeepSquare registry](publishing-deepsquare).
- [About environment variables](environment-variables) and [bind mounts](mounts).
- [About using the available caches](how-to-cache)
- [About mapping the UID/GID of the container](mapuid)

## Connectivity

Looking for a solution to connect your workload? Pull or push data?

These guides will help you:

- Expose a workload with [a Wireguard interface (TCP,UDP)](connecting-wireguard) or [a Bore proxy (TCP,HTTPS)](connecting-bore)
- Isolate workloads with [network namespaces](network-namespace) and interconnecting them with [Virtual Networks](connecting-interhost).
- [Use the DeepSquare Logger](logging)
- [Interact with the workload](interactive-mode)
- [Download the output of the jobs](fetching-output)

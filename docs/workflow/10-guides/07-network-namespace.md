# Avoiding port conflict with neighbors by using network namespaces

As we're in the HPC ecosystem, there's no default network isolation, so MPI can work as fast as possible. Consequently, opening a port **risky**. It is not just a question of privilege, but also of security.

To avoid port conflicts and data exposure, it is possible to use an **unprivileged Linux network namespace**.

## Network namespace

In the global namespace, a Linux installation shares a single set of network interfaces and routing table entries. Not only is this set shared across users, but it can only be modified by the root user.

With network namespaces, we can have our own network interfaces, IP routing tables, firewall rules, `/prod/net` directory (which is a symbolic link to `/proc/<ns pid>/net`), and so on.

This allows to have a namespace where you can host your own services and use your own network interfaces like Wireguard without having to worry about neighbors.

## **Unprivileged** network namespace

When we talk about **unprivileged**, we're talking about **root-less** and **capabilities-less**, i.e. you need to be able to create a functional network namespace with Internet connection without needing root privileges.

Usually, when creating a network namespace (for example, with `unshare --net --user`), the namespace has **zero** connectivity. It is fully isolated and fully contained. That creates one issue: there is no Internet connection. Normally, with Docker or Kubernetes, a bridge or a virtual Ethernet device (veth) is created to handle to traffic forwarding and Internet connectivity. To create that interface, you need root privileged, which is not possible on DeepSquare. Instead, we have to rely on TCP/IP stack emulation.

Currently on DeepSquare, there are two solutions for unprivileged network namespaces:

- [`slirp4netns`](https://github.com/rootless-containers/slirp4netns) which uses SliRP.
- [`pasta`](https://passt.top/passt/about/) which uses `passt`.

[SLiRP](https://gitlab.freedesktop.org/slirp/libslirp) is the common way of transmitting traffic from a network namespace to the host by creating a TAP device and emulating the TCP/IP stack. It is used by [Podman](https://github.com/containers/podman/blob/d161cf32fbbe25dc6496cd07cb3359eefca929c8/docs/tutorials/basic_networking.md#slirp4netns), QEMU and VirtualBox.

`passt` is very similar to SLiRP, but does not emulate the complete TCP/IP stack. In terms of performance, it is slightly inferior to SLiRP. However, from a functionality point of view, it may be promising for DeepSquare, since it allows precise control of NAT and port forwarding.

## TL;DR

To use a network namespace:

```yaml
steps:
  - name: 'interactive'
    run:
      # You may need to configure the DNS.
      dns:
        - 8.8.8.8
      # Set the network namespace. You need to be root to setup WireGuard.
      mapUid: 0
      mapGid: 0
      network: slirp4netns # or "pasta"
      customNetworkInterfaces:
        - wireguard:
            # ...
```

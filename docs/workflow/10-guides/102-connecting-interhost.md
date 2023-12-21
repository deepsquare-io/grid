# Interconnecting the network namespaces with a Virtual Network

:::warning

This feature is experimental and may show some instabilities.

Please report any feedback on GitHub or Discord.

:::

If you are using `slirp4netns` or `pasta`, your steps are naturally isolated and cannot communicate due to the unprivileged network namespace.

While network namespaces provide isolation, many workflows require tasks to communicate with each other. Consider a scenario where a workflow involves multiple microservices, each running in its own network namespace. These microservices may need to exchange data or collaborate on a larger computation. In such cases, an effective interconnection mechanism becomes essential.

To address the challenge of interconnecting tasks in isolated network namespaces, the concept of a "Virtual Network" is introduced. This solution leverages WireGuard, a L3 virtual private network that utilizes state-of-the-art cryptography. By creating a virtual network that spans across network namespaces, tasks can seamlessly communicate while maintaining the benefits of isolation.

## How to use

Start by **creating a virtual network** at the top level of the workflow definition:

```yaml title="Workflow"
enableLogging: true

resources:
  tasks: 1
  cpusPerTask: 1
  memPerCpu: 512
  gpus: 0

virtualNetworks:
  - name: my-network
    gatewayAddress: 10.0.0.1/24

steps:
  # TODO
```

The `name` will be used by steps to reference which virtual network to use. The `gatewayAddress` is nothing more than the address of the router. You can choose any address, but it must be correct.

To avoid choosing a possible conflicting address, we recommend `10.0.0.1/24`, or `10.0.1.1/24`, ... and so on, which are addresses often used by VPNs.

After creating the network, **add a network interface to the steps**. In this example, we are running a `netcat` server-client:

```yaml title="Workflow"
#...
virtualNetworks:
  - name: my-network
    gatewayAddress: 10.0.0.1/24

steps:
  - name: do-async
    launch:
      steps:
        - name: server
          run:
            network: slirp4netns
            mapUid: 0
            mapGid: 0
            customNetworkInterfaces:
              - vnet:
                  name: my-network
                  address: 10.0.0.2/24
            command: |-
              nc -l -p 12345

  - name: client
    run:
      network: slirp4netns
      mapUid: 0
      mapGid: 0
      customNetworkInterfaces:
        - vnet:
            name: my-network
            address: 10.0.0.3/24
      command: |-
        while true; do
            echo "Trying to connect..."
            if echo "hello world" | nc 10.0.0.2 12345; then
                echo "Connection successful"
                break
            else
                echo "Connection failed. Retrying in 5 seconds..."
                sleep 5
            fi
        done
        echo "Success"
```

You **have** to **map to root** (`mapUid: 0`) so that it is possible to create `vnet` network interfaces.

The `name` field of the network interface references the Virtual Network previously declared. The `address` is the IP used by the step. Note that there is a mask (`/24`), this can be changed in case you which to change the route mask.

Remainder:

- `/24` routes the IP range from `10.0.0.1` to `10.0.0.255`, to the virtual network.
- `/16` routes the IP range from `10.0.0.1` to `10.0.255.255`, to the virtual network.

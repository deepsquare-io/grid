# Exposing all the ports using a Wireguard tunnel

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

Exposing a workload is quite complicated. You can use `ngrok` to expose your service on a ngrok proxy. We offer a way to connect a wireguard tunnel to the jobs.

This solution is well adapted when you need to forward UDP traffic too.

## Schema

<div style={{textAlign: 'center'}}>

![wireguard.drawio](./100-connecting-wireguard.assets/wireguard.drawio.svg#invert-on-dark)

</div>

## How to use

### 1. Setting up a wireguard cloud router

#### 1.a. Deploying a server on the cloud

We recommend using a cloud provider which allows fast connectivity. We recommend [Exoscale](https://www.exoscale.com) for this.

For the rest of this guide, we assume that you have deployed an instance with these specifications:

- 1 vCPU, 1GB of RAM, 10 GB of persistent storage
- 1 public IP

Since we are in a server-client setup, **Network Address Translation (NAT)** becomes necessary to enable communication between the clients and the external network. To achieve this, the server uses **masquerading**. Specifically, the network interface dedicated to WireGuard utilizes the Linux IP masquerade feature to perform NAT on the packets passing through that interface. This involves modifying the source IP of the packets to match the IP address of the WireGuard server. By doing so, the routing on the client-side can be configured by setting the `allowedIPs` field to specify the VPN subnet they can access. This NAT setup ensures proper communication and routing between the clients and the external network.

For the sake of the example, we are also deploying NGINX servers on the compute nodes, which will open the port 8080/tcp. We will either load-balance with HAProxy or via iptables.

If you can deploy a [VyOS](https://vyos.io) as an OS image, it will be a lot simpler since it is an OS optimized for routers. You can also use an [OPNSense](https://opnsense.org) OS image, which is an OS for firewalls with routing capabilities, and has a web dashboard.

Alternatively, you can just use a Linux OS Image with iptables/nftables.

**If your cloud uses a firewall, make sure to open ports 22/tcp and the wireguard port (often 51820/udp).**

#### 1.b. Setup wireguard

<Tabs groupId="os">
<TabItem label="Linux with iptables" value="iptable">

1. Start by [installing Wireguard](https://www.wireguard.com/install/)

```shell title="root@~/"
sudo dnf install wireguard-tools
```

2. For old kernels (< 5.6), you must load the kernel module manually and make it persists:

   ```shell title="root@~/"
   sudo modprobe wireguard
   echo wireguard > /etc/modules-load.d/wireguard
   ```

   Check with:

   ```shell
   lsmod | grep wireguard
   ```

3. Enable packet forwarding by editing the `/etc/sysctl.conf` file:

   ```shell title="root@/etc/sysctl.conf"
   net.ipv4.ip_forward = 1
   ```

   And apply it:

   ```shell title="root@~/"
   sysctl -p /etc/sysctl.conf
   ```

4. Create the server and peer private key and public key with:

   ```shell title="root@~/"
   PK="$(wg genkey)"
   PUB="$(echo "$PK" | wg pubkey)"
   echo "PrivateKey = $PK"
   echo "PublicKey  = $PUB"  # You will send the public key to the job. Don't place it in the server configuration.
   ```

   For this example:

   ```shell
   ServerPrivateKey = 2GuuSBxL0pd1Mdv7sstzg2IYi5SO6TuuCEp+cDW8r0c=
   ServerPublicKey  = uF7mD0B9CxMVBY+1tn+bBHu/QTBBYIjw5l/92vgF/yE=
   ```

   To set up the wireguard virtual interface, create a file `/etc/wireguard/wg0.conf`. The file name indicates the name of the network interface. In this file add:

   ```shell title="root@/etc/wireguard/wg0.conf"
   [Interface]
   Address = 10.0.0.1/24
   ListenPort = 51820
   MTU = 1420
   PrivateKey = 2GuuSBxL0pd1Mdv7sstzg2IYi5SO6TuuCEp+cDW8r0c=
   # PublicKey = uF7mD0B9CxMVBY+1tn+bBHu/QTBBYIjw5l/92vgF/yE=
   # NAT: wg0 -> eth0 and any -> wg0
   PostUp = iptables -t nat -A POSTROUTING -s 10.0.0.0/24 -j MASQUERADE -o eth0
   PostUp = iptables -t nat -A POSTROUTING -o wg0 -j MASQUERADE
   PostDown = iptables -t nat -D POSTROUTING -s 10.0.0.0/24 -j MASQUERADE -o eth0
   PostDown = iptables -t nat -D POSTROUTING -o wg0 -j MASQUERADE
   ```

5. To add peers, create the private and public keys of the peers using the same command you used to create the server keys. Assuming you want to load balance between 2 tasks:

   ```shell title="root@/etc/wireguard/wg0.conf"
   [Interface]
   ...

   [Peer]
   # PrivateKey = iBSW/dma96OjnH098U5BEWzNcHIbsZqsCPiQ1DfP11c=
   PublicKey = nnhYC6rdEYyNqpTC9n0Q1ubnL5rmvbZIQ1IA75A1chk=
   AllowedIPs = 10.0.0.2/32

   [Peer]
   # PrivateKey = wIQkWhjPnIg9GUg1dhH6FmEIftKuxZdkvaD9VjOWH1Q=
   PublicKey = hCIM4037XrySn6BSLKz194X+ulqE4+PTVg2il1W12TU=
   AllowedIPs = 10.0.0.3/32
   ```

6. Use iptables to loadbalance and port forward. You can add `PostUp` and `PostDown` rules to add your iptables rules safely:

   ```shell title="root@/etc/wireguard/wg0.conf"
   [Interface]
   ...
   # NAT: wg0 -> eth0
   PostUp = iptables -t nat -A POSTROUTING -s 10.0.0.0/24 -j MASQUERADE -o eth0
   # NAT: any -> wg0
   PostUp = iptables -t nat -A POSTROUTING -o wg0 -j MASQUERADE
   # Accept traffic to 10.0.0.2:80 10.0.0.3:80
   PostUp = iptables -A FORWARD -p tcp -d 10.0.0.2 --dport 8080 -m state --state NEW,ESTABLISHED,RELATED -j ACCEPT
   PostUp = iptables -A FORWARD -p tcp -d 10.0.0.3 --dport 8080 -m state --state NEW,ESTABLISHED,RELATED -j ACCEPT
   # Forward packets and loadbalance
   PostUp = iptables -t nat -A PREROUTING -p tcp -i eth0 --dport 80 -j DNAT --to-destination 10.0.0.2:8080 -m statistic --mode random --probability 0.5
   PostUp = iptables -t nat -A PREROUTING -p tcp -i eth0 --dport 80 -j DNAT --to-destination 10.0.0.3:8080 -m statistic --mode random --probability 0.5
   # Cleanup
   PostDown = iptables -t nat -D POSTROUTING -s 10.0.0.0/24 -j MASQUERADE -o eth0
   PostDown = iptables -t nat -D POSTROUTING -o wg0 -j MASQUERADE
   PostDown = iptables -D FORWARD -p tcp -d 10.0.0.2 --dport 8080 -m state --state NEW,ESTABLISHED,RELATED -j ACCEPT
   PostDown = iptables -D FORWARD -p tcp -d 10.0.0.3 --dport 8080 -m state --state NEW,ESTABLISHED,RELATED -j ACCEPT
   PostDown = iptables -t nat -D PREROUTING -p tcp -i eth0 --dport 80 -j DNAT --to-destination 10.0.0.2:8080 -m statistic --mode random --probability 0.5
   PostDown = iptables -t nat -D PREROUTING -p tcp -i eth0 --dport 80 -j DNAT --to-destination 10.0.0.3:8080 -m statistic --mode random --probability 0.5

   [Peer]
   ...

   [Peer]
   ...
   ```

   Depending on the OS that you are using, eth0 could be renamed with `en*`.

   Note that this load balancing doesn't healthcheck the servers. You better use a [HAProxy](https://www.haproxy.org) as a Ingress for this.

   <details>
   <summary>HAProxy configuration example</summary>

   ```cfg title="/etc/haproxy/haproxy.cfg"
   global
       log         127.0.0.1 local2

       chroot      /var/lib/haproxy
       pidfile     /var/run/haproxy.pid
       maxconn     50000
       user        haproxy
       group       haproxy
       daemon

       # turn on stats unix socket
       stats socket /var/lib/haproxy/stats

       # utilize system-wide crypto-policies
       ssl-default-bind-ciphers PROFILE=SYSTEM
       ssl-default-server-ciphers PROFILE=SYSTEM

   defaults
       mode                    http
       log                     global
       option                  httplog
       option                  dontlognull
       timeout connect         10s
       timeout client          30s
       timeout server          30s
       timeout tunnel          30s
       maxconn                 50000

   # Frontends
   frontend my_frontend
       bind :80
       mode tcp
       default_backend my_backend

   # Backends
   backend my_backend
       mode tcp
       balance roundrobin
       server job1 10.0.0.2:8080 check
       server job2 10.0.0.3:8080 check

   ```

   Also, if you want to disable the load balancing, just remove the `-m statistic --mode random --probability 0.5` parameters and remove the extras port forwarding rules.

   :::warning

   Make sure that iptables does not conflict with other firewall software (like FirewallD or UFW)!

   Also note that there are no rules allowing traffic on ports. If your iptables forbid traffic by default, you can add these rules:

   ```shell
   # Allow HTTP (you can add this one in PostUp)
   iptables -A INPUT -p tcp -m tcp --dport 80 -j ACCEPT

   # Allow Wireguard (you can add this one in PostUp)
   iptables -A INPUT -p udp -m udp --dport 51820 -j ACCEPT

   # Allow SSH (you should have already opened this port)
   iptables -A INPUT -p tcp -m tcp --dport 22 -j ACCEPT

   # Allow ICMP (you should have already opened this port)
   iptables -A INPUT -p icmp ACCEPT
   ```

   :::

   You can remove the iptables forwarding rules.

   </details>

   You can also use IPVS for load balancing, which offer better performance than iptables (but doesn't offer healthchecks):

   <details>
   <summary>IPVS configuration</summary>

   ```shell
   # Add a virtual service
   ipvsadm -A -t <public IP>:80 -s rr
   # Forward packets
   ipvsadm -a -t <public IP>:80 -r 10.0.0.2:80 -m
   ipvsadm -a -t <public IP>:80 -r 10.0.0.3:80 -m
   ```

   You can add these rules to `PostUp`.

   </details>

   One last alternative, if you want to have high performance and high customization, try to look at the latest technologies: eBPF/XDP-based load balancing! BPF load balancing provides high performance, flexibility, and programmability in a safe and efficient way, allowing for the manipulation and analysis of network packets at a low level with minimal overhead.

   - [Cilium standalone](https://cilium.io/blog/2022/04/12/cilium-standalone-L4LB-XDP/)
   - [Envoy with HTTP3](https://www.envoyproxy.io/docs/envoy/latest/intro/arch_overview/http/http3.html#http-3-overview)
   - [LoxiLB](https://github.com/loxilb-io/loxilb)
   - [Developping with Cilium eBPF](https://docs.cilium.io/en/latest/bpf/index.html)
   - [Developping with Katran](https://github.com/facebookincubator/katran)

7. Enable the interface and start it!

   ```shell title="root@~/"
   chmod 600 /etc/wireguard/wg0.conf
   systemctl enable wg-quick@wg0
   systemctl start wg-quick@wg0
   ```

   Your cloud router is now configured!

</TabItem>
<TabItem label="VyOS 1.4 (17-03-2023)" value="vyos">

Login on the instance and update the VyOS to a version at least greater than `202303170317`:

```shell title="vyos@~/"
# Update OS
add system image https://s3-us.vyos.io/rolling/current/vyos-1.4-rolling-202303170317-amd64.iso
# Reboot
reboot
```

Then write these commands, be sure to modify accordingly:

```shell title="vyos@~/"
show version
# Go in config mode
config
# Configure firewall
# "in" means "incoming packets to the network interface"
set firewall interface eth0 in name 'OUTSIDE-IN'
# "local" is equivalent to the INPUT chain for iptables
# Basically local means "packets destined to the router"
set firewall interface eth0 local name 'OUTSIDE-LOCAL'
# Adding rules on incoming
set firewall name OUTSIDE-IN default-action 'drop'
set firewall name OUTSIDE-IN rule 10 action 'accept'
set firewall name OUTSIDE-IN rule 10 description 'Allow established/related'
set firewall name OUTSIDE-IN rule 10 state established 'enable'
set firewall name OUTSIDE-IN rule 10 state related 'enable'
set firewall name OUTSIDE-LOCAL default-action 'drop'
set firewall name OUTSIDE-LOCAL rule 10 action 'accept'
set firewall name OUTSIDE-LOCAL rule 10 description 'Allow established/related'
set firewall name OUTSIDE-LOCAL rule 10 state established 'enable'
set firewall name OUTSIDE-LOCAL rule 10 state related 'enable'
set firewall name OUTSIDE-LOCAL rule 20 action 'accept'
set firewall name OUTSIDE-LOCAL rule 20 description 'Allow ICMP'
set firewall name OUTSIDE-LOCAL rule 20 icmp type-name 'echo-request'
set firewall name OUTSIDE-LOCAL rule 20 protocol 'icmp'
set firewall name OUTSIDE-LOCAL rule 20 state new 'enable'
set firewall name OUTSIDE-LOCAL rule 30 action 'accept'
set firewall name OUTSIDE-LOCAL rule 30 description 'Allow SSH'
set firewall name OUTSIDE-LOCAL rule 30 destination port '22'
set firewall name OUTSIDE-LOCAL rule 30 protocol 'tcp'
set firewall name OUTSIDE-LOCAL rule 30 state new 'enable'
set firewall name OUTSIDE-LOCAL rule 100 action 'accept'
set firewall name OUTSIDE-LOCAL rule 100 description 'Allow Wireguard'
set firewall name OUTSIDE-LOCAL rule 100 destination port '51820'
set firewall name OUTSIDE-LOCAL rule 100 log 'enable'
set firewall name OUTSIDE-LOCAL rule 100 protocol 'udp'
set firewall name OUTSIDE-LOCAL rule 100 source
# If you prefer port forwarding over HAProxy, change OUTSIDE-LOCAL to OUTSIDE-IN
set firewall name OUTSIDE-LOCAL rule 120 action 'accept'
set firewall name OUTSIDE-LOCAL rule 120 description 'Allow HTTP'
set firewall name OUTSIDE-LOCAL rule 120 destination port '80'
set firewall name OUTSIDE-LOCAL rule 120 log 'enable'
set firewall name OUTSIDE-LOCAL rule 120 protocol 'tcp'
set firewall name OUTSIDE-LOCAL rule 120 state new 'enable'

# Configure wireguard
set interfaces wireguard wg0 address '10.0.0.1/24'
set interfaces wireguard wg0 description 'Wireguard to jobs'
set interfaces wireguard wg0 private-key '2GuuSBxL0pd1Mdv7sstzg2IYi5SO6TuuCEp+cDW8r0c='
set interfaces wireguard wg0 mtu '1420'
set interfaces wireguard wg0 peer job1 allowed-ips '10.0.0.2/32'
set interfaces wireguard wg0 peer job1 public-key 'nnhYC6rdEYyNqpTC9n0Q1ubnL5rmvbZIQ1IA75A1chk='
set interfaces wireguard wg0 peer job2 allowed-ips '10.0.0.3/32'
set interfaces wireguard wg0 peer job2 public-key 'hCIM4037XrySn6BSLKz194X+ulqE4+PTVg2il1W12TU='
set interfaces wireguard wg0 port '51820'

set nat source rule 20 description "any -> wg0"
set nat source rule 20 outbound-interface wg0
set nat source rule 20 translation address 'masquerade'

set nat source rule 30 description "wg0 -> eth0"
set nat source rule 30 outbound-interface eth0
set nat source rule 30 source address 10.0.0.0/24
set nat source rule 30 translation address 'masquerade'

# If you prefer port forwarding over HAProxy, uncomment the following rule:
# set nat destination rule 201 description 'http-forward'
# set nat destination rule 201 destination port '80'
# set nat destination rule 201 inbound-interface 'eth0'
# set nat destination rule 201 protocol 'tcp'
# set nat destination rule 201 translation address '10.0.0.2'
# set nat destination rule 201 translation port '8080'
# VyOS doesn't offer TCP load-balacing with the NAT rules.

commit-confirm
confirm
save
```

As root, add to `/etc/haproxy/haproxy.cfg`:

```cfg title="/etc/haproxy/haproxy.cfg"
# ...

# Frontends
frontend my_frontend
    bind :80
    mode tcp
    default_backend my_backend

# Backends
backend my_backend
    mode tcp
    balance roundrobin
    server job1 10.0.0.2:8080 check
    server job2 10.0.0.3:8080 check

```

Start and enable HAProxy:

```cfg title="vyos@~/"
sudo systemctl enable haproxy
sudo systemctl start haproxy
```

Your cloud router is now configured!

</TabItem>
</Tabs>

### 2. Sending jobs and connect them Wireguard VPN

To connect your job to the wireguard server, you can specify the client configuration in the workflow. Like this:

<Tabs groupId="client">
  <TabItem label="Peer 1" value="peer1">

```json title="Workflow"
{
  "enableLogging": true,
  "resources": {
    "tasks": 1,
    "cpusPerTask": 1,
    "memPerCpu": 1000,
    "gpusPerTask": 0
  },
  "steps": [
    {
      "name": "start-nginx",
      "run": {
        "container": {
          "registry": "registry-1.docker.io",
          "image": "nginxinc/nginx-unprivileged"
        },
        "dns": ["8.8.8.8"],
        "network": "slirp4netns",
        "customNetworkInterfaces": [
          {
            "wireguard": {
              "address": ["10.0.0.2/24"],
              "privateKey": "iBSW/dma96OjnH098U5BEWzNcHIbsZqsCPiQ1DfP11c=",
              "peers": [
                {
                  "publicKey": "uF7mD0B9CxMVBY+1tn+bBHu/QTBBYIjw5l/92vgF/yE=",
                  "allowedIPs": ["10.0.0.1/32"],
                  "persistentKeepalive": 10,
                  "endpoint": "<public IP>:51820"
                }
              ]
            }
          }
        ],
        "command": "/docker-entrypoint.sh nginx -g \"daemon off;\""
      }
    }
  ]
}
```

  </TabItem>
  <TabItem label="Peer 2" value="peer2">

```json title="Workflow"
{
  "enableLogging": true,
  "resources": {
    "tasks": 1,
    "cpusPerTask": 1,
    "memPerCpu": 1000,
    "gpusPerTask": 0
  },
  "steps": [
    {
      "name": "start-nginx",
      "run": {
        "container": {
          "registry": "registry-1.docker.io",
          "image": "nginxinc/nginx-unprivileged"
        },
        "dns": ["8.8.8.8"],
        "network": "slirp4netns",
        "customNetworkInterfaces": [
          {
            "wireguard": {
              "address": ["10.0.0.3/24"],
              "privateKey": "wIQkWhjPnIg9GUg1dhH6FmEIftKuxZdkvaD9VjOWH1Q=",
              "peers": [
                {
                  "publicKey": "uF7mD0B9CxMVBY+1tn+bBHu/QTBBYIjw5l/92vgF/yE=",
                  "allowedIPs": ["10.0.0.1/32"],
                  "persistentKeepalive": 10,
                  "endpoint": "<public IP>:51820"
                }
              ]
            }
          }
        ],
        "command": "/docker-entrypoint.sh nginx -g \"daemon off;\""
      }
    }
  ]
}
```

  </TabItem>
</Tabs>

A second job will have `wIQkWhjPnIg9GUg1dhH6FmEIftKuxZdkvaD9VjOWH1Q=` as private key and the `10.0.0.3` address.

## Why not IPsec or OpenVPN ?

WireGuard is seamlessly integrated into the Linux kernel, delivering superior performance and employing advanced encryption techniques. Its lightweight configuration makes it particularly well-suited for deployment in DeepSquare, providing a more straightforward setup compared to OpenVPN.

Thanks to its UDP-based nature and support for dynamic IPs, WireGuard inherently facilitates NAT traversal, rendering it highly suitable for server-client setups in a decentralized ecosystem like DeepSquare. As a result, traditional point-to-point VPN solutions like IPsec are not supported.

Overall, Wireguard is well-suited in our environment. It simplifies the process of establishing and maintaining connections across NAT devices, ensuring smooth and reliable communication between the server and clients, even in dynamic and diverse network environments like DeepSquare.

## Limitations

If you enable the `slirp4netns` container networking will automatically **re-map the user as root**.

Although you are remapped as root, this is still in an unprivileged container, which means that it is not possible to bind restricted ports (like 80).

Unprivileged containers have some limitations in terms of capabilities and network access. Unprivileged containers are typically limited in their network access, as they are usually isolated from the host network and can only communicate with other containers or the outside world through specific network interfaces or bridges set up by the container runtime.

These limitations are intended to provide better security and isolation for the container and its contents, but may pose challenges for certain types of applications or workloads that require more extensive system access or network connectivity.

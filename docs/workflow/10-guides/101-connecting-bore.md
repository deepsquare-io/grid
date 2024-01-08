# Exposing HTTPS using a Bore tunnel without overhead

While a WireGuard tunnel allow UDP and TCP ports forwarding, it's quite complicated and heavy to set up.

Instead of using WireGuard, it is possible to use [jkuri/bore](https://github.dev/jkuri/bore/), which is an implementation of a bore proxy based on Go and SSH port forwarding.

Advantages of jkuri/bore over WireGuard are:

- Easy to set up
- HTTPS/HTTP URL generation

## How to use

Similar to WireGuard, you MUST set `network` to `slirp4netns` or `pasta`.

A bore server has been deployed at [bore.deepsquare.run:2200](https://bore.deepsquare.run), with HTTPS already set up. If you are concerned about authority, you can choose to host your own bore server by following the README in [the official repository of jkuri/bore](https://github.dev/jkuri/bore/). We also recommend deploying [Caddy](https://caddyserver.com) reverse proxy for easy configuration (auto TLS, HTTP/3 support).

Using DeepSquare to expose a port is quite easy:

```yaml title="Workflow"
enableLogging: true

resources:
  tasks: 1
  cpusPerTask: 1
  memPerCpu: 512
  gpus: 0

steps:
  - name: start-nginx
    run:
      command: nginx -g "daemon off;"
      container:
        registry: registry-1.docker.io
        image: nginxinc/nginx-unprivileged:latest
      ## Use the container network interface slirp4netns (or pasta) to create a network namespace.
      network: slirp4netns # or pasta
      customNetworkInterfaces:
        ## Forward TCP/UDP traffic from port 8080 to bore.deepsquare.run:2200.
        - bore:
            boreAddress: bore.deepsquare.run:2200
            targetPort: 8080
```

Remember that we are still running in an unprivileged container, so it is impossible to bind to restricted ports.

We use `nginxinc/nginx-unprivileged:latest` in our example, which is a simple web server that binds to port 8080. The bore client connects to the `bore.deepsquare.run:2200` server and redirects the local port 8080 to the bore proxy.

You can then fetch the generated URL and port in the logs. It should look like this:

```shell
Generated HTTP URL:  http://3d6393aa.bore.deepsquare.run
Generated HTTPS URL: https://3d6393aa.bore.deepsquare.run
Direct TCP:          tcp://bore.deepsquare.run:63206
```

## Allocate a Bore address

This feature is unique to DeepSquare. If you go to the  [DeepSquare's Bore Proxy](https://bore.deepsquare.run) website, you can fetch a unique route:

![image-20240108143148167](./101-connecting-bore.assets/image-20240108143148167.png)

Use a **long-lived** token if you plan to reuse the same authentication token. **Fetch route** will fetch an existing route. If the route doesn't exists yet, a new route will be generated.

Use MetaMask to authenticate. No payment is required, you only have to **sign a message**, which doesn't use any blockchain.

This will be returned:

> Allocated HTTP URL: `http://f43677a4-6acb-4d81-b917-85915e3f3073.bore.deepsquare.run`
>
> Allocated HTTPS URL: `https://f43677a4-6acb-4d81-b917-85915e3f3073.bore.deepsquare.run`
>
> Allocated TCP URL: `tcp://bore.deepsquare.run:39447`
>
> Token: `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ4MDcxODIsInVzZXJfaWQiOiIweDc1NzYxYjE3YzMwODhjZTVjZDhlMDI1NzVjNmRhYTQzOGZmYTZlMTIiLCJ1c2VyX25hbWUiOiIweDc1NzYxYjE3YzMwODhjZTVjZDhlMDI1NzVjNmRhYTQzOGZmYTZlMTIiLCJub25jZSI6IlFrODE5a2tLQU9DM0FhR0I2bXdrMXc9PSJ9.vcmUYuS-okXVpCQcVhJWBD8-Tq4yv1r3SpaK1secFcg`
>
> Token expiration: 2024-01-09 14:33:02 +0100 CET

The URL and port is allocated to you and you can use it for your jobs. Right now, only one URL is allowed per user.

Use the **token** to fill the `secret` field of the bore proxy:

```yaml
bore:
  boreAddress: bore.deepsquare.run:2200
  targetPort: 8080
  secret: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDQ4MDcxODIsInVzZXJfaWQiOiIweDc1NzYxYjE3YzMwODhjZTVjZDhlMDI1NzVjNmRhYTQzOGZmYTZlMTIiLCJ1c2VyX25hbWUiOiIweDc1NzYxYjE3YzMwODhjZTVjZDhlMDI1NzVjNmRhYTQzOGZmYTZlMTIiLCJub25jZSI6IlFrODE5a2tLQU9DM0FhR0I2bXdrMXc9PSJ9.vcmUYuS-okXVpCQcVhJWBD8-Tq4yv1r3SpaK1secFcg
```

The authentication token will automatically expire after a certain time, but **the route will remain accessible.**

If you've lost your token, you can recover the route and the old token will be invalidated.

:::note

If you self-host your own [Bore](https://github.com/jkuri/bore) proxy, the `secret` will correspond to the `id` parameter.

:::

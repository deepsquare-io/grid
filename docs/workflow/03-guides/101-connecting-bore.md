# Exposing HTTPS using a Bore tunnel without overhead

While a wireguard tunnel allow UDP and TCP ports forwarding, it's quite complicated and heavy to setup.

Instead of using Wireguard, it is possible to use [jkuri/bore](https://github.dev/jkuri/bore/), which is an implementation of a bore proxy based on Go and SSH port forwarding.

Advantages of jkuri/bore over wireguard are:

- Easy to setup
- HTTPS/HTTP url generation

## How to use

Similar to wireguard, you MUST set `network` to `slirp4netns`.

A bore server has been deployed at [bore.deepsquare.run:2200](https://bore.deepsquare.run), with HTTPS already set up. If you are concerned about authority, you can choose to host your own bore server by following the README in [the official repository of jkuri/bore](https://github.dev/jkuri/bore/). We also recommend to deploy [Caddy](https://caddyserver.com) reverse proxy for easy configuration (auto TLS, HTTP/3 support).

Using DeepSquare to expose a port is quite easy:

```json
{
  "enableLogging": true,
  "resources": {
    "tasks": 1,
    "cpusPerTask": 1,
    "memPerCpu": 512,
    "gpusPerTask": 0
  },
  "steps": [
    {
      "name": "start-nginx",
      "run": {
        "command": "nginx -g \"daemon off;\"",
        "container": {
          "registry": "registry-1.docker.io",
          "image": "nginxinc/nginx-unprivileged:latest"
        },
        "network": "slirp4netns",
        "customNetworkInterfaces": [
          {
            "bore": {
              "address": "bore.deepsquare.run",
              "port": 2200,
              "targetPort": 8080
            }
          }
        ]
      }
    }
  ]
}
```

Remember that we are still running in an unprivileged container, so it is impossible to bind to restricted ports.

We use `nginxinc/nginx-unprivileged:latest` in our example, which is a simple web server that binds to port 8080. The bore client connects to the bore.deepsquare.run:2200 server and redirects the local port 8080 to the bore proxy.

You can then fetch the generated URL and port in the logs. It should look like this:

```shell
Generated HTTP URL:  http://3d6393aa.bore.deepsquare.run
Generated HTTPS URL: https://3d6393aa.bore.deepsquare.run
Direct TCP:          tcp://bore.deepsquare.run:63206
```

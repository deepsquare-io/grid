# Exposing container logs

Because DeepSquare is decentralized, we cannot open ports or expect infrastructure provider to open ports. Instead the job is aggregating the logs and pushing them to a remote server.

The [portal](https://app.deepsquare.run) can access to that logging server. The access is secured with web3 authentication and encrypted with SSL thanks to HTTPS.

## Usage

```json title="Workflow"
{
  "resources": {
    "tasks": 1,
    "gpusPerTask": 0,
    "cpusPerTask": 1,
    "memPerCpu": 1024
  },
  "enableLogging": true,
  "steps": [
    {
      "name": "hello world",
      "run": {
        "command": "echo \"Hello World\""
      }
    }
  ]
}
```

## Can I use my own logging service?

Not yet, but it is planned.

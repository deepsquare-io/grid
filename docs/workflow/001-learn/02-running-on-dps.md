# Running jobs on DeepSquare

DeepSquare allows you to run your computational tasks on any compute provider within the DeepSquare Grid through job scheduling, container technologies, and Web3. Container technologies ensure compatibility across different clusters, while web3 provides transparency, availability, and scalability as the backbone of a global job scheduler.

## Workflow Files: Your Starting Point

A workflow file is your primary tool for running applications on DeepSquare. It outlines resource allocation and the sequence of instructions for your application. Here's an example:

```json
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

## Executing a workflow

To launch a workflow file, you have a couple of choices:

1\. Visit [DeepSquare Dev App](https://app.deepsquare.run/sandbox), paste the workflow, and run it. You might need to convert YAML to JSON until we make the portal compatible with YAML.

2\. Use the [DeepSquare SDK](https://www.npmjs.com/package/@deepsquare/deepsquare-client) which provides a simple and abstracted interface from web3 to the DeepSquare Grid.

## Credits and SQUARE Tokens

Regardless of how you choose to run the job, whether through the portal or SDK, you need to have credits to pay for the jobs and SQUARE tokens to pay for gas fees. SQUARE tokens function similarly to ether in the Ethereum network, while credits are akin to USDC on the Ethereum blockchain. You can request free credits [here](https://share-eu1.hsforms.com/18lhtQBNNTVWVRXCm7t-83Aev6gi).

### Credit Allocation

You'll be prompted to specify the number of credit tokens to allocate when initiating a job via the SDK. Allocate extra tokens to prevent premature termination due to insufficient credits. Tokens are locked during the job, with remaining tokens returned to your account upon completion.

### Submitting Your Job

Once you've prepared your workflow file, acquired credits, and obtained SQUARE tokens, you're ready to submit your job to the DeepSquare Grid. Either press 'Submit' on the portal or use the `submitJob` function in the SDK.

### Monitoring Job Status

Monitor your [job status](/workflow/learn/core-concepts#job-status) until it's finished to retrieve your job results, if any.

### Job Completion

Jobs end when they've naturally completed (success, failure, canceled), or there are no credits left for the job.

### Pricing

To gain a practical understanding of executing workloads on our platform, we recommend following our [Getting Started](/workflow/getting-started/introduction) guide. This guide provides a step-by-step tutorial on submitting your first job on the DeepSquare Grid, thereby familiarizing you with our tools and processes.

If you have further inquiries or require assistance, our team is readily available on our [Discord server](https://discord.gg/rDaWwNfxfg) to provide support. As you delve deeper into the platform, don't hesitate to reach out to our community and experts who can help you navigate any complexities.

We look forward to supporting your computational work on the DeepSquare platform.

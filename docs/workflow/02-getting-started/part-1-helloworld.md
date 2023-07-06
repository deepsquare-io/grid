---
title: 'Part 1: Running a simple helloworld'
---

## Setting up a Crypto Wallet

Before using the platform, you need to set up a crypto wallet to manage your credits. We recommend using either [MetaMask](https://support.metamask.io/hc/en-us/articles/360015489531-Getting-started-with-MetaMask) or [Core Wallet](https://chrome.google.com/webstore/detail/core-crypto-wallet-nft-ex/agoakfejjabomempkjlepdflaleeobhb).

## Acquiring Free Credits

To get started, you can acquire some free credits by applying through our [form](https://share-eu1.hsforms.com/18lhtQBNNTVWVRXCm7t-83Aev6gi).

## Hello World workflow

A DeepSquare **workflow file** serves as the blueprint of your operations, detailing the resources your application needs and the sequence of steps it must follow to execute its tasks. As the core of running workloads on DeepSquare, these workflow files form the backbone of any operation within the platform.

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

1. **`resources`**: This initial block in the workflow file dictates the resource allocation for the entire job. It represents a high-level request for computational resources you'll need for all your tasks.
2. **`tasks`**: Within the **`resources`** block, the tasks field specifies the number of independent units of work, or tasks, your application needs to run. Each task runs in parallel, and the scheduler assigns it to the available computational resources.
3. **`gpusPerTask`**: This field designates the number of GPU resources each task can utilize. If your tasks don't require GPU processing power, this can be set to 0.
4. **`cpusPerTask`**: Similar to `gpusPerTask`, this field determines the number of CPU resources each task will have access to.
5. **`memPerCpu`**: This represents the amount of memory (in MB) allocated for each CPU.
6. **`steps`**: This block is where you outline the specific actions your tasks will take. Steps can be thought of as individual instructions that will be executed by your tasks, utilizing the resources defined above.
7. **`command`**: Within each step, you specify a command. This is the exact operation that should be performed during the step. It could be a script to run, a function to call, or any other operation your application requires.

The **`enableLogging`** field, when set to true, sends the application logs to the DeepSquare logging system. This allows you to monitor the progress of your tasks.

So, in a nutshell, a workflow file in DeepSquare provides a structured way for you to define what resources you need, how many parallel tasks you want to run, and what each task should do.

## Executing the workflow

To launch a workflow file, you have a couple of choices:

1. Visit [DeepSquare Dev App](https://app.deepsquare.run/sandbox), paste the workflow, and run it. You might need to convert YAML to JSON until we make the portal compatible with YAML.
2. Use the [DeepSquare SDK](https://www.npmjs.com/package/@deepsquare/deepsquare-client) which provides a simple and abstracted interface from web3 to the DeepSquare Grid.

### With the DeepSquare portal

#### Use the Dev App

1. Visit [DeepSquare Dev App](https://app.deepsquare.run/sandbox).
2. Log-in using your crypto wallet.
3. Make sure you have some credits (100 is the minimum entry cost) and some SQUARE tokens to pay for transactions.
4. In the **Workflow Editor**, write or paste your workflow JSON.
5. If your workflow is valid, click **Submit** to run the job.

#### Downloading Results and Review Logs

1. In the **Job Status** page, locate your job based on its name.
2. Monitor the progress of your job by checking its status.
3. Click on the **Logs** button in order to read the realtime logs.
4. Once your job completes, you can view the results, logs, and other output files in the designated output storage.

With DeepSquare's Dev application, you can efficiently create, manage, and run custom workflows to cater to a wide range of high-performance computing demands.

### With the SDK

Visit the [DeepSquare SDK documentation page](https://www.npmjs.com/package/@deepsquare/deepsquare-client) for a detailed walkthrough on setting up and getting started with your first SDK-based application.

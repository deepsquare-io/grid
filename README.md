# The DeepSquare Grid

A decentralized network of compute resources.

## Monorepo dependency tree

```mermaid
%%{
  init: {
    "theme": "dark",
    "logLevel": "info",
    "flowchart": {
      "htmlLabels": true
    }
  }
}%%
flowchart LR
  subgraph Specifications
    schemas/sbatchapi
    protos/loggerapis
    smart-contracts
    protos/sbatchapis
    protos/supervisorapis
  end

  subgraph Provider
    subgraph Provider Services
      ldap-connector
      supervisor
    end


    subgraph CLIs
      grid-logger-writer
      provider-ssh-authorized-keys
    end

    subgraph Slurm plugins
      provider-job-completion-plugin
      provider-spank-plugin
    end
  end

  subgraph Deepsquare
    subgraph Deepsquare Services
      sbatch-service
      grid-logger-server
      meta-scheduler
      oracle-scheduler
    end
  end

  protos/loggerapis --> grid-logger-server
  protos/loggerapis --> grid-logger-writer
  schemas/sbatchapi --> sbatch-service
  protos/sbatchapis --> sbatch-service
  protos/sbatchapis --> supervisor
  protos/supervisorapis --> supervisor
  protos/supervisorapis --> provider-ssh-authorized-keys
  protos/supervisorapis --> provider-job-completion-plugin
  protos/supervisorapis --> provider-spank-plugin
  smart-contracts --> meta-scheduler
  smart-contracts --> oracle-scheduler
  smart-contracts --> ldap-connector
  smart-contracts --> supervisor
```

## What is the DeepSquare Grid

TODO

## Why use the DeepSquare Grid

TODO

## Upgrade the smart-contracts

Go to the [Smart Contracts CI page](https://github.com/deepsquare-io/the-grid/actions/workflows/smart-contracts.yaml),
and run the workflow with the Release flag checked and Initial Deploy unchecked.

## Documentation

TODO

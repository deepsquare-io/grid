# The DeepSquare Grid

A decentralized network of compute resources.

## Monorepo dependency tree

```mermaid
flowchart TD
  grid-logger --> protos/loggerapis --> customer-portal
  sbatch-service --> protos/sbatchapis --> supervisor
  sbatch-service --> schemas/sbatchapi --> customer-portal
  supervisor --> protos/supervisorapis
  protos/supervisorapis --> provider-job-completion-plugin
  protos/supervisorapis --> provider-spank-plugin
  protos/supervisorapis --> provider-ssh-authorized-keys-plugin
  smart-contracts --> ldap-connector
  smart-contracts --> meta-scheduler
  smart-contracts --> oracle-scheduler
  smart-contracts --> customer-portal
  smart-contracts --> supervisor
```

## What is the DeepSquare Grid

TODO

## Why use the DeepSquare Grid

TODO

## Upgrade the smart-contracts

Go to the [Smart Contracts CI page](https://github.com/deepsquare-io/the-grid/actions/workflows/smart-contracts.yaml), and run the workflow with the Release flag checked and Initial Deploy unchecked.

## Documentation

TODO

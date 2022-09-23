#!/bin/sh

set -ex

SCRIPTPATH=$(dirname "$(realpath "$0")")
PROJECTPATH="${SCRIPTPATH}/.."
CONTRACTSPATH="${SCRIPTPATH}/../../smart-contracts"

cd "${CONTRACTSPATH}"

pnpm exec hardhat compile

mkdir -p "${PROJECTPATH}/gen/go/contracts/credit"
solc ./contracts/Credit.sol \
  --base-path . \
  --include-path "node_modules/" \
  --combined-json abi,bin | abigen --pkg credit \
  --combined-json - \
  --out "${PROJECTPATH}/gen/go/contracts/credit/credit.go"

mkdir -p "${PROJECTPATH}/gen/go/contracts/datastruct"
solc ./contracts/DataStruct.sol \
  --base-path . \
  --include-path "node_modules/" \
  --combined-json abi,bin | abigen --pkg datastruct \
  --combined-json - \
  --out "${PROJECTPATH}/gen/go/contracts/datastruct/datastruct.go"

mkdir -p "${PROJECTPATH}/gen/go/contracts/faucet"
solc ./contracts/Faucet.sol \
  --base-path . \
  --include-path "node_modules/" \
  --combined-json abi,bin | abigen --pkg faucet \
  --combined-json - \
  --out "${PROJECTPATH}/gen/go/contracts/faucet/faucet.go"

mkdir -p "${PROJECTPATH}/gen/go/contracts/grid"
solc ./contracts/Grid.sol \
  --base-path . \
  --include-path "node_modules/" \
  --combined-json abi,bin | abigen --pkg grid \
  --combined-json - \
  --out "${PROJECTPATH}/gen/go/contracts/grid/grid.go"

mkdir -p "${PROJECTPATH}/gen/go/contracts/job_manager"
solc ./contracts/JobManager.sol \
  --base-path . \
  --include-path "node_modules/" \
  --combined-json abi,bin | abigen --pkg job_manager \
  --combined-json - \
  --out "${PROJECTPATH}/gen/go/contracts/job_manager/job_manager.go"

mkdir -p "${PROJECTPATH}/gen/go/contracts/lock"
solc ./contracts/Lock.sol \
  --base-path . \
  --include-path "node_modules/" \
  --combined-json abi,bin | abigen --pkg lock \
  --combined-json - \
  --out "${PROJECTPATH}/gen/go/contracts/lock/lock.go"

mkdir -p "${PROJECTPATH}/gen/go/contracts/metascheduler"
solc ./contracts/Metascheduler.sol \
  --base-path . \
  --include-path "node_modules/" \
  --combined-json abi,bin | abigen --pkg metascheduler \
  --combined-json - \
  --out "${PROJECTPATH}/gen/go/contracts/metascheduler/metascheduler.go"

mkdir -p "${PROJECTPATH}/gen/go/contracts/provider_manager"
solc ./contracts/ProviderManager.sol \
  --base-path . \
  --include-path "node_modules/" \
  --combined-json abi,bin | abigen --pkg provider_manager \
  --combined-json - \
  --out "${PROJECTPATH}/gen/go/contracts/provider_manager/provider_manager.go"

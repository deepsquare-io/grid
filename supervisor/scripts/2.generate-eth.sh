#!/bin/sh

set -ex

SCRIPTPATH=$(dirname "$(realpath "$0")")
PROJECTPATH="${SCRIPTPATH}/.."
CONTRACTSPATH="${SCRIPTPATH}/../../smart-contracts"

cd "${CONTRACTSPATH}"

pnpm exec hardhat compile

mkdir -p "${PROJECTPATH}/gen/go/contracts/metascheduler"
solc ./contracts/Metascheduler.sol \
  --base-path . \
  --include-path "node_modules/" \
  --combined-json abi,bin | abigen --pkg metascheduler \
  --combined-json - \
  --out "${PROJECTPATH}/gen/go/contracts/metascheduler/metascheduler.go"

#!/bin/sh

set -ex

SCRIPTPATH=$(dirname "$(realpath "$0")")
PROJECTPATH="${SCRIPTPATH}/.."
CONTRACTSPATH="${SCRIPTPATH}/../../smart-contracts"

cd "${CONTRACTSPATH}"

mkdir -p "${PROJECTPATH}/gen/go/contracts/metascheduler"
solc --optimize --optimize-runs=200 ./contracts/Metascheduler.sol \
  --base-path . \
  --include-path "node_modules/" \
  --combined-json abi | abigen --pkg metascheduler \
  --combined-json - \
  --exc "contracts/Tools.sol:Tools" \
  --out "${PROJECTPATH}/gen/go/contracts/metascheduler/metascheduler.go"

cd "${PROJECTPATH}"

solc --optimize --optimize-runs=200 ./pkg/eth/ErrorContract.sol \
  --base-path . \
  --include-path "${CONTRACTSPATH}/contracts/" \
  --include-path "${CONTRACTSPATH}/node_modules/" \
  --combined-json abi,bin | abigen --pkg eth \
  --combined-json - \
  --out "${PROJECTPATH}/pkg/eth/eth_error_abi.go"

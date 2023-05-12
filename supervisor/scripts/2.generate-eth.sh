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

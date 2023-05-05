#!/bin/sh

set -ex

SCRIPTPATH=$(dirname "$(realpath "$0")")
PROJECTPATH="${SCRIPTPATH}/.."
CONTRACTSPATH="${SCRIPTPATH}/../../smart-contracts"

cd "${CONTRACTSPATH}"

mkdir -p "${PROJECTPATH}/contracts/metascheduler"
solc --optimize --optimize-runs=200 ./contracts/Metascheduler.sol \
  --base-path . \
  --include-path "node_modules/" \
  --combined-json abi | abigen --pkg metascheduler \
  --combined-json - \
  --out "${PROJECTPATH}/contracts/metascheduler/metascheduler.go"

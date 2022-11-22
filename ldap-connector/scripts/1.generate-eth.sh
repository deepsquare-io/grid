#!/bin/sh

set -ex

SCRIPTPATH=$(dirname "$(realpath "$0")")
PROJECTPATH="${SCRIPTPATH}/.."
CONTRACTSPATH="${SCRIPTPATH}/../../smart-contracts"

cd "${CONTRACTSPATH}"

mkdir -p "${PROJECTPATH}/gen/go/contracts/jobmanager"
solc ./contracts/JobManager.sol \
  --base-path . \
  --include-path "node_modules/" \
  --combined-json abi,bin | abigen --pkg jobmanager \
  --combined-json - \
  --out "${PROJECTPATH}/gen/go/contracts/jobmanager/jobmanager.go"

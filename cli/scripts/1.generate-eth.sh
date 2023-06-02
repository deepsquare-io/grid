#!/bin/sh

set -ex

SCRIPTPATH=$(dirname "$(realpath "$0")")
PROJECTPATH="${SCRIPTPATH}/.."
CONTRACTSPATH="${SCRIPTPATH}/../../smart-contracts"

mkdir -p "${PROJECTPATH}/deepsquare/generated/abi/metascheduler/"
solc --optimize --optimize-runs=200 "${CONTRACTSPATH}/contracts/Metascheduler.sol" \
  --base-path . \
  --include-path "${CONTRACTSPATH}/node_modules/" \
  --include-path "${CONTRACTSPATH}/contracts/" \
  --combined-json abi | abigen --pkg metaschedulerabi \
  --combined-json - \
  --exc "Tools.sol:Tools" \
  --out "${PROJECTPATH}/deepsquare/generated/abi/metascheduler/metascheduler.go"

mkdir -p "${PROJECTPATH}"/deepsquare/generated/abi/errors
solc --optimize --optimize-runs=200 "${PROJECTPATH}/deepsquare/metascheduler/ErrorContract.sol" \
  --base-path . \
  --include-path "${CONTRACTSPATH}/contracts/" \
  --include-path "${CONTRACTSPATH}/node_modules/" \
  --combined-json abi,bin | abigen --pkg errorsabi \
  --combined-json - \
  --out "${PROJECTPATH}/deepsquare/generated/abi/errors/errors.go"

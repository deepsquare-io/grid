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
  --combined-json abi | abigen --pkg metashechdulerabi \
  --combined-json - \
  --exc "contracts/Tools.sol:Tools" \
  --out "${PROJECTPATH}/generated/metascheduler/abi/metascheduler.go"

cd "${PROJECTPATH}"

solc --optimize --optimize-runs=200 ./pkg/metascheduler/ErrorContract.sol \
  --base-path . \
  --include-path "${CONTRACTSPATH}/contracts/" \
  --include-path "${CONTRACTSPATH}/node_modules/" \
  --combined-json abi,bin | abigen --pkg errorsabi \
  --combined-json - \
  --out "${PROJECTPATH}/generated/errors/abi/errors.go"

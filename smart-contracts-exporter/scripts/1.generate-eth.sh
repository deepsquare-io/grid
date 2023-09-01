#!/bin/sh

set -ex

SCRIPTPATH=$(dirname "$(realpath "$0")")
PROJECTPATH="${SCRIPTPATH}/.."
CONTRACTSPATH="${SCRIPTPATH}/../../smart-contracts"

cd "${CONTRACTSPATH}"

mkdir -p "${PROJECTPATH}/contracts/metascheduler"
solc --optimize --optimize-runs=200 "${CONTRACTSPATH}/contracts/Metascheduler.sol" \
  --base-path . \
  --include-path "${CONTRACTSPATH}/node_modules/" \
  --include-path "${CONTRACTSPATH}/contracts/" \
  --combined-json abi >"${PROJECTPATH}/contracts/metascheduler/metascheduler.json"
sed -Ei 's/"type":"JobStatus"/"type": "uint8"/g' "${PROJECTPATH}/contracts/metascheduler/metascheduler.json"
sed -Ei 's/"type":"StorageType"/"type": "uint8"/g' "${PROJECTPATH}/contracts/metascheduler/metascheduler.json"
abigen --pkg metascheduler \
  --combined-json "${PROJECTPATH}/contracts/metascheduler/metascheduler.json" \
  --out "${PROJECTPATH}/contracts/metascheduler/metascheduler.go"
rm "${PROJECTPATH}/contracts/metascheduler/metascheduler.json"
echo "WARNING: enum JobStatus was mapped to uint8"
echo "WARNING: enum StorageType was mapped to uint8"

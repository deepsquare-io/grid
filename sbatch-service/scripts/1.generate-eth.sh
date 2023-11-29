#!/bin/sh

set -ex

SCRIPTPATH=$(dirname "$(realpath "$0")")
PROJECTPATH="${SCRIPTPATH}/.."
CONTRACTSPATH="${SCRIPTPATH}/../../smart-contracts"

mkdir -p "${PROJECTPATH}/abi/metascheduler"
solc --evm-version paris --optimize --optimize-runs=200 "${CONTRACTSPATH}/contracts/Metascheduler.sol" \
  --base-path . \
  --include-path "${CONTRACTSPATH}/node_modules/" \
  --include-path "${CONTRACTSPATH}/contracts/" \
  --combined-json abi >"${PROJECTPATH}/abi/metascheduler/metascheduler.json"
sed -Ei 's/"type":"JobStatus"/"type": "uint8"/g' "${PROJECTPATH}/abi/metascheduler/metascheduler.json"
sed -Ei 's/"type":"StorageType"/"type": "uint8"/g' "${PROJECTPATH}/abi/metascheduler/metascheduler.json"
abigen --pkg metaschedulerabi \
  --combined-json "${PROJECTPATH}/abi/metascheduler/metascheduler.json" \
  --out "${PROJECTPATH}/abi/metascheduler/metascheduler.go"
rm "${PROJECTPATH}/abi/metascheduler/metascheduler.json"
echo "WARNING: enum JobStatus was mapped to uint8"
echo "WARNING: enum StorageType was mapped to uint8"

mkdir -p "${PROJECTPATH}/abi/error"
solc --evm-version paris --optimize --optimize-runs=200 "${PROJECTPATH}/types/ErrorContract.sol" \
  --base-path . \
  --include-path "${CONTRACTSPATH}/node_modules/" \
  --include-path "${CONTRACTSPATH}/contracts/" \
  --combined-json abi,bin >"${PROJECTPATH}/abi/error/error.json"
sed -Ei 's/"type":"JobStatus"/"type": "uint8"/g' "${PROJECTPATH}/abi/error/error.json"
sed -Ei 's/"type":"StorageType"/"type": "uint8"/g' "${PROJECTPATH}/abi/error/error.json"
abigen --pkg errorabi \
  --combined-json "${PROJECTPATH}/abi/error/error.json" \
  --out "${PROJECTPATH}/abi/error/error.go"
rm "${PROJECTPATH}/abi/error/error.json"
echo "WARNING: enum JobStatus was mapped to uint8"
echo "WARNING: enum StorageType was mapped to uint8"

#!/bin/sh

set -ex

SCRIPTPATH=$(dirname "$(realpath "$0")")
PROJECTPATH="${SCRIPTPATH}/.."
CONTRACTSPATH="${SCRIPTPATH}/../../smart-contracts"

mkdir -p "${PROJECTPATH}/types/abi/metascheduler/"
solc --evm-version paris --optimize --optimize-runs=200 "${CONTRACTSPATH}/contracts/Metascheduler.sol" \
  --base-path . \
  --include-path "${CONTRACTSPATH}/node_modules/" \
  --include-path "${CONTRACTSPATH}/contracts/" \
  --combined-json abi >"${PROJECTPATH}/types/abi/metascheduler/metascheduler.json"
sed -Ei 's/"type":"JobStatus"/"type": "uint8"/g' "${PROJECTPATH}/types/abi/metascheduler/metascheduler.json"
sed -Ei 's/"type":"StorageType"/"type": "uint8"/g' "${PROJECTPATH}/types/abi/metascheduler/metascheduler.json"
abigen --pkg metaschedulerabi \
  --combined-json "${PROJECTPATH}/types/abi/metascheduler/metascheduler.json" \
  --out "${PROJECTPATH}/types/abi/metascheduler/metascheduler.go"
rm "${PROJECTPATH}/types/abi/metascheduler/metascheduler.json"
echo "WARNING: enum JobStatus was mapped to uint8"
echo "WARNING: enum StorageType was mapped to uint8"

mkdir -p "${PROJECTPATH}"/types/abi/errors
solc --evm-version paris --optimize --optimize-runs=200 "${PROJECTPATH}/metascheduler/ErrorContract.sol" \
  --base-path . \
  --include-path "${CONTRACTSPATH}/contracts/" \
  --include-path "${CONTRACTSPATH}/node_modules/" \
  --combined-json abi,bin >"${PROJECTPATH}/types/abi/errors/errors.json"
sed -Ei 's/"type":"JobStatus"/"type": "uint8"/g' "${PROJECTPATH}/types/abi/errors/errors.json"
sed -Ei 's/"type":"StorageType"/"type": "uint8"/g' "${PROJECTPATH}/types/abi/errors/errors.json"
abigen --pkg errorsabi \
  --combined-json "${PROJECTPATH}/types/abi/errors/errors.json" \
  --out "${PROJECTPATH}/types/abi/errors/errors.go"
rm "${PROJECTPATH}/types/abi/errors/errors.json"
echo "WARNING: enum JobStatus was mapped to uint8"
echo "WARNING: enum StorageType was mapped to uint8"

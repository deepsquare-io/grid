#!/bin/sh

set -ex

SCRIPTPATH=$(dirname "$(realpath "$0")")
PROJECTPATH="${SCRIPTPATH}/.."
CONTRACTSPATH="${SCRIPTPATH}/../../smart-contracts"

cd "${CONTRACTSPATH}"

mkdir -p "${PROJECTPATH}/generated/abi/metascheduler"
solc --optimize --optimize-runs=200 "${CONTRACTSPATH}/contracts/Metascheduler.sol" \
  --base-path . \
  --include-path "${CONTRACTSPATH}/node_modules/" \
  --include-path "${CONTRACTSPATH}/contracts/" \
  --combined-json abi >"${PROJECTPATH}/generated/abi/metascheduler/metascheduler.json"
sed -Ei 's/"type":"JobStatus"/"type": "uint8"/g' "${PROJECTPATH}/generated/abi/metascheduler/metascheduler.json"
sed -Ei 's/"type":"StorageType"/"type": "uint8"/g' "${PROJECTPATH}/generated/abi/metascheduler/metascheduler.json"
abigen --pkg metaschedulerabi \
  --combined-json "${PROJECTPATH}/generated/abi/metascheduler/metascheduler.json" \
  --out "${PROJECTPATH}/generated/abi/metascheduler/metascheduler.go"
rm "${PROJECTPATH}/generated/abi/metascheduler/metascheduler.json"
echo "WARNING: enum JobStatus was mapped to uint8"
echo "WARNING: enum StorageType was mapped to uint8"

cd "${PROJECTPATH}"

solc --optimize --optimize-runs=200 ./pkg/metascheduler/ErrorContract.sol \
  --base-path . \
  --include-path "${CONTRACTSPATH}/node_modules/" \
  --include-path "${CONTRACTSPATH}/contracts/" \
  --combined-json abi,bin >"${PROJECTPATH}/generated/abi/errors/errors.json"
sed -Ei 's/"type":"JobStatus"/"type": "uint8"/g' "${PROJECTPATH}/generated/abi/errors/errors.json"
sed -Ei 's/"type":"StorageType"/"type": "uint8"/g' "${PROJECTPATH}/generated/abi/errors/errors.json"
abigen --pkg errorsabi \
  --combined-json "${PROJECTPATH}/generated/abi/errors/errors.json" \
  --out "${PROJECTPATH}/generated/abi/errors/errors.go"
rm "${PROJECTPATH}/generated/abi/errors/errors.json"
echo "WARNING: enum JobStatus was mapped to uint8"
echo "WARNING: enum StorageType was mapped to uint8"

#!/bin/sh

set -e

SCRIPTPATH=$(dirname "$(realpath "$0")")
PROJECTPATH="${SCRIPTPATH}/../"

cd "${PROJECTPATH}"

go install \
  github.com/ethereum/go-ethereum/cmd/geth@latest

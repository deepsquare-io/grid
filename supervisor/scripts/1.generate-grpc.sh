#!/bin/sh

set -e

SCRIPTPATH=$(dirname "$(realpath "$0")")
PROTOSPATH="${SCRIPTPATH}/../../protos"

cd "${PROTOSPATH}"

buf generate

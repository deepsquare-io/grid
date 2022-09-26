#!/bin/sh

set -e

SCRIPTPATH=$(dirname "$(realpath "$0")")

cd "${SCRIPTPATH}/../../protos/supervisorapis"

buf generate

cd "${SCRIPTPATH}/../../protos/oracleapis"

buf generate

#!/bin/sh

set -e

SCRIPTPATH=$(dirname "$(realpath "$0")")

cd "${SCRIPTPATH}/../../protos/supervisorapis"

buf generate

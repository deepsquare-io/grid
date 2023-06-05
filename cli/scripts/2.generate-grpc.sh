#!/bin/sh

set -e

SCRIPTPATH=$(dirname "$(realpath "$0")")

cd "${SCRIPTPATH}/../../protos/loggerapis"

buf generate

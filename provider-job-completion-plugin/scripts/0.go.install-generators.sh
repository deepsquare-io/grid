#!/bin/sh

set -e

SCRIPTPATH=$(dirname "$(realpath "$0")")
PROJECTPATH="${SCRIPTPATH}/../"

cd "${PROJECTPATH}"

mkdir -p build
cd build
cmake ..
make -j "$(nproc)" grpc++

cp ./_deps/grpc-build/grpc_*_plugin ~/.local/bin

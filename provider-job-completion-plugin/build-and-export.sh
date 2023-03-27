#!/bin/sh

set -e

mkdir -p cache/8.6 cache/9.0

podman build \
  -t jobcomp-builder:8.6 \
  -f Dockerfile.rpm \
  -v "$(pwd)/cache/8.6:/work/build" \
  --build-arg rocky_version=8.6\
  .

podman build \
  -t jobcomp-builder:9.0 \
  -f Dockerfile.rpm \
  -v "$(pwd)/cache/9.0:/work/build" \
  --build-arg rocky_version=9.0\
  .

mkdir -p artifacts/8.6/ artifacts/9.0/
podman run -v "$(pwd)/artifacts/8.6:/out" --rm jobcomp-builder:8.6 sh -c "cp /artifacts/* /out/"
podman run -v "$(pwd)/artifacts/9.0:/out" --rm jobcomp-builder:9.0 sh -c "cp /artifacts/* /out/"

s3cmd sync -v --acl-public --config="$(pwd)/.s3cfg" "$(pwd)/artifacts/8.6/" "s3://yum-repository/8/x86_64/"
s3cmd sync -v --acl-public --config="$(pwd)/.s3cfg" "$(pwd)/artifacts/9.0/" "s3://yum-repository/9/x86_64/"

#!/bin/sh
set -ex

slurm_version=$(curl --silent https://api.github.com/repos/schedmd/slurm/tags | jq '.[1].name' | awk '{ print substr( $0, 2, length($0)-2 ) }')
sed -i "s/slurm=.*\$/slurm=${slurm_version}/g" ./versions.lock
cd ./externals/slurm
git checkout "${slurm_version}"

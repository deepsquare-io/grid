#!/bin/sh
#
##SBATCH --gpus=1
##SBATCH --mem=10G

ml stdenv Tdp/1.1.0-beta

export DISPLAY=:99

srun --export=ALL \
  /home/ldap-users/marc/tdp/DeepSquareURSSample.x86_64 \
  -maxplayers 8 \
  -renderfps 60 \
  -webserverurl "wss://tdp.deepsquare.run" \
  -displayfpscounter \
  --logFile -

#!/bin/bash

export JOB_NAME=name

./bin/grid-logger-writer-linux-amd64 \
  --server.tls \
  --server.tls.server-host-override=grid-logger.deepsquare.run \
  --server.tls.ca=/etc/ssl/certs/ca-certificates.crt \
  --server.endpoint=grid-logger.deepsquare.run:443 \
  --pipe.path=$JOB_NAME \
  --log-name=$JOB_NAME \
  --user=0x75761B17c3088ce5Cd8e02575c6DAa438FFA6e12 \
  >/dev/stdout 2>/dev/stderr &
LOGGER_PID=$!

echo "logger running at $LOGGER_PID"

sleep 1

exec &>>"$(pwd)/pipe"

echo "$(date): Running some workload!"
echo "$(date): Huh!"

sleep 3

echo "$(date): srun --container-image=test"

echo >&2 "$(date): fake error"

touch /test

kill $LOGGER_PID

wait

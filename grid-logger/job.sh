#!/bin/bash

./bin/grid-logger-writer \
  --server.tls=false \
  --pipe.path=pipe \
  --log-name=test \
  --user=user \
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

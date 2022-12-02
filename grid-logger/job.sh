#!/bin/bash

./bin/client --server.tls=false --pipe-file=pipe >/dev/stdout 2>/dev/stderr &
LOGGER_PID=$!

echo "logger running at $LOGGER_PID"

sleep 1

exec >>"$(pwd)/pipe"

echo "$(date): Running some workload!"
echo "$(date): Huh!"

sleep 5

echo "$(date): srun --container-image=test"

kill $LOGGER_PID

wait

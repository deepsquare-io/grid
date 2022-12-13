#!/bin/bash -l

set -e

export STORAGE_PATH="/opt/cache/shared/$UID/$SLURM_JOB_NAME"
mkdir -p "$STORAGE_PATH"
chmod -R 700 "$STORAGE_PATH"
chown -R "$UID:cluster-users" "$STORAGE_PATH"

{{- if and .EnableLogging (derefBool .EnableLogging ) }}
/usr/local/bin/grid-logger-writer \
  --server.tls \
  --server.tls.ca=/etc/ssl/certs/ca-certificates.crt \
  --server.tls.server-host-override=grid-logger.deepsquare.run \
  --server.endpoint=grid-logger.deepsquare.run:443 \
  --pipe.path="/tmp/$SLURM_JOB_NAME-pipe" \
  --log-name="$SLURM_JOB_NAME" \
  --user="$USER" \
  >/dev/stdout 2>/dev/stderr &
LOGGER_PID="$!"
sleep 1
exec &>>"/tmp/$SLURM_JOB_NAME-pipe"
{{- end }}

{{- range $env := .Env }}
export {{ $env.Key | squote }}={{ $env.Value | squote }}
{{- end }}

{{- range $step := .Steps }}
{{ $step | renderStep }}
{{- end }}

{{- if and .EnableLogging (derefBool .EnableLogging ) }}
kill $LOGGER_PID
wait $LOGGER_PID
{{- end }}

#!/bin/bash -l

set -e

export NTASKS='{{ .Resources.Tasks }}'
export CPUS_PER_TASK='{{ .Resources.CpusPerTask }}'
export MEM_PER_CPU='{{ .Resources.MemPerCPU }}'
export GPUS_PER_TASK='{{ .Resources.GpusPerTask }}'
export GPUS='{{ mul .Resources.GpusPerTask .Resources.Tasks }}'
export CPUS='{{ mul .Resources.CpusPerTask .Resources.Tasks }}'
export MEM='{{ mul .Resources.MemPerCPU .Resources.CpusPerTask .Resources.Tasks }}'

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
/usr/bin/sleep 1
exec 3>&1
exec 4>&2
exec 1>>"/tmp/$SLURM_JOB_NAME-pipe"
exec 2>&1

disposeLogs() {
  echo cleaning up
  /usr/bin/sleep 15
  exec 1>&3
  exec 2>&4
  echo killing logger
  kill $LOGGER_PID || true
  wait $LOGGER_PID || true
  echo cleaned
}
trap disposeLogs EXIT
{{- end }}
export STORAGE_PATH="/opt/cache/shared/$UID/$SLURM_JOB_NAME"
/usr/bin/mkdir -p "$STORAGE_PATH" "$STORAGE_PATH/output/" "$STORAGE_PATH/input/"
/usr/bin/touch "$STORAGE_PATH/env"
/usr/bin/chmod -R 700 "$STORAGE_PATH"
/usr/bin/chown -R "$UID:cluster-users" "$STORAGE_PATH"
loadDeepsquareEnv() {
  /usr/bin/grep -v '^#' "$STORAGE_PATH/env" | /usr/bin/grep '=' | /usr/bin/sed -Ez '$ s/\n+$//' | tr '\n' ','
}
{{- if and .Input .Input.HTTP }}
cd $STORAGE_PATH/input/
/usr/bin/curl -fsORSL {{ .Input.HTTP.URL | squote }}
for filepath in "$STORAGE_PATH/input/"*; do
  /usr/bin/tar -xvaf "$filepath" 2>/dev/null && continue
  case $(file "$filepath") in
      *bzip2*) bzip2 -fdk "$filepath";;
      *gzip*) gunzip -df "$filepath";;
      *zip*) ;&
      *Zip*) unzip -o "$filepath";;
      *xz*) ;&
      *XZ*) unxz -f "$filepath";;
      *'7-zip'*) 7z x "$filepath" -aoa;;
      *) 1>&2 /usr/bin/echo "Unknown archive '$filepath'";;
  esac
done
cd -
/usr/bin/echo "Input contains:"
/usr/bin/find "$STORAGE_PATH/input/" -exec realpath --relative-to "$STORAGE_PATH/input/" {} \;
{{- else if and .Input .Input.S3 }}
export AWS_ACCESS_KEY_ID={{ .Input.S3.AccessKeyID | squote }}
export AWS_SECRET_ACCESS_KEY={{ .Input.S3.SecretAccessKey | squote }}
export S3_ENDPOINT_URL={{ .Input.S3.EndpointURL | squote }}

s5cmd cp --source-region {{ .Input.S3.Region | squote }} {{ .Input.S3.BucketURL | squote }}{{ .Input.S3.Path | squote }}'*' "$STORAGE_PATH/input/"
/usr/bin/echo "Input contains:"
/usr/bin/find "$STORAGE_PATH/input/" -exec realpath --relative-to "$STORAGE_PATH/input/" {} \;
{{- end }}
{{- if .InputMode }}
/usr/bin/chmod -R {{ .InputMode | derefInt | octal }} "$STORAGE_PATH/input/"
{{- end }}

{{- if and .Output .Output.HTTP .ContinuousOutputSync (derefBool .ContinuousOutputSync) }}
/usr/bin/echo "Continous output sync is not avaible with HTTP. Will use simple output."
{{- else if and .Output .Output.S3 .ContinuousOutputSync (derefBool .ContinuousOutputSync) }}
ContinuousOutputSync() {
  export AWS_ACCESS_KEY_ID={{ .Output.S3.AccessKeyID | squote }}
  export AWS_SECRET_ACCESS_KEY={{ .Output.S3.SecretAccessKey | squote }}
  export S3_ENDPOINT_URL={{ .Output.S3.EndpointURL | squote }}
  while true; do
    s5cmd sync --destination-region {{ .Output.S3.Region | squote }} "$STORAGE_PATH/output/" {{ .Output.S3.BucketURL | squote }}{{ .Output.S3.Path | squote }}
    /usr/bin/sleep 5
  done
}
ContinuousOutputSync &
CONTINUOUS_SYNC_PID="$!"
{{- end }}

{{- range $env := .Env }}
export {{ $env.Key | squote }}={{ $env.Value | squote }}
{{- end }}

{{- range $step := .Steps }}
{{ $step | renderStep }}
{{- end }}

{{- if and .Output .Output.HTTP }}
/usr/bin/echo "Output contains:"
/usr/bin/find "$STORAGE_PATH/output/" -exec realpath --relative-to "$STORAGE_PATH/output/" {} \;
/usr/bin/tar -cvf "$STORAGE_PATH/output.tar" "$STORAGE_PATH/output/"
/usr/bin/curl --upload-file "$STORAGE_PATH/output.tar" {{ .Output.HTTP.URL | squote }}
{{- else if and .Output .Output.S3 }}
{{- if and .ContinuousOutputSync (derefBool .ContinuousOutputSync) }}
kill $CONTINUOUS_SYNC_PID || true
wait $CONTINUOUS_SYNC_PID || true
{{- end }}
/usr/bin/echo "Output contains:"
/usr/bin/find "$STORAGE_PATH/output/" -exec realpath --relative-to "$STORAGE_PATH/output/" {} \;
export AWS_ACCESS_KEY_ID={{ .Output.S3.AccessKeyID | squote }}
export AWS_SECRET_ACCESS_KEY={{ .Output.S3.SecretAccessKey | squote }}
export S3_ENDPOINT_URL={{ .Output.S3.EndpointURL | squote }}

s5cmd sync --destination-region {{ .Output.S3.Region | squote }} "$STORAGE_PATH/output/" {{ .Output.S3.BucketURL | squote }}{{ .Output.S3.Path | squote }}
{{- end }}

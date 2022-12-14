#!/bin/bash -l

set -e

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
export STORAGE_PATH="/opt/cache/shared/$UID/$SLURM_JOB_NAME"
mkdir -p "$STORAGE_PATH" "$STORAGE_PATH/output/" "$STORAGE_PATH/input/"
chmod -R 700 "$STORAGE_PATH"
chown -R "$UID:cluster-users" "$STORAGE_PATH"

{{- if and .Input .Input.HTTP }}
curl -JORSL {{ .Input.HTTP.URL | squote }} -o "$STORAGE_PATH/input/"
for filepath in "$STORAGE_PATH/input/"*; do
  tar -xvaf "$filepath" 2>/dev/null && continue
  case $(file "$filepath") in
      *bzip2*) bzip2 -dk "$filepath";;
      *gzip*) gunzip "$filepath";;
      *zip*) ;&
      *Zip*) unzip "$filepath";;
      *xz*) ;&
      *XZ*) unxz  "$filepath";;
      *'7-zip'*) 7z x "$filepath";;
      *) 1>&2 echo "Unknown archive '$filepath'";;
  esac
done
echo "Input contains:"
find "$STORAGE_PATH/input/"
{{- else if and .Input .Input.S3 }}
export AWS_ACCESS_KEY_ID={{ .Input.S3.AccessKey | squote }}
export AWS_SECRET_ACCESS_KEY={{ .Input.S3.SecretAccessKey | squote }}
export S3_ENDPOINT_URL={{ .Input.S3.EndpointURL | squote }}

s5cmd sync --source-region {{ .Input.S3.Region | squote }} {{ .Input.S3.BucketURL | squote }}{{ .Input.S3.Path | squote }} "$STORAGE_PATH/input/"
echo "Input contains:"
find "$STORAGE_PATH/input/"
{{- end }}

{{- if and .Output .Output.HTTP .ContinuousOutputSync (derefBool .ContinuousOutputSync) }}
echo "Continous output sync is not avaible with HTTP. Will use simple output."
{{- else if and .Output .Output.S3 .ContinuousOutputSync (derefBool .ContinuousOutputSync) }}
ContinuousOutputSync() {
  export AWS_ACCESS_KEY_ID={{ .Output.S3.AccessKey | squote }}
  export AWS_SECRET_ACCESS_KEY={{ .Output.S3.SecretAccessKey | squote }}
  export S3_ENDPOINT_URL={{ .Output.S3.EndpointURL | squote }}
  while true; do
    s5cmd sync --destination-region {{ .Output.S3.Region | squote }} "$STORAGE_PATH/outputs/" {{ .Output.S3.BucketURL | squote }}{{ .Output.S3.Path | squote }}
    sleep 5
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
echo "Output contains:"
find "$STORAGE_PATH/output/"
tar -cvf "$STORAGE_PATH/output.tar" "$STORAGE_PATH/output/"
curl --upload-file "$STORAGE_PATH/output.tar" {{ .Output.HTTP.URL | squote }}
{{- else if and .Output .Output.S3 }}
{{- if and .ContinuousOutputSync (derefBool .ContinuousOutputSync) }}
kill $CONTINUOUS_SYNC_PID
wait $CONTINUOUS_SYNC_PID
{{- end }}
echo "Output contains:"
find "$STORAGE_PATH/output/"
export AWS_ACCESS_KEY_ID={{ .Output.S3.AccessKey | squote }}
export AWS_SECRET_ACCESS_KEY={{ .Output.S3.SecretAccessKey | squote }}
export S3_ENDPOINT_URL={{ .Output.S3.EndpointURL | squote }}

s5cmd sync --destination-region {{ .Output.S3.Region | squote }} "$STORAGE_PATH/outputs/" {{ .Output.S3.BucketURL | squote }}{{ .Output.S3.Path | squote }}
{{- end }}

{{- if and .EnableLogging (derefBool .EnableLogging ) }}
kill $LOGGER_PID
wait $LOGGER_PID
{{- end }}

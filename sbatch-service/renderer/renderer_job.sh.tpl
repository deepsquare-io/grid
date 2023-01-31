#!/bin/bash -l

set -e

export NTASKS='{{ .Job.Resources.Tasks }}'
export CPUS_PER_TASK='{{ .Job.Resources.CpusPerTask }}'
export MEM_PER_CPU='{{ .Job.Resources.MemPerCPU }}'
export GPUS_PER_TASK='{{ .Job.Resources.GpusPerTask }}'
export GPUS='{{ mul .Job.Resources.GpusPerTask .Job.Resources.Tasks }}'
export CPUS='{{ mul .Job.Resources.CpusPerTask .Job.Resources.Tasks }}'
export MEM='{{ mul .Job.Resources.MemPerCPU .Job.Resources.CpusPerTask .Job.Resources.Tasks }}'

{{- if and .Job.EnableLogging (derefBool .Job.EnableLogging ) }}
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
(
{{- end }}
export STORAGE_PATH="/opt/cache/shared/$UID/$SLURM_JOB_NAME"
export DEEPSQUARE_INPUT="$STORAGE_PATH/input"
export DEEPSQUARE_OUTPUT="$STORAGE_PATH/output"
export DEEPSQUARE_ENV="$STORAGE_PATH/env"
/usr/bin/mkdir -p "$STORAGE_PATH" "$DEEPSQUARE_OUTPUT" "$DEEPSQUARE_INPUT"
/usr/bin/touch "$DEEPSQUARE_ENV"
/usr/bin/chmod -R 700 "$STORAGE_PATH"
/usr/bin/chown -R "$UID:cluster-users" "$STORAGE_PATH"
loadDeepsquareEnv() {
  /usr/bin/grep -v '^#' "$DEEPSQUARE_ENV" | /usr/bin/grep '=' | /usr/bin/sed -Ez '$ s/\n+$//' | tr '\n' ','
}
{{- if and .Job.Input .Job.Input.HTTP }}
cd $DEEPSQUARE_INPUT/
/usr/bin/curl -fsORSL {{ .Job.Input.HTTP.URL | squote }}
for filepath in "$DEEPSQUARE_INPUT/"*; do
  /usr/bin/tar -xvaf "$filepath" 2>/dev/null && continue
  case $(file "$filepath") in
      *bzip2*) /usr/bin/bzip2 -fdk "$filepath";;
      *gzip*) /usr/bin/gunzip -df "$filepath";;
      *zip*) ;&
      *Zip*) /usr/bin/unzip -o "$filepath";;
      *RAR*) /usr/bin/unrar x -o+ "$filepath";;
      *xz*) ;&
      *XZ*) /usr/bin/unxz -f "$filepath";;
      *'7-zip'*) /usr/bin/7z x "$filepath" -aoa;;
      *) 1>&2 /usr/bin/echo "Unknown archive '$filepath'";;
  esac
done
cd -
/usr/bin/chmod -R 700 "$DEEPSQUARE_INPUT/" || echo "chmod failed, but we are ignoring it"
/usr/bin/echo "Input contains:"
/usr/bin/find "$DEEPSQUARE_INPUT/" -exec realpath --relative-to "$DEEPSQUARE_INPUT/" {} \;
{{- else if and .Job.Input .Job.Input.S3 }}
export AWS_ACCESS_KEY_ID={{ .Job.Input.S3.AccessKeyID | squote }}
export AWS_SECRET_ACCESS_KEY={{ .Job.Input.S3.SecretAccessKey | squote }}
export S3_ENDPOINT_URL={{ .Job.Input.S3.EndpointURL | squote }}

s5cmd cp --source-region {{ .Job.Input.S3.Region | squote }} {{ .Job.Input.S3.BucketURL | squote }}{{ .Job.Input.S3.Path | squote }}'*' "$DEEPSQUARE_INPUT/"
/usr/bin/chmod -R 700 "$DEEPSQUARE_INPUT/" || echo "chmod failed, but we are ignoring it"
/usr/bin/echo "Input contains:"
/usr/bin/find "$DEEPSQUARE_INPUT/" -exec realpath --relative-to "$DEEPSQUARE_INPUT/" {} \;
{{- end }}
{{- if .Job.InputMode }}
/usr/bin/chmod -R {{ .Job.InputMode | derefInt | octal }} "$DEEPSQUARE_INPUT/"
{{- end }}

{{- if and .Job.Output .Job.Output.HTTP .Job.ContinuousOutputSync (derefBool .Job.ContinuousOutputSync) }}
/usr/bin/echo "Continous output sync is not avaible with HTTP. Will use simple output."
{{- else if and .Job.Output .Job.Output.S3 .Job.ContinuousOutputSync (derefBool .Job.ContinuousOutputSync) }}
ContinuousOutputSync() {
  export AWS_ACCESS_KEY_ID={{ .Job.Output.S3.AccessKeyID | squote }}
  export AWS_SECRET_ACCESS_KEY={{ .Job.Output.S3.SecretAccessKey | squote }}
  export S3_ENDPOINT_URL={{ .Job.Output.S3.EndpointURL | squote }}
  while true; do
    s5cmd sync {{ if and .Job.Output.S3.DeleteSync (derefBool .Job.Output.S3.DeleteSync) }}--delete {{ end }}--destination-region {{ .Job.Output.S3.Region | squote }} "$DEEPSQUARE_OUTPUT/" {{ .Job.Output.S3.BucketURL | squote }}{{ .Job.Output.S3.Path | squote }}
    /usr/bin/sleep 5
  done
}
ContinuousOutputSync &
CONTINUOUS_SYNC_PID="$!"
(
{{- end }}

{{- range $env := .Job.Env }}
export {{ $env.Key | squote }}={{ $env.Value | squote }}
{{- end }}

{{- range $step := .Job.Steps }}
{{ renderStep $.Job $step }}
{{- end }}

{{- if and .Job.Output .Job.Output.HTTP }}
/usr/bin/echo "Output contains:"
/usr/bin/find "$DEEPSQUARE_OUTPUT/" -exec realpath --relative-to "$DEEPSQUARE_OUTPUT/" {} \;
cd $DEEPSQUARE_OUTPUT/..
if [ "$(find output/ -type f | wc -l)" -eq 1 ]; then
/usr/bin/curl --upload-file "$(find output/ -type f | wc -l)" {{ .Job.Output.HTTP.URL | squote }}
else
/usr/bin/zip -r "output.zip" "output/"
/usr/bin/curl --upload-file "output.zip" {{ .Job.Output.HTTP.URL | squote }}
fi
/usr/bin/echo
cd -
{{- else if and .Job.Output .Job.Output.S3 }}
{{- if and .Job.ContinuousOutputSync (derefBool .Job.ContinuousOutputSync) }}
)
kill $CONTINUOUS_SYNC_PID || true
wait $CONTINUOUS_SYNC_PID || true
{{- end }}
/usr/bin/echo "Output contains:"
/usr/bin/find "$DEEPSQUARE_OUTPUT/" -exec realpath --relative-to "$DEEPSQUARE_OUTPUT/" {} \;
export AWS_ACCESS_KEY_ID={{ .Job.Output.S3.AccessKeyID | squote }}
export AWS_SECRET_ACCESS_KEY={{ .Job.Output.S3.SecretAccessKey | squote }}
export S3_ENDPOINT_URL={{ .Job.Output.S3.EndpointURL | squote }}

s5cmd sync {{ if and .Job.Output.S3.DeleteSync (derefBool .Job.Output.S3.DeleteSync) }}--delete {{ end }}--destination-region {{ .Job.Output.S3.Region | squote }} "$DEEPSQUARE_OUTPUT/" {{ .Job.Output.S3.BucketURL | squote }}{{ .Job.Output.S3.Path | squote }}
{{- end }}
{{- if and .Job.EnableLogging (derefBool .Job.EnableLogging ) }}
)
{{- end }}

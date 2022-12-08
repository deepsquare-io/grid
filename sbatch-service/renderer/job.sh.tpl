set -e

export STORAGE_PATH="/opt/cache/shared/$UID/$SLURM_JOB_NAME"
mkdir -p "$STORAGE_PATH"
chmod -R 700 "$STORAGE_PATH"
chown -R "$UID:cluster-users" "$STORAGE_PATH"

{{- range $env := .Env }}
export {{ $env.Key | squote }}={{ $env.Value | squote }}
{{- end }}

{{- range $step := .Steps }}
{{ $step | renderStep }}
{{- end }}

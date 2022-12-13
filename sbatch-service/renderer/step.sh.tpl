{{- if .Run -}}
{{- if and .Run.X11 (derefBool .Run.X11 ) .Run.Image -}}
MOUNTS="/tmp/.X11-unix:/tmp/.X11-unix:ro"
{{- end }}
srun --job-name={{ .Name | squote }} \
  --export=ALL{{ range $env := .Run.Env }},{{ $env.Key | squote }}={{ $env.Value | squote }}{{ end }} \
  --cpus-per-task={{ .Run.Resources.CpusPerTask }} \
  --mem-per-cpu={{ .Run.Resources.MemPerCPU }} \
  --gpus-per-task={{ .Run.Resources.GpusPerTask }} \
  --ntasks={{ .Run.Resources.Tasks }} \
{{- if and .Run.Image (derefStr .Run.Image ) }}
{{- if and .Run.X11 (derefBool .Run.X11 ) }}
  --container-mounts="${MOUNTS}" \
{{- end }}
  --container-image={{ .Run.Image | derefStr | squote }} \
{{- end }}
  sh -c {{ .Run.Command | squote }}
{{- else if .For -}}
doFor() {
{{- if .For.Range }}
  export index="$1"
{{- else if .For.Items }}
  export item="$1"
{{- end }}

{{- range $step := .For.Steps -}}
  {{ $step | renderStep | nindent 2 }}
{{- end }}
}

{{- if .For.Parallel }}
pids=()
{{- end }}
{{- if .For.Range }}
for index in $(seq {{ .For.Range.Begin }} {{ .For.Range.Increment }} {{ .For.Range.End }}); do
  doFor "$index" {{if .For.Parallel}}&{{end}}
  {{- if .For.Parallel }}
  pids+=("$!")
  {{- end }}
done
{{- else if .For.Items }}
items=({{ range $item := .For.Items }}{{ $item | squote }} {{ end }})
for item in "${items[@]}"; do
  doFor "$item" {{ if .For.Parallel }}&{{ end }}
  {{- if .For.Parallel }}
  pids+=("$!")
  {{- end }}
done
{{- end }}
{{- if .For.Parallel }}
for pid in "${pids[@]}"; do
  wait "$pid"
done
{{- end }}
{{- end -}}

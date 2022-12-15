{{- if .Run -}}
{{- if .Run.Image -}}
MOUNTS="$STORAGE_PATH:/deepsquare:rw{{ if and .Run.X11 (derefBool .Run.X11 ) }},/tmp/.X11-unix:/tmp/.X11-unix:ro{{ end }}"
{{- end }}
{{ if and .Run.Image (derefStr .Run.Image ) }}STORAGE_PATH=/deepsquare {{ end }}/usr/bin/srun --job-name={{ .Name | squote }} \
  --export=ALL,"$(loadDeepsquareEnv)"{{ range $env := .Run.Env }},{{ $env.Key | squote }}={{ $env.Value | squote }}{{ end }} \
  --cpus-per-task={{ .Run.Resources.CpusPerTask }} \
  --mem-per-cpu={{ .Run.Resources.MemPerCPU }} \
  --gpus-per-task={{ .Run.Resources.GpusPerTask }} \
  --ntasks={{ .Run.Resources.Tasks }} \
{{- if and .Run.Image (derefStr .Run.Image ) }}
  --container-mounts="${MOUNTS}" \
  --container-image={{ .Run.Image | derefStr | squote }} \
{{- end }}
  {{ if .Run.Shell }}{{ derefStr .Run.Shell }}{{ else }}/bin/sh{{ end }} -c {{ .Run.Command | squote }}
{{- else if .For -}}
doFor() {
{{- if .For.Range }}
export index="$1"
{{- else if .For.Items }}
export item="$1"
{{- end }}

{{- range $step := .For.Steps }}
{{ $step | renderStep }}
{{- end }}
}

{{- if .For.Parallel }}
pids=()
{{- end }}
{{- if .For.Range }}
for index in $(seq {{ .For.Range.Begin }} {{ if ne .For.Range.Increment 0 }}{{ .For.Range.Increment }}{{ else }}1{{ end }} {{ .For.Range.End }}); do
  doFor "$index" {{ if .For.Parallel }}&{{ end }}
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
  /usr/bin/wait "$pid"
done
{{- end }}
{{- end -}}

{{- if .Run -}}
srun --job-name={{ .Name | squote }} \
  --cpus-per-task={{ .Run.Resources.CpusPerTask }} \
  --mem-per-cpu={{ .Run.Resources.MemPerCPU }} \
  --gpus-per-task={{ .Run.Resources.GpusPerTask }} \
  --ntasks={{ .Run.Resources.Tasks }} \
{{- if and .Run.Image (deref .Run.Image ) }}
  --container-image={{ .Run.Image | deref | squote }} \
{{- end }}
  sh -c {{ .Run.Command | squote }}
{{- else if .For }}

{{- end }}

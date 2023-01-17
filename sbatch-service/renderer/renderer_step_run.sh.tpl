{{- if .Step.Run.Image -}}
MOUNTS="$STORAGE_PATH:/deepsquare:rw{{ if and .Step.Run.X11 (derefBool .Step.Run.X11 ) }},/tmp/.X11-unix:/tmp/.X11-unix:ro{{ end }}"
{{ end }}
{{- if and .Step.Run.Image (derefStr .Step.Run.Image ) }}STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env' {{ end }}/usr/bin/srun --job-name={{ .Step.Name | squote }} \
  --export=ALL,"$(loadDeepsquareEnv)"{{ range $env := .Step.Run.Env }},{{ $env.Key | squote }}={{ $env.Value | squote }}{{ end }} \
  --cpus-per-task={{ default .Step.Run.Resources.CpusPerTask (derefInt .Step.Run.Resources.CpusPerTask) .Job.Resources.CpusPerTask }} \
  --mem-per-cpu={{ default .Step.Run.Resources.MemPerCPU (derefInt .Step.Run.Resources.MemPerCPU) .Job.Resources.MemPerCPU }} \
  --gpus-per-task={{ default .Step.Run.Resources.GpusPerTask (derefInt .Step.Run.Resources.GpusPerTask) .Job.Resources.GpusPerTask }} \
  --ntasks={{ default .Step.Run.Resources.Tasks (derefInt .Step.Run.Resources.Tasks) 1 }} \
{{- if and .Step.Run.DisableCPUBinding (derefBool .Step.Run.DisableCPUBinding ) }}
  --cpu-bind=none \
{{- end }}
{{- if and .Step.Run.Image (derefStr .Step.Run.Image ) }}
  --gpu-bind=none \
  --container-mounts="${MOUNTS}" \
  --container-image={{ .Step.Run.Image | derefStr | squote }} \
{{- end }}
  {{ if .Step.Run.Shell }}{{ derefStr .Step.Run.Shell }}{{ else }}/bin/sh{{ end }} -c {{ .Step.Run.Command | squote -}}

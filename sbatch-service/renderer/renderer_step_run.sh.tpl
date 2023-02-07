{{- if .Step.Run.Container -}}
{{- if or (and .Step.Run.Container.Apptainer (derefBool .Step.Run.Container.Apptainer)) (and .Step.Run.Container.DeepsquareHosted (derefBool .Step.Run.Container.DeepsquareHosted)) -}}
export APPTAINER_BIND="$STORAGE_PATH:/deepsquare:rw{{ if and .Step.Run.Container.X11 (derefBool .Step.Run.Container.X11 ) }},/tmp/.X11-unix:/tmp/.X11-unix:ro{{ end }}"{{ range $mount := .Step.Run.Container.Mounts }},{{ $mount.HostDir | squote }}:{{ $mount.ContainerDir | squote }}:{{ $mount.Options | squote }}{{ end }}
{{- if and .Step.Run.Container.Registry .Step.Run.Container.Username .Step.Run.Container.Password }}
export APPTAINER_DOCKER_USERNAME={{ derefStr .Step.Run.Container.Username | squote }}
export APPTAINER_DOCKER_PASSWORD={{ derefStr .Step.Run.Container.Password | squote }}
{{- end }}
# shellcheck disable=SC2097,SC2098
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env'{{ range $env := .Step.Run.Env }} {{ $env.Key | squote }}={{ $env.Value | squote }}{{ end }} /usr/bin/srun --job-name={{ .Step.Name | squote }} \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task={{ default .Step.Run.Resources.CpusPerTask (derefInt .Step.Run.Resources.CpusPerTask) .Job.Resources.CpusPerTask }} \
  --mem-per-cpu={{ default .Step.Run.Resources.MemPerCPU (derefInt .Step.Run.Resources.MemPerCPU) .Job.Resources.MemPerCPU }}M \
  --gpus-per-task={{ default .Step.Run.Resources.GpusPerTask (derefInt .Step.Run.Resources.GpusPerTask) .Job.Resources.GpusPerTask }} \
  --ntasks={{ default .Step.Run.Resources.Tasks (derefInt .Step.Run.Resources.Tasks) 1 }} \
{{- if and .Step.Run.DisableCPUBinding (derefBool .Step.Run.DisableCPUBinding ) }}
  --cpu-bind=none \
{{- end }}
  --gpu-bind=none \
  /usr/bin/apptainer --silent exec \
  --nv \
  {{ if and (hasPrefix "/" .Step.Run.Container.Image) (not (and .Step.Run.Container.DeepsquareHosted (derefBool .Step.Run.Container.DeepsquareHosted))) }}"$STORAGE_PATH"{{ end }}{{ formatImageURL .Step.Run.Container.Registry .Step.Run.Container.Image .Step.Run.Container.Apptainer .Step.Run.Container.DeepsquareHosted | squote }} \
  {{ if .Step.Run.Shell }}{{ derefStr .Step.Run.Shell }}{{ else }}/bin/sh{{ end }} -c {{ .Step.Run.Command | squote -}}
{{- else -}}
{{- if and .Step.Run.Container.Registry .Step.Run.Container.Username .Step.Run.Container.Password -}}
/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine {{ derefStr .Step.Run.Container.Registry | quoteEscape }} login {{ derefStr .Step.Run.Container.Username | quoteEscape }} password {{ derefStr .Step.Run.Container.Password | quoteEscape }}
EOFnetrc
{{- end }}
MOUNTS="$STORAGE_PATH:/deepsquare:rw{{ if and .Step.Run.Container.X11 (derefBool .Step.Run.Container.X11 ) }},/tmp/.X11-unix:/tmp/.X11-unix:ro{{ end }}"{{ range $mount := .Step.Run.Container.Mounts }},{{ $mount.HostDir | squote }}:{{ $mount.ContainerDir | squote }}:{{ $mount.Options | squote }}{{ end }}
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env'{{ range $env := .Step.Run.Env }} {{ $env.Key | squote }}={{ $env.Value | squote }}{{ end }} /usr/bin/srun --job-name={{ .Step.Name | squote }} \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task={{ default .Step.Run.Resources.CpusPerTask (derefInt .Step.Run.Resources.CpusPerTask) .Job.Resources.CpusPerTask }} \
  --mem-per-cpu={{ default .Step.Run.Resources.MemPerCPU (derefInt .Step.Run.Resources.MemPerCPU) .Job.Resources.MemPerCPU }}M \
  --gpus-per-task={{ default .Step.Run.Resources.GpusPerTask (derefInt .Step.Run.Resources.GpusPerTask) .Job.Resources.GpusPerTask }} \
  --ntasks={{ default .Step.Run.Resources.Tasks (derefInt .Step.Run.Resources.Tasks) 1 }} \
{{- if and .Step.Run.DisableCPUBinding (derefBool .Step.Run.DisableCPUBinding ) }}
  --cpu-bind=none \
{{- end }}
  --gpu-bind=none \
  --container-mounts="${MOUNTS}" \
  --container-image={{ formatImageURL .Step.Run.Container.Registry .Step.Run.Container.Image .Step.Run.Container.Apptainer .Step.Run.Container.DeepsquareHosted | squote }} \
  {{ if .Step.Run.Shell }}{{ derefStr .Step.Run.Shell }}{{ else }}/bin/sh{{ end }} -c {{ .Step.Run.Command | squote -}}
{{- end -}}
{{- else -}}
{{ range $env := .Step.Run.Env }}{{ $env.Key | squote }}={{ $env.Value | squote }} {{ end }}/usr/bin/srun --job-name={{ .Step.Name | squote }} \
  --export=ALL"$(loadDeepsquareEnv)" \
  --cpus-per-task={{ default .Step.Run.Resources.CpusPerTask (derefInt .Step.Run.Resources.CpusPerTask) .Job.Resources.CpusPerTask }} \
  --mem-per-cpu={{ default .Step.Run.Resources.MemPerCPU (derefInt .Step.Run.Resources.MemPerCPU) .Job.Resources.MemPerCPU }}M \
  --gpus-per-task={{ default .Step.Run.Resources.GpusPerTask (derefInt .Step.Run.Resources.GpusPerTask) .Job.Resources.GpusPerTask }} \
  --ntasks={{ default .Step.Run.Resources.Tasks (derefInt .Step.Run.Resources.Tasks) 1 }} \
{{- if and .Step.Run.DisableCPUBinding (derefBool .Step.Run.DisableCPUBinding ) }}
  --cpu-bind=none \
{{- end }}
  {{ if .Step.Run.Shell }}{{ derefStr .Step.Run.Shell }}{{ else }}/bin/sh{{ end }} -c {{ .Step.Run.Command | squote -}}
{{- end -}}

{{- if .Step.Run.Container -}}
{{- if or (and .Step.Run.Container.Apptainer (derefBool .Step.Run.Container.Apptainer)) (and .Step.Run.Container.DeepsquareHosted (derefBool .Step.Run.Container.DeepsquareHosted)) -}}
export APPTAINER_BIND="$STORAGE_PATH:/deepsquare:rw{{ if and .Step.Run.Container.X11 (derefBool .Step.Run.Container.X11 ) }},/tmp/.X11-unix:/tmp/.X11-unix:ro{{ end }}"{{ range $mount := .Step.Run.Container.Mounts }},{{ $mount.HostDir | squote }}:{{ $mount.ContainerDir | squote }}:{{ $mount.Options | squote }}{{ end }}
{{- if and .Step.Run.Container.Registry .Step.Run.Container.Username .Step.Run.Container.Password }}
export APPTAINER_DOCKER_USERNAME={{ derefStr .Step.Run.Container.Username | squote }}
export APPTAINER_DOCKER_PASSWORD={{ derefStr .Step.Run.Container.Password | squote }}
{{- end }}
# shellcheck disable=SC2097,SC2098
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env'{{ range $env := .Step.Run.Env }} {{ $env.Key }}={{ $env.Value | squote }}{{ end }} /usr/bin/srun --job-name={{ .Step.Name | squote }} \
  --export=ALL"$(loadDeepsquareEnv)" \
{{- if and .Step.Run.Resources .Step.Run.Resources.CpusPerTask (derefInt .Step.Run.Resources.CpusPerTask) }}
  --cpus-per-task={{ derefInt .Step.Run.Resources.CpusPerTask }} \
{{- else }}
  --cpus-per-task={{ .Job.Resources.CpusPerTask }} \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.MemPerCPU (derefInt .Step.Run.Resources.MemPerCPU) }}
  --mem-per-cpu={{ derefInt .Step.Run.Resources.MemPerCPU }}M \
{{- else }}
  --mem-per-cpu={{ .Job.Resources.MemPerCPU }}M \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.GpusPerTask }}
  --gpus-per-task={{ derefInt .Step.Run.Resources.GpusPerTask }} \
{{- else }}
  --gpus-per-task={{ .Job.Resources.GpusPerTask }} \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.Tasks (derefInt .Step.Run.Resources.Tasks) }}
  --ntasks={{ derefInt .Step.Run.Resources.Tasks }} \
{{- else }}
  --ntasks=1 \
{{- end }}
{{- if and .Step.Run.Mpi (derefStr .Step.Run.Mpi) }}
  --mpi={{ derefStr .Step.Run.Mpi }} \
{{- end }}
{{- if and .Step.Run.DisableCPUBinding (derefBool .Step.Run.DisableCPUBinding ) }}
  --cpu-bind=none \
{{- end }}
  --gpu-bind=none \
  {{ if and .Step.Run.Network (eq (derefStr .Step.Run.Network) "slirp4netns") }}/bin/sh -c {{ renderSlirp4NetNS .Step.Run.CustomNetworkInterfaces .Step.Run.DNS (renderApptainerCommand .Step.Run) .Step.Run.Shell | squote -}}{{ else }}{{ renderApptainerCommand .Step.Run }}{{ end }}
{{- else -}}
{{- if and .Step.Run.Container.Registry .Step.Run.Container.Username .Step.Run.Container.Password -}}
/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine {{ derefStr .Step.Run.Container.Registry | quoteEscape }} login {{ derefStr .Step.Run.Container.Username | quoteEscape }} password {{ derefStr .Step.Run.Container.Password | quoteEscape }}
EOFnetrc
{{- end }}
{{- if and .Step.Run.Network (eq (derefStr .Step.Run.Network) "slirp4netns") }}
/usr/bin/srun --job-name={{ .Step.Name | squote }} \
  --export=ALL"$(loadDeepsquareEnv)" \
{{- if and .Step.Run.Resources .Step.Run.Resources.CpusPerTask (derefInt .Step.Run.Resources.CpusPerTask) }}
  --cpus-per-task={{ derefInt .Step.Run.Resources.CpusPerTask }} \
{{- else }}
  --cpus-per-task={{ .Job.Resources.CpusPerTask }} \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.MemPerCPU (derefInt .Step.Run.Resources.MemPerCPU) }}
  --mem-per-cpu={{ derefInt .Step.Run.Resources.MemPerCPU }}M \
{{- else }}
  --mem-per-cpu={{ .Job.Resources.MemPerCPU }}M \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.GpusPerTask }}
  --gpus-per-task={{ derefInt .Step.Run.Resources.GpusPerTask }} \
{{- else }}
  --gpus-per-task={{ .Job.Resources.GpusPerTask }} \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.Tasks (derefInt .Step.Run.Resources.Tasks) }}
  --ntasks={{ derefInt .Step.Run.Resources.Tasks }} \
{{- else }}
  --ntasks=1 \
{{- end }}
{{- if and .Step.Run.Mpi (derefStr .Step.Run.Mpi) }}
  --mpi={{ derefStr .Step.Run.Mpi }} \
{{- end }}
{{- if and .Step.Run.DisableCPUBinding (derefBool .Step.Run.DisableCPUBinding ) }}
  --cpu-bind=none \
{{- end }}
  --gpu-bind=none \
  /bin/sh -c '
{{- $image := formatImageURL .Step.Run.Container.Registry .Step.Run.Container.Image .Step.Run.Container.Apptainer .Step.Run.Container.DeepsquareHosted -}}
{{- if isAbs $image -}}
/usr/bin/enroot create --name "container-$SLURM_JOB_ID" -- "/tmp/$SLURM_JOB_ID.sqsh"
{{- else -}}
/usr/bin/enroot import -o "/tmp/$SLURM_JOB_ID.sqsh" -- "docker://"{{ $image | squote | escapeSQuote }}
/usr/bin/enroot create --name "container-$SLURM_JOB_ID" -- "/tmp/$SLURM_JOB_ID.sqsh"
/usr/bin/rm -f "/tmp/$SLURM_JOB_ID.sqsh"
{{- end }}
enrootClean() {
  /usr/bin/enroot remove -f "container-$SLURM_JOB_ID"
}
trap enrootClean EXIT INT TERM
'{{ renderSlirp4NetNS .Step.Run.CustomNetworkInterfaces .Step.Run.DNS (renderEnrootCommand .Step.Run) .Step.Run.Shell | squote -}}
{{- else }}
MOUNTS="$STORAGE_PATH:/deepsquare:rw{{ if and .Step.Run.Container.X11 (derefBool .Step.Run.Container.X11 ) }},/tmp/.X11-unix:/tmp/.X11-unix:ro{{ end }}"{{ range $mount := .Step.Run.Container.Mounts }},{{ $mount.HostDir | squote }}:{{ $mount.ContainerDir | squote }}:{{ $mount.Options | squote }}{{ end }}
STORAGE_PATH='/deepsquare' DEEPSQUARE_INPUT='/deepsquare/input' DEEPSQUARE_OUTPUT='/deepsquare/output' DEEPSQUARE_ENV='/deepsquare/env'{{ range $env := .Step.Run.Env }} {{ $env.Key }}={{ $env.Value | squote }}{{ end }} /usr/bin/srun --job-name={{ .Step.Name | squote }} \
  --export=ALL"$(loadDeepsquareEnv)" \
{{- if and .Step.Run.Resources .Step.Run.Resources.CpusPerTask (derefInt .Step.Run.Resources.CpusPerTask) }}
  --cpus-per-task={{ derefInt .Step.Run.Resources.CpusPerTask }} \
{{- else }}
  --cpus-per-task={{ .Job.Resources.CpusPerTask }} \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.MemPerCPU (derefInt .Step.Run.Resources.MemPerCPU) }}
  --mem-per-cpu={{ derefInt .Step.Run.Resources.MemPerCPU }}M \
{{- else }}
  --mem-per-cpu={{ .Job.Resources.MemPerCPU }}M \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.GpusPerTask }}
  --gpus-per-task={{ derefInt .Step.Run.Resources.GpusPerTask }} \
{{- else }}
  --gpus-per-task={{ .Job.Resources.GpusPerTask }} \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.Tasks (derefInt .Step.Run.Resources.Tasks) }}
  --ntasks={{ derefInt .Step.Run.Resources.Tasks }} \
{{- else }}
  --ntasks=1 \
{{- end }}
{{- if and .Step.Run.Mpi (derefStr .Step.Run.Mpi) }}
  --mpi={{ derefStr .Step.Run.Mpi }} \
{{- end }}
{{- if and .Step.Run.DisableCPUBinding (derefBool .Step.Run.DisableCPUBinding ) }}
  --cpu-bind=none \
{{- end }}
  --gpu-bind=none \
{{- if and .Step.Run.MapRoot (derefBool .Step.Run.MapRoot ) }}
  --container-remap-root \
{{- else }}
  --no-container-remap-root \
{{- end }}
  --no-container-mount-home \
  --container-mounts="${MOUNTS}" \
{{- if and .Step.Run.WorkDir (derefStr .Step.Run.WorkDir) }}
  --container-workdir={{ derefStr .Step.Run.WorkDir | squote }} \
{{- end }}
  --container-image={{ formatImageURL .Step.Run.Container.Registry .Step.Run.Container.Image .Step.Run.Container.Apptainer .Step.Run.Container.DeepsquareHosted | squote }} \
  {{ if .Step.Run.Shell }}{{ derefStr .Step.Run.Shell }}{{ else }}/bin/sh{{ end }} -c {{ .Step.Run.Command | squote -}}
{{- end -}}
{{- end }}
{{- else -}}
{{ range $env := .Step.Run.Env }}{{ $env.Key }}={{ $env.Value | squote }} {{ end }}/usr/bin/srun --job-name={{ .Step.Name | squote }} \
  --export=ALL"$(loadDeepsquareEnv)" \
{{- if and .Step.Run.Resources .Step.Run.Resources.CpusPerTask (derefInt .Step.Run.Resources.CpusPerTask) }}
  --cpus-per-task={{ derefInt .Step.Run.Resources.CpusPerTask }} \
{{- else }}
  --cpus-per-task={{ .Job.Resources.CpusPerTask }} \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.MemPerCPU (derefInt .Step.Run.Resources.MemPerCPU) }}
  --mem-per-cpu={{ derefInt .Step.Run.Resources.MemPerCPU }}M \
{{- else }}
  --mem-per-cpu={{ .Job.Resources.MemPerCPU }}M \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.GpusPerTask }}
  --gpus-per-task={{ derefInt .Step.Run.Resources.GpusPerTask }} \
{{- else }}
  --gpus-per-task={{ .Job.Resources.GpusPerTask }} \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.Tasks (derefInt .Step.Run.Resources.Tasks) }}
  --ntasks={{ derefInt .Step.Run.Resources.Tasks }} \
{{- else }}
  --ntasks=1 \
{{- end }}
{{- if and .Step.Run.Mpi (derefStr .Step.Run.Mpi) }}
  --mpi={{ derefStr .Step.Run.Mpi }} \
{{- end }}
{{- if and .Step.Run.DisableCPUBinding (derefBool .Step.Run.DisableCPUBinding ) }}
  --cpu-bind=none \
{{- end }}
{{- if and .Step.Run.MapRoot (derefBool .Step.Run.MapRoot ) }}
  /usr/bin/unshare --user --map-root-user --mount \
{{- end }}
  {{ if .Step.Run.Shell }}{{ derefStr .Step.Run.Shell }}{{ else }}/bin/sh{{ end }} -c {{ if and .Step.Run.Network (eq (derefStr .Step.Run.Network) "slirp4netns") }}{{ renderSlirp4NetNS .Step.Run.CustomNetworkInterfaces .Step.Run.DNS .Step.Run.Command .Step.Run.Shell | squote -}}{{ else }}
{{- if and .Step.Run.WorkDir (derefStr .Step.Run.WorkDir) -}}
  'cd {{ derefStr .Step.Run.WorkDir | squote | escapeSQuote }} || { echo "change dir to working directory failed"; exit 1; };'{{ end -}}
  {{ .Step.Run.Command | squote -}}{{ end }}
{{- end -}}

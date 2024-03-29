{{- /* In this file contains 3 launch modes: Apptainer, Enroot and none. */ -}}
{{- $ntasks := 1 }}
{{- if and .Step.Run.Resources .Step.Run.Resources.Tasks (derefInt .Step.Run.Resources.Tasks) }}
{{- $ntasks = derefInt .Step.Run.Resources.Tasks }}
{{- end }}
{{- if .Step.Run.Container -}}
{{- if .Step.Run.Container.Mounts -}}
/usr/bin/cat << 'EOFmounterror'
WARNING: Mounts is now deprecated.
If you need a cache (disk, shared, per-user or global), please read https://docs.deepsquare.run/workflow/guides/environment-variables.
The cache is cleared periodically and only persists on the site.
EOFmounterror
{{ end -}}
{{- $image := formatImageURL .Step.Run.Container.Registry .Step.Run.Container.Image .Step.Run.Container.Apptainer .Step.Run.Container.DeepsquareHosted -}}
{{- if or (and .Step.Run.Container.Apptainer (derefBool .Step.Run.Container.Apptainer)) (and .Step.Run.Container.DeepsquareHosted (derefBool .Step.Run.Container.DeepsquareHosted)) -}}
{{- /* Apptainer */ -}}
{{- if and .Step.Run.Container.Registry .Step.Run.Container.Username .Step.Run.Container.Password -}}
export APPTAINER_DOCKER_USERNAME={{ derefStr .Step.Run.Container.Username | squote }}
export APPTAINER_DOCKER_PASSWORD={{ derefStr .Step.Run.Container.Password | squote }}
{{ end -}}
{{- if not (or (isAbs $image) (and .Step.Run.Container.DeepsquareHosted (derefBool .Step.Run.Container.DeepsquareHosted))) -}}
IMAGE_PATH="$STORAGE_PATH/$SLURM_JOB_ID-$(echo $RANDOM | md5sum | head -c 20).sif"
export IMAGE_PATH
/usr/bin/echo "Importing image..."
/usr/bin/apptainer --silent pull --disable-cache "$IMAGE_PATH" {{ $image | squote }}
/usr/bin/echo "Image successfully imported!"
{{ end -}}
/usr/bin/srun {{ if and .Step.Name (derefStr .Step.Name) }}--job-name={{ derefStr .Step.Name | squote }}{{ end }} \
  --export=ALL"$(loadDeepsquareEnv)" \
{{- if and .Step.Run.Resources .Step.Run.Resources.CPUsPerTask (derefInt .Step.Run.Resources.CPUsPerTask) }}
  --cpus-per-task={{ derefInt .Step.Run.Resources.CPUsPerTask }} \
{{- else }}
  --cpus-per-task={{ .Job.Resources.CPUsPerTask }} \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.MemPerCPU (derefInt .Step.Run.Resources.MemPerCPU) }}
  --mem-per-cpu={{ derefInt .Step.Run.Resources.MemPerCPU }}M \
{{- else }}
  --mem-per-cpu={{ .Job.Resources.MemPerCPU }}M \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.GPUsPerTask }}
  --gpus-per-task={{ derefInt .Step.Run.Resources.GPUsPerTask }} \
  --gpus={{ mul (derefInt .Step.Run.Resources.GPUsPerTask) $ntasks }} \
{{- else }}
  --gpus-per-task=0 \
  --gpus=0 \
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
  /bin/bash -c {{ if and .Step.Run.Network (eq (derefStr .Step.Run.Network) "slirp4netns") }}{{ renderSlirp4NetNS .Step.Run .Job (renderApptainerCommand .Step.Run) | squote -}}{{ else if and .Step.Run.Network (eq (derefStr .Step.Run.Network) "pasta") }}{{ renderPastaNS .Step.Run .Job (renderApptainerCommand .Step.Run) | squote -}}{{ else }}{{ renderApptainerCommand .Step.Run | squote }}{{ end }}
{{- else -}}
{{- /* Enroot */ -}}
{{- if and .Step.Run.Container.Registry .Step.Run.Container.Username .Step.Run.Container.Password -}}
/usr/bin/mkdir -p "$HOME/.config/enroot/"
/usr/bin/cat << 'EOFnetrc' > "$HOME/.config/enroot/.credentials"
machine {{ derefStr .Step.Run.Container.Registry }} login {{ derefStr .Step.Run.Container.Username | quoteEscape }} password {{ derefStr .Step.Run.Container.Password | quoteEscape }}
EOFnetrc
{{- end }}
{{ if not (isAbs $image) -}}
IMAGE_PATH="$STORAGE_PATH/$SLURM_JOB_ID-$(echo $RANDOM | md5sum | head -c 20).sqsh"
export IMAGE_PATH
/usr/bin/echo "Importing image..."
set +e
/usr/bin/enroot import -o "$IMAGE_PATH" -- {{ $image | squote }} &> "/tmp/enroot.import.$SLURM_JOB_ID.log"
if [ $? -ne 0 ]; then
  cat "/tmp/enroot.import.$SLURM_JOB_ID.log"
fi
set -e
tries=1; while [ "$tries" -lt 10 ]; do
  if /usr/bin/file "$IMAGE_PATH" | /usr/bin/grep -q "Squashfs filesystem"; then
    break
  fi
  /usr/bin/echo "Image is not complete. Wait a few seconds... ($tries/10)"
  /usr/bin/sleep 10
  tries=$((tries+1))
done
if [ "$tries" -ge 10 ]; then
  /usr/bin/echo "Image import failure (corrupted image). Please try again."
  exit 1
fi
/usr/bin/echo "Image successfully imported!"
{{- end }}
# shellcheck disable=SC2097,SC2098,SC1078
/usr/bin/srun {{ if and .Step.Name (derefStr .Step.Name) }}--job-name={{ derefStr .Step.Name | squote }}{{ end }} \
  --export=ALL"$(loadDeepsquareEnv)" \
{{- if and .Step.Run.Resources .Step.Run.Resources.CPUsPerTask (derefInt .Step.Run.Resources.CPUsPerTask) }}
  --cpus-per-task={{ derefInt .Step.Run.Resources.CPUsPerTask }} \
{{- else }}
  --cpus-per-task={{ .Job.Resources.CPUsPerTask }} \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.MemPerCPU (derefInt .Step.Run.Resources.MemPerCPU) }}
  --mem-per-cpu={{ derefInt .Step.Run.Resources.MemPerCPU }}M \
{{- else }}
  --mem-per-cpu={{ .Job.Resources.MemPerCPU }}M \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.GPUsPerTask }}
  --gpus-per-task={{ derefInt .Step.Run.Resources.GPUsPerTask }} \
  --gpus={{ mul (derefInt .Step.Run.Resources.GPUsPerTask) $ntasks }} \
{{- else }}
  --gpus-per-task=0 \
  --gpus=0 \
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
  /bin/bash -c '
{{- if isAbs $image -}}
/usr/bin/enroot create --name "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" -- "$STORAGE_PATH"{{ $image | squote }} >/dev/null 2>&1
{{- else -}}
/usr/bin/enroot create --name "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" -- "$IMAGE_PATH" >/dev/null 2>&1
{{- end }}
enrootClean() {
  /usr/bin/enroot remove -f "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID"
}
trap enrootClean EXIT INT TERM
'{{ if and .Step.Run.Network (eq (derefStr .Step.Run.Network) "slirp4netns") }}{{ renderSlirp4NetNS .Step.Run .Job (renderEnrootCommand .Step.Run) | squote -}}{{ else if and .Step.Run.Network (eq (derefStr .Step.Run.Network) "pasta") }}{{ renderPastaNS .Step.Run .Job (renderEnrootCommand .Step.Run) | squote -}}{{ else }}{{ renderEnrootCommand .Step.Run | squote }}{{ end }}
{{- end }}
{{- else -}}
{{ range $env := .Step.Run.Env }}{{ $env.Key }}={{ $env.Value | squote }} {{ end }}/usr/bin/srun {{ if and .Step.Name (derefStr .Step.Name) }}--job-name={{ derefStr .Step.Name | squote }}{{ end }} \
  --export=ALL"$(loadDeepsquareEnv)" \
{{- if and .Step.Run.Resources .Step.Run.Resources.CPUsPerTask (derefInt .Step.Run.Resources.CPUsPerTask) }}
  --cpus-per-task={{ derefInt .Step.Run.Resources.CPUsPerTask }} \
{{- else }}
  --cpus-per-task={{ .Job.Resources.CPUsPerTask }} \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.MemPerCPU (derefInt .Step.Run.Resources.MemPerCPU) }}
  --mem-per-cpu={{ derefInt .Step.Run.Resources.MemPerCPU }}M \
{{- else }}
  --mem-per-cpu={{ .Job.Resources.MemPerCPU }}M \
{{- end }}
{{- if and .Step.Run.Resources .Step.Run.Resources.GPUsPerTask }}
  --gpus-per-task={{ derefInt .Step.Run.Resources.GPUsPerTask }} \
  --gpus={{ mul (derefInt .Step.Run.Resources.GPUsPerTask) $ntasks }} \
{{- else }}
  --gpus-per-task=0 \
  --gpus=0 \
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
{{- if or .Step.Run.MapUID .Step.Run.MapGID }}
  /usr/bin/unshare --map-current-user{{ if .Step.Run.MapUID }} --map-user={{ .Step.Run.MapUID }}{{ end }}{{ if .Step.Run.MapGID }} --map-group={{ .Step.Run.MapGID }}{{ end }} --mount \
{{- end }}
  {{ if and .Step.Run.Shell (not eq (derefStr .Step.Run.Shell) "ENTRYPOINT") }}{{ derefStr .Step.Run.Shell }}{{ else }}/bin/sh{{ end }} -c {{ if and .Step.Run.Network (eq (derefStr .Step.Run.Network) "slirp4netns") }}{{ renderSlirp4NetNS .Step.Run .Job .Step.Run.Command | squote -}}{{ else if and .Step.Run.Network (eq (derefStr .Step.Run.Network) "pasta") }}{{ renderPastaNS .Step.Run .Job .Step.Run.Command | squote -}}{{ else }}
{{- if and .Step.Run.WorkDir (derefStr .Step.Run.WorkDir) -}}
  'mkdir -p {{ derefStr .Step.Run.WorkDir | squote | escapeSQuote }} && cd {{ derefStr .Step.Run.WorkDir | squote | escapeSQuote }} || { echo "change dir to working directory failed"; exit 1; };'{{ end -}}
  {{ .Step.Run.Command | squote -}}{{ end }}
{{- end -}}

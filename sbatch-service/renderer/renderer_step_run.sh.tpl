{{- /* In this file contains 3 launch modes: Apptainer, Enroot and none. */ -}}
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
export APPTAINER_BIND="$STORAGE_PATH:/deepsquare:rw,$DEEPSQUARE_SHARED_TMP:/deepsquare/tmp:rw,$DEEPSQUARE_SHARED_WORLD_TMP:/deepsquare/world-tmp:rw,$DEEPSQUARE_DISK_TMP:/deepsquare/disk/tmp:rw,$DEEPSQUARE_DISK_WORLD_TMP:/deepsquare/disk/world-tmp:rw{{ if and .Step.Run.Container.X11 (derefBool .Step.Run.Container.X11 ) }},/tmp/.X11-unix:/tmp/.X11-unix:ro{{ end }}"{{ range $mount := .Step.Run.Container.Mounts }},{{ $mount.HostDir | squote }}:{{ $mount.ContainerDir | squote }}:{{ $mount.Options | squote }}{{ end }}
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' \
DEEPSQUARE_TMP='/deepsquare/tmp' \
DEEPSQUARE_SHARED_TMP='/deepsquare/tmp' \
DEEPSQUARE_SHARED_WORLD_TMP='/deepsquare/world-tmp' \
DEEPSQUARE_DISK_TMP='/deepsquare/disk/tmp' \
DEEPSQUARE_DISK_WORLD_TMP='/deepsquare/disk/world-tmp' \
DEEPSQUARE_INPUT='/deepsquare/input' \
DEEPSQUARE_OUTPUT='/deepsquare/output' \
DEEPSQUARE_ENV="/deepsquare/$(basename $DEEPSQUARE_ENV)"{{ range $env := .Step.Run.Env }} {{ $env.Key }}={{ $env.Value | squote }}{{ end }} /usr/bin/srun {{ if and .Step.Name (derefStr .Step.Name) }}--job-name={{ derefStr .Step.Name | squote }}{{ end }} \
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
  --gpus-per-task=0 \
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
  {{ if and .Step.Run.Network (eq (derefStr .Step.Run.Network) "slirp4netns") }}/bin/sh -c {{ renderSlirp4NetNS .Step.Run (renderApptainerCommand .Step.Run) | squote -}}{{ else }}{{ renderApptainerCommand .Step.Run }}{{ end }}
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
  --gpus-per-task=0 \
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
{{- if isAbs $image -}}
/usr/bin/enroot create --name "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" -- "$STORAGE_PATH"{{ $image | squote }} >/dev/null 2>&1
{{- else -}}
/usr/bin/enroot create --name "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" -- "$IMAGE_PATH" >/dev/null 2>&1
{{- end }}
enrootClean() {
  /usr/bin/enroot remove -f "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID"
}
trap enrootClean EXIT INT TERM
'{{ if and .Step.Run.Network (eq (derefStr .Step.Run.Network) "slirp4netns") }}{{ renderSlirp4NetNS .Step.Run (renderEnrootCommand .Step.Run) | squote -}}{{ else }}{{ renderEnrootCommand .Step.Run | squote }}{{ end }}
{{- end }}
{{- else -}}
{{ range $env := .Step.Run.Env }}{{ $env.Key }}={{ $env.Value | squote }} {{ end }}/usr/bin/srun {{ if and .Step.Name (derefStr .Step.Name) }}--job-name={{ derefStr .Step.Name | squote }}{{ end }} \
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
  --gpus-per-task=0 \
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
  {{ if .Step.Run.Shell }}{{ derefStr .Step.Run.Shell }}{{ else }}/bin/sh{{ end }} -c {{ if and .Step.Run.Network (eq (derefStr .Step.Run.Network) "slirp4netns") }}{{ renderSlirp4NetNS .Step.Run .Step.Run.Command | squote -}}{{ else }}
{{- if and .Step.Run.WorkDir (derefStr .Step.Run.WorkDir) -}}
  'cd {{ derefStr .Step.Run.WorkDir | squote | escapeSQuote }} || { echo "change dir to working directory failed"; exit 1; };'{{ end -}}
  {{ .Step.Run.Command | squote -}}{{ end }}
{{- end -}}

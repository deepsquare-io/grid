/usr/bin/cat <<'EOFenroot' >"$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf"
#ENROOT_REMAP_ROOT=n
{{- if and .Run.Container.ReadOnlyRootFS (derefBool .Run.Container.ReadOnlyRootFS) }}
#ENROOT_ROOTFS_WRITABLE=n
{{- else }}
#ENROOT_ROOTFS_WRITABLE=y
{{- end }}

{{- if and .Run.Container.MountHome (derefBool .Run.Container.MountHome) }}
#ENROOT_MOUNT_HOME=y
{{- else }}
#ENROOT_MOUNT_HOME=n
{{- end }}

environ() {
  # Keep all the environment from the host
  env

  cat "${ENROOT_ROOTFS}/etc/environment"

  echo "STORAGE_PATH=/deepsquare"
  echo "DEEPSQUARE_TMP=/deepsquare/tmp"
  echo "DEEPSQUARE_SHARED_TMP=/deepsquare/tmp"
  echo "DEEPSQUARE_SHARED_WORLD_TMP=/deepsquare/world-tmp"
  echo "DEEPSQUARE_DISK_TMP=/deepsquare/disk/tmp"
  echo "DEEPSQUARE_DISK_WORLD_TMP=/deepsquare/disk/world-tmp"
  echo "DEEPSQUARE_INPUT=/deepsquare/input"
  echo "DEEPSQUARE_OUTPUT=/deepsquare/output"
  echo "DEEPSQUARE_ENV=/deepsquare/$(basename $DEEPSQUARE_ENV)"
}

mounts() {
  echo "$STORAGE_PATH /deepsquare none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_SHARED_TMP /deepsquare/tmp none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_SHARED_WORLD_TMP /deepsquare/world-tmp none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_DISK_TMP /deepsquare/disk/tmp none x-create=dir,bind,rw"
  echo "$DEEPSQUARE_DISK_WORLD_TMP /deepsquare/disk/world-tmp none x-create=dir,bind,rw"
{{- if and .Run.Container.X11 (derefBool .Run.Container.X11 ) }}
  echo "/tmp/.X11-unix /tmp/.X11-unix none x-create=dir,bind,ro"
{{- end }}
{{- range $mount := .Run.Container.Mounts }}
  echo '{{ $mount.HostDir }} {{ $mount.ContainerDir }} none x-create=auto,bind,{{ $mount.Options }}'
{{- end }}
}

hooks() {
  cat << 'EOFrclocal' > "${ENROOT_ROOTFS}/etc/rc.local"
{{- range $env := .Run.Env }}{{- $v := randomString 8 }}
{{ $env.Key }}="$(cat << 'EOF{{ $v }}'
{{ $env.Value }}
EOF{{ $v }}
)"
export {{ $env.Key }}
{{- end }}
{{- if and .Run.WorkDir (derefStr .Run.WorkDir) }}
mkdir -p {{ derefStr .Run.WorkDir | squote }} && cd {{ derefStr .Run.WorkDir | squote }} || { echo "change dir to working directory failed"; exit 1; }
{{- end }}
{{- if not (and .Run.Shell (eq (derefStr .Run.Shell) "ENTRYPOINT")) }}
exec "$@"
{{- end }}
EOFrclocal
}
EOFenroot
{{- if and (not (and .Run.Network (eq (derefStr .Run.Network) "slirp4netns"))) (not (and .Run.Network (eq (derefStr .Run.Network) "pasta"))) (or .Run.MapUID .Run.MapGID) }}
/usr/bin/unshare --map-current-user{{ if .Run.MapUID }} --map-user={{ .Run.MapUID }}{{ end }}{{ if .Run.MapGID }} --map-group={{ .Run.MapGID }}{{ end }} --mount \
{{- end }}
/usr/bin/enroot start \
  --conf "$STORAGE_PATH/enroot-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID.conf" \
  "container-$SLURM_JOB_ID.$SLURM_STEP_ID.$SLURM_PROCID" {{ if .Run.Command }}\
  {{ if and .Run.Shell (eq (derefStr .Run.Shell) "ENTRYPOINT") }}{{ .Run.Command | escapeCommand }}{{ else if .Run.Shell}}{{ derefStr .Run.Shell }} -c {{ .Run.Command | squote }}{{ else }}/bin/sh -c {{ .Run.Command | squote }}{{ end }}{{ end -}}

/usr/bin/cat <<'EOFenroot' >"$STORAGE_PATH/enroot.conf"
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
  /usr/bin/env

  /usr/bin/cat "${ENROOT_ROOTFS}/etc/environment"

  /usr/bin/echo "STORAGE_PATH=/deepsquare"
  /usr/bin/echo "DEEPSQUARE_TMP=/deepsquare/tmp"
  /usr/bin/echo "DEEPSQUARE_SHARED_TMP=/deepsquare/tmp"
  /usr/bin/echo "DEEPSQUARE_SHARED_WORLD_TMP=/deepsquare/world-tmp"
  /usr/bin/echo "DEEPSQUARE_DISK_TMP=/deepsquare/disk/tmp"
  /usr/bin/echo "DEEPSQUARE_DISK_WORLD_TMP=/deepsquare/disk/world-tmp"
  /usr/bin/echo "DEEPSQUARE_INPUT=/deepsquare/input"
  /usr/bin/echo "DEEPSQUARE_OUTPUT=/deepsquare/output"
  /usr/bin/echo "DEEPSQUARE_ENV=/deepsquare/$(basename $DEEPSQUARE_ENV)"
{{- range $env := .Run.Env }}
  /usr/bin/echo "{{ $env.Key }}={{ $env.Value | squote }}"
{{- end }}
}

mounts() {
  /usr/bin/echo "$STORAGE_PATH /deepsquare none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_SHARED_TMP /deepsquare/tmp none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_SHARED_WORLD_TMP /deepsquare/world-tmp none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_DISK_TMP /deepsquare/disk/tmp none x-create=dir,bind,rw"
  /usr/bin/echo "$DEEPSQUARE_DISK_WORLD_TMP /deepsquare/disk/world-tmp none x-create=dir,bind,rw"
{{- if and .Run.Container.X11 (derefBool .Run.Container.X11 ) }}
  /usr/bin/echo "/tmp/.X11-unix /tmp/.X11-unix none x-create=dir,bind,ro"
{{- end }}
{{- range $mount := .Run.Container.Mounts }}
  /usr/bin/echo '{{ $mount.HostDir }} {{ $mount.ContainerDir }} none x-create=auto,bind,{{ $mount.Options }}'
{{- end }}
}

hooks() {
  /usr/bin/cat << 'EOFrclocal' > "${ENROOT_ROOTFS}/etc/rc.local"
{{- if and .Run.WorkDir (derefStr .Run.WorkDir) }}
cd {{ derefStr .Run.WorkDir | squote }} || { echo "change dir to working directory failed"; exit 1; }
{{- else }}
cd "/deepsquare" || { echo "change dir to working directory failed"; exit 1; }
{{- end }}
exec "$@"
EOFrclocal
}
EOFenroot
{{- if and (not (and .Run.Network (eq (derefStr .Run.Network) "slirp4netns"))) (or .Run.MapUID .Run.MapGID) }}
/usr/bin/unshare --map-current-user{{ if .Run.MapUID }} --map-user={{ .Run.MapUID }}{{ end }}{{ if .Run.MapGID }} --map-group={{ .Run.MapGID }}{{ end }} --mount \
{{- end }}
/usr/bin/enroot start \
  --conf "$STORAGE_PATH/enroot.conf" \
  "container-$SLURM_JOB_ID.$SLURM_STEP_ID" \
  {{ if .Run.Shell }}{{ derefStr .Run.Shell }}{{ else }}/bin/sh{{ end }} -c {{ .Run.Command | squote -}}

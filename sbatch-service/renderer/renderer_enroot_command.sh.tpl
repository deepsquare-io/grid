/usr/bin/cat << 'EOFenroot' > "$STORAGE_PATH/enroot.conf"
{{- if and .Run.MapRoot (derefBool .Run.MapRoot ) }}
#ENROOT_REMAP_ROOT=y
{{- else }}
#ENROOT_REMAP_ROOT=n
{{- end }}
#ENROOT_ROOTFS_WRITABLE=y
#ENROOT_MOUNT_HOME=y

environ() {
  # Keep all the environment from the host
  env

  cat "${ENROOT_ROOTFS}/etc/environment"

  echo "STORAGE_PATH=/deepsquare"
  echo "DEEPSQUARE_INPUT=/deepsquare/input"
  echo "DEEPSQUARE_OUTPUT=/deepsquare/output"
  echo "DEEPSQUARE_ENV=/deepsquare/env"
{{- range $env := .Run.Env }}
  echo "{{ $env.Key }}={{ $env.Value | squote }}"
{{- end }}
}

mounts() {
  echo "$STORAGE_PATH /deepsquare none x-create=dir,bind,rw"
{{- if and .Run.Container.X11 (derefBool .Run.Container.X11 ) }}
  echo "/tmp/.X11-unix /tmp/.X11-unix none x-create=dir,bind,ro"
{{- end }}
{{- range $mount := .Run.Container.Mounts }}
  echo '{{ $mount.HostDir }} {{ $mount.ContainerDir }} bind,{{ $mount.Options }}'
{{- end }}
}

hooks() {
  echo 'exec "$@"' > "${ENROOT_ROOTFS}/etc/rc.local"
}
EOFenroot
/usr/bin/enroot start \
  --conf "$STORAGE_PATH/enroot.conf" \
  "container-$SLURM_JOB_ID" \
  {{ if .Run.Shell }}{{ derefStr .Run.Shell }}{{ else }}/bin/sh{{ end }} -c {{ .Run.Command | squote -}}

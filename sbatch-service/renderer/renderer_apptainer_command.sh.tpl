{{- $image := formatImageURL .Run.Container.Registry .Run.Container.Image .Run.Container.Apptainer .Run.Container.DeepsquareHosted -}}
export APPTAINER_BIND="$STORAGE_PATH:/deepsquare:rw,$DEEPSQUARE_SHARED_TMP:/deepsquare/tmp:rw,$DEEPSQUARE_SHARED_WORLD_TMP:/deepsquare/world-tmp:rw,$DEEPSQUARE_DISK_TMP:/deepsquare/disk/tmp:rw,$DEEPSQUARE_DISK_WORLD_TMP:/deepsquare/disk/world-tmp:rw{{ if and .Run.Container.X11 (derefBool .Run.Container.X11 ) }},/tmp/.X11-unix:/tmp/.X11-unix:ro{{ end }}"{{ range $mount := .Run.Container.Mounts }},{{ $mount.HostDir | squote }}:{{ $mount.ContainerDir | squote }}:{{ $mount.Options | squote }}{{ end }}
# shellcheck disable=SC2097,SC2098,SC1078
STORAGE_PATH='/deepsquare' \
DEEPSQUARE_TMP='/deepsquare/tmp' \
DEEPSQUARE_SHARED_TMP='/deepsquare/tmp' \
DEEPSQUARE_SHARED_WORLD_TMP='/deepsquare/world-tmp' \
DEEPSQUARE_DISK_TMP='/deepsquare/disk/tmp' \
DEEPSQUARE_DISK_WORLD_TMP='/deepsquare/disk/world-tmp' \
DEEPSQUARE_INPUT='/deepsquare/input' \
DEEPSQUARE_OUTPUT='/deepsquare/output' \
DEEPSQUARE_ENV="/deepsquare/$(basename $DEEPSQUARE_ENV)"{{ range $env := .Run.Env }} {{ $env.Key }}={{ $env.Value | squote }}{{ end }} \
{{- if and (not (and .Run.Network (eq (derefStr .Run.Network) "slirp4netns"))) (not (and .Run.Network (eq (derefStr .Run.Network) "pasta"))) (or .Run.MapUID .Run.MapGID) -}}
/usr/bin/unshare --map-current-user{{ if .Run.MapUID }} --map-user={{ .Run.MapUID }}{{ end }}{{ if .Run.MapGID }} --map-group={{ .Run.MapGID }}{{ end }} --mount \
{{- end }}
/usr/bin/apptainer --silent {{ if or (not .Run.Command) (and .Run.Shell (eq (derefStr .Run.Shell) "ENTRYPOINT")) }}run{{ else }}exec{{ end }} \
  --disable-cache \
  --contain \
{{- if not (and .Run.Container.ReadOnlyRootFS (derefBool .Run.Container.ReadOnlyRootFS)) }}
  --writable-tmpfs \
{{- end }}
{{- if not (and .Run.Container.MountHome (derefBool .Run.Container.MountHome)) }}
  --no-home \
{{- end }}
  --nv \
{{- if and .Run.WorkDir (derefStr .Run.WorkDir) }}
  --pwd {{ derefStr .Run.WorkDir | squote }} \
{{- else }}
{{- if and .Image .Image.Config.WorkingDir }}
  --pwd {{ .Image.Config.WorkingDir | squote }} \
{{- else }}
  --pwd "/" \
{{- end }}
{{- end }}
{{- if isAbs $image }}
  {{ if not (and .Run.Container.DeepsquareHosted (derefBool .Run.Container.DeepsquareHosted)) }}"$STORAGE_PATH"{{ end }}{{ $image | squote }} {{ if .Run.Command }}\{{ end }}
{{- else }}
  "$IMAGE_PATH" {{ if .Run.Command }}\{{ end }}
{{- end }}
  {{ if and .Run.Shell (eq (derefStr .Run.Shell) "ENTRYPOINT") }}{{ .Run.Command | escapeCommand }}{{ else if .Run.Shell}}{{ derefStr .Run.Shell }} -c {{ .Run.Command | squote }}{{ else }}/bin/sh -c {{ .Run.Command | squote }}{{ end -}}

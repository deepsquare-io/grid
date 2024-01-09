{{- $image := formatImageURL .Run.Container.Registry .Run.Container.Image .Run.Container.Apptainer .Run.Container.DeepsquareHosted -}}
{{- if and (not (and .Run.Network (eq (derefStr .Run.Network) "slirp4netns"))) (not (and .Run.Network (eq (derefStr .Run.Network) "pasta"))) (or .Run.MapUID .Run.MapGID) -}}
/usr/bin/unshare --map-current-user{{ if .Run.MapUID }} --map-user={{ .Run.MapUID }}{{ end }}{{ if .Run.MapGID }} --map-group={{ .Run.MapGID }}{{ end }} --mount \
{{- end }}/usr/bin/apptainer --silent {{ if .Run.Command }}exec{{ else }}run{{ end }} \
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
  --pwd "/deepsquare" \
{{- end }}
{{- if isAbs $image }}
  {{ if not (and .Run.Container.DeepsquareHosted (derefBool .Run.Container.DeepsquareHosted)) }}"$STORAGE_PATH"{{ end }}{{ $image | squote }} {{ if .Run.Command }}\{{ end }}
{{- else }}
  "$IMAGE_PATH" {{ if .Run.Command }}\{{ end }}
{{- end }}
  {{ if .Run.Command }}{{ if .Run.Shell }}{{ derefStr .Run.Shell }}{{ else }}/bin/sh{{ end }} -c {{ .Run.Command | squote }}{{ end -}}

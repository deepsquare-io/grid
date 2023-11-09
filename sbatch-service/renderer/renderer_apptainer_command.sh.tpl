{{- $image := formatImageURL .Run.Container.Registry .Run.Container.Image .Run.Container.Apptainer .Run.Container.DeepsquareHosted -}}
{{- if and (not (and .Run.Network (eq (derefStr .Run.Network) "slirp4netns"))) (or .Run.MapUID .Run.MapGID) -}}
/usr/bin/unshare --map-current-user{{ if .Run.MapUID }} --map-user={{ .Run.MapUID }}{{ end }}{{ if .Run.MapGID }} --map-group={{ .Run.MapGID }}{{ end }} --mount \
{{- end }}/usr/bin/apptainer --silent exec \
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
  {{ if not (and .Run.Container.DeepsquareHosted (derefBool .Run.Container.DeepsquareHosted)) }}"$STORAGE_PATH"{{ end }}{{ $image | squote }} \
{{- else }}
  "$IMAGE_PATH" \
{{- end }}
  {{ if .Run.Shell }}{{ derefStr .Run.Shell }}{{ else }}/bin/sh{{ end }} -c {{ .Run.Command | squote -}}

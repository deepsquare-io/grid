{{- $image := formatImageURL .Run.Container.Registry .Run.Container.Image .Run.Container.Apptainer .Run.Container.DeepsquareHosted -}}
{{- if and (not (and .Run.Network (eq (derefStr .Run.Network) "slirp4netns"))) (not (and .Run.Network (eq (derefStr .Run.Network) "pasta"))) (or .Run.MapUID .Run.MapGID) -}}
/usr/bin/unshare --map-current-user{{ if .Run.MapUID }} --map-user={{ .Run.MapUID }}{{ end }}{{ if .Run.MapGID }} --map-group={{ .Run.MapGID }}{{ end }} --mount \
{{- end }}/usr/bin/apptainer --silent {{ if or (not .Run.Command) (and .Run.Shell (eq (derefStr .Run.Shell) "ENTRYPOINT")) }}run{{ else }}exec{{ end }} \
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
{{- end }}
{{- if isAbs $image }}
  {{ if not (and .Run.Container.DeepsquareHosted (derefBool .Run.Container.DeepsquareHosted)) }}"$STORAGE_PATH"{{ end }}{{ $image | squote }} {{ if .Run.Command }}\{{ end }}
{{- else }}
  "$IMAGE_PATH" {{ if .Run.Command }}\{{ end }}
{{- end }}
  {{ if and .Run.Shell (eq (derefStr .Run.Shell) "ENTRYPOINT") }}{{ .Run.Command | escapeCommand }}{{ else if .Run.Shell}}{{ derefStr .Run.Shell }} -c {{ .Run.Command | squote }}{{ else }}/bin/sh -c {{ .Run.Command | squote }}{{ end -}}

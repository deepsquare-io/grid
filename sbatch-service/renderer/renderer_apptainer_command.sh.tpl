{{- $image := formatImageURL .Run.Container.Registry .Run.Container.Image .Run.Container.Apptainer .Run.Container.DeepsquareHosted -}}
/usr/bin/apptainer --silent exec \
  --disable-cache \
  --writable-tmpfs \
{{- if not (and .Run.Container.MountHome (derefBool .Run.Container.MountHome)) }}
  --no-home \
{{- end }}
  --nv \
{{- if and .Run.WorkDir (derefStr .Run.WorkDir) }}
  --pwd {{ derefStr .Run.WorkDir | squote }} \
{{- else }}
  --pwd "/deepsquare" \
{{- end }}
{{- if and .Run.MapRoot (derefBool .Run.MapRoot ) }}
  --fakeroot \
{{- end }}
{{- if isAbs $image }}
  {{ if not (and .Run.Container.DeepsquareHosted (derefBool .Run.Container.DeepsquareHosted)) }}"$STORAGE_PATH"{{ end }}{{ $image | squote }} \
{{- else }}
  "$IMAGE_PATH" \
{{- end }}
  {{ if .Run.Shell }}{{ derefStr .Run.Shell }}{{ else }}/bin/sh{{ end }} -c {{ .Run.Command | squote -}}

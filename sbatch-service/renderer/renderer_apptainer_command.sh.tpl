/usr/bin/apptainer --silent exec \
  --disable-cache \
  --nv \
{{- if and .Run.WorkDir (derefStr .Run.WorkDir) }}
  --pwd {{ derefStr .Run.WorkDir | squote }} \
{{- end }}
{{- if and .Run.MapRoot (derefBool .Run.MapRoot ) }}
  --fakeroot \
{{- end }}
  {{ if and (hasPrefix "/" .Run.Container.Image) (not (and .Run.Container.DeepsquareHosted (derefBool .Run.Container.DeepsquareHosted))) }}"$STORAGE_PATH"{{ end }}{{ formatImageURL .Run.Container.Registry .Run.Container.Image .Run.Container.Apptainer .Run.Container.DeepsquareHosted | squote }} \
  {{ if .Run.Shell }}{{ derefStr .Run.Shell }}{{ else }}/bin/sh{{ end }} -c {{ .Run.Command | squote -}}

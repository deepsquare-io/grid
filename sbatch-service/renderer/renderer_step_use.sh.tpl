(
{{- range $input := .Module.Inputs }}
{{- if $input.Default }}
export {{ $input.Key }}={{ derefStr $input.Default | squote }}
{{- end }}
{{- range $env := $.Use.Args }}
{{- if eq $input.Key $env.Key }}
export {{ $env.Key }}={{ $env.Value | squote }}
{{- end }}
{{- end }}
{{- end }}

{{- /* Temporary store env path so we can store the module variables */ -}}
{{- if and .Use.ExportEnvAs (derefStr .Use.ExportEnvAs) }}
DEEPSQUARE_{{ .UUID }}_OLD_ENV="$DEEPSQUARE_ENV"
export DEEPSQUARE_ENV="$STORAGE_PATH/DEEPSQUARE_{{ .UUID }}_env"
/usr/bin/touch $DEEPSQUARE_ENV
{{- end }}

{{- range $step := .Module.Steps }}
{{ renderStep $.Job $step }}
{{- end }}

{{- /* Extract the output variables */ -}}
{{- if and .Use.ExportEnvAs (derefStr .Use.ExportEnvAs) }}
{{- range $output := .Module.Outputs }}
echo "{{ derefStr $.Use.ExportEnvAs }}_$(/usr/bin/grep "^{{- $output.Key -}}" "$DEEPSQUARE_ENV")" >> "${DEEPSQUARE_{{ $.UUID }}_OLD_ENV}"
{{- end }}
export DEEPSQUARE_ENV="${DEEPSQUARE_{{ $.UUID }}_OLD_ENV}"
{{- end }}
)

{{- /* Extract the output variables */ -}}
{{- if and .Use.ExportEnvAs (derefStr .Use.ExportEnvAs) }}
{{- range $output := .Module.Outputs }}
export "{{ derefStr $.Use.ExportEnvAs }}_$(/usr/bin/grep "^{{- $output.Key -}}" "$DEEPSQUARE_ENV")"
{{- end }}
{{- end }}

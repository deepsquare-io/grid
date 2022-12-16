/usr/bin/echo 'Running: '{{ .Step.Name | squote }}
{{- if .Step.Run }}
{{ renderStepRun .Job .Step }}
{{- else if .Step.For }}
{{ renderStepFor .Job .Step.For }}
{{- end -}}

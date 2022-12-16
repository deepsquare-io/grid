doFor() {
{{- if .For.Range }}
export index="$1"
{{- else if .For.Items }}
export item="$1"
{{- end }}

{{- range $step := .For.Steps }}
{{ renderStep $.Job $step }}
{{- end }}
}

{{- if .For.Parallel }}
pids=()
{{- end }}
{{- if .For.Range }}
for index in $(seq {{ .For.Range.Begin }} {{ if and .For.Range.Increment (ne (derefInt .For.Range.Increment) 0) }}{{ derefInt .For.Range.Increment }}{{ else }}1{{ end }} {{ .For.Range.End }}); do
  doFor "$index" {{ if .For.Parallel }}&{{ end }}
  {{- if .For.Parallel }}
  pids+=("$!")
  {{- end }}
done
{{- else if .For.Items }}
items=({{ range $item := .For.Items }}{{ $item | squote }} {{ end }})
for item in "${items[@]}"; do
  doFor "$item" {{ if .For.Parallel }}&{{ end }}
  {{- if .For.Parallel }}
  pids+=("$!")
  {{- end }}
done
{{- end }}
{{- if .For.Parallel }}
for pid in "${pids[@]}"; do
  wait "$pid"
done
{{- end -}}

{{- range $depend := .Step.DependsOn -}}
if [ -n "${PID_{{ $depend | upper }}+x}" ]; then
/usr/bin/echo 'Waiting for: {{ $depend | upper }}.'
wait "${PID_{{ $depend | upper }}}"
else
/usr/bin/echo 'Cannot await: {{ $depend | upper }} is undefined. Exiting to avoid undefined behavior.'
/usr/bin/echo 'Is the {{ $depend | upper }} handleName set and is defined at the same scope ?'
exit 1
fi
{{ end -}}
{{- if and .Step.If (derefStr .Step.If) }}
# shellcheck disable=SC2016,SC2089
CONDITION={{ derefStr .Step.If | squote }}
# shellcheck disable=SC2090
export CONDITION
CONDITION_RESULT="$(eval "if [[ $CONDITION ]]; then echo 'true' ; else echo 'false' ; fi")"
if [ $CONDITION_RESULT = "true" ]; then
{{ end -}}
{{ if or .Step.Catch .Step.Finally }}
( # CATCH FINALLY
{{- end }}
{{- if .Step.Catch }}
set +e
{{ end -}}
{{- if .Step.Finally }}
finally() {
set -e
{{- range $step := .Step.Finally }}
{{ renderStep $.Job $step }}
{{- end }}
}
trap finally EXIT INT TERM
{{ end -}}
{{- if and .Step.Name (derefStr .Step.Name) -}}
/usr/bin/echo 'Running: '{{ derefStr .Step.Name | squote }}
{{- end -}}
{{- if .Step.Catch }}
( # CATCH
set -e
{{- end -}}
{{ if .Step.Steps }}
{{- range $step := .Step.Steps }}
{{ renderStep $.Job $step }}
{{- end }}
{{- else if .Step.Run }}
{{ renderStepRun .Job .Step }}
{{- else if .Step.For }}
{{ renderStepFor .Job .Step.For }}
{{- else if .Step.Launch }}
{{ renderStepAsyncLaunch .Job .Step.Launch }}
{{- else if .Step.Use }}
{{ renderStepUse .Job .Step .Step.Use }}
{{- end -}}
{{- if .Step.Catch }}
) # CATCH
DEEPSQUARE_ERROR_CODE=$?
export DEEPSQUARE_ERROR_CODE
if [ $DEEPSQUARE_ERROR_CODE -ne 0 ]; then
set -e
{{- range $step := .Step.Catch }}
{{ renderStep $.Job $step }}
{{- end }}
fi
{{- end -}}
{{- if or .Step.Catch .Step.Finally }}
) # CATCH FINALLY
{{- end -}}
{{ if and .Step.If (derefStr .Step.If) }}
:
fi
{{ end -}}

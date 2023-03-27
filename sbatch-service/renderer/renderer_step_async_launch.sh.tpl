(
declare -A EXIT_SIGNALS

{{- range $step := .Launch.Steps }}
{{ renderStep $.Job $step }}
{{- end }}

for pid in "${!EXIT_SIGNALS[@]}"; do
  /usr/bin/echo "Process $$ sending signal ${EXIT_SIGNALS[$pid]} to $pid"
  kill -s "${EXIT_SIGNALS[$pid]}" "$pid" || echo "Sending signal ${EXIT_SIGNALS[$pid]} to $pid failed, continuing..."
done
) &
asynclaunchpid="$!"
{{- if and .Launch.HandleName (derefStr .Launch.HandleName) }}
export PID_{{ .Launch.HandleName | upper }}="$asynclaunchpid"
{{- end -}}
{{- if and .Launch.SignalOnParentStepExit (derefInt .Launch.SignalOnParentStepExit) }}
EXIT_SIGNALS[$asynclaunchpid]={{ derefInt .Launch.SignalOnParentStepExit }}
{{- else if not .Launch.SignalOnParentStepExit }}
EXIT_SIGNALS[$asynclaunchpid]=15
{{- end -}}

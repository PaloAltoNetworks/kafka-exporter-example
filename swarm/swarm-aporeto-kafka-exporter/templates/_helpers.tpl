##
## Placement contrainsts and pref helpers
##
{{- define "general.placement" }}
{{- if .Values.placement }}
placement:
  {{- if .Values.placement.constraints }}
  constraints:
  {{- range .Values.placement.constraints}}
    - {{.}}
  {{- end}}
  {{- end}}
  {{- if .Values.placement.preferences }}
  preferences:
  {{- range .Values.placement.preferences}}
    - {{. | toYaml }}
  {{- end}}
  {{- end}}
{{- end }}
{{- end }}


{{- define "general.replicas" }}
{{- $r := (max .Values.replicas 1) }}
{{- if .Values.global.scaling.enabled }}
{{- $f := (required "global.scaling.factor is required" .Values.global.scaling.factor) }}
replicas: {{ max (mul $r $f) 1 }}
{{- else }}
replicas: {{ $r }}
{{- end }}
{{- end }}

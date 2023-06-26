package global

{{- if .HasGlobal }}

import "appnest/server/plugin/{{ .Snake}}/config"

var GlobalConfig = new(config.{{ .PlugName}})
{{ end -}}
// Code generated by ZERP Codegen Util. DO NOT EDIT.
// Last updated at {{.Timestamp}}

package {{ .Package }}

import (
	"mverp/data/model/{{ .Package }}"
)
{{- $package := .Package }}
{{- $name := .Name }}
{{- range .Objects }}
func (r *{{ $name }}Repository) GetAll{{ .Name }}Enums() []{{ $package }}.{{ .Name }} {
	return {{ $package }}.All{{ .Name }}()
}
{{- end }}
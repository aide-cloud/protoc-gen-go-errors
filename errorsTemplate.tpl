{{ range .Errors }}

{{ if .HasComment }}// Is{{.CamelValue}} {{ .Comment }}{{ end -}}
func Is{{.CamelValue}}(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == {{ .Name }}_{{ .Value }}.String() && e.Code == {{ .HTTPCode }}
}

{{ if .HasComment }}// Error{{ .CamelValue }} {{ .Comment }}{{ end -}}
func Error{{ .CamelValue }}(format string, args ...interface{}) *errors.Error {
    return errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), fmt.Sprintf(format, args...))
}

{{ if .HasComment }}// Error{{ .CamelValue }}WithContext {{ .Comment }}//  带上下文，支持国际化输出元数据
{{ end -}}
func Error{{ .CamelValue }}WithContext({{ if .HasMetadata }}ctx{{else}}_{{ end }} context.Context, format string, args ...interface{}) *errors.Error {
    return errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), fmt.Sprintf(format, args...)){{ if .HasMetadata }}.WithMetadata(map[string]string{{ .Metadata }}){{ end }}
}

{{ if .HasI18n }}
    const ErrorI18n{{ .CamelValue }}ID = "{{ .ID }}"

    {{ if .HasComment }}// ErrorI18n{{ .CamelValue }} {{ .Comment }}//  支持国际化输出
    {{ end -}}
    func ErrorI18n{{ .CamelValue }}(ctx context.Context, args ...interface{}) *errors.Error {
        config := &i18n.LocalizeConfig{
            MessageID: ErrorI18n{{ .CamelValue }}ID,
        }
        if len(args) > 0 {
            config.TemplateData = args[0]
        }
        err := errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), fmt.Sprintf("{{ .Message }}", args...))
        local, ok := FromContext(ctx)
        if ok {
            localize, err1 := local.Localize(config)
            if err1 != nil {
                err = errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), fmt.Sprintf("{{ .Message }}", args...)).WithCause(err1)
            } else {
                err = errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), localize)
            }
        }

        return err{{ if .HasMetadata }}.WithMetadata(map[string]string{{ .Metadata }}){{ end }}
    }

{{ end }}
{{- end }}

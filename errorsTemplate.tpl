{{ range .Errors }}

{{ if .HasComment }}{{ .Comment }}{{ end -}}
func Is{{.CamelValue}}(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == {{ .Name }}_{{ .Value }}.String() && e.Code == {{ .HTTPCode }}
}

{{ if .HasComment }}{{ .Comment }}{{ end -}}
func Error{{ .CamelValue }}(format string, args ...interface{}) *errors.Error {
	 return errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), fmt.Sprintf(format, args...))
}

{{ if .HasI18n }}
    const ErrorI18n{{ .CamelValue }}ID = "{{ .ID }}"

    {{ if .HasComment }}{{ .Comment }}{{ end -}}
    func ErrorI18n{{ .CamelValue }}(ctx context.Context, args ...interface{}) *errors.Error {
        config := &i18n.LocalizeConfig{
            MessageID: ErrorI18n{{ .CamelValue }}ID,
        }
        if len(args) > 0 {
            config.TemplateData = args[0]
        }

        localize, err := FromContext(ctx).Localize(config)
        if err != nil {
            return errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), fmt.Sprintf("{{ .Message }}", args...))
        }
        return errors.New({{ .HTTPCode }}, {{ .Name }}_{{ .Value }}.String(), localize)
    }

{{ end }}
{{- end }}

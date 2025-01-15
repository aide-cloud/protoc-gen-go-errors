{{/* gotype: github.com/aide-cloud/protoc-gen-go-errors.errorWrapper*/}}
{{ range .Errors }}
const Error{{ .CamelValue }}ID = "{{ .ID }}"

{{ if .HasComment }}// Is{{.CamelValue}} {{ .Comment }}{{ end -}}
func Is{{.CamelValue}}(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == Error{{ .CamelValue }}ID && e.Code == {{ .HTTPCode }}
}

{{ if .HasComment }}// Error{{ .CamelValue }} {{ .Comment }}{{ end -}}
func Error{{ .CamelValue }}(format string, args ...interface{}) *errors.Error {
    return errors.New({{ .HTTPCode }}, Error{{ .CamelValue }}ID, fmt.Sprintf(format, args...))
}

{{ if .HasComment }}// Error{{ .CamelValue }}WithContext {{ .Comment }}//  带上下文，支持国际化输出元数据
{{ end -}}
func Error{{ .CamelValue }}WithContext({{ if .HasMetadata }}ctx{{else}}_{{ end }} context.Context, format string, args ...interface{}) *errors.Error {
    return errors.New({{ .HTTPCode }}, Error{{ .CamelValue }}ID, fmt.Sprintf(format, args...)){{ if .HasMetadata }}.WithMetadata(map[string]string{{ .Metadata }}){{ end }}
}

{{ if .HasI18n }}
    var _{{ .CamelValue }}Msg = &i18n.Message{
        ID:    Error{{ .CamelValue }}ID,
        One:   "{{ .Message }}",
        Other: "{{ .Message }}",
    }
    {{ if .HasComment }}// ErrorI18n{{ .CamelValue }} {{ .Comment }}//  支持国际化输出
    {{ end -}}
    func ErrorI18n{{ .CamelValue }}(ctx context.Context, args ...interface{}) *errors.Error {
        msg := "{{ .Message }}"
        defaultMessage := _{{ .CamelValue }}Msg
        if len(args) > 0 {
            msg = fmt.Sprintf(msg, args...)
            defaultMessage.One = fmt.Sprintf(defaultMessage.One, args...)
            defaultMessage.Other = fmt.Sprintf(defaultMessage.Other, args...)
        }
        err := errors.New({{ .HTTPCode }}, Error{{ .CamelValue }}ID, msg)
        local, ok := FromContext(ctx)
        if ok {
            config := &i18n.LocalizeConfig{
                MessageID: Error{{ .CamelValue }}ID,
                DefaultMessage: defaultMessage,
            }
            localize, err1 := local.Localize(config)
            if err1 != nil {
                err = errors.New({{ .HTTPCode }}, Error{{ .CamelValue }}ID, msg).WithCause(err1)
            } else {
                err = errors.New({{ .HTTPCode }}, Error{{ .CamelValue }}ID, localize)
            }
        }

        return err{{ if .HasMetadata }}.WithMetadata(map[string]string{{ .Metadata }}){{ end }}
    }

{{ end }}
{{- end }}

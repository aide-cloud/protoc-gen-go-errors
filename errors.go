package main

import (
	"fmt"
	"strings"
	"unicode"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"

	"github.com/aide-cloud/protoc-gen-go-errors/errors"
)

const (
	errorsPackage  = protogen.GoImportPath("github.com/go-kratos/kratos/v2/errors")
	fmtPackage     = protogen.GoImportPath("fmt")
	contextPackage = protogen.GoImportPath("context")
	i18nPackage    = protogen.GoImportPath("github.com/nicksnyder/go-i18n/v2/i18n")
)

var enCases = cases.Title(language.AmericanEnglish, cases.NoLower)

// generateFile generates a _errors.pb.go file containing kratos errors definitions.
func generateFile(gen *protogen.Plugin, file *protogen.File) *protogen.GeneratedFile {
	if len(file.Enums) == 0 {
		return nil
	}
	filename := file.GeneratedFilenamePrefix + "_errors.pb.go"
	g := gen.NewGeneratedFile(filename, file.GoImportPath)
	g.P("// Code generated by protoc-gen-go-errors. DO NOT EDIT.")
	g.P()
	g.P("package ", file.GoPackageName)
	g.P()
	g.QualifiedGoIdent(contextPackage.Ident(""))
	g.QualifiedGoIdent(fmtPackage.Ident(""))
	g.QualifiedGoIdent(i18nPackage.Ident(""))
	generateFileContent(gen, file, g)
	return g
}

// generateFileContent generates the kratos errors definitions, excluding the package statement.
func generateFileContent(gen *protogen.Plugin, file *protogen.File, g *protogen.GeneratedFile) {
	if len(file.Enums) == 0 {
		return
	}

	g.P("// This is a compile-time assertion to ensure that this generated file")
	g.P("// is compatible with the kratos package it is being compiled against.")
	g.P("const _ = ", errorsPackage.Ident("SupportPackageIsVersion1"))
	g.P("type localizeKey struct{}")
	g.P("func FromContext(ctx context.Context) (*i18n.Localizer, bool) {\n\tlocal, ok:= ctx.Value(localizeKey{}).(*i18n.Localizer)\n\treturn local, ok\n}")
	g.P()
	g.P("func WithLocalize(ctx context.Context, localize *i18n.Localizer) context.Context {\n\treturn context.WithValue(ctx, localizeKey{}, localize)\n}")
	g.P()
	g.P("// GetI18nMessage 获取错误信息\nfunc GetI18nMessage(ctx context.Context, id string, args ...interface{}) string {\n\tif id == \"\" {\n\t\treturn id\n\t}\n\tconfig := &i18n.LocalizeConfig{\n\t\tMessageID: id,\n\t}\n\tif len(args) > 0 {\n\t\tconfig.TemplateData = args[0]\n\t}\n\tlocal, ok := FromContext(ctx)\n\tif !ok {\n\t\treturn id\n\t}\n\tlocalize, err := local.Localize(config)\n\tif err != nil {\n\t\treturn id\n\t}\n\treturn localize\n}")
	g.P()
	index := 0
	for _, enum := range file.Enums {
		if !genErrorsReason(gen, file, g, enum) {
			index++
		}
	}
	// If all enums do not contain 'errors.code', the current file is skipped
	if index == 0 {
		g.Skip()
	}
}

func genErrorsReason(_ *protogen.Plugin, _ *protogen.File, g *protogen.GeneratedFile, enum *protogen.Enum) bool {
	defaultCode := proto.GetExtension(enum.Desc.Options(), errors.E_DefaultCode)
	code := 0
	if ok := defaultCode.(int32); ok != 0 {
		code = int(ok)
	}
	if code > 600 || code < 0 {
		panic(fmt.Sprintf("Enum '%s' range must be greater than 0 and less than or equal to 600", string(enum.Desc.Name())))
	}
	var ew errorWrapper
	for _, v := range enum.Values {
		enumCode := code

		eCode := proto.GetExtension(v.Desc.Options(), errors.E_Code)
		if ok := eCode.(int32); ok != 0 {
			enumCode = int(ok)
		}
		// If the current enumeration does not contain 'errors.code'
		// or the code value exceeds the range, the current enum will be skipped
		if enumCode > 600 || enumCode < 0 {
			panic(fmt.Sprintf("Enum '%s' range must be greater than 0 and less than or equal to 600", string(v.Desc.Name())))
		}
		if enumCode == 0 {
			continue
		}

		comment := v.Comments.Leading.String()
		if comment == "" {
			comment = v.Comments.Trailing.String()
		}
		comment = strings.TrimPrefix(comment, "// ")
		id := proto.GetExtension(v.Desc.Options(), errors.E_Id).(string)
		message := proto.GetExtension(v.Desc.Options(), errors.E_Message).(string)
		metadata := proto.GetExtension(v.Desc.Options(), errors.E_Metadata).([]*errors.Metadata)
		bizReasons := proto.GetExtension(v.Desc.Options(), errors.E_BizReason).([]*errors.BizReason)

		metadataMap := make(map[string]string)
		for _, m := range metadata {
			metadataMap[m.Key] = fmt.Sprintf(`GetI18nMessage(ctx, "%s")`, m.Value)
		}
		metadataMapBsStringBuilder := strings.Builder{}
		metadataMapBsStringBuilder.WriteString("{\n")
		for k, val := range metadataMap {
			metadataMapBsStringBuilder.WriteString(fmt.Sprintf(`"%s": %s,`, k, val))
			metadataMapBsStringBuilder.WriteString("\n")
		}
		metadataMapBsStringBuilder.WriteString("\n}")

		err := errorInfo{
			Name:        string(enum.Desc.Name()),
			Value:       string(v.Desc.Name()),
			HTTPCode:    enumCode,
			CamelValue:  case2Camel(string(v.Desc.Name())),
			Comment:     comment,
			HasComment:  len(comment) > 0,
			ID:          id,
			Message:     message,
			HasI18n:     id != "" || message != "",
			Metadata:    metadataMapBsStringBuilder.String(),
			HasMetadata: metadataMap != nil && len(metadataMap) > 0,
		}
		if err.ID == "" {
			err.ID = err.Name + "_" + err.Value
		}
		ew.Errors = append(ew.Errors, &err)
		if len(bizReasons) > 0 {
			for _, br := range bizReasons {
				bizErr := err
				if br.GetReason() != "" {
					bizErr.CamelValue += case2Camel(br.GetReason())
					bizErr.ID += "__" + br.GetReason()
					bizErr.Comment += "//  " + br.GetReason() + "\n"
				}

				if br.GetMessage() != "" {
					bizErr.Message = br.GetMessage()
					bizErr.Comment += "//  " + br.GetMessage() + "\n"
				}

				if len(br.GetMetadata()) > 0 {
					bizMetadataMapBsStringBuilder := strings.Builder{}
					bizMetadataMapBsStringBuilder.WriteString("{\n")
					for _, val := range br.GetMetadata() {
						bizMetadataMapBsStringBuilder.WriteString(fmt.Sprintf(`"%s": %s,`, val.GetKey(), fmt.Sprintf(`GetI18nMessage(ctx, "%s")`, val.GetValue())))
						bizMetadataMapBsStringBuilder.WriteString("\n")
					}
					bizMetadataMapBsStringBuilder.WriteString("\n}")
					bizErr.Metadata = bizMetadataMapBsStringBuilder.String()
				}
				ew.Errors = append(ew.Errors, &bizErr)
			}
		}
	}
	if len(ew.Errors) == 0 {
		return true
	}
	g.P(ew.execute())

	return false
}

func case2Camel(name string) string {
	if !strings.Contains(name, "_") {
		if name == strings.ToUpper(name) {
			name = strings.ToLower(name)
		}
		return enCases.String(name)
	}
	strs := strings.Split(name, "_")
	words := make([]string, 0, len(strs))
	for _, w := range strs {
		hasLower := false
		for _, r := range w {
			if unicode.IsLower(r) {
				hasLower = true
				break
			}
		}
		if !hasLower {
			w = strings.ToLower(w)
		}
		w = enCases.String(w)
		words = append(words, w)
	}

	return strings.Join(words, "")
}

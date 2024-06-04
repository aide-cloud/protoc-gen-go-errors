package main

import (
	"context"
	"embed"
	"fmt"

	"github.com/BurntSushi/toml"
	"github.com/aide-cloud/protoc-gen-go-errors/example/api"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
)

//go:embed i18n/active.*.toml
var LocaleFS embed.FS

func main() {
	// 中文
	bundle := i18n.NewBundle(language.Chinese)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, err := bundle.LoadMessageFileFS(LocaleFS, "i18n/active.zh.toml")
	if err != nil {
		panic(err)
	}
	ctx := api.WithLocalize(context.Background(), i18n.NewLocalizer(bundle, "zh-CN"))

	fmt.Println(api.ErrorI18nSystemError(ctx))
	fmt.Println(api.ErrorI18nUserAlreadyExists(ctx))
	fmt.Println(api.ErrorI18nUserNotFound(ctx))

	// 英文
	bundle = i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	_, err = bundle.LoadMessageFileFS(LocaleFS, "i18n/active.en.toml")
	if err != nil {
		panic(err)
	}
	ctx = api.WithLocalize(context.Background(), i18n.NewLocalizer(bundle, "en-US"))
	fmt.Println(api.ErrorI18nSystemError(ctx))
	fmt.Println(api.ErrorI18nUserAlreadyExists(ctx))
	fmt.Println(api.ErrorI18nUserNotFound(ctx))
}

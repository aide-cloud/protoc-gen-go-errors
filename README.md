# kratos errror生成器

## 功能点

1. 支持默认message
2. 支持国际化

## 安装

```sh
go install github.com/aide-cloud/protoc-gen-go-errors@v1.0.1
# 或者
go install github.com/aide-cloud/protoc-gen-go-errors@latest
```

## 使用

```makefile
GOHOSTOS:=$(shell go env GOHOSTOS)
ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	#Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	Git_Bash=$(subst \,/,$(subst cmd\,bin\bash.exe,$(dir $(shell where git))))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find cmd -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find cmd -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: error
# generate api proto
error:
	protoc --proto_path=./api \
	       --proto_path=../third_party \
 	       --go_out=paths=source_relative:./api \
 	       --go-errors_out=paths=source_relative:./api \
	       $(API_PROTO_FILES)
```

## 示例

```proto
syntax = "proto3";

package example.api;

import "errors/errors.proto";

option go_package = "github.com/aide-cloud/protoc-gen-go-errors/example/api;api";
option java_multiple_files = true;
option java_package = "example.api";

enum ErrorReason {
	option (errors.default_code) = 500;

	SYSTEM_ERROR = 0 [
		(errors.code) = 500,
		(errors.id) = "SYSTEM_ERROR",
		(errors.message) = "系统错误"
	];

	USER_NOT_FOUND = 1 [
		(errors.code) = 404,
		(errors.id) = "USER_NOT_FOUND",
		(errors.message) = "用户不存在"
	];

	USER_ALREADY_EXISTS = 2 [
		(errors.code) = 400,
		(errors.id) = "USER_ALREADY_EXISTS",
		(errors.message) = "用户已存在"
	];
}
```

## 测试代码

```go
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

```

## 运行结果

```sh
error: code = 500 reason = SYSTEM_ERROR message = 系统错误 metadata = map[] cause = <nil>
error: code = 400 reason = USER_ALREADY_EXISTS message = 用户已存在 metadata = map[] cause = <nil>
error: code = 404 reason = USER_NOT_FOUND message = 用户不存在 metadata = map[] cause = <nil>
error: code = 500 reason = SYSTEM_ERROR message = System error metadata = map[] cause = <nil>
error: code = 400 reason = USER_ALREADY_EXISTS message = User already exists metadata = map[] cause = <nil>
error: code = 404 reason = USER_NOT_FOUND message = User not found metadata = map[] cause = <nil>

```

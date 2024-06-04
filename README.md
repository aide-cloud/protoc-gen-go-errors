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

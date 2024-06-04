# kratos errror生成器

## 功能点

1. 支持默认message
2. 支持国际化

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

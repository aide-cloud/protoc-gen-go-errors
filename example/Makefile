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
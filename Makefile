# usage: make help

# 多行注释
# 	- 每行开头第一个字符为注释符
# 	- 首行开头第一个字符为注释符，后面每行结尾加上反斜杠（\），则下一行也被注释
# 在Shell编程中，notdir是一个函数，用于获取一个文件路径中的文件名部分，即去除路径前缀后的部分。notdir函数的语法如下：\
$(notdir names...) \
其中，names是一个或多个文件路径，可以使用通配符。\
notdir函数会返回names中每个文件路径的文件名部分，多个文件名之间用空格分隔。如果中的某个元素不包含路径，则返回该元素本身。\
例如，假设当前目录下有两个文件：/home/user/foo.txt和/home/user/bar.txt，我们可以使用notdir函数获取它们的文件名：\
$ echo $(notdir /home/user/foo.txt /home/user/bar.txt) \
foo.txt bar.txt \
notdir函数可以方便地将文件路径转换为文件名，常用于Makefile等Shell脚本中。 \

projectName := $(notdir ${PWD})


# --------------------
# 构建与运行
# --------------------

.PHONY: build
build: ## go build
	$(call logInfo, building $(projectName))
	@go build -o $(projectName)
	$(call logInfo, build done)


.PHONY: run
run: ## go run
	$(call logInfo, starting $(projectName))
	./$(projectName) -conf conf/conf.yaml



# --------------------
# 静态代码分析
# --------------------

.PHONY: sa
sa: vet errcheck revive lint build  ## [static analysis]: vet, errcheck, revive, lint

.PHONY: vet
vet: ## [static analysis]: go vet ./... 
	$(call logInfo, go vet ./...)
	@go vet ./... || ($(call logErrorP, go vet: not pass) && exit 1)
	$(call logInfo, go vet: pass)


.PHONY: errcheck
errcheck: ## [static analysis]: errcheck
	$(call logInfo, errcheck ./...)
	@hash errcheck 2>&- || go get -u -v github.com/kisielk/errcheck
	@errcheck ./... || ($(call logErrorP, errcheck: not pass) && exit 1)
	$(call logInfo, errcheck: pass)


.PHONY: revive
revive: ## [static analysis]: revive
	$(call logInfo, revive)
	@hash revive 2>&- || go get -u -v github.com/mgechev/revive
	revive -formatter stylish ./... || ($(call logErrorP, revive: not pass) && exit 1)
	$(call logInfo, revive: pass)


.PHONY: lint
lint: ## [static analysis]: golangci-lint
	$(call logInfo, golangci-lint)
	@hash golangci-lint 2>&- || go get -u -v github.com/golangci/golangci-lint/cmd/golangci-lint
	@golangci-lint run || ($(call logErrorP, golangci-lint: not pass) && exit 1)
	$(call logInfo, golangci-lint: pass)



# --------------------
# 代码格式化相关
# --------------------

.PHONY: complete
complete: fmt cmt  ## [formatting]: 运行代码格式化: fmt, cmt

.PHONY: fmt
fmt: ## [formatting]: gofumpt 代码格式化
	$(call logInfo, gofumpt)
	@hash gofumpt 2>&- || go install mvdan.cc/gofumpt@latest
	gofumpt -l -w .
	$(call logInfo, gofumpt: done)

.PHONY: cmt
cmt: ## [formatting]: 添加函数/结构体头部注释
	$(call logInfo, gocmt)
	@hash gocmt 2>&- || go get -u github.com/cuonglm/gocmt
	gocmt -i -t "TODO: comments" `find . -name '*.go' | grep -v _test.go`
	@grep -R "TODO: comments" * | grep -v makefile
	ag "TODO: comments"
	$(call logInfo, gocmt: done)



define logInfo
	@echo ["\033[32mINFO\033[0m"] [`date +'%Y-%m-%d %H:%M:%S'`] $(1)
endef

define logErrorP
	echo ["\033[31mERROR\033[0m"] [`date +'%Y-%m-%d %H:%M:%S'`] $(1)
endef


# --------------------
# 其他
# --------------------

help: ## 查看帮助文档
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

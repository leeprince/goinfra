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
	./$(projectName)

# --------------------
# 静态代码分析
# --------------------

.PHONY: sa
sa: vet errcheck golangci-lint revive build  ## [static analysis]: vet, errcheck, revive, lint

.PHONY: vet
# Go vet是Go语言自带的一个代码静态检查工具，它可以帮助我们在编译阶段发现代码中的一些常见错误，提高代码质量。
#
## 作用：
#    检查代码中是否存在可疑的构造，例如Printf调用的参数是否与格式字符串匹配。
#    检查代码中是否存在常见的错误，例如是否使用了未使用的变量或者包。
#    检查代码中是否存在不符合Go语言规范的地方，例如是否存在死代码或者无效的构造。
#
## 使用方法：
#在命令行中使用go vet命令，后面跟上要检查的文件名或者包名。例如：
#go vet main.go
#或者：
#go vet ./...
#
#这将会检查当前目录及其所有子目录下的所有Go文件。
#
## 注意事项：
#    Go vet是一个静态检查工具，它并不能保证程序的正确性。即使程序通过了go vet的检查，也可能存在运行时错误。
#    Go vet的检查并不全面，它只能检查出一部分常见的错误。对于一些复杂的逻辑错误，我们还需要借助其他的工具或者手动检查。
#    Go vet的检查可能会有误报或者漏报的情况，我们需要根据实际情况判断。
#    在使用go vet时，我们需要确保代码能够正确编译，否则go vet可能无法正确工作。
vet: ## [static analysis]: go vet ./...
	$(call logInfo, go vet)
	@go vet ./... || ($(call logErrorP, go vet: not pass) && exit 1) 
	$(call logInfo, go vet: pass)


.PHONY: errcheck
# golang errcheck 命令的作用，安装方式，使用方式和注意事项
## 作用：
#errcheck是一个用于Go语言的程序错误检查工具。它的主要作用是检查Go代码中所有函数调用的错误返回值是否都被正确处理。在Go语言中，错误处理是非常重要的一部分，如果忽视了错误的处理，可能会导致程序在运行时出现严重问题。
#
## 安装方式：
#你可以通过go get命令来装errcheck，如下所示：
#go install github.com/kisielk/errcheck
#
## 使用方式：
#安装完成后，你可以在命令行中使用errcheck命令来检查你的Go代码，如下所示：
# errcheck ./...
#
#这将会检查当前目录及其所有子目录下的所有Go文件。
#
## 注意事项：
#    errcheck只能检查函数调用的错误返回值是否被处理，它不能保证错误被正确处理。你需要确保你的代码中正确地处理了所有可能的错误情况。
#    errcheck可能会有误报或者漏报的情况，你需要根据实际情况判断。
#    在使用errcheck时，你需要确保你的代码能够正确编译，否则errcheck可能无法正确工作。
#    errcheck的检查可能会比较耗时，特别是在大型项目中。你可以考虑在提交代码之前或者定期运行errcheck，而不是在每次编译时都运行。
errcheck: ## [static analysis]: errcheck ./...
	$(call logInfo, errcheck)
	@hash errcheck 2>&- || go install github.com/kisielk/errcheck
	@errcheck ./... || ($(call logErrorP, errcheck: not pass) && exit 1)
	$(call logInfo, errcheck: pass)

.PHONY: golangci-lint
## 作用：
#golangci-lint是一个用于Go语言的Linters Runner，它可以运行多个linters并对结果进行汇总。它的目标是创建一个集成了所有可能需要的linters的工具，以提高代码质量和开发者的效率。
#官方：golangci-lint 是一个快速的 Go linters 运行器。它并行运行 linter，使用缓存，支持 yaml 配置，与所有主要的 IDE 集成，并包含数十个 linter。
#
## 安装方式：
#你可以通过以下命令来安装golangci-lint：
#go get github.com/golangci/golangci-lint/cmd/golangci-lint
#
#或者，你可以下载预编译的二进制文件，从golangci-lint的GitHub releases页面下载对应的版本。
#
## 使用方式：
#安装完成后，你可以在命令行中使用golangci-lint命令来检查你的Go代码，如下所示：
#golangci-lint run
#
#这将会在当前目录下运行所有的linters。
#
#你也可以通过配置文件来自定义golangci-lint的行为，例如：
#golangci-lint run --config .golangci.yml
#
#其中，.golangci.yml是你的配置文件，你可以在其中定义你需要的linters和规则。
#
## 注意事项：
#    golangci-lint只能帮助你发现代码中的问题，你需要根据它的建议来修复这些问题。
#    golangci-lint的检查可能会有误报或者漏报的情况，你需要根据实际情况判断。
#    在使用golangci-lint时，你需要确保你的代码能够正确编译，否则golangci-lint可能无法正确工作。
#    golangci-lint的检查可能会比较耗时，特别是在大型项目中。你可以考虑在提交代码之前或者定期运行golangci-lint，而不是在每次编译时都运行。
golangci-lint: ## [static analysis]: golangci-lint run
	$(call logInfo, golangci-lint)
	@hash golangci-lint 2>&- || go install github.com/golangci/golangci-lint/cmd/golangci-lint
	@golangci-lint run || ($(call logErrorP, golangci-lint: not pass) && exit 1)
	$(call logInfo, golangci-lint: pass)

.PHONY: golint
## 作用：
#golint是一个用于Go语言的代码审查工具，它可以帮助你发现代码中可能存在的样式问题。golint遵循Go语言的官方代码规范，并提供了一些额外的规则。
#
## 安装方式：
#你可以通过go get命令来安装golint，如下所示：
#go install golang.org/x/lint/golint
#
#使用方式：
#安装完成后，你可以在命令行中使用golint命令来检查你的Go代码，如下所示：
#golint ./...
#
#这将会检查当前目录及其所有子目录下的所有Go文件。
#
## 注意事项：
#    golint只能帮助你发现代码中的样式问题，它不能保证代码的正确性。你需要确保你的代码不仅遵循了Go语言的代码规范，而且也通过了充分的测试。
#    golint的检查可能会有误报或者漏报的情况，你需要根据实际情况判断。
#    在使用golint时，你需要确保你的代码能够正确编译，否则golint可能无法正确工作。
#    golint的检查可能会比较耗时，特别是在大型项目中。你可以考虑在提交代码之前或者定期运行golint，而不是在每次编译时都运行。
golint: ## [static analysis]: golint ./...
	$(call logInfo, golint)
	@hash golint 2>&- || go install golang.org/x/lint/golint
	@golint ./... || ($(call logErrorP, golint: not pass) && exit 1)
	$(call logInfo, golint: pass)

.PHONY: revive
# golang revive 命令的作用，安装方式，使用方式和注意事项
## 作用：
#Revive是一个用于Go语言的代码审查工具，它可以帮助开发者发现代码中的问题并提供修复建议。Revive的规则比go lint更丰富，执行速度更快，配置更灵活。
#官网：快速、可配置、可扩展、灵活且美观的 Go linter。golint 的直接替代品。复活提供用于开发自定义规则的框架，并允许您定义一个严格的预设来增强您的开发和代码审查过程。
#
## 安装方式：
#你可以通过go get命令来安装revive，如下所示：
#go install github.com/mgechev/revive
#
## 使用方式：
#安装完成后，你可以在命令行中使用revive命令来检查你的Go代码，如下所示：
#revive ./...
#
#这将会检查当前目录及其所有子目录下的所有Go文件。
#
#你也可以通过配置文件来自定义revive的规则，例如：
#revive -config revive.toml ./...
#
#其中，revive.toml是你的配置文件，你可以在其中定义你需要的规则。
#
## revive 接受以下命令行参数：
#-config[PATH]-TOML格式的配置文件的路径，默认为$HOME/reve.TOML（如果存在）。
#-exclude〔PATTERN〕：要为linting排除的文件/目录/包的模式。您可以将要排除的文件指定为程序包名称（即github.com/mgechev/revese），将它们列为单独的文件（即file.go）、目录（即./foo/…），或这三者的任意组合。
#
#- formatter [NAME]：用于输出的格式化程序。当前可用的格式化程序有：
#- - default：将以与golint相同的方式输出故障。
#- - json：以json格式输出故障。
#- - ndjson：以换行分隔的JSON（ndjson）格式将失败作为流输出。
#- - friendly：发现故障时输出。显示所有故障的摘要。
#- - stylish：将故障格式化在表格中。请记住，它不会流式传输输出，因此与其他输出相比，它可能会被认为较慢。
#- - checkstyle：以与Java的checkstyle兼容的XML格式输出故障。
#
#- max_open_files：同时打开的最大文件数。默认为无限制。
#- set_exit_status：如果发现任何问题，则将退出状态设置为1，覆盖配置中的errorCode和warningCode。
#- version：版本
#
## 注意事项：
#    revive只能帮助你发现代码中的问题，你需要根据它的建议来修复这些问题。
#    revive的检查可能会有误报或者漏报的情况，你需要根据实际情况判断。
#    在使用revive时，你需要确保你的代码能够正确编译，否则revive可能无法正确工作。
#    revive的检查可能会比较耗时，特别是在大型项目中。你可以考虑在提交代码之前或者定期运行revive，而不是在每次编译时都运行。
revive: ## [static analysis]: revive -formatter stylish ./...
	$(call logInfo, revive)
	@hash revive 2>&- || go install github.com/mgechev/revive
	revive -formatter stylish ./... || ($(call logErrorP, revive: not pass) && exit 1)
	$(call logInfo, revive: pass)

# --------------------
# 代码格式化相关
# --------------------

.PHONY: complete
complete: gofumpt gocmt  ## [formatting]: 运行代码格式化: fmt, cmt

.PHONY: fmt
## 作用：
#gofmt是Go语言的官方代码格式化工具。它可以自动格式化Go代码，使代码风格保持一致。这对于代码的可读性和维护性非常重要。
#
#安装方式：
#gofmt是Go语言的标准工具，当你安装Go语言时，gofmt会自动安装，无需额外安装。
#
## 使用方式：
#你可以在命令行中使用gofmt命令来格式化你的Go代码，如下所示：
#gofmt -w yourcode.go
#
#这将会格式化yourcode.go这个文件，并将修改直接写入文件。
#-w：将结果写入（源）文件而不是stdout。
#-l：列出格式与gofmt不同的文件。
#
#你也可以使用gofmt来格式化一个目录下的所有Go文件，如下所示：
#gofmt -l -w .
#
#这将会格式化当前目录及其所有子目录下的所有Go文件。
#
#注意事项：
#    gofmt的修改是无法撤销的，所以在使用gofmt之前，你应该确保你的代码已经被正确地保存和备份。
#    gofmt只能帮助你格式化代码，它不能保证代码的正确性。你需要确保你的代码不仅遵循了Go语言的代码规范，而且也通过了充分的测试。
#    在使用gofmt时，你需要确保你的代码能够正确编译，否则gofmt可能无法正确工作。
gofmt: ## [formatting]: gofmt -l -w .
	$(call logInfo, gofmt)
	gofmt -l -w .
	$(call logInfo, gofmt: done)

.PHONY: gofumpt
## 作用：
#gofumpt是一个Go语言的代码格式化工具，它是在gofmt的基础上进行了一些扩展和增强。gofumpt不仅会格式化代码，还会对代码进行一些额外的检查和修改，以使代码更加符合Go语言的规范和习惯。
#
## 安装方式：
#你可以通过go get命令来安装gofumpt，如下所示：
#go install mvdan.cc/gofumpt
#
## 使用方式：
#安装完成后，你可以在命令行中使用gofumpt命令来格式化你的Go代码，如下所示：
#gofumpt -l -w .
#-w：将结果写入（源）文件而不是stdout。
#-l：列出格式与gofumpt不同的文件。
#
#这将会格式化当前目录及其所有子目录下的所有Go文件。
#
## 注意事项：
#    gofumpt的修改是无法撤销的，所以在使用gofumpt之前，你应该确保你的代码已经被正确地保存和备份。
#    gofumpt只能帮助你格式化代码和进行一些基本的检查，它不能保证代码的正确性。你需要确保你的代码不仅遵循了Go语言的代码规范，而且也通过了充分的测试。
#    在使用gofumpt时，你需要确保你的代码能够正确编译，否则gofumpt可能无法正确工作。
gofumpt: ## [formatting]: gofumpt -l -w .
	$(call logInfo, gofumpt)
	@hash gofumpt 2>&- || go install mvdan.cc/gofumpt@latest
	gofumpt -l -w .
	$(call logInfo, gofumpt: done)

.PHONY: gocmt
## 作用：
#在 go 文件中为导出（允许外部访问的，即方法首字母大写的）的函数、方法、类型、常量、变量添加缺少的注释
#
## 安装：
#go install github.com/cuonglm/gocmt
#
## 输出哪些包含"TODO: comments"的文件
#@grep -R "TODO: comments" * | grep -v Makefile
gocmt: ## [formatting]: 添加函数/结构体头部注释
	$(call logInfo, gocmt)
	@hash gocmt 2>&- || go install github.com/cuonglm/gocmt
	gocmt -i -t "TODO: comments" `find . -name '*.go' | grep -v _test.go`
	#@grep -R "TODO: comments" * | grep -v Makefile
	$(call logInfo, gocmt: done)

#使用echo命令来输出文本，并使用转义序列来设置文本的样式
#-e选项用于启用转义序列的解释。\033[1m是转义序列，用于设置文本为粗体; \033[0m是另一个转义序列，用于重置文本样式到默认值。
#
#这仅在支持ANSI转义序列的终端中有效。
define logInfo
	@echo -e ["\033[32mINFO\033[0m" [`date +'%Y-%m-%d %H:%M:%S'`] $(1)
endef

define logErrorP
	@echo -e ["\033[31mERROR\033[0m" [`date +'%Y-%m-%d %H:%M:%S'`] $(1)
endef

# --------------------
# 一键安装所有工具
# --------------------
.PHONY: tool
tool: ## tool
	$(call logInfo, errcheck)
	@go install github.com/kisielk/errcheck@latest

	$(call logInfo, golangci-lint)
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

	$(call logInfo, golint)
	@go install golang.org/x/lint/golint@latest

	$(call logInfo, revive)
	@go install github.com/mgechev/revive@latest

	$(call logInfo, gofumpt)
	@go install mvdan.cc/gofumpt@latest

	$(call logInfo, gocmt)
	@go install github.com/cuonglm/gocmt

# --------------------
# 其他
# --------------------

help: ## 查看帮助文档
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) \
		| awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

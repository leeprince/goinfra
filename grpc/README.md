# protoc

protoc 是 Protocol Buffers（简称 Protobuf）的编译器，用于将 .proto 文件编译为各种语言的代码。Protobuf 是 Google 开发的一种高效的数据序列化协议，广泛用于网络通信、数据存储等场景。

# protoc插件

## **protoc-gen-go v1.25**

proto编译生成go文件

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

## **protoc-gen-micro v1.0.0**

protoc go-micro 插件安装

```
go install github.com/micro/protoc-gen-micro@latest
```

## **protoc-gen-go-grpc**

```
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

# 使用

## proto/user/user.proto

```
protoc --go_out=./proto/user --go-grpc_out=./proto/user --micro_out=./proto/user proto/user/user.proto
```

# protoc

## 一、生成 xxx.pb.go、xxx_grpc.pb.go、xxx.pb.gw.go

### （一）进入protocol 项目的 *.proto 所有目录【推荐】

1. 进入 *.proto 所有目录
2. 在当前目录中新建 *.proto 对应包名的文件夹`tscproxyjy`(便于整体移动)
3. 执行

```

protoc
--go-grpc_opt=require_unimplemented_servers=false
--grpc-gateway_opt=logtostderr=true
--grpc-gateway_opt=paths=source_relative
-I .
-I ../../../vendors
--go_out=./tscproxyjy
--go-grpc_out=./tscproxyjy
--grpc-gateway_out=./tscproxyjy
./*.proto

```

### （二）进入protocol 项目的 *.proto 所有目录

1. 进入 *.proto 所有目录

2. 复制 vendors 到当前目录中

 `-I vendors`可以替换为相对路径 `-I ../../../vendors` 就可以不移动 vendors 到当前目录中

3. 执行

```

protoc \
--go-grpc_opt=require_unimplemented_servers=false \
--grpc-gateway_opt=logtostderr=true \
--grpc-gateway_opt=paths=source_relative \
-I . \
-I vendors \
--go_out=. \
--go-grpc_out=. \
--grpc-gateway_out=. \
./*.proto

```

### （三）进入 protocol 项目根目录

> 仅没有通过 import 引入其他 proto 文件时生效，否则报错：proto/xxxx.proto:41:60: "xxx" seems to be defined in "proto/xxx.proto", which is not imported by "proto/xxx.proto".  To use it here, please add the necessary import.

```

protoc \
--go-grpc_opt=require_unimplemented_servers=false \
--grpc-gateway_opt=logtostderr=true \
--grpc-gateway_opt=paths=source_relative \
-I proto \
-I vendors \
--go_out=. \
--go-grpc_out=. \
--grpc-gateway_out=. \
./proto/tax-service-center/tsc-open-api/*.proto

```

## 二、grpc协议生成php代码

### （一）**方式一**

```
<https://xxx.xxx.com/xxx-cloud/protocol/-/blob/master/gen_php.sh>

脚本说明：

 proto_path="finance-center"
 out_path="../gc-admin/app/Protos/"

 protocol/ 和 gc-admin/ 在同级文件目录

 proto_path的对应的路径为：  <https://gitlab.xxx.com/xxx-cloud/protocol/-/tree/master/proto/finance-center>

 out_path对应本地项目路径     gc-admin/app/Protos/

 即：将proto/finance-center中的协议生成到gc-admin项目的app/Protos/目录下

再在docker的php容器中执行以下命令，加载项目新增的协议文件

cd /usr/share/nginx/html/ && composer dump-autoload

```

### （二）**方式二**

- 将协议生成php代码

 ```

 protoc -I protocol/proto/usercenter \
  --php_out=./gc-admin/common/proto \
  --grpc_out=./gc-admin/common/proto \
  --plugin=protoc-gen-grpc=manager-center/grpc_php_plugin  
  protocol/proto/usercenter/userInfo.proto

 ```

- composer.json配置

 ```

 "autoload": {
     "classmap": [
         ...
         "common/proto/"
     ]
 },

 ```

- 执行以下命令自动加载class composer dump-autoload

## 三、gxxxn_cloud protocol 项目示例

### （一）多级目录生成proto协议需注意

#### 报错 `has inconsistent names`

> ---------- 注意场景。
>
> 目录结构为：
>
> - industry-commerce-center
>
>   - individual-biz
>
>   • biz.proto
>
>   - ics_open_api.proto
>
> 因为 `ics_open_api.proto`  包含 `import "individual-biz/biz.proto";`。`ics_open_api.proto`  定义了 `option go_package = ".;icscenterapi";`， 而 `individual-biz/biz.proto` 定义了 `option go_package = ".;individualbiz";` 则会报错 `protoc-gen-go: Go package "." has inconsistent names icscenterapi (ics_open_api.proto) and individualbiz (individual-biz/biz.proto)`。

#### 解决

需要在 `individual-biz/biz.proto` 修改 **go_package** 的值。方法有两种：

**1. go_package 定义修改为完整的包路径【推荐】。**如：`github.com/leeprince/protocol/grpc/industrycommerce/individualbiz;individualbiz`。

**2. go_package 临时定义替代**`**.**`**定义。**

使用临时定义替代`.`定义，不过这种方法需要注意的是根据  `individual-biz/biz.proto`  生成的 `*.go` 之后与`biz.proto`的相关依赖的包路径为临时 go_package 定义，需要手动修改，所以还是推荐被依赖的包 `individual-biz/biz.proto` 尽量使用完整的包路径，即方法1.

### （二）`industry-commerce-center`示例

#### 目录结构

```

- grpc
- proto
    industry-commerce-center
        - individual-biz
          - biz.proto
        - ics_open_api.proto
- vendor
- Makefile

```

#### 通过 `biz.proto` 生成新的协议

1. cd 到 `proto/industry-commerce-center`

 `proto/industry-commerce-center`目录下手动新建`tmp_local`文件夹

2. 执行

 ```

 protoc \
 --go-grpc_opt=require_unimplemented_servers=false \
 --grpc-gateway_opt=logtostderr=true \
 --grpc-gateway_opt=paths=source_relative \
 -I . \
 -I ../../vendors \
 --go_out=./tmp_local \
 --go-grpc_out=./tmp_local \
 --grpc-gateway_out=./tmp_local \
 ./individual-biz/*.proto

 ```

#### 通过 `ics_open_api.proto`生成新的协议

1. cd 到 `proto/industry-commerce-center`

 > `proto/industry-commerce-center`目录下手动新建`tmp_local`文件夹

2. 执行所有.proto

 ```

 protoc \
 --go-grpc_opt=require_unimplemented_servers=false \
 --grpc-gateway_opt=logtostderr=true \
 --grpc-gateway_opt=paths=source_relative \
 -I . \
 -I ../../vendors \
 --go_out=./tmp_local \
 --go-grpc_out=./tmp_local \
 --grpc-gateway_out=./tmp_local \
 ./*.proto

 ```

 **指定的proto**

 ```

 protoc \
 --go-grpc_opt=require_unimplemented_servers=false \
 --grpc-gateway_opt=logtostderr=true \
 --grpc-gateway_opt=paths=source_relative \
 -I . \
 -I ../../vendors \
 --go_out=./tmp_local \
 --go-grpc_out=./tmp_local \
 --grpc-gateway_out=./tmp_local \
 ./ics_open_api.proto

 ```

# gozero
官方文档：https://go-zero.dev/docs/tasks
---

## demogrpc

### 启动
go run demogrpc.go

### 测试
#### grpcurl
测试命令
```
# 安装grpcurl命令
#   mac: brew install grpcurl
grpcurl -plaintext  127.0.0.1:8080 demogrpc.Demogrpc.Ping
```

报错
```
Error invoking method "demogrpc.Demogrpc.Ping": failed to query for service descriptor "demogrpc.Demogrpc": server does not support the reflection API
```

原因：未在给定的gRPC服务器上注册服务器反射服务。

分析：在`demogrpc.go`中有如下代码
```
if c.Mode == service.DevMode || c.Mode == service.TestMode {
    reflection.Register(grpcServer)
}
```
所以解决办法及时在配置文件`gozero-demo/demogrpc/etc/demogrpc.yaml` 中加上
```
Mode: dev # 或者 test 即可
```


#### grpcui
测试命令
```
# 安装grpcurl命令
#   mac: brew install grpcui
grpcui -plaintext 127.0.0.1:8080
```

报错（报错的原因与 `grpcurl` 一致，按照上面即可解决）
```
Failed to compute set of methods to expose: server does not support the reflection API
```

#### Goland 中使用 `GRPC` 命令
在 Goland 的中安装`Endpoints`插件，即可在Goland中点开`Endpoints`插件的控制面板，并使用`GRPC`命令来测试

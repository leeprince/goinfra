# 分布式链路追踪

## 目录结构
```
.
├── README.md
├── opentracing
│   ├── jaeger_client // 封装 Jaeger 客户端
│   │   ├── ...
│   └── opentelemetry_client  // 封装 Opentelemetry 客户端
│       ├── ...
└── opentracing_tutorial // 测试
    ├── jaeger_client // 参考:https://github.com/leeprince/opentracing-tutorial/tree/master/go
    │   ├── lesson01 // 测试1
    │   ├── lesson02 // 测试2
    │   ├── lesson03 // 测试3
    │   └── opentracing_test // 结合 ../../opentracing/jaeger_client/ 的测试
    └── opentelemetry_client // 参考:https://opentelemetry.io/docs/instrumentation/go/getting-started/
        ├── 1get_started // 测试1
        ├── 2libraries // 测试2
        ├── 3manual // 测试3
        ├── 4export_data // 测试4
        ├── 5passthrough // 测试5
        ├── opentracing_test // 结合../../opentracing/opentelemetry_client/ 的测试
        └── opentracing_test_app // 结合 ../../opentracing/jaeger_client/ 的应用测试
```


## 客户端
### [Jaeger](https://www.jaegertracing.io) 的作为 opentracing 客户端

github:[https://github.com/jaegertracing/jaeger-client-go](https://github.com/jaegertracing/jaeger-client-go)


#### 日志

jaegerLogger 基于 github.com/leeprince/goinfra/plog 实现 github.com/uber/jaeger-client-go@logger.go 的 Logger 接口

### [opentelemetry](https://opentelemetry.io/) 的作为 opentracing 客户端
github:[https://github.com/open-telemetry/opentelemetry-go](https://github.com/open-telemetry/opentelemetry-go)

> 注意：
> opentelemetry 已兼容 Jaeger 客户端的功能：在 opentelemetry 中 `Exporter` 导出的数据使用 Jaeger 的收集器路由地址(默认：`http://localhost:14268/api/traces`)即可.
> 虽然 opentelemetry 已兼容 Jaeger 客户端的功能，但是数据收集及展示UI这块还是 Jaeger 为主

## 部署 Jaeger 看板
### all-in-one
```
## make sure to expose only the ports you use in your deployment scenario!
docker run -d --name jaeger \
  -e COLLECTOR_ZIPKIN_HOST_PORT=:9411 \
  -p 5775:5775/udp \
  -p 6831:6831/udp \
  -p 6832:6832/udp \
  -p 5778:5778 \
  -p 16686:16686 \
  -p 14250:14250 \
  -p 14268:14268 \
  -p 14269:14269 \
  -p 9411:9411 \
  jaegertracing/all-in-one:1.32
```

> [Jaeger UI](http://localhost:16686):http://localhost:16686

- all-in-one 容器暴露的端口说明
| Port  | Protocol | Component | Function                                                     |
| ----- | -------- | --------- | ------------------------------------------------------------ |
| 5775  | UDP      | agent     | accept zipkin.thrift over compact thrift protocol (deprecated, used by legacy clients only) |
| 6831  | UDP      | agent     | accept jaeger.thrift over compact thrift protocol            |
| 6832  | UDP      | agent     | accept jaeger.thrift over binary thrift protocol             |
| 5778  | HTTP     | agent     | serve configs                                                |
| 16686 | HTTP     | query     | serve frontend                                               |
| 14268 | HTTP     | collector | accept jaeger.thrift directly from clients                   |
| 14250 | HTTP     | collector | accept model.proto                                           |
| 9411  | HTTP     | collector | Zipkin compatible endpoint (optional)                        |


## 应用架构中，在 http、gRPC 请求中设置中间件（拦截器）开启分布式调用链追踪

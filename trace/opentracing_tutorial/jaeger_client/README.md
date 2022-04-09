# jaeger 的作为 opentracing 客户端
github.com/uber/jaeger-client-go 

# 参考
- https://github.com/leeprince/opentracing-tutorial/tree/master/go

# 目录结构
├── lesson01 // 课程1
├── lesson02 // 课程2
├── lesson03 // 课程3
└── opentracing_test // 结合 ../../opentracing/jaeger_client.go 的测试

# 日志
## 实现 github.com/uber/jaeger-client-go@logger.go 的 Logger 接口
jaegerLogger 基于 github.com/leeprince/goinfra/plog 实现 github.com/uber/jaeger-client-go@logger.go 的 Logger 接口

# 分布式调用链追踪

## 参考
- https://github.com/leeprince/opentracing-tutorial/tree/master/go


## 日志
### 实现 github.com/uber/jaeger-client-go@logger.go 的 Logger 接口
jaegerLogger 基于 github.com/leeprince/goinfra/plog 实现 github.com/uber/jaeger-client-go@logger.go 的 Logger 接口

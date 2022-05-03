# 分布式链路追踪


## jaeger
### 部署 jaeger 后端环境
#### all-in-one
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

##### all-in-one 容器暴露的端口说明
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

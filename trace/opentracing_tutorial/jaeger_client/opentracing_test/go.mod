module opentracing_test

go 1.16

require (
	github.com/leeprince/goinfra v0.0.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
	golang.org/x/tools v0.8.0 // indirect
)

// 替换为本地包
replace github.com/leeprince/goinfra => ../../../../

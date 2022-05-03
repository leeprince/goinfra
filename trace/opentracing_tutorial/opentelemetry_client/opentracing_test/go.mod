module opentelemetry_opentracing_test

go 1.16

require (
	github.com/leeprince/goinfra v0.0.0
	go.opentelemetry.io/otel v1.6.3
	go.opentelemetry.io/otel/trace v1.6.3
)

// 替换为本地包
replace github.com/leeprince/goinfra => ../../../../

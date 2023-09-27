module lesson03

go 1.16

require (
	github.com/leeprince/goinfra v0.0.0
	github.com/opentracing/opentracing-go v1.2.0
	github.com/uber/jaeger-client-go v2.30.0+incompatible
)

replace (
	github.com/leeprince/goinfra => ../../../../../goinfra
)
module go.opentelemetry.io/otel/example/passthrough

go 1.16

require (
	github.com/leeprince/goinfra v0.0.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.31.0
	go.opentelemetry.io/otel v1.6.3
	go.opentelemetry.io/otel/exporters/jaeger v1.6.3
	go.opentelemetry.io/otel/sdk v1.6.3
	go.opentelemetry.io/otel/trace v1.6.3
)

//replace github.com/leeprince/goinfra => /Users/leeprince/www/go/goinfra
replace github.com/leeprince/goinfra => F:/www/goinfra

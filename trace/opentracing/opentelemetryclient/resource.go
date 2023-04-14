package opentelemetryclient

import (
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/15 下午6:17
 * @Desc:   资源：描述应用的资源
 */

func newResource(serviceName, env string) (*resource.Resource, error) {
	return resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
			attribute.String("environment", env),
		),
	)
}

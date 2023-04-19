package rabbitmq

import (
	"context"
	"github.com/leeprince/goinfra/perror"
	"github.com/leeprince/goinfra/trace/opentracing/jaegerclient"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/18 11:29
 * @Desc:	对应交换机类型的发布消息+链路追踪封装方法
 */

// 简单队列（一次向一个消费者发送消息）
func (cli *RabbitMQClient) PublishSimpleTrace(ctx context.Context, body []byte, opts ...PropertiesOpt) (err error) {
	if jaegerclient.SpanFromContext(ctx) == nil {
		return perror.BizErrObjNil
	}

	err = cli.Publish(
		WithPublishParamProperties(
			body,
			opts...,
		),
	)
	return
}

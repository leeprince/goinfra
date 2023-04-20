package rabbitmqgdtest

import (
	"context"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/leeprince/goinfra/mq/rabbitmq/rabbitmqtest/rabbitmqgdtest/pbfile"
	"github.com/opentracing/opentracing-go"
	"github.com/streadway/amqp"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/12/7 17:28
 * @Desc:
 */

// ---------------------------- 生产者、消费者 PublishTracerV2:测试span互通,logID互通,jaeger上报为完整数据链-v2 -------------------------------------
// 测试发布消息
func TestMQservicePublishTracerV2(t *testing.T) {
	fmt.Println("初始化链路追踪...")
	tracer.InitTracer(Config.JaegerAgentUri, "prince-jaeger-mq-test-publish")

	ctx := context.Background()
	logID := common.LogIdByCtx(ctx)
	fmt.Println("logID------0", logID)

	mqReq := pbfile.ClientReportBussinessResult{
		RequestSeq:   logID,
		EnterpriseId: "xxxxxxxxx",
		ResultType:   pbfile.ClientReportBussinessResultResultType(1),
		Body:         nil,
	}
	// 获取配置
	rabbitConf := GetRabbitConf(RABBIT_CONFKEY)
	fmt.Println(rabbitConf)
	if rabbitConf == nil {
		panic("GetRabbitConf(RABBIT_CONFKEY)")
	}

	conn := GetMqConn(RABBIT_CONFKEY)
	err := conn.DeclareQueue(rabbitConf.QueueName, rabbitConf.Key)
	if err != nil {
		gclog.WithField("mqReq", &mqReq).WithError(err).Error(logID, " MQservice.TestMQservicePublishTracer DeclareQueue err")
		panic("Publish(rabbitConf.Key, mqReq)")
	}

	// 模拟应用启动的时候已开启的**根span**
	span := tracer.GetTracer().StartSpan("PublishTracer-op")
	defer span.Finish()
	fmt.Println("context-----1", ctx)
	logID = common.LogIdByCtx(ctx)
	fmt.Println("logID-----1", logID)
	ctx = opentracing.ContextWithSpan(ctx, span)
	fmt.Println("context-----2", ctx)
	logID = common.LogIdByCtx(ctx)
	fmt.Println("logID-----2", logID)

	err = conn.PublishTracerV2(ctx, rabbitConf.Key, &mqReq)
	if err != nil {
		panic("conn.PublishTracer(context.Background(), rabbitConf.Key, mqReq)")
	}
	logID = common.LogIdByCtx(ctx)
	fmt.Println("logID-----3", logID)

	fmt.Println(" MQservice.TestMQservicePublishTracer successfuly info")

	//需要添加select{} 否则无法上报当前服务的span
	select {}
}

// 测试消费消息
func TestMQserviceComsumptionTraceAndParseDataV2(t *testing.T) {
	fmt.Println("初始化链路追踪...")
	tracer.InitTracer(Config.JaegerAgentUri, "prince-jaeger-mq-test-comsumption")

	// 获取配置
	rabbitConf := GetRabbitConf(RABBIT_CONFKEY)
	if rabbitConf == nil {
		println(fmt.Sprintf(" mq.GetRabbitConf rabbitConf == nil"))
		return
	}
	fmt.Println(" rabbit rabbitConf:", rabbitConf)

	conn := GetMqConn(RABBIT_CONFKEY)

	consumeHandler := func(msg *amqp.Delivery) {
		fmt.Println(" -msg.Body：", string(msg.Body))
		ctx, dataBytes, err := conn.ParsePublishTracerDataV2(msg)
		if err != nil {
			fmt.Println("consumeHandler ParsePublishTracerData err", err.Error())
			msg.Reject(false)
			return
		}
		fmt.Println(ctx, dataBytes)
		fmt.Println("------------")

		logID := common.LogIdByCtx(ctx)
		fmt.Println("logID++++++++++", logID)

		var msgData pbfile.ClientReportBussinessResult
		err = jsonpb.UnmarshalString(string(dataBytes), &msgData)
		if err != nil {
			fmt.Println("jsonpb.UnmarshalString(string(dataBytes), &msgData) err", err)
			msg.Reject(false)
			return
		}
		fmt.Println("******* pbfile.ClientReportBussinessResult *******")
		fmt.Printf("msgData %+v \n", &msgData)
		fmt.Printf("msgData %#v \n", &msgData)
		// --- 原数据解析 -end

		msg.Ack(true)

		return
	}
	// 消费协程
	conn.Consume(consumeHandler, rabbitConf.QueueName, 2, rabbitConf.Key)
}

// ---------------------------- 生产者、消费者 PublishTracer:测试span互通,logID互通,jaeger上报为完整数据链-v2 -end -------------------------------------

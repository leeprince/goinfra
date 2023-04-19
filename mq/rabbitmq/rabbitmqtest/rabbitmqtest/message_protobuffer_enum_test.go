package rabbitmqtest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/leeprince/goinfra/mq/rabbitmq/rabbitmqtest/rabbitmqtest/pbfile"
	"github.com/streadway/amqp"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/17 17:32
 * @Desc:
 */

// ---------------------------- 生产者、消费者 proto 中定义 enum 类型 -------------------------------------
// 测试发布消息
func TestMQservicePublishEnum(t *testing.T) {
	logID := "prince-test-TestMQservicePublishEnum"
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
		panic("jaegermq.GetRabbitConf(RABBIT_CONFKEY)")
	}

	conn := GetMqConn(RABBIT_CONFKEY)
	err := conn.DeclareQueue(rabbitConf.QueueName, rabbitConf.Key)
	if err != nil {
		panic("conn.DeclareQueue(rabbitConf.QueueName, rabbitConf.Key)")
	}
	// 发布成功的消息为：{"request_seq":"prince-test-TestMQservicePublishEnum","enterprise_id":"xxxxxxxxx","result_type":"QrCode","body":null}
	err = conn.Publish(rabbitConf.Key, &mqReq)
	if err != nil {
		gclog.WithField("mqReq", &mqReq).WithError(err).Error(logID, " MQservice.TestMQservicePublish mq.GetRabbitConf err")
		panic("Publish(rabbitConf.Key, mqReq)")
	}

	fmt.Println(logID, " MQservice.TestMQservicePublish successfuly info")
	return
}

// 测试消费消息
func TestMQserviceComsumptionPublishEnum(t *testing.T) {
	consumeHandler := func(msg *amqp.Delivery) {
		var msgData pbfile.ClientReportBussinessResult
		err := json.Unmarshal(msg.Body, &msgData)
		if err != nil {
			// 正确的解析方式
			fmt.Println("json.Unmarshal err: ", err, "-msg.Body：", string(msg.Body))
			reader := bytes.NewReader(msg.Body)
			err := jsonpb.Unmarshal(reader, &msgData)
			if err != nil {
				fmt.Println(" jsonpb.Unmarshal(reader, &msgData) err: ", err, "-msg.Body：", string(msg.Body))
				msg.Reject(false)
				return
			}
			fmt.Println(">>>>>>>>>>>>>>>>++++++++")
		}

		fmt.Println("******* pbfile.ClientReportBussinessResult *******", &msgData)
		fmt.Printf("TestMQserviceComsumptionPublishEnum msgData %+v \n", &msgData)
		fmt.Printf("TestMQserviceComsumptionPublishEnum msgData %#v \n", &msgData)
		msg.Ack(true)

		return
	}
	// 获取配置
	rabbitConf := GetRabbitConf(RABBIT_CONFKEY)
	if rabbitConf == nil {
		println(fmt.Sprintf(" TestMQserviceComsumptionPublishEnum mq.GetRabbitConf rabbitConf == nil"))
		return
	}
	fmt.Println(" TestMQserviceComsumptionPublishEnum rabbit rabbitConf:", rabbitConf)

	conn := GetMqConn(RABBIT_CONFKEY)

	// 消费协程
	conn.Consume(consumeHandler, rabbitConf.QueueName, 2, rabbitConf.Key)
}

// ---------------------------- 生产者、消费者 proto 中定义 enum 类型 -end -------------------------------------

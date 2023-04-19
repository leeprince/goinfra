package rabbitmqtest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/jsonpb"
	"github.com/streadway/amqp"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/17 17:30
 * @Desc:
 */

// ---------------------------- 生产者、消费者 proto 中定义的结构体 -------------------------------------
// 测试发布消息
func TestMQservicePublish(t *testing.T) {
	logID := "prince-test-TestMQservicePublishEnum"
	mqReq := pbfile.CommonConnectInfo{
		ClientAddr:   "prince-ClientAddr",
		CreatedAt:    0,
		UpdatedAt:    0,
		EnterpriseId: "prince-EnterpriseId",
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
	err = conn.Publish(rabbitConf.Key, &mqReq)
	if err != nil {
		gclog.WithField("mqReq", &mqReq).WithError(err).Error(logID, " MQservice.TestMQservicePublish mq.GetRabbitConf err")
		panic("Publish(rabbitConf.Key, mqReq)")
	}

	fmt.Println(logID, " MQservice.TestMQservicePublish successfuly info")
	return
}

// 测试消费消息
func TestMQserviceComsumptionPublish(t *testing.T) {
	consumeHandler := func(msg *amqp.Delivery) {
		var msgData pbfile.CommonConnectInfo
		err := json.Unmarshal(msg.Body, &msgData)
		if err != nil {
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

// ---------------------------- 生产者、消费者 proto 中定义的结构体 -end -------------------------------------

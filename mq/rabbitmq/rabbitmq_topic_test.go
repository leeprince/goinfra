package rabbitmq

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/streadway/amqp"
	"sync"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/17 15:37
 * @Desc:
 */

func TestRabbitMQClient_PublishTopic(t *testing.T) {
	type fields struct {
		conf     *rabbitMQConf
		conn     *amqp.Connection
		connChan *amqp.Channel
		queue    amqp.Queue
		mt       sync.Mutex
	}
	type args struct {
		body  []byte
		topic string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			args: args{
				body:  []byte(fmt.Sprintf("TestRabbitMQClient_PublishTopic-%s-%s", routingKeyTopic, cast.ToString(time.Now().UnixNano()))),
				topic: routingKeyTopic,
			},
		},
		{
			args: args{
				body:  []byte(fmt.Sprintf("TestRabbitMQClient_PublishTopic-%s-%s", routingKeyTopic01, cast.ToString(time.Now().UnixNano()))),
				topic: routingKeyTopic01,
			},
		},
		{
			args: args{
				body:  []byte(fmt.Sprintf("TestRabbitMQClient_PublishTopic-%s-%s", routingKeyTopic02, cast.ToString(time.Now().UnixNano()))),
				topic: routingKeyTopic02,
			},
		},
		{
			args: args{
				body:  []byte(fmt.Sprintf("TestRabbitMQClient_PublishTopic-%s-%s", routingKeyTopic03, cast.ToString(time.Now().UnixNano()))),
				topic: routingKeyTopic03,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, err := NewRabbitMQClient(
				WithUrl(myUrl),
				WithExchangeDeclare(
					exchangeNameTopic,
					ExchangeTypeTopic,
					WithExchangeDeclareDurable(true),
				),
				WithQueueDeclare(
					queueNameTopic,
					WithQueueDeclareDurable(false),
				),
				WithRoutingKey(tt.args.topic),
			)
			if err != nil {
				fmt.Println("NewRabbitMQClient err:", err)
				return
			}
			fmt.Println("Boby:", string(tt.args.body))
			if err := cli.PublishTopic(
				tt.args.body,
			); (err != nil) != tt.wantErr {
				t.Errorf("PublishDirect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRabbitMQClient_ConsumeTopic(t *testing.T) {
	handle := func(msg amqp.Delivery) {
		fmt.Printf("msg:%+v \n", msg)
		fmt.Printf("msg.Headers::%+v \n", msg.Headers)
		fmt.Printf("msg.Headers::%#v \n", msg.Headers)
		fmt.Println("string(msg.requestBody):", string(msg.Body))

		time.Sleep(time.Second * 2)

		// 设置自动确认后，不可以手动确认，否则报错
		// --- 手动回复
		// msg.Reject(false)
		// msg.Reject(true)
		// ---0
		// msg.Ack(false)
		// msg.Ack(true)
		// ---1
		// msg.Nack(false, false)
		// msg.Nack(false, true)
		// msg.Nack(true, true)
		// msg.Nack(true, false)
		// --- 手动回复

		fmt.Println("msg --- end")
	}

	topic := "prince.RoutingKey.*"
	cli, err := NewRabbitMQClient(
		WithUrl(myUrl),
		WithExchangeDeclare(
			exchangeNameTopic,
			ExchangeTypeTopic,
			WithExchangeDeclareDurable(true),
		),
		WithQueueDeclare(
			"",
			WithQueueDeclareDurable(false),
			WithQueueDeclareExclusive(true),
			WithQueueDeclareAutoDelete(false),
		),
		WithRoutingKey(topic),
	)
	if err != nil {
		fmt.Println("NewRabbitMQClient err:", err)
		return
	}

	fmt.Println("topic::::::::::::::", topic)
	err = cli.ConsumeTopic(
		handle,
		WithConsumeParamOptExclusive(true),
	)
	if err != nil {
		t.Errorf("ConsumeTopic() error = %v", err)
	}
}

func TestRabbitMQClient_ConsumeTopic00(t *testing.T) {
	handle := func(msg amqp.Delivery) {
		fmt.Printf("msg:%+v \n", msg)
		fmt.Printf("msg.Headers::%+v \n", msg.Headers)
		fmt.Printf("msg.Headers::%#v \n", msg.Headers)
		fmt.Println("string(msg.requestBody):", string(msg.Body))

		time.Sleep(time.Second * 2)

		// 设置自动确认后，不可以手动确认，否则报错
		// --- 手动回复
		// msg.Reject(false)
		// msg.Reject(true)
		// ---0
		// msg.Ack(false)
		// msg.Ack(true)
		// ---1
		// msg.Nack(false, false)
		// msg.Nack(false, true)
		// msg.Nack(true, true)
		// msg.Nack(true, false)
		// --- 手动回复

		fmt.Println("msg --- end")
	}

	topic := "prince.RoutingKey.*"
	cli, err := NewRabbitMQClient(
		WithUrl(myUrl),
		WithExchangeDeclare(
			exchangeNameTopic,
			ExchangeTypeTopic,
			WithExchangeDeclareDurable(true),
		),
		WithQueueDeclare(
			"",
			WithQueueDeclareDurable(false),
			WithQueueDeclareExclusive(true),
			WithQueueDeclareAutoDelete(false),
		),
		WithRoutingKey(topic),
	)
	if err != nil {
		fmt.Println("NewRabbitMQClient err:", err)
		return
	}

	fmt.Println("topic::::::::::::::", topic)
	err = cli.ConsumeTopic(
		handle,
		WithConsumeParamOptExclusive(true),
	)
	if err != nil {
		t.Errorf("ConsumeTopic() error = %v", err)
	}
}

func TestRabbitMQClient_ConsumeTopic01(t *testing.T) {
	handle := func(msg amqp.Delivery) {
		fmt.Printf("msg:%+v \n", msg)
		fmt.Printf("msg.Headers::%+v \n", msg.Headers)
		fmt.Printf("msg.Headers::%#v \n", msg.Headers)
		fmt.Println("string(msg.requestBody):", string(msg.Body))

		time.Sleep(time.Second * 2)

		// 设置自动确认后，不可以手动确认，否则报错
		// --- 手动回复
		// msg.Reject(false)
		// msg.Reject(true)
		// ---0
		// msg.Ack(false)
		// msg.Ack(true)
		// ---1
		// msg.Nack(false, false)
		// msg.Nack(false, true)
		// msg.Nack(true, true)
		// msg.Nack(true, false)
		// --- 手动回复

		fmt.Println("msg --- end")
	}

	topic := "prince.*.topic"
	cli, err := NewRabbitMQClient(
		WithUrl(myUrl),
		WithExchangeDeclare(
			exchangeNameTopic,
			ExchangeTypeTopic,
			WithExchangeDeclareDurable(true),
		),
		WithQueueDeclare(
			"",
			WithQueueDeclareDurable(false),
			WithQueueDeclareExclusive(true),
			WithQueueDeclareAutoDelete(false),
		),
		WithRoutingKey(topic),
	)
	if err != nil {
		fmt.Println("NewRabbitMQClient err:", err)
		return
	}

	fmt.Println("topic::::::::::::::", topic)
	err = cli.ConsumeTopic(handle)
	if err != nil {
		t.Errorf("ConsumeTopic() error = %v", err)
	}
}

func TestRabbitMQClient_ConsumeTopic02(t *testing.T) {
	handle := func(msg amqp.Delivery) {
		fmt.Printf("msg:%+v \n", msg)
		fmt.Printf("msg.Headers::%+v \n", msg.Headers)
		fmt.Printf("msg.Headers::%#v \n", msg.Headers)
		fmt.Println("string(msg.requestBody):", string(msg.Body))

		time.Sleep(time.Second * 2)

		// 设置自动确认后，不可以手动确认，否则报错
		// --- 手动回复
		// msg.Reject(false)
		// msg.Reject(true)
		// ---0
		// msg.Ack(false)
		// msg.Ack(true)
		// ---1
		// msg.Nack(false, false)
		// msg.Nack(false, true)
		// msg.Nack(true, true)
		// msg.Nack(true, false)
		// --- 手动回复

		fmt.Println("msg --- end")
	}

	topic := "prince.#"
	cli, err := NewRabbitMQClient(
		WithUrl(myUrl),
		WithExchangeDeclare(
			exchangeNameTopic,
			ExchangeTypeTopic,
			WithExchangeDeclareDurable(true),
		),
		WithQueueDeclare(
			"",
			WithQueueDeclareDurable(false),
			WithQueueDeclareExclusive(true),
			WithQueueDeclareAutoDelete(false),
		),
		WithRoutingKey(topic),
	)
	if err != nil {
		fmt.Println("NewRabbitMQClient err:", err)
		return
	}

	fmt.Println("topic::::::::::::::", topic)
	err = cli.ConsumeTopic(handle)
	if err != nil {
		t.Errorf("ConsumeTopic() error = %v", err)
	}
}

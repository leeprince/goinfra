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

func TestRabbitMQClient_PublishDirect(t *testing.T) {
	type fields struct {
		conf     *rabbitMQConf
		conn     *amqp.Connection
		connChan *amqp.Channel
		queue    amqp.Queue
		mt       sync.Mutex
	}
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			args: args{
				body: []byte(fmt.Sprintf("TestRabbitMQClient_PublishDirect-%s", cast.ToString(time.Now().UnixNano()))),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, err := NewRabbitMQClient(
				WithUrl(myUrl),
				WithExchangeDeclare(
					exchangeNameDirect,
					ExchangeTypeDirect,
					WithExchangeDeclareDurable(true),
				),
				WithQueueDeclare(
					queueNameDirect,
				),
				WithRoutingKey(routingKeyDirect),
			)
			if err != nil {
				fmt.Println("NewRabbitMQClient err:", err)
				return
			}
			fmt.Println("Boby:", string(tt.args.body))
			if err := cli.PublishDirect(tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("PublishDirect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRabbitMQClient_PublishDirect01(t *testing.T) {
	type fields struct {
		conf     *rabbitMQConf
		conn     *amqp.Connection
		connChan *amqp.Channel
		queue    amqp.Queue
		mt       sync.Mutex
	}
	type args struct {
		body []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			args: args{
				body: []byte(fmt.Sprintf("TestRabbitMQClient_PublishDirect-%s", cast.ToString(time.Now().UnixNano()))),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, err := NewRabbitMQClient(
				WithUrl(myUrl),
				WithExchangeDeclare(
					exchangeNameDirect,
					ExchangeTypeDirect,
					WithExchangeDeclareDurable(true),
				),
				WithQueueDeclare(
					queueNameDirect01,
					WithQueueDeclareDurable(false),
				),
				WithRoutingKey(routingKeyDirect),
			)
			if err != nil {
				fmt.Println("NewRabbitMQClient err:", err)
				return
			}
			fmt.Println("Boby:", string(tt.args.body))
			if err := cli.PublishDirect(
				tt.args.body,
				WithPropertiesDeliveryMode(PropertiesDeliveryModePersistent),
			); (err != nil) != tt.wantErr {
				t.Errorf("PublishDirect() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRabbitMQClient_ConsumeDirect(t *testing.T) {
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

	cli, err := NewRabbitMQClient(
		WithUrl(myUrl),
		WithExchangeDeclare(
			exchangeNameDirect,
			ExchangeTypeDirect,
			WithExchangeDeclareDurable(true),
		),
		WithQueueDeclare(
			"",
			WithQueueDeclareDurable(false),
			WithQueueDeclareExclusive(true),
			WithQueueDeclareAutoDelete(false),
		),
		WithRoutingKey(routingKeyDirect),
	)
	if err != nil {
		fmt.Println("NewRabbitMQClient err:", err)
		return
	}

	err = cli.ConsumeDirect(handle)
	if err != nil {
		t.Errorf("ConsumeDirect() error = %v", err)
	}
}

func TestRabbitMQClient_ConsumeDirect01(t *testing.T) {
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

	cli, err := NewRabbitMQClient(
		WithUrl(myUrl),
		WithExchangeDeclare(
			exchangeNameDirect,
			ExchangeTypeDirect,
			WithExchangeDeclareDurable(true),
		),
		WithQueueDeclare(
			"",
			WithQueueDeclareDurable(false),
			WithQueueDeclareExclusive(true),
			WithQueueDeclareAutoDelete(false),
		),
		WithRoutingKey(routingKeyDirect),
	)
	if err != nil {
		fmt.Println("NewRabbitMQClient err:", err)
		return
	}

	err = cli.ConsumeDirect(handle)
	if err != nil {
		t.Errorf("ConsumeDirect() error = %v", err)
	}
}

package rabbitmq

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/streadway/amqp"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/17 15:36
 * @Desc:
 */

func TestRabbitMQClient_PublishFanout(t *testing.T) {
	type fields struct {
		conf     *rabbitMQConf
		conn     *amqp.Connection
		connChan *amqp.Channel
		queue    amqp.Queue
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
				body: []byte(fmt.Sprintf("TestRabbitMQClient_PublishFanout-%s", cast.ToString(time.Now().UnixNano()))),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, err := NewRabbitMQClient(
				WithUrl(myUrl),
				WithExchangeDeclare(
					exchangeNameFanout,
					ExchangeTypeFanout,
					WithExchangeDeclareDurable(false),
				),
				WithQueueDeclare(
					"",
					WithQueueDeclareDurable(false),
					WithQueueDeclareExclusive(true),
					WithQueueDeclareAutoDelete(true),
				),
			)
			if err != nil {
				fmt.Println("NewRabbitMQClient err:", err)
				return
			}
			fmt.Println("Boby:", string(tt.args.body))
			if err := cli.PublishFanout(tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("PublishFanout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRabbitMQClient_PublishFanout01(t *testing.T) {
	type fields struct {
		conf     *rabbitMQConf
		conn     *amqp.Connection
		connChan *amqp.Channel
		queue    amqp.Queue
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
				body: []byte(fmt.Sprintf("TestRabbitMQClient_PublishFanout-%s", cast.ToString(time.Now().UnixNano()))),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, err := NewRabbitMQClient(
				WithUrl(myUrl),
				WithExchangeDeclare(
					exchangeNameFanout,
					ExchangeTypeFanout,
					WithExchangeDeclareDurable(false),
				),
				WithQueueDeclare(
					"prince.queueName.Exclusive.tmp", // 声明队列为独占后，断开连接则会自动删除，所以声明队列名是没有意义的
					WithQueueDeclareDurable(true),
					WithQueueDeclareExclusive(true),
					WithQueueDeclareAutoDelete(false),
				),
			)
			if err != nil {
				fmt.Println("NewRabbitMQClient err:", err)
				return
			}
			fmt.Println("Boby:", string(tt.args.body))
			if err := cli.PublishFanout(tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("PublishFanout() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRabbitMQClient_ConsumeFanout(t *testing.T) {
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
			exchangeNameFanout,
			ExchangeTypeFanout,
			WithExchangeDeclareDurable(false),
		),
		WithQueueDeclare(
			"",
			WithQueueDeclareDurable(false),
			WithQueueDeclareExclusive(true),
			WithQueueDeclareAutoDelete(true),
		),
	)
	if err != nil {
		fmt.Println("NewRabbitMQClient err:", err)
		return
	}

	err = cli.ConsumeFanout(handle)
	if err != nil {
		t.Errorf("ConsumeFanout() error = %v", err)
	}
}

func TestRabbitMQClient_ConsumeFanout01(t *testing.T) {
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
			exchangeNameFanout,
			ExchangeTypeFanout,
			WithExchangeDeclareDurable(false),
		),
		WithQueueDeclare(
			"",
			WithQueueDeclareDurable(false),
			WithQueueDeclareExclusive(true),
			WithQueueDeclareAutoDelete(true),
		),
	)
	if err != nil {
		fmt.Println("NewRabbitMQClient err:", err)
		return
	}

	err = cli.ConsumeFanout(handle)
	if err != nil {
		t.Errorf("ConsumeFanout() error = %v", err)
	}
}

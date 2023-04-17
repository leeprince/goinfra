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
 * @Date:   2023/4/17 15:34
 * @Desc:
 */

func TestRabbitMQClient_PublishSimple(t *testing.T) {
	type args struct {
		body []byte
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			args: args{
				body: []byte(fmt.Sprintf("PublishSimple-%s", cast.ToString(time.Now().UnixNano()/1e3))),
			},
		},
		{
			args: args{
				body: []byte(fmt.Sprintf("PublishSimple-%s", cast.ToString(time.Now().UnixNano()/1e3))),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, err := NewRabbitMQClient(
				WithUrl(myUrl),
				WithQueueDeclare(queueNameSimple),
			)
			if err != nil {
				fmt.Println("NewRabbitMQClient err:", err)
				return
			}
			fmt.Println("Boby:", string(tt.args.body))
			if err := cli.PublishSimple(tt.args.body); err != nil {
				t.Errorf("PublishSimple() error = %v", err)
			}
			fmt.Println("test:PublishSimple:msg:", string(tt.args.body))
		})
	}
}

func TestRabbitMQClient_ConsumeSimple(t *testing.T) {
	handle := func(msg amqp.Delivery) {
		fmt.Printf("msg:%+v \n", msg)
		fmt.Printf("msg.Headers::%+v \n", msg.Headers)
		fmt.Printf("msg.Headers::%#v \n", msg.Headers)
		fmt.Println("string(msg.requestBody):", string(msg.Body))

		// --- 手动回复
		msg.Reject(false)
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
		WithQueueDeclare(queueNameSimple),
	)
	if err != nil {
		fmt.Println("NewRabbitMQClient err:", err)
		return
	}

	err = cli.ConsumeSimple(handle)
	if err != nil {
		t.Errorf("ConsumeSimple() error = %v", err)
	}
}

func TestRabbitMQClient_ConsumeSimple01(t *testing.T) {
	handle := func(msg amqp.Delivery) {
		fmt.Printf("msg:%+v \n", msg)
		fmt.Printf("msg.Headers::%+v \n", msg.Headers)
		fmt.Printf("msg.Headers::%#v \n", msg.Headers)
		fmt.Println("string(msg.requestBody):", string(msg.Body))

		// --- 手动回复
		msg.Reject(false)
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
		WithQueueDeclare(queueNameSimple),
	)
	if err != nil {
		fmt.Println("NewRabbitMQClient err:", err)
		return
	}

	err = cli.ConsumeSimple(handle)
	if err != nil {
		t.Errorf("ConsumeSimple() error = %v", err)
	}
}

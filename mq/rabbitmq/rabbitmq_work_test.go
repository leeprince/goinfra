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
 * @Date:   2023/4/17 15:35
 * @Desc:
 */

func TestRabbitMQClient_PublishWork(t *testing.T) {
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
				body: []byte(fmt.Sprintf("PublishWork-%s", cast.ToString(time.Now().UnixNano()))),
			},
		},
		{
			args: args{
				body: []byte(fmt.Sprintf("PublishWork-%s", cast.ToString(time.Now().UnixNano()))),
			},
		},
		{
			args: args{
				body: []byte(fmt.Sprintf("PublishWork-%s", cast.ToString(time.Now().UnixNano()))),
			},
		},
		{
			args: args{
				body: []byte(fmt.Sprintf("PublishWork-%s", cast.ToString(time.Now().UnixNano()))),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cli, err := NewRabbitMQClient(
				WithUrl(myUrl),
				WithQueueDeclare(
					queueNameWork,
					WithQueueDeclareDurable(true),
				),
			)
			if err != nil {
				fmt.Println("NewRabbitMQClient err:", err)
				return
			}
			fmt.Println("Boby:", string(tt.args.body))
			if err := cli.PublishWork(tt.args.body); (err != nil) != tt.wantErr {
				t.Errorf("PublishWork() error = %v, wantErr %v", err, tt.wantErr)
			}
			fmt.Println("test:PublishWork:msg:", string(tt.args.body))
		})
	}
}

func TestRabbitMQClient_ConsumeWork(t *testing.T) {
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
		WithQueueDeclare(
			queueNameWork,
			WithQueueDeclareDurable(true),
		),
		WithQos(2, 0, false),
	)
	if err != nil {
		fmt.Println("NewRabbitMQClient err:", err)
		return
	}

	err = cli.ConsumeWork(handle)
	if err != nil {
		t.Errorf("ConsumeWork() error = %v", err)
	}
}

func TestRabbitMQClient_ConsumeWork01(t *testing.T) {
	handle := func(msg amqp.Delivery) {
		fmt.Printf("msg:%+v \n", msg)
		fmt.Printf("msg.Headers::%+v \n", msg.Headers)
		fmt.Printf("msg.Headers::%#v \n", msg.Headers)
		fmt.Println("string(msg.requestBody):", string(msg.Body))

		time.Sleep(time.Second * 2)

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
		WithQueueDeclare(
			queueNameWork,
			WithQueueDeclareDurable(true),
		),
		WithQos(1, 0, false),
	)
	if err != nil {
		fmt.Println("NewRabbitMQClient err:", err)
		return
	}

	err = cli.ConsumeWork(
		handle,
	)
	if err != nil {
		t.Errorf("ConsumeWork() error = %v", err)
	}
}

func TestRabbitMQClient_ConsumeWork02(t *testing.T) {
	handle := func(msg amqp.Delivery) {
		fmt.Printf("msg:%+v \n", msg)
		fmt.Printf("msg.Headers::%+v \n", msg.Headers)
		fmt.Printf("msg.Headers::%#v \n", msg.Headers)
		fmt.Println("string(msg.requestBody):", string(msg.Body))

		time.Sleep(time.Second * 2)

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
		WithQueueDeclare(
			queueNameWork,
			WithQueueDeclareDurable(true),
		),
		WithQos(1, 0, false),
	)
	if err != nil {
		fmt.Println("NewRabbitMQClient err:", err)
		return
	}

	err = cli.ConsumeWork(
		handle,
		WithConsumeParamOptExclusive(true),
	)
	if err != nil {
		t.Errorf("ConsumeWork() error = %v", err)
	}
}

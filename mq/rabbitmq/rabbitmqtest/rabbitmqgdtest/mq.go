package rabbitmqgdtest

import (
	"fmt"
	"github.com/leeprince/goinfra/mq/rabbitmq"
	"github.com/leeprince/goinfra/utils/yamlutil"
)

var Mqs map[string]*rabbitmq.RabbitMQClient

const RABBIT_CONFKEY = "prince_test"

func init() {
	fmt.Println("初始化配置...")
	yamlutil.ParseFileToConfig("./conf.yaml", &Config)
	fmt.Println(Config)

	fmt.Println("连接RabbitMq")
	InitMQ()
}

func InitMQ() {
	fmt.Println("init rabbit")
	Mqs = make(map[string]*rabbitmq.RabbitMQClient)

	for name, rbConf := range Config.Rabbit {
		fmt.Println("queue name: ", name)
		fmt.Println("ExchangeName name: ", rbConf.ExchangeName)
		fmt.Println("ExchangeType: ", rbConf.ExchangeType)
		conn, err := rabbitmq.NewRabbitMQClient(
			rabbitmq.WithUrl(rbConf.Url),
			rabbitmq.WithExchangeDeclare(rbConf.ExchangeName, rbConf.ExchangeType),
			rabbitmq.WithQueueDeclare(rbConf.QueueName),
			rabbitmq.WithRoutingKey(rbConf.Key),
			rabbitmq.WithDel
		)
		if err != nil {
			panic("连接rabbitmq失败:" + err.Error())
		}

		Mqs[name] = conn
	}

	fmt.Println("Mqs : ", Mqs)
}

func GetRabbitConf(name string) *RabbitConf {
	return Config.Rabbit[name]
}

func GetMqConn(name string) *rabbitmq.RabbitMQClient {
	return Mqs[name]
}

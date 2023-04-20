package rabbitmqgdtest

var Config *Conf

type Conf struct {
	JaegerAgentUri string                 `yaml:"jaeger_agent_uri"`
	Rabbit         map[string]*RabbitConf `yaml:"rabbit"`
}
type RabbitConf struct { // rabbit配置
	Url            string `yaml:"url"`
	ExchangeName   string `yaml:"exchange_name"`
	ExchangeType   string `yaml:"queue_type"`
	QueueName      string `yaml:"queue_name"`
	DelayQueueName string `yaml:"delay_queue_name"`
	Key            string `yaml:"key"`
	DelayInterval  int    `yaml:"delay_interval"`
	Prefech        int    `yaml:"Prefech"`
}

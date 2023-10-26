package config

var Config *Conf

type Conf struct {
	Env     string `yaml:"env"`
	AppName string `yaml:"app_name"`

	Port           int    `yaml:"port"`
	LogPath        string `yaml:"log_path"`
	LogLevel       int    `yaml:"log_level"`
	SentryDsn      string `yaml:"sentry_dsn"`
	Mode           string `yaml:"mode"`
	JaegerAgentUri string `yaml:"jaeger_agent_uri"`

	Rabbit struct {
		ImportInoviceFileTask RabbitMQConfig
	}

	Redis struct {
		Cache Redis
	}

	MongoDB struct {
		Log MongoDB
	}

	SMTP map[string]*SMTPConf `yaml:"smtp"`

	GRPC *GRPCConf `yaml:"grpc"`

	// --- 动态配置
	Nacos   *Nacos      `yaml:"nacos"`
	RConfig *RemoteConf `yaml:"remote_conf"` // 远端配置. env为local,读取本地配置，其他环境都读取远程配置
}

// ---

type TencentCloudCOSConfig struct {
	SecretID  string
	SecretKey string
	Bucket    string
	AppID     string
	Region    string
}

type APIHost struct {
	Host string
}

type MongoDB struct {
	Uri string `yaml:"uri"`
}

// 远端配置
type RemoteConf struct {
	// 风控指标阈值
	RiskControl RiskControl

	// 告警
	Notice Notice
}

type RedisConfig struct {
	Uri      string `yaml:"uri" json:"uri"`
	Password string `yaml:"password" json:"password"`
	Db       int    `yaml:"db" json:"db"`
}

type RabbitMQConfig struct {
	RabbitMQURL    string
	VHost          string
	ExchangeName   string
	ExchangeType   string
	QueueName      string
	RouteKey       string
	DelayQueueName string // 实际也当作是延迟队列的 route_key
	DelayInterval  int32
	Prefech        int32 // 一次消费的消息总数
}

type RiskControl struct {
	IsOpen                             bool             // 是否检查风控
	SingleDayAllChannelTotalAmountYuan int64            // 单日出款总额（所有通道汇总金额），单位元
	SingleDayOneChannelTotalAmountYuan map[string]int64 // 日出款总额度（按通道汇总），单位元
	SingleDayTotalCount                map[string]int64 // 日出款总笔数（按通道汇总）
}

type SMTPConf struct { // SMTP 配置
	Host           string `yaml:"host"`
	Port           int32  `yaml:"port"`
	Sender         string `yaml:"sender"`
	Password       string `yaml:"password"`
	EncryptionType string `yaml:"encryption_type"`
}

type GRPCConf struct { // GRPC 配置
	PublicToolkitAddr string `yaml:"publictoolkitaddr"`
	PublicGcMailAddr  string `yaml:"publicgcmailaddr"`
}

type Notice struct {
	PhoneList             []string // 手机号
	ReceiverAddressesList []string // 接收方邮件地址
	CcAddressesList       []string // 抄送方邮件地址
}

type Redis struct {
	Host     string
	Port     int
	Password string
	DB       int
}

type Nacos struct {
	Ip          string
	Port        uint64
	Username    string
	Password    string
	NamespaceId string
	DataId      string
	Group       string
	ContextPath string

	LogDir     string // 日志存储路径
	RotateTime string // 日志轮转周期
	MaxAge     int64  // 最大日志文件数
	LogLevel   string // 日志默认级别
}

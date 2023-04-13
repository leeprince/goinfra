package rabbithttp

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/12 15:16
 * @Desc:
 */

type QueueMessageReq struct {
	Vhost    string `json:"vhost"`
	Name     string `json:"name"`
	Truncate string `json:"truncate"`
	Ackmode  string `json:"ackmode"`
	Encoding string `json:"encoding"`
	Count    string `json:"count"`
}
type QueueMessageResp struct {
	PayloadBytes    int           `json:"payload_bytes"`
	Redelivered     bool          `json:"redelivered"`
	Exchange        string        `json:"exchange"`
	RoutingKey      string        `json:"routing_key"`
	MessageCount    int           `json:"message_count"`
	Properties      []interface{} `json:"properties"`
	Payload         string        `json:"payload"`
	PayloadEncoding string        `json:"payload_encoding"`
}

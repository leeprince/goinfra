package rabbithttp

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/12 15:38
 * @Desc:
 */
type TestQueueMessage struct {
	No   int    `json:"no"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

func TestRabbitHttp_GetQueueMessage(t *testing.T) {
	r := &RabbitHttp{
		ctx:      context.Background(),
		logID:    "prince-logID-01",
		host:     "http://10.21.40.11:15672",
		username: "rabbit",
		password: "aTjHMj7opZ3d5Kw6",
		vhost:    "/",
	}

	queueName := "prince_test_queue"
	count := 3

	gotMessage, err := r.GetQueueMessage(queueName, count, ACKMODE_ACK_REQUEUE_TRUE)
	fmt.Println(err)
	fmt.Println(gotMessage)

	var messages []TestQueueMessage
	if err == nil {
		for _, s := range gotMessage {
			message := TestQueueMessage{}
			err = json.Unmarshal([]byte(s), &message)
			if err != nil {
				t.Errorf("json.Unmarshal([]byte(s), &message) err：%+v", err)
			}
			messages = append(messages, message)
		}

	}
	fmt.Println(messages)

}

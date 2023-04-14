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
	mqClient := NewRabbitHttp(context.Background(),
		"prince-logID-001",
		"http://10.21.40.11:15672",
		"/",
		"rabbit",
		"xxxxxxxxx",
	)

	queueName := "prince_test_queue"
	count := 3

	gotMessage, err := mqClient.GetQueueMessage(queueName, count, ACKMODE_ACK_REQUEUE_TRUE)
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

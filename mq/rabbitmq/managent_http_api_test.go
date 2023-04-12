package rabbitmq

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/12 11:16
 * @Desc:	[rabbitmq_management 中的HTTP API](https://rawcdn.githack.com/rabbitmq/rabbitmq-server/v3.11.13/deps/rabbitmq_management/priv/www/api/index.html)
 */

func TestRabbitMQHTTPAPIGetQueueMessage(t *testing.T) {
	url := "http://127.0.0.1:15672/api/queues/%2F/prince_test_queue/get"
	method := "POST"

	//payload := strings.NewReader(`{
	//   "vhost": "/",
	//   "name": "prince_test_queue",
	//   "truncate": "50000",
	//   "ackmode": "ack_requeue_false",
	//   "encoding": "auto",
	//   "count": "1"
	//}`)
	payload := bytes.NewBufferString(`{
	   "vhost": "/",
	   "name": "prince_test_queue",
	   "truncate": "50000",
	   "ackmode": "ack_requeue_false",
	   "encoding": "auto",
	   "count": "1"
	}`)

	client := http.DefaultClient
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.SetBasicAuth("xxxx", "xxx")

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "*/*")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

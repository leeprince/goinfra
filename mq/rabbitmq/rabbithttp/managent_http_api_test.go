package rabbithttp

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/leeprince/goinfra/http/httpcli"
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
	//url := "http://127.0.0.1:15672/api/queues/%2F/prince_test_queue/get"
	url := "http://10.21.40.11:15672/api/queues/%2F/prince_test_queue/get"
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

	// Basic Auth 认证方式的两种写法
	// req.SetBasicAuth底层代码还是设置到Header中(r.header.Set("Authorization", "Basic "+basicAuth(username, password)))
	// `Basic xxxxxxxxxxxxxxxxxxx=`中的`xxxxxxxxxxxxxxxxxxx=`是由`Basic Auth`认证时输入的"rabbit-username"和"rabbit-password"生成出来的。具体算法为：base64.StdEncoding.EncodeToString([]byte("rabbit-username+":"+"rabbitpassword"))
	req.SetBasicAuth("rabbit", "pppppxxxxxxxx")
	//req.header.Add("Authorization", "Basic xxxxxxxxxxxxxxxxxxx=")

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

func TestRabbitMQHTTPAPIGetQueueMessageV1(t *testing.T) {
	url := "http://10.21.40.11:15672/api/queues/%2F/prince_test_queue/get"

	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte("rabbit"+":"+"xxxxxxxxxxxxxx"))

	body := []byte(`{
	   "vhost": "/",
	   "name": "prince_test_queue",
	   "truncate": "50000",
	   "ackmode": "ack_requeue_true",
	   "encoding": "auto",
	   "count": "1"
	}`)

	httpClient := httpcli.NewHttpClient().
		WithURL(url).
		WithIsHttpTrace(false).
		WithLogID("prince-TestRabbitMQHTTPAPIGetQueueMessage-01").
		WithMethod(http.MethodPost).
		WithHeader(header).
		WithBody(body)

	respBobyByte, _, err := httpClient.Do()
	if err != nil {
		t.Errorf("err != nil。 %+v", err)
	}
	fmt.Println("respBobyByte:", string(respBobyByte))
}

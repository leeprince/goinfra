package rabbithttp

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/leeprince/goinfra/http/httpcli"
	"github.com/leeprince/goinfra/plog"
	"net/http"
	"strconv"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/12 15:11
 * @Desc:
 */

type Ackmode string

const (
	ACKMODE_ACK_REQUEUE_TRUE     Ackmode = "ack_requeue_true"
	ACKMODE_ACK_REQUEUE_FALSE    Ackmode = "ack_requeue_false"
	ACKMODE_REJECT_REQUEUE_TRUE  Ackmode = "reject_requeue_true"
	ACKMODE_REJECT_REQUEUE_FALSE Ackmode = "reject_requeue_false"
)

type RabbitHttp struct {
	ctx        context.Context
	logID      string
	host       string
	username   string
	password   string
	vhost      string
	httpClient *httpcli.HttpClient
}

// host=config.Config.RabbitHttpUri
func NewRabbitHttp(ctx context.Context, logID, host, vhost, username, password string) *RabbitHttp {
	return &RabbitHttp{
		ctx:        ctx,
		logID:      logID,
		host:       host,
		username:   username,
		password:   password,
		vhost:      vhost,
		httpClient: httpcli.NewHttpClient().WithLogID(logID).WithIsHttpTrace(false),
	}
}

func (r *RabbitHttp) header() map[string]string {
	header := make(map[string]string)
	header["Content-Type"] = "application/json"
	header["Authorization"] = "Basic " + base64.StdEncoding.EncodeToString([]byte(r.username+":"+r.password))

	return header
}

func (r *RabbitHttp) GetQueueMessage(queueName string, count int, ackmode Ackmode) (messages []string, err error) {
	url := r.host + "/api/queues/%2F/prince_test_queue/get"

	header := r.header()

	body := QueueMessageReq{
		Name:     queueName,
		Count:    strconv.Itoa(count),
		Vhost:    r.vhost,
		Truncate: "50000",
		Ackmode:  string(ackmode),
		Encoding: "auto",
	}

	httpClient := r.httpClient.
		WithURL(url).
		WithMethod(http.MethodPost).
		WithHeader(header).
		WithBody(body)

	respBobyByte, _, err := httpClient.Do()
	if err != nil {
		plog.WithError(err).Error(r.logID, "http do err")
		return nil, err
	}

	var resp []QueueMessageResp
	err = json.Unmarshal(respBobyByte, &resp)
	if err != nil {
		plog.WithError(err).Error(r.logID, "json.Unmarshal err")
		return nil, errors.New("数据解析错误")
	}

	messages = []string{}
	for _, i2 := range resp {
		if i2.PayloadEncoding != "string" {
			plog.WithError(err).Error(r.logID, "i2.PayloadEncoding != string err")
			return nil, errors.New("负载的数据编码格式错误")
		}
		messages = append(messages, i2.Payload)
	}

	return messages, nil
}

package cos

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/3/24 18:54
 * @Desc:
 */

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"sync"
)

type TxCos struct {
	secretID  string
	secretKey string
	bucket    string
	region    string
	appId     string
}

var o sync.Once
var cosClient *cos.Client

func NewTxCos(secretID string, secretKey string, appId string, bucket string, region string) *TxCos {
	return &TxCos{secretID: secretID,
		secretKey: secretKey,
		appId:     appId,
		bucket:    bucket,
		region:    region,
	}
}

func (t *TxCos) GetCosClient() *cos.Client {
	o.Do(func() {
		u, _ := url.Parse(t.getBucketUrl())
		b := &cos.BaseURL{BucketURL: u}

		cosClient = cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  t.secretID,
				SecretKey: t.secretKey,
			},
		})
	})
	return cosClient
}

func (t *TxCos) getBucketUrl() string {
	return "https://" + t.bucket + "-" + t.appId + ".cos." + t.region + ".myqcloud.com"
}

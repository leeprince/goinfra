package cos

import (
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"sync"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/1/17 11:45
 * @Desc:
 */

var o sync.Once
var client *cos.Client

type CosClient struct {
	secretID, secretKey, appId, bucket, region string
}

func NewCosClient(secretID, secretKey, appId, bucket, region string) *CosClient {
	o.Do(func() {
		u, _ := url.Parse(getBucketUrl(appId, bucket, region))
		b := &cos.BaseURL{BucketURL: u}

		client = cos.NewClient(b, &http.Client{
			Transport: &cos.AuthorizationTransport{
				SecretID:  secretID,
				SecretKey: secretKey,
			},
		})
	})

	return &CosClient{
		secretID:  secretID,
		secretKey: secretKey,
		appId:     appId,
		bucket:    bucket,
		region:    region,
	}
}

func getBucketUrl(appId, bucket, region string) string {
	return "https://" + bucket + "-" + appId + ".cos." + region + ".myqcloud.com"
}

package elasticsearch

import (
    elasticsearch8 "github.com/elastic/go-elasticsearch/v8"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/7/10 下午4:16
 * @Desc:   初始化
 */

// 全局变量
var elasticClients map[string]elasticClient

type elasticClient struct {
    client *elasticsearch8.Client
}

type ElasticConfs map[string]*ElasticConf
type ElasticConf struct {
    Url      string `yaml:"url" json:"url"`
    Username string `yaml:"username" json:"username"`
    Password string `yaml:"password" json:"password"`
    CACert   string `yaml:"ca_cert" json:"ca_cert"`
}

func InitElasticClient(confs ElasticConfs) error {
    elasticClients = make(map[string]elasticClient)
    
    for name, conf := range confs {
        esConf := elasticsearch8.Config{
            Addresses:               []string{conf.Url},
            Username:                conf.Username,
            Password:                conf.Password,
            CloudID:                 "",
            APIKey:                  "",
            ServiceToken:            "",
            CertificateFingerprint:  "",
            Header:                  nil,
            CACert:                  nil,
            RetryOnStatus:           nil,
            DisableRetry:            false,
            MaxRetries:              0,
            RetryOnError:            nil,
            CompressRequestBody:     false,
            DiscoverNodesOnStart:    false,
            DiscoverNodesInterval:   0,
            EnableMetrics:           false,
            EnableDebugLogger:       false,
            EnableCompatibilityMode: false,
            DisableMetaHeader:       false,
            RetryBackoff:            nil,
            Transport:               nil,
            Logger:                  nil,
            Selector:                nil,
            ConnectionPoolFunc:      nil,
        }
        
        if conf.CACert != "" {
            esConf.CACert = []byte(conf.CACert)
        }
        
        client, err := elasticsearch8.NewClient(esConf)
        if err != nil {
            return err
        }
        
        elasticClients[name] = elasticClient{
            client,
        }
    }
    return nil
}

func GetElasticClient(name string) *elasticsearch8.Client {
    elasticClient, ok := elasticClients[name]
    if !ok {
        return nil
    }
    return elasticClient.client
}

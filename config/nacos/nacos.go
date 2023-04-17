package nacos

import (
	"errors"
	"github.com/leeprince/goinfra/plog"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/4 下午3:55
 * @Desc:
 */

type NacosClient struct {
	configClient      config_client.IConfigClient
	nacosClientParams *NacosClientParams
}

func MustNewNacosClient(namespaceId, group, dataID string, opts ...NacosClienParamsOpt) (cli *NacosClient) {
	client, err := NewNacosClient(namespaceId, group, dataID, opts...)
	if err != nil {
		panic(err)
	}
	return client
}

func NewNacosClient(namespaceId, group, dataID string, opts ...NacosClienParamsOpt) (cli *NacosClient, err error) {
	params := NacosClientParams{
		namespaceID: namespaceId,
		dataID:      dataID,
		group:       group,
		timeoutMs:   5000,
		logDir:      "./logs",
		cacheDir:    "./cache",
		logLevel:    "info",
		ipAddr:      "127.0.0.1",
		port:        8848,
		scheme:      "",
		contextPath: "",
	}
	for _, opt := range opts {
		opt(&params)
	}

	err = checkNacosClienParams(params)
	if err != nil {
		return
	}

	// 初始化Nacos客户端
	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(params.namespaceID), // When namespace is public, fill in the blank string here.
		constant.WithNotLoadCacheAtStart(true),
		constant.WithTimeoutMs(params.timeoutMs),
		constant.WithLogDir(params.logDir),
		constant.WithCacheDir(params.cacheDir),
		constant.WithLogLevel(params.logLevel),
	)
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      params.ipAddr,
			Port:        params.port,
			ContextPath: params.contextPath,
			Scheme:      params.scheme,
		},
	}
	configClient, errn := clients.NewConfigClient(
		vo.NacosClientParam{
			ClientConfig:  &clientConfig,
			ServerConfigs: serverConfigs,
		},
	)
	if errn != nil {
		err = errn
		return
	}

	cli = &NacosClient{
		configClient:      configClient,
		nacosClientParams: &params,
	}
	return
}

type dynamicConfigHandle func(conf []byte)

func (c *NacosClient) ListenConfig(handle dynamicConfigHandle) (err error) {
	// 先立即同步获取配置
	conf, errn := c.configClient.GetConfig(vo.ConfigParam{
		DataId: c.nacosClientParams.dataID,
		Group:  c.nacosClientParams.group,
	})
	if errn != nil {
		return errn
	}
	handle([]byte(conf))

	// 继续监听配置变更：监听配置变更，动态更新配置
	err = c.configClient.ListenConfig(vo.ConfigParam{
		DataId: c.nacosClientParams.dataID,
		Group:  c.nacosClientParams.group,
		OnChange: func(namespace, group, dataId, data string) {
			plog.Info(">>> 监听到配置变更，动态更新配置")
			plog.Info("group:" + group + ", dataID:" + dataId + ", data:" + data)
			handle([]byte(data))
			plog.Info("<<< 监听到配置变更，动态更新配置")
		},
	})
	return err
}

func checkNacosClienParams(nacosClienParams NacosClientParams) error {
	if nacosClienParams.namespaceID == "" {
		return errors.New("namespaceID must not empty")
	}
	if nacosClienParams.dataID == "" {
		return errors.New("dataID must not empty")
	}
	if nacosClienParams.group == "" {
		return errors.New("group must not empty")
	}
	return nil
}

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
	nacosClienParams := NacosClientParams{
		namespaceID: namespaceId,
		group:       group,
		dataID:      dataID,
		timeoutMs:   5000,
		logDir:      "./log",
		cacheDir:    "./cache",
		logLevel:    "error",
		ipAddr:      "127.0.0.1",
		port:        8848,
	}
	for _, opt := range opts {
		opt(&nacosClienParams)
	}

	err = checkNacosClienParams(nacosClienParams)
	if err != nil {
		return
	}

	clientConfig := *constant.NewClientConfig(
		constant.WithNamespaceId(nacosClienParams.namespaceID), // When namespace is public, fill in the blank string here.
		constant.WithNotLoadCacheAtStart(true),
		constant.WithTimeoutMs(nacosClienParams.timeoutMs),
		constant.WithLogDir(nacosClienParams.logDir),
		constant.WithCacheDir(nacosClienParams.cacheDir),
		constant.WithLogLevel(nacosClienParams.logLevel),
	)
	serverConfigs := []constant.ServerConfig{
		{
			IpAddr:      nacosClienParams.ipAddr,
			Port:        nacosClienParams.port,
			ContextPath: "/nacos",
			Scheme:      "http",
		},
	}
	// 初始化Nacos客户端
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
		nacosClientParams: &nacosClienParams,
	}
	return
}

type dynamicConfigHandle func(conf []byte)

func (c *NacosClient) ListenConfig(handle dynamicConfigHandle) (err error) {
	// 先同步获取配置
	conf, errn := c.configClient.GetConfig(vo.ConfigParam{
		DataId: c.nacosClientParams.dataID,
		Group:  c.nacosClientParams.group,
	})
	if errn != nil {
		return errn
	}
	handle([]byte(conf))

	// 监听配置变更：监听配置变更，动态更新配置
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
	if nacosClienParams.group == "" {
		return errors.New("group must not empty")
	}
	if nacosClienParams.ipAddr == "" {
		return errors.New("ipAddr must not empty")
	}
	if nacosClienParams.port <= 0 {
		return errors.New("port must not empty")
	}
	return nil
}

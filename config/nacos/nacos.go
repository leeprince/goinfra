package nacos

import (
    "errors"
    "fmt"
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
    configClient     config_client.IConfigClient
    nacosClienParams *NacosClienParams
}

func NewNacosClient(opts ...NacosClienParamsOpt) (cli *NacosClient, err error) {
    nacosClienParams := NacosClienParams{
        namespaceId: "",
        dataId:      "",
        group:       "",
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
        constant.WithNamespaceId(nacosClienParams.namespaceId), // When namespace is public, fill in the blank string here.
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
    configClient, err := clients.NewConfigClient(
        vo.NacosClientParam{
            ClientConfig:  &clientConfig,
            ServerConfigs: serverConfigs,
        },
    )
    if err != nil {
        return
    }
    
    cli = &NacosClient{
        configClient:     configClient,
        nacosClienParams: &nacosClienParams,
    }
    return
}

type dynamicConfigHandle func(conf []byte)

func (c *NacosClient) ListenConfig(handle dynamicConfigHandle) (err error) {
    // 先获取配置
    conf, err := c.configClient.GetConfig(vo.ConfigParam{
        DataId: c.nacosClienParams.dataId,
        Group:  c.nacosClienParams.group,
    })
    if err != nil {
        return
    }
    handle([]byte(conf))
    
    // 动态获取配置：监听配置变更，更新配置
    err = c.configClient.ListenConfig(vo.ConfigParam{
        DataId: c.nacosClienParams.dataId,
        Group:  c.nacosClienParams.group,
        OnChange: func(namespace, group, dataId, data string) {
            fmt.Println(">>> 监听到配置变更，动态更新配置")
            fmt.Println("group:" + group + ", dataId:" + dataId + ", data:" + data)
            handle([]byte(data))
            fmt.Println("<<< 监听到配置变更，动态更新配置")
        },
    })
    return
}

func checkNacosClienParams(nacosClienParams NacosClienParams) error {
    if nacosClienParams.namespaceId == "" {
        return errors.New("namespaceId must not empty")
    }
    if nacosClienParams.dataId == "" {
        return errors.New("dataId must not empty")
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

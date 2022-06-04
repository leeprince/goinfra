package nacos

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/4 下午5:09
 * @Desc:
 */

type NacosClienParams struct {
    // --- vo.NacosClientParam.ClientConfig
    namespaceId string
    dataId      string
    group       string
    timeoutMs   uint64
    logDir      string
    cacheDir    string
    logLevel    string // debug,info,warn,error, default value is info
    
    // --- vo.NacosClientParam.ServerConfigs
    ipAddr string
    port   uint64
}

type NacosClienParamsOpt func(params *NacosClienParams)

func WithNamespaceId(v string) NacosClienParamsOpt {
    return func(params *NacosClienParams) {
        params.namespaceId = v
    }
}

func WithDataId(v string) NacosClienParamsOpt {
    return func(params *NacosClienParams) {
        params.dataId = v
    }
}

func WithGoup(v string) NacosClienParamsOpt {
    return func(params *NacosClienParams) {
        params.group = v
    }
}

func WithTimeoutMs(v uint64) NacosClienParamsOpt {
    return func(params *NacosClienParams) {
        params.timeoutMs = v
    }
}

func WithLogDir(v string) NacosClienParamsOpt {
    return func(params *NacosClienParams) {
        params.group = v
    }
}

func WithCacheDir(v string) NacosClienParamsOpt {
    return func(params *NacosClienParams) {
        params.cacheDir = v
    }
}

func WithLogLevel(v string) NacosClienParamsOpt {
    return func(params *NacosClienParams) {
        params.logLevel = v
    }
}

func WithIpAddr(v string) NacosClienParamsOpt {
    return func(params *NacosClienParams) {
        params.ipAddr = v
    }
}

func WithPort(v uint64) NacosClienParamsOpt {
    return func(params *NacosClienParams) {
        params.port = v
    }
}

package nacos

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/4 下午5:09
 * @Desc:
 */

type NacosClientParams struct {
	// --- vo.NacosClientParam.ClientConfig
	namespaceID string
	dataID      string
	group       string
	timeoutMs   uint64
	logDir      string
	cacheDir    string
	logLevel    string // debug,info,warn,error, default value is info

	// --- vo.NacosClientParam.ServerConfigs
	ipAddr string
	port   uint64
}

type NacosClienParamsOpt func(params *NacosClientParams)

func WithTimeoutMs(v uint64) NacosClienParamsOpt {
	return func(params *NacosClientParams) {
		params.timeoutMs = v
	}
}

func WithLogDir(v string) NacosClienParamsOpt {
	return func(params *NacosClientParams) {
		params.group = v
	}
}

func WithCacheDir(v string) NacosClienParamsOpt {
	return func(params *NacosClientParams) {
		params.cacheDir = v
	}
}

func WithLogLevel(v string) NacosClienParamsOpt {
	return func(params *NacosClientParams) {
		params.logLevel = v
	}
}

func WithIpAddr(v string) NacosClienParamsOpt {
	return func(params *NacosClientParams) {
		params.ipAddr = v
	}
}

func WithPort(v uint64) NacosClienParamsOpt {
	return func(params *NacosClientParams) {
		params.port = v
	}
}

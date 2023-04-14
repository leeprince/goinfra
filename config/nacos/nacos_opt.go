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
	ipAddr      string // 127.0.0.1
	port        uint64 // 8848
	scheme      string // constant.DEFAULT_SERVER_SCHEME
	contextPath string // constant.DEFAULT_CONTEXT_PATH
}

type NacosClienParamsOpt func(params *NacosClientParams)

func WithTimeoutMs(v uint64) NacosClienParamsOpt {
	return func(params *NacosClientParams) {
		params.timeoutMs = v
	}
}

func WithLogDir(v string) NacosClienParamsOpt {
	return func(params *NacosClientParams) {
		params.logDir = v
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

func WithScheme(v string) NacosClienParamsOpt {
	return func(params *NacosClientParams) {
		params.scheme = v
	}
}

func WithContextPath(v string) NacosClienParamsOpt {
	return func(params *NacosClientParams) {
		params.contextPath = v
	}
}

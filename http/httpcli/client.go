package httpcli

import (
	"bytes"
	"context"
	"crypto/tls"
	jsoniter "github.com/json-iterator/go"
	"github.com/leeprince/goinfra/perror"
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/trace/opentracing/jaegerclient"
	"github.com/leeprince/goinfra/utils"
	"github.com/leeprince/goinfra/utils/pstring"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"time"
)

//1.基于httpclient_v2版本实现
//2.变更了span上报字段名为小写
//3.修改了类名(HttpCtlService->HttpClient)
//4.去除了无用的New部分代码片段
//5.增加了调用方法（如WithContext、WithLogID、DoRequest）
//6.优化了报错信息
//7.去除clone逻辑
//8.添加是否打印日志
//9.添加是否添加http链路追踪
//10.添加是否使用代理
//11.添加重定向检查

const HttpDefaultTimeout = time.Second * 8

var defaultHeader = map[string]string{"Content-Type": "application/json"}

type HttpClient struct {
	ctx           context.Context
	logID         string
	timeout       time.Duration
	method        string // http包中的字符串。如：http.MethodGet
	url           string
	header        map[string]string
	requestBody   interface{}
	skipVerify    bool // 是否跳过安全校验
	notLogging    bool // 不打印日志标记
	isHttpTrace   bool // 是否注入调用链跟踪标记到请求头中 能开启的必要条件是：s.isHttpTrace && jaeger_client.SpanFromContext(ctx) != nil，也就是需要初始化jaeger获得span的context后才能正常开启链路追踪
	proxyURL      *url.URL
	checkRedirect func(req *http.Request, via []*http.Request) error // 检查重定向的方法
}

// TODO: 优化为传递必传参数 prince@todo 2023/4/12 16:56
func NewHttpClient() *HttpClient {
	hc := &HttpClient{
		timeout:       HttpDefaultTimeout,
		method:        "",
		url:           "",
		logID:         utils.UniqID(),
		header:        nil,
		requestBody:   nil,
		ctx:           context.Background(),
		skipVerify:    false,
		notLogging:    false,
		isHttpTrace:   true,
		proxyURL:      nil,
		checkRedirect: nil,
	}

	return hc
}

func (s *HttpClient) WithURL(url string) *HttpClient {
	s.url = url
	return s
}

func (s *HttpClient) WithLogID(logID string) *HttpClient {
	s.logID = logID
	return s
}

func (s *HttpClient) WithContext(ctx context.Context) *HttpClient {
	s.ctx = ctx
	return s
}

func (s *HttpClient) WithMethod(method string) *HttpClient {
	s.method = method
	return s
}

func (s *HttpClient) WithTimeout(timeout time.Duration) *HttpClient {
	s.timeout = timeout
	return s
}

func (s *HttpClient) WithHeader(header map[string]string) *HttpClient {
	s.header = header
	return s
}

func (s *HttpClient) WithBody(body interface{}) *HttpClient {
	s.requestBody = body
	return s
}

//不打印日志
func (s *HttpClient) WithNotLogging(notLogging bool) *HttpClient {
	s.notLogging = notLogging
	return s
}

func (s *HttpClient) WithSkipVerify(skipVerify bool) *HttpClient {
	s.skipVerify = skipVerify
	return s
}

// 链路追踪。这是WithIsHttpTrace后可以省略 WithContext()
func (s *HttpClient) WithIsHttpTrace(isHttpTrace bool, ctx context.Context) *HttpClient {
	s.isHttpTrace = isHttpTrace
	s.ctx = ctx
	if isHttpTrace {
		if jaegerclient.SpanFromContext(ctx) == nil {
			s.isHttpTrace = false
		}
	}
	return s
}

func (s *HttpClient) WithProxyURL(proxyURL *url.URL) *HttpClient {
	s.proxyURL = proxyURL
	return s
}

func (s *HttpClient) WithCheckRedirect(checkRedirect func(req *http.Request, via []*http.Request) error) *HttpClient {
	s.checkRedirect = checkRedirect
	return s
}

func (s *HttpClient) Do() ([]byte, *http.Response, error) {
	return s.do()
}

func (s *HttpClient) DoRequest(ctx context.Context, logID string, url string, method string, headers map[string]string, body interface{}) ([]byte, *http.Response, error) {
	return s.WithContext(ctx).WithLogID(logID).WithURL(url).WithMethod(method).WithHeader(headers).WithBody(body).Do()
}

func (s *HttpClient) do() ([]byte, *http.Response, error) {
	if s.url == "" {
		return nil, nil, errors.New("无效的URL")
	}

	if s.header == nil {
		s.header = defaultHeader
	}

	// 链路追踪
	if s.isHttpTrace {
		//// 新增埋点
		//routerSuffix := "-"
		//pos := strings.LastIndex(s.url, "/")
		//if pos+1 < len(s.url) {
		//	routerSuffix = s.url[pos+1:] //以接口路由的最后一级节点名作为span的操作名称
		//}
		//operationName := "http-" + routerSuffix
		//pos = strings.Index(operationName, "?")
		//if pos != -1 {
		//	operationName = operationName[:pos]
		//}
		//s.ctx = jaeger_client.StartSpan(s.ctx, operationName)
		//defer jaeger_client.Finish(s.ctx)
	}

	var (
		reqBytes []byte
		ok       bool
		err      error
	)

	if s.requestBody != nil {
		if reqBytes, ok = s.requestBody.([]byte); !ok {
			reqBytes, err = jsoniter.Marshal(s.requestBody)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	method := s.method
	fields := logrus.Fields{}
	fields["http.req.url"] = s.url
	fields["http.req.method"] = method

	fields["http.req.body"] = pstring.Bytes2String(reqBytes)

	req, err := http.NewRequest(method, s.url, bytes.NewReader(reqBytes))
	if err != nil {
		return nil, nil, err
	}

	hasContentType := false
	for k, v := range s.header {
		req.Header.Set(k, v)
		if k == "Content-Type" {
			hasContentType = true
		}
	}

	if method == http.MethodPost && !hasContentType {
		req.Header.Set("Content-Type", "application/json")
	}

	req = req.WithContext(s.ctx)

	// 链路追踪。
	fields["http.req.isHttpTrace"] = s.isHttpTrace
	if s.isHttpTrace {
		jaegerclient.LogKV(s.ctx, "url", s.url)
		jaegerclient.LogKV(s.ctx, "method", method)
		jaegerclient.LogKV(s.ctx, "log_id", s.logID)
		jaegerclient.LogKV(s.ctx, "body", req.Body)
		// 注入 tracer 信息到 req.header 中。`req.header`是指针，所以`req.header`的修改会影响到`fields["http.req.header"] = req.header`
		err = jaegerclient.InjectTraceHTTPClient(s.ctx, s.url, method, req.Header)
		if err != nil {
			// 链路追踪错误，不中断，仅记录日志
			plog.LogID(s.logID).WithError(err).Error("HttpClient.do InjectTraceHTTPClient")
		}
	}
	fields["http.req.header"] = req.Header

	fields["http.req.isProxy"] = false
	if s.proxyURL != nil {
		fields["http.req.isProxy"] = true
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(s.proxyURL),
		DialContext: (&net.Dialer{
			Timeout: s.timeout,
		}).DialContext,
		TLSHandshakeTimeout: 0,
		IdleConnTimeout:     0,
		ProxyConnectHeader:  nil,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.skipVerify,
		},
	}

	client := &http.Client{
		Transport:     transport,
		CheckRedirect: s.checkRedirect,
		Timeout:       0,
	}

	if !s.notLogging {
		plog.LogID(s.logID).WithFields(fields).Info("发起Http请求")
	}

	resp, err := client.Do(req)
	if err != nil {
		if !s.notLogging {
			plog.LogID(s.logID).WithError(err).
				WithField("url", s.url).
				Error("发起Http请求失败")
		}
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		if !s.notLogging {
			plog.LogID(s.logID).WithError(err).
				WithField("url", s.url).
				Error("读取http响应结果失败")
		}
		return nil, nil, perror.ReplaceIPErr(err)
	}

	if resp.StatusCode != http.StatusOK {
		if !s.notLogging {
			plog.LogID(s.logID).WithError(err).
				WithFields(fields).
				WithField("Resp.StatusCode", resp.StatusCode).
				WithField("Resp.Status", resp.Status).
				Error("Http 响应码异常")
		}
		return nil, nil, errors.Errorf("上游服务报错,http status code:%d", resp.StatusCode)
	}

	if !s.notLogging {
		respBodyLog := body
		if len(body) > 1024 {
			respBodyLog = body[:1023]
		}
		if !s.notLogging {
			plog.LogID(s.logID).WithField("http.req.url", s.url).
				WithField("http.resp.status_code", resp.StatusCode).
				WithField("http.resp.body", pstring.Bytes2String(respBodyLog)).
				Info("接收http响应")
		}
	}

	return body, resp, nil
}

func Do(req *http.Request) (respByte []byte, resp *http.Response, err error) {
	resp, err = http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	respByte, err = ResponseToBytes(resp)
	return
}

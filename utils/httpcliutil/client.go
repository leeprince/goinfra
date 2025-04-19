package httpcliutil

import (
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"net"
	"net/http"
	"net/url"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/leeprince/goinfra/consts"
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/utils/contextutil"
	"github.com/leeprince/goinfra/utils/idutil"
	"github.com/leeprince/goinfra/utils/stringutil"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// 1.基于httpclient_v2版本实现
// 2.变更了span上报字段名为小写
// 3.修改了类名(HttpCtlService->HttpClient)
// 4.去除了无用的New部分代码片段
// 5.增加了调用方法（如WithContext、WithLogID、DoRequest）
// 6.优化了报错信息
// 7.去除clone逻辑
// 8.添加是否打印日志
// 9.添加是否添加http链路追踪
// 10.添加是否使用代理
// 11.添加重定向检查

const (
	HttpDefaultTimeout = time.Second * 3
)

var defaultHeader = map[string]string{"Content-Type": "application/json"}

type Logger interface {
	Info(...any)
	Error(...any)
}

type HttpClient struct {
	url           string
	method        string          // http包中的字符串。如：http.MethodGet
	ctx           context.Context // 默认：context.Background()
	logID         string
	logger        Logger
	timeout       time.Duration
	header        map[string]string
	requestBody   interface{}
	resp          interface{} // 响应。不能为nil
	skipVerify    bool        // 是否跳过安全校验
	notLogging    bool        // 不打印日志标记
	proxyURL      *url.URL
	checkRedirect func(req *http.Request, via []*http.Request) error // 检查重定向的方法
}

// NewHttpClient
//
//	@Description: http 客户端
//	@return *HttpClient
func NewHttpClient() *HttpClient {
	hc := &HttpClient{
		url:           "",
		timeout:       HttpDefaultTimeout,
		method:        http.MethodGet,
		logID:         "",
		header:        nil,
		requestBody:   nil,
		ctx:           nil,
		skipVerify:    false,
		notLogging:    false,
		proxyURL:      nil,
		checkRedirect: nil,
	}

	return hc
}

// WithURL 请求的主机地址+路由前缀(http://xxx.xx.xx/api-route-prefix)
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

func (s *HttpClient) WithResponse(resp interface{}) *HttpClient {
	s.resp = resp
	return s
}

func (s *HttpClient) WithLogger(logger Logger) *HttpClient {
	s.logger = logger
	return s
}

// 不打印日志
func (s *HttpClient) WithNotLogging(notLogging bool) *HttpClient {
	s.notLogging = notLogging
	return s
}

func (s *HttpClient) WithSkipVerify(skipVerify bool) *HttpClient {
	s.skipVerify = skipVerify
	return s
}

// 链路追踪。这是WithIsHttpTrace后可以省略 WithContext()
func (s *HttpClient) WithIsHttpTrace(isHttpTrace bool, ctx ...context.Context) *HttpClient {
	if len(ctx) >= 1 {
		s.ctx = ctx[0]
	}
	if !isHttpTrace {
		return s
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
	bodyBytes, httpResp, err := s.do()
	if err != nil {
		return nil, nil, err
	}

	if s.resp == nil {
		return bodyBytes, httpResp, nil
	}

	if len(bodyBytes) <= 0 {
		return nil, nil, errors.New("响应为空")
	}
	err = jsoniter.Unmarshal(bodyBytes, s.resp)
	if err != nil {
		return bodyBytes, httpResp, err
	}

	return bodyBytes, httpResp, nil
}

func (s *HttpClient) DoRequest(ctx context.Context, logID string, url string, method string, headers map[string]string, body interface{}) ([]byte, *http.Response, error) {
	return s.WithContext(ctx).WithLogID(logID).WithURL(url).WithMethod(method).WithHeader(headers).WithBody(body).Do()
}

func (s *HttpClient) do() ([]byte, *http.Response, error) {
	if s.url == "" {
		return nil, nil, errors.New("无效的URL")
	}

	if s.ctx != nil {
		if s.logID == "" {
			s.logID = contextutil.LogIdByContext(&s.ctx)
		}
	} else {
		s.ctx = context.Background()
		s.logID = idutil.UniqIDV3()
	}

	if s.header == nil {
		s.header = defaultHeader
	}
	if _, ok := s.header[consts.HeaderXLogID]; !ok {
		s.header[consts.HeaderXLogID] = s.logID
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

	fields["http.req.body"] = stringutil.Bytes2String(reqBytes)

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
		Timeout:       s.timeout,
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

	body, err := io.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		if !s.notLogging {
			plog.LogID(s.logID).WithError(err).
				WithField("url", s.url).
				Error("读取http响应结果失败")
		}
		return nil, nil, err
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
		plog.LogID(s.logID).WithField("http.req.url", s.url).
			WithField("http.resp.status_code", resp.StatusCode).
			WithField("http.resp.body", stringutil.Bytes2String(respBodyLog)).
			Info("接收http响应")
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

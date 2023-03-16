package httpcli

import (
	"bytes"
	"context"
	"crypto/tls"
	"github.com/leeprince/goinfra/plog"
	"github.com/leeprince/goinfra/utils"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"strings"
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
	Timeout       time.Duration
	Method        string // http包中的字符串。如：http.MethodGet
	URL           string
	LogID         string
	Header        map[string]string
	Body          interface{}
	Ctx           context.Context
	SkipVerify    bool // 是否跳过安全校验
	NotLogging    bool // 不打印日志标记
	isHttpTrace   bool // 是否注入调用链跟踪标记到请求头中
	proxyURL      *url.URL
	checkRedirect func(req *http.Request, via []*http.Request) error // 检查重定向的方法
}

func NewHttpClient() *HttpClient {
	hc := &HttpClient{
		Timeout:       HttpDefaultTimeout,
		Method:        "",
		URL:           "",
		LogID:         utils.UniqID(),
		Header:        nil,
		Body:          nil,
		Ctx:           context.Background(),
		SkipVerify:    false,
		NotLogging:    false,
		isHttpTrace:   true,
		proxyURL:      nil,
		checkRedirect: nil,
	}

	return hc
}

func (s *HttpClient) WithURL(url string) *HttpClient {
	s.URL = url
	return s
}

func (s *HttpClient) WithLogID(logID string) *HttpClient {
	s.LogID = logID
	return s
}

func (s *HttpClient) WithContext(ctx context.Context) *HttpClient {
	s.Ctx = ctx
	return s
}

func (s *HttpClient) WithMethod(method string) *HttpClient {
	s.Method = method
	return s
}

func (s *HttpClient) WithTimeout(timeout time.Duration) *HttpClient {
	s.Timeout = timeout
	return s
}

func (s *HttpClient) WithHeader(header map[string]string) *HttpClient {
	s.Header = header
	return s
}

func (s *HttpClient) WithBody(body interface{}) *HttpClient {
	s.Body = body
	return s
}

//不打印日志
func (s *HttpClient) WithNotLogging(notLogging bool) *HttpClient {
	s.NotLogging = notLogging
	return s
}

func (s *HttpClient) WithSkipVerify(skipVerify bool) *HttpClient {
	s.SkipVerify = skipVerify
	return s
}

func (s *HttpClient) WithIsHttpTrace(isHttpTrace bool) *HttpClient {
	s.isHttpTrace = isHttpTrace
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
	if s.URL == "" {
		return nil, nil, errors.New("无效的URL")
	}

	if s.Header == nil {
		s.Header = defaultHeader
	}

	var span opentracing.Span
	// 链路追踪
	if s.isHttpTrace {
		// 新增埋点
		routerSuffix := "-"
		pos := strings.LastIndex(s.URL, "/")
		if pos+1 < len(s.URL) {
			routerSuffix = s.URL[pos+1:] //以接口路由的最后一级节点名作为span的操作名称
		}
		operationName := "http-" + routerSuffix
		pos = strings.Index(operationName, "?")
		if pos != -1 {
			operationName = operationName[:pos]
		}
		span = tracer.Start(&s.Ctx, operationName)
		defer span.Finish()
	}

	var (
		reqBytes []byte
		ok       bool
		err      error
	)

	if s.Body != nil {
		if reqBytes, ok = s.Body.([]byte); !ok {
			reqBytes, err = util.Json.Marshal(s.Body)
			if err != nil {
				return nil, nil, err
			}
		}
	}

	method := s.Method
	fields := logrus.Fields{}
	fields["http.req.url"] = s.URL
	fields["http.req.method"] = method

	fields["http.req.body"] = byteutil.Bytes2String(reqBytes)

	req, err := http.NewRequest(method, s.URL, bytes.NewReader(reqBytes))
	if err != nil {
		return nil, nil, err
	}

	hasContentType := false
	for k, v := range s.Header {
		req.Header.Set(k, v)
		if k == "Content-Type" {
			hasContentType = true
		}
	}

	if method == http.MethodPost && !hasContentType {
		req.Header.Set("Content-Type", "application/json")
	}

	// 链路追踪
	fields["http.req.isHttpTrace"] = s.isHttpTrace
	if s.isHttpTrace {
		span.LogKV("url", s.URL)
		span.LogKV("method", method)
		// span.LogKV("req", req) //不打印
		span.LogKV("log_id", s.LogID)
		req = req.WithContext(s.Ctx)
		// 注入 tracer 信息到 req.Header 中。`req.Header`是指针，所以`req.Header`的修改会影响到`fields["http.req.header"] = req.Header`
		tracer.InjectHTTPHeader(req.Context(), req.Header)
	}
	fields["http.req.header"] = req.Header

	fields["http.req.isProxy"] = false
	if s.proxyURL != nil {
		fields["http.req.isProxy"] = true
	}
	transport := &http.Transport{
		Proxy: http.ProxyURL(s.proxyURL),
		DialContext: (&net.Dialer{
			Timeout: s.Timeout,
		}).DialContext,
		TLSHandshakeTimeout: 0,
		IdleConnTimeout:     0,
		ProxyConnectHeader:  nil,
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: s.SkipVerify,
		},
	}

	client := &http.Client{
		Transport:     transport,
		CheckRedirect: s.checkRedirect,
		Timeout:       0,
	}

	if !s.NotLogging {
		plog.WithFields(fields).Info(s.LogID, "发起Http请求")
	}

	resp, err := client.Do(req)
	if err != nil {
		if !s.NotLogging {
			plog.WithError(err).
				WithField("url", s.URL).
				Error(s.LogID, "发起Http请求失败")
		}
		return nil, nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	_ = resp.Body.Close()
	if err != nil {
		if !s.NotLogging {
			plog.WithError(err).
				WithField("url", s.URL).
				Error(s.LogID, "读取http响应结果失败")
		}
		return nil, nil, common.FilterErr(err)
	}

	if resp.StatusCode != http.StatusOK {
		if !s.NotLogging {
			plog.WithError(err).
				WithFields(fields).
				WithField("Resp.StatusCode", resp.StatusCode).
				WithField("Resp.Status", resp.Status).
				Error(s.LogID, "Http 响应码异常")
		}
		return nil, nil, errors.Errorf("上游服务报错,http status code:%d", resp.StatusCode)
	}

	if !s.NotLogging {
		respBodyLog := body
		if len(body) > 1024 {
			respBodyLog = body[:1023]
		}
		if !s.NotLogging {
			plog.WithField("http.req.url", s.URL).
				WithField("http.resp.status_code", resp.StatusCode).
				WithField("http.resp.body", byteutil.Bytes2String(respBodyLog)).
				Info(s.LogID, "接收http响应")
		}
	}

	return body, resp, nil
}

func Do(req *http.Request) (respByte []byte, err error) {
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return
	}
	return ResponseToBytes(resp)
}

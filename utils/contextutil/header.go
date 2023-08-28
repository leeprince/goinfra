package contextutil

import (
	"context"
	"github.com/leeprince/goinfra/consts"
	"github.com/leeprince/goinfra/perror"
	"google.golang.org/grpc/metadata"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/19 10:07
 * @Desc:
 */

// Header 外部请求中的头部信息
type Header struct {
	UberTraceID string
	XRealIp     string
	LogId       string
	Token       string
	AccessToken string
}

// WriteHeaderContext 将 header 写入到当前进程的上下文 && 调用外部系统(Grpc调用)的上下文
func WriteHeaderContext(ctx context.Context, data *Header) context.Context {
	// 写入到当前进程的上下文
	ctx = context.WithValue(ctx, Header{}, data)
	
	// 写入调用外部系统(Grpc调用)的上下文
	return metadata.AppendToOutgoingContext(ctx,
		consts.HeaderXLogID, data.LogId,
		consts.HeaderToken, data.Token,
		consts.HeaderAccessToken, data.AccessToken,
	)
}

// ReadHeaderByContext 从当前进程的上下文中获取 header || 从外部系统(Grpc调用)的上下文中获取 header
func ReadHeaderByContext(ctx context.Context) (*Header, error) {
	// 从当前进程的上下文中获取 header
	header, err := ReadHeaderByCurrentContext(ctx)
	if err == nil {
		return header, nil
	}
	
	header, err = ReadHeaderByIncomingContext(ctx)
	if err != nil {
		return nil, err
	}
	return header, nil
}

// MustReadHeaderContext (从当前进程的上下文中获取 header || 从外部系统(Grpc调用)的上下文中获取 header), 并忽略错误
func MustReadHeaderContext(ctx context.Context) (*Header, error) {
	header, err := ReadHeaderByContext(ctx)
	if err != nil {
		return &Header{}, nil
	}
	return header, nil
}

// ReadHeaderByCurrentContext 从当前进程的上下文中获取 header
func ReadHeaderByCurrentContext(ctx context.Context) (*Header, error) {
	val := ctx.Value(Header{})
	header, ok := val.(*Header)
	if !ok {
		return nil, perror.BizErrDataNil
	}
	
	return header, nil
}

// MustReadHeaderByCurrentContext 从当前进程的上下文中获取 header，并忽略错误
func MustReadHeaderByCurrentContext(ctx context.Context) *Header {
	h, err := ReadHeaderByCurrentContext(ctx)
	if err != nil {
		return &Header{}
	}
	return h
}

// ReadHeaderByIncomingContext 从外部系统(Grpc调用)的上下文中获取 header
func ReadHeaderByIncomingContext(ctx context.Context) (*Header, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, perror.BizErrDataNil
	}
	
	header := &Header{
		UberTraceID: "",
		XRealIp:     "",
		LogId:       "",
		Token:       "",
		AccessToken: "",
	}
	for key, strings := range md {
		var value string
		if len(strings) > 0 {
			value = strings[0]
		}
		switch key {
		case consts.HeaderUberTraceID:
			header.UberTraceID = value
		case consts.HeaderXRealIp:
			header.XRealIp = value
		case consts.HeaderXLogID:
			header.LogId = value
		case consts.HeaderToken:
			header.Token = value
		case consts.HeaderAccessToken:
			header.AccessToken = value
		}
	}
	
	if header.UberTraceID == "" &&
		header.XRealIp == "" &&
		header.LogId == "" &&
		header.Token == "" &&
		header.AccessToken == "" {
		return nil, perror.BizErrDataNil
	}
	
	return header, nil
}

// MustReadHeaderByIncomingContext 从外部系统(Grpc调用)的上下文中获取 header，并忽略错误
func MustReadHeaderByIncomingContext(ctx context.Context) (*Header, error) {
	header, err := ReadHeaderByIncomingContext(ctx)
	if err != nil {
		return &Header{}, nil
	}
	
	return header, nil
}

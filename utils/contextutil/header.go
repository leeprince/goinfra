package gdheader

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
	LogID       string
	Token       string
	AccessToken string
}

// UpdateCurrentAndOutgoingContext 将 header 写入到当前进程的上下文 && 调用外部系统(Grpc调用)的上下文
func UpdateCurrentAndOutgoingContext(ctx context.Context, data Header) context.Context {
	ctx = context.WithValue(ctx, Header{}, &data)
	
	return metadata.AppendToOutgoingContext(ctx,
		consts.HeaderToken, data.Token,
		consts.HeaderAccessToken, data.AccessToken,
	)
}

// ReadHeaderByCurrentContext 从当前进程的上下文中获取 header
func ReadHeaderByCurrentContext(ctx context.Context) (*Header, error) {
	val := ctx.Value(Header{})
	header, ok := val.(*Header)
	if !ok {
		return nil, perror.BizErrTypeAsserts
	}
	
	return header, nil
}

// ReadHeaderByIncomingContext 从外部系统(Grpc调用)的上下文中获取 header
func ReadHeaderByIncomingContext(ctx context.Context) (*Header, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, perror.BizErrDataNil
	}
	
	header := &Header{
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
		case consts.HeaderLogID:
			header.LogID = value
		case consts.HeaderToken:
			header.Token = value
		case consts.HeaderAccessToken:
			header.AccessToken = value
		}
	}
	
	return header, nil
}

func MustReadHeaderByCurrentContext(ctx context.Context) *Header {
	h, err := ReadHeaderByCurrentContext(ctx)
	if err != nil {
		return &Header{}
	}
	return h
}

package security

import (
	"encoding/base64"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/28 下午10:30
 * @Desc:
 */

// HAMAC基于的`Hash函数`类型
type HMACHashType int32

const (
	HMACHashTypeMd5 HMACHashType = iota
	HMACHashTypeSha1
	HMACHashTypeSha256
)

// AES的加密模式
type AESBlockModeType int32

const (
	AESBlockModeTypeECB AESBlockModeType = iota
	AESBlockModeTypeCBC
	AESBlockModeTypeCTR
	AESBlockModeTypeCFB
	AESBlockModeTypeOFB
)

const (
	aesDefaultIV = "0000000000000000"
)

type OutInputType int32

const (
	OutputTypeBase64 OutInputType = iota
	OutputTypeHex
)

type Option struct {
	outputType       OutInputType     // 加密输出的字符串类型：base64/十六字符串。默认：base64
	inputType        OutInputType     // 解密输入的字符串类型：base64/十六字符串。默认：base64
	hmacHashType     HMACHashType     // HAMAc基于的`Hash函数`类型
	aesIV            string           // AES iv
	aesBlockModeType AESBlockModeType // AES的加密模式
	bcryptCost       int              // bcrypt 的工作因子
}

type OptionFunc func(opt *Option)

func OutputFormat(src []byte, outputType OutInputType) string {
	switch outputType {
	case OutputTypeBase64:
		return base64.StdEncoding.EncodeToString(src)
	case OutputTypeHex:
		return hex.EncodeToString(src)
	default:
		return base64.StdEncoding.EncodeToString(src)
	}
}

func InputFormat(src string, inputType OutInputType) ([]byte, error) {
	switch inputType {
	case OutputTypeBase64:
		return base64.StdEncoding.DecodeString(src)
	case OutputTypeHex:
		return hex.DecodeString(src)
	default:
		return base64.StdEncoding.DecodeString(src)
	}
}

// 初始化可选项
func initOption(opts ...OptionFunc) *Option {
	opt := &Option{
		outputType:       OutputTypeBase64,
		inputType:        OutputTypeBase64,
		hmacHashType:     HMACHashTypeMd5,
		aesIV:            aesDefaultIV,
		aesBlockModeType: AESBlockModeTypeCBC,
		bcryptCost:       bcrypt.MinCost,
	}
	for _, optFunc := range opts {
		optFunc(opt)
	}
	return opt
}

func WithOutputType(v OutInputType) OptionFunc {
	return func(opt *Option) {
		opt.outputType = v
	}
}

func WithInputType(v OutInputType) OptionFunc {
	return func(opt *Option) {
		opt.inputType = v
	}
}

func WithHMACHashType(v HMACHashType) OptionFunc {
	return func(opt *Option) {
		opt.hmacHashType = v
	}
}

func WithAESIV(v string) OptionFunc {
	return func(opt *Option) {
		opt.aesIV = v
	}
}

func WithAESBlockModeType(v AESBlockModeType) OptionFunc {
	return func(opt *Option) {
		opt.aesBlockModeType = v
	}
}

func WithBcryptCost(v int) OptionFunc {
	return func(opt *Option) {
		if v > bcrypt.MinCost || v > bcrypt.MaxCost {
			v = bcrypt.MinCost
		}
		opt.bcryptCost = v
	}
}

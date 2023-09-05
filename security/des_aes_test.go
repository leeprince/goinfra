package security_test

import (
	"fmt"
	"github.com/leeprince/goinfra/security"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/7/3 上午12:33
 * @Desc:
 */

const (
	DESAESSrcStr1  = "DESAESSrc01"
	DESAESSrcStr2  = "DESAESSrc02 我爱中国！"
	DESAESSrcStr3  = "DESAESSrc03 我爱中国！我爱中国！我爱中国！"
	DESAESSrcStr4  = "DESAESSrc03 我爱中国！我爱中国！我爱中国！DESAESSrc03 我爱中国！我爱中国！我爱中国！DESAESSrc03 我爱中国！我爱中国！我爱中国！DESAESSrc03 我爱中国！我爱中国！我爱中国！DESAESSrc03 我爱中国！我爱中国！我爱中国！DESAESSrc03 我爱中国！我爱中国！我爱中国！DESAESSrc03 我爱中国！我爱中国！我爱中国！"
	DESAESKey8Str  = "DESAESK1"
	DESAESKey16Str = "DESAESK2DESAESK2"
	DESAESKey24Str = "DESAESK3DESAESK3DESAESK3"
	DESAESKey32Str = "DESAESK4DESAESK4DESAESK4DESAESK4"
	
	/*DESAESKeyAppTicketStr = "LHZZXW---AA:BB:CC:中国---1692507180"
	DESAESKeyAppTicketKey = "zz2023082000yyds"
	DESAESKeyAppTicketIv  = "yyds2023082000zz"*/
	
	DESAESKeyAppTicketStr = "AA:BB:CC:中国---1692507180"
	DESAESKeyAppTicketKey = "ApplyCodeKey-xxx"
	DESAESKeyAppTicketIv  = "ApplyCodeIv-xxxx"
	
	AESIv   = "1000000000011111"
	AESIvv  = "abcdef0000000000"
	AESIvvv = "abcdefghi0000000"
)

func TestDESEncryptDecrypt(t *testing.T) {
	type args struct {
		src            string
		key            string
		optsDESEncrypt []security.OptionFunc
		optsDecrypt    []security.OptionFunc
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "DESAESSrcStr1-DESAESKey8Str",
			args: args{
				src:            DESAESSrcStr1,
				key:            DESAESKey8Str,
				optsDESEncrypt: []security.OptionFunc{},
				optsDecrypt:    []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr1-DESAESKey8Str",
			args: args{
				src:            DESAESSrcStr1,
				key:            DESAESKey8Str,
				optsDESEncrypt: []security.OptionFunc{},
				optsDecrypt:    []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr2-DESAESKey8Str",
			args: args{
				src:            DESAESSrcStr2,
				key:            DESAESKey8Str,
				optsDESEncrypt: []security.OptionFunc{},
				optsDecrypt:    []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr3-DESAESKey8Str",
			args: args{
				src:            DESAESSrcStr3,
				key:            DESAESKey8Str,
				optsDESEncrypt: []security.OptionFunc{},
				optsDecrypt:    []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr1-DESAESKey24Str",
			args: args{
				src:            DESAESSrcStr1,
				key:            DESAESKey24Str,
				optsDESEncrypt: []security.OptionFunc{},
				optsDecrypt:    []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr2-DESAESKey24Str",
			args: args{
				src:            DESAESSrcStr2,
				key:            DESAESKey24Str,
				optsDESEncrypt: []security.OptionFunc{},
				optsDecrypt:    []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr3-DESAESKey24Str",
			args: args{
				src:            DESAESSrcStr3,
				key:            DESAESKey24Str,
				optsDESEncrypt: []security.OptionFunc{},
				optsDecrypt:    []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr3-DESAESKey24Str",
			args: args{
				src:            DESAESSrcStr3,
				key:            DESAESKey24Str,
				optsDESEncrypt: []security.OptionFunc{},
				optsDecrypt:    []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr3-DESAESKey24Str",
			args: args{
				src:            DESAESSrcStr3,
				key:            DESAESKey24Str,
				optsDESEncrypt: []security.OptionFunc{},
				optsDecrypt:    []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr4-DESAESKey8Str",
			args: args{
				src:            DESAESSrcStr4,
				key:            DESAESKey8Str,
				optsDESEncrypt: []security.OptionFunc{},
				optsDecrypt:    []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr4-DESAESKey24Str",
			args: args{
				src:            DESAESSrcStr4,
				key:            DESAESKey24Str,
				optsDESEncrypt: []security.OptionFunc{},
				optsDecrypt:    []security.OptionFunc{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			desEncrypt, err := security.DESEncrypt(tt.args.src, tt.args.key, tt.args.optsDESEncrypt...)
			fmt.Println("DESEncrypt:", desEncrypt, err)
			if err != nil {
				fmt.Println()
				fmt.Println("----")
				return
			}
			desDecrypt, err := security.DESDecrypt(desEncrypt, tt.args.key, tt.args.optsDecrypt...)
			fmt.Println("DESDecrypt:", desDecrypt, err)
		})
	}
}

func TestAESEncryptDecrypt(t *testing.T) {
	type args struct {
		text           string
		key            string
		optsAESEncrypt []security.OptionFunc
		optsAESDecrypt []security.OptionFunc
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "DESAESSrcStr1-DESAESKey16Str",
			args: args{
				text:           DESAESSrcStr1,
				key:            DESAESKey16Str,
				optsAESEncrypt: []security.OptionFunc{},
				optsAESDecrypt: []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr1-DESAESKey16Str",
			args: args{
				text:           DESAESSrcStr1,
				key:            DESAESKey16Str,
				optsAESEncrypt: []security.OptionFunc{},
				optsAESDecrypt: []security.OptionFunc{},
			},
		},
		{
			name: "Decrypt_not_AESIv",
			args: args{
				text:           DESAESSrcStr1,
				key:            DESAESKey16Str,
				optsAESEncrypt: []security.OptionFunc{security.WithAESIV(AESIv)},
				optsAESDecrypt: []security.OptionFunc{},
			},
		},
		{
			name: "",
			args: args{
				text:           DESAESSrcStr1,
				key:            DESAESKey16Str,
				optsAESEncrypt: []security.OptionFunc{security.WithAESIV(AESIv)},
				optsAESDecrypt: []security.OptionFunc{security.WithAESIV(AESIv)},
			},
		},
		{
			name: "AESIvv",
			args: args{
				text:           DESAESSrcStr1,
				key:            DESAESKey16Str,
				optsAESEncrypt: []security.OptionFunc{security.WithAESIV(AESIvv)},
				optsAESDecrypt: []security.OptionFunc{security.WithAESIV(AESIvv)},
			},
		},
		{
			name: "AESIvvv",
			args: args{
				text:           DESAESSrcStr1,
				key:            DESAESKey16Str,
				optsAESEncrypt: []security.OptionFunc{security.WithAESIV(AESIvvv)},
				optsAESDecrypt: []security.OptionFunc{security.WithAESIV(AESIvvv)},
			},
		},
		{
			name: "DESAESSrcStr2-DESAESKey16Str",
			args: args{
				text:           DESAESSrcStr2,
				key:            DESAESKey16Str,
				optsAESEncrypt: []security.OptionFunc{},
				optsAESDecrypt: []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr3-DESAESKey16Str",
			args: args{
				text:           DESAESSrcStr3,
				key:            DESAESKey16Str,
				optsAESEncrypt: []security.OptionFunc{},
				optsAESDecrypt: []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr1-DESAESKey24Str",
			args: args{
				text:           DESAESSrcStr1,
				key:            DESAESKey24Str,
				optsAESEncrypt: []security.OptionFunc{},
				optsAESDecrypt: []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr2-DESAESKey24Str",
			args: args{
				text:           DESAESSrcStr2,
				key:            DESAESKey24Str,
				optsAESEncrypt: []security.OptionFunc{},
				optsAESDecrypt: []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr3-DESAESKey24Str",
			args: args{
				text:           DESAESSrcStr3,
				key:            DESAESKey24Str,
				optsAESEncrypt: []security.OptionFunc{},
				optsAESDecrypt: []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr1-DESAESKey32Str",
			args: args{
				text:           DESAESSrcStr1,
				key:            DESAESKey32Str,
				optsAESEncrypt: []security.OptionFunc{},
				optsAESDecrypt: []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr2-DESAESKey32Str",
			args: args{
				text:           DESAESSrcStr2,
				key:            DESAESKey32Str,
				optsAESEncrypt: []security.OptionFunc{},
				optsAESDecrypt: []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr3-DESAESKey32Str",
			args: args{
				text:           DESAESSrcStr3,
				key:            DESAESKey32Str,
				optsAESEncrypt: []security.OptionFunc{},
				optsAESDecrypt: []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr3-DESAESKey32Str",
			args: args{
				text:           DESAESSrcStr3,
				key:            DESAESKey32Str,
				optsAESEncrypt: []security.OptionFunc{},
				optsAESDecrypt: []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr4-DESAESKey16Str",
			args: args{
				text:           DESAESSrcStr4,
				key:            DESAESKey16Str,
				optsAESEncrypt: []security.OptionFunc{},
				optsAESDecrypt: []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr4-DESAESKey24Str",
			args: args{
				text:           DESAESSrcStr4,
				key:            DESAESKey24Str,
				optsAESEncrypt: []security.OptionFunc{},
				optsAESDecrypt: []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr4-DESAESKey32Str",
			args: args{
				text:           DESAESSrcStr4,
				key:            DESAESKey32Str,
				optsAESEncrypt: []security.OptionFunc{},
				optsAESDecrypt: []security.OptionFunc{},
			},
		},
		{
			name: "DESAESSrcStr4-DESAESKey32Str",
			args: args{
				text:           DESAESSrcStr4,
				key:            DESAESKey32Str,
				optsAESEncrypt: []security.OptionFunc{},
				optsAESDecrypt: []security.OptionFunc{},
			},
		},
		{
			name: "DESAESKeyAppTicketStr-DESAESKeyAppTicketKey",
			args: args{
				text:           DESAESKeyAppTicketStr,
				key:            DESAESKeyAppTicketKey,
				optsAESEncrypt: []security.OptionFunc{security.WithAESIV(DESAESKeyAppTicketIv)},
				optsAESDecrypt: []security.OptionFunc{security.WithAESIV(DESAESKeyAppTicketIv)},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			desEncrypt, err := security.AESEncrypt(tt.args.text, tt.args.key, tt.args.optsAESEncrypt...)
			fmt.Println("DESEncrypt:", desEncrypt, err)
			if err != nil {
				fmt.Println()
				fmt.Println("----")
				return
			}
			desDecrypt, err := security.AESDecrypt(desEncrypt, tt.args.key, tt.args.optsAESDecrypt...)
			fmt.Println("DESDecrypt:", desDecrypt, err)
		})
	}
}

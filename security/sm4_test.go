package security_test

import (
    "fmt"
    "github.com/leeprince/goinfra/security"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/7/5 上午9:31
 * @Desc:
 */

const (
    SM4SrcStr1 = "SM4Src01"
    SM4SrcStr2 = "SM4Src02 我爱中国！"
    SM4SrcStr3 = "SM4Src03 我爱中国！我爱中国！我爱中国！"
    SM4SrcStr4 = "SM4Src04 我爱中国！我爱中国！我爱中国！SM4Src04 我爱中国！我爱中国！我爱中国！SM4Src04 我爱中国！我爱中国！我爱中国！SM4Src04 我爱中国！我爱中国！我爱中国！"
    SM4SrcStr5 = "SM4Src04 我爱中国！我爱中国！我爱中国！SM4Src04 我爱中国！我爱中国！我爱中国！SM4Src04 我爱中国！我爱中国！我爱中国！SM4Src04 我爱中国！我爱中国！我爱中国！SM4Src04 我爱中国！我爱中国！我爱中国！"
    SM4Key1    = "fsafdsf898sdfa31"
)

func TestSM4EncryptSM4Decrypt(t *testing.T) {
    type args struct {
        src         string
        key         string
        encryptOpts []security.OptionFunc
        decryptOpts []security.OptionFunc
    }
    tests := []struct {
        name    string
        args    args
        want    string
        wantErr bool
    }{
        {
            args: args{
                src:         SM4SrcStr1,
                key:         SM4Key1,
                encryptOpts: nil,
            },
        },
        {
            args: args{
                src:         SM4SrcStr1,
                key:         SM4Key1,
                encryptOpts: nil,
            },
        },
        {
            args: args{
                src:         SM4SrcStr2,
                key:         SM4Key1,
                encryptOpts: nil,
            },
        },
        {
            args: args{
                src:         SM4SrcStr3,
                key:         SM4Key1,
                encryptOpts: nil,
            },
        },
        {
            args: args{
                src:         SM4SrcStr4,
                key:         SM4Key1,
                encryptOpts: nil,
            },
        },
        {
            args: args{
                src:         SM4SrcStr5,
                key:         SM4Key1,
                encryptOpts: nil,
            },
        },
        {
            args: args{
                src:         SM4SrcStr5,
                key:         SM4Key1,
                encryptOpts: []security.OptionFunc{security.WithOutputType(security.OutputTypeHex)},
                decryptOpts: []security.OptionFunc{security.WithInputType(security.OutputTypeHex)},
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            sm4Encrypt, err := security.SM4Encrypt(tt.args.src, tt.args.key, tt.args.encryptOpts...)
            fmt.Println("SM4Encrypt", sm4Encrypt, err)
            if err != nil {
                fmt.Println()
                fmt.Println("---")
                return
            }
            
            sm4decrypt, err := security.SM4Decrypt(sm4Encrypt, tt.args.key, tt.args.decryptOpts...)
            fmt.Println("SM4Encrypt", sm4decrypt, err)
            if err != nil {
                return
            }
        })
    }
}

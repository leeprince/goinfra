package security_test

import (
    "fmt"
    "github.com/leeprince/goinfra/security"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/7/3 下午7:06
 * @Desc:
 */

const (
    RSASrcStr1        = "RSASrc01"
    RSASrcStr2        = "RSASrc02 我爱中国！"
    RSASrcStr3        = "RSASrc03 我爱中国！我爱中国！我爱中国！"
    RSASrcStr4        = "RSASrc04 我爱中国！我爱中国！我爱中国！RSASrc04 我爱中国！我爱中国！我爱中国！RSASrc04 我爱中国！我爱中国！我爱中国！RSASrc04 我爱中国！我爱中国！我爱中国！"
    RSASrcStr5        = "RSASrc04 我爱中国！我爱中国！我爱中国！RSASrc04 我爱中国！我爱中国！我爱中国！RSASrc04 我爱中国！我爱中国！我爱中国！RSASrc04 我爱中国！我爱中国！我爱中国！RSASrc04 我爱中国！我爱中国！我爱中国！"
    RSAPrivateKey1024 = `-----BEGIN RSA PRIVATE KEY-----
MIICXQIBAAKBgQCnJO93pHgEsebeAyxUi6M4EQacjV3inwYzdxo/XWx9IWu78sou
EaTikl4p3HeKgC3R2SnxSpDyhBUrm0vlnDkuc51wQRHK1UlFHKM0nhEDU6WGJoDK
q5sNj4s2g+Pcg6RzjADWzzYmsNlRIlsLWTRs8/w/+dqPWv0bFOMRZQcbMQIDAQAB
AoGAYveux0PeQMFp9ukgFYt9AJSsOoRGJAqPLGgIZZ6Wv1zLosT2y+JspC+Qi+7b
5WlSOCADArlpK//jXSed//3JqO+ayNnst9BbtUiFNy2400+DARpTIHY0LbOwCylG
btHVxC+Qe37NMx2LZxl8MeZQn+WFht+6QXs+E5xT3xw4BLECQQDPqROHg6oD9TQQ
DQWHGDBfC8aW5hS70oHSlqGB18KwEhEetlXH0ipcYz0La1bIXvkmXaKpu8txyzFW
JAxLifDdAkEAzg1qOXFM9wOcYn9LRd0x+kqyqwbkIkX1BZxd5oHRlimJbT2NtOOv
oJTlVIZamRBB5TAtOFz7JSHnJKehXuwkZQJAKF1iKW5DZv+Lvi75yxe9l0wPrxdM
InI5v/h9rmKFOnpYj5K7u9qzV0AHBqg3tz0WywlabAnP8u+fSHI7XZeTlQJBALU2
3RX5zAtt1IpXkza1Sy/po+p/3AE8bznpBDgmMdITY7Z8LPVKTPo4GNxWcLUutBVB
YnOwanuErojxGB2oJnUCQQC7nyHxGh2DsMqHwYyofcOOQJvLzog3lz0k69KvErnR
BjQ6otfKqE3KH6LGdzZywNOMMgWyhQYyMWkikgxyrX9E
-----END RSA PRIVATE KEY-----`
    RSAPrivateKey2048 = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAyifjJuo1UYB3Is4pLYk4Q8CpQAXMNfsx7Klczc2uDOi5HKpU
nScZKKaHfME/BJsqjvXmt8QO1+wLO2Poizf+hlmw+iCU47sWdzkpuEdAu4GnsdHB
K0CjLCxPSuR8on8986j0L/rdWbU9mZCXJJxGkZF+qui2cGcIWh7B1bX+Hk5ijgJA
hIzpYwEAqMIzDqsftpRw8nzGs+RmTOR6vB3JVcFHinwRBfw6MAm8PZc/7ZQT9MbY
7FBBixcdrUQBuSdYwd1h7Q1SoGjYLAoaW+QM7Xwx2dR374zWrf9a12kmimOJHJE+
v1IL5fgm/CDYeDVgR40ygZK8aLIf/TKvWoXnpwIDAQABAoIBAHRNTlntFI+3INNd
WENiVoRkKvsyWrITuj01kra0RhYXLahYNiXTgJ3qcLhNVTwJKQrmXb4LhZ6jpjKc
4AR4Sm5w22iLphz+XYZp1J64H33AsgGTc1a28SlQtK2ZljrGiZXM9e9EnGQn6TCY
BtyBK/hDhGQ/TfaM0DlIyKty8Kft2vD974oqH3skBRd47onLGPDyIFFIdNODYcW3
/wUdrawjVQswo7Ik8SaKcq9g6AoFLwGBtQGAMmxg7NMYzNqMmTYhWMlGQ5wpFXJo
q0nXoVNL+DLB5eXEwZL6zH70CpMRE5iq82Esd2hisS+UigLh0XOgjEsK2sPs54rb
KC9xtUECgYEA5V8AKPKA2wwKt2FP1S4Eqs9YRf7usmK80vnq26EoThhX74FeN35b
WkRf6gROYwSLnM4mLbnR1XFwXgv484zHsXKqzjXWgnwR3QzQhNjoxzm21HNoC0t8
DiSss6dvseneW0XXSjWvicjSBBLYdieU76OiSUnuQ5iVnRKziaHc60UCgYEA4aAJ
o34Z5VkqUjfDJp4ncWZDffrv3UQd7OMDlXTSsvUpdUTF+JSpDed02lvqPsVrXwrr
sRiabHMJKhDS55W9iyd5TmNoexBRpZoazSueXgRJ0oScKlIwx/DSrw01FhDTIl2p
akVN0qJgWrwNKrKJKZwi97IijvufSQPKW16kf/sCgYAcf9xWOiN2lB10wZuYwloE
GzU9pTplYc1SbYkX9wM6CN5MPQfG32VcZuh//D79IKB0QE0QG2mOGsU6ekRZhqF9
U+ETNC9OETpq+9+g0g7CSlKEPT6tQJjObRIkVGaVdZiSQLBKYTdJaHFn3iuVKr/f
srZEYvI+5eOZG6zBKiJ3/QKBgHl7qRTttdXGf8ILIjlt2ID3dgmkDnjNz2sYBHr5
juUqmer5X7rrmGbTJBjaerLXq2tePu949tTDz8BllJl7B7agR3GMltoEPGH1Ks8j
2D55AqKmIkurO3a8VURJ0TaTUotjcO+2ZyOtqEHSlShTPwU3e6BwuqjQFMDEmLU/
rUT3AoGAF0GwN+Oa0yxjMiQub1qaaNh1sDsljn83XOVP96/vmMNaqcdbQ1U0LFNo
t/rqwLhGpCp2Zbv4ZgX5O5ObWsHlTeuTn3Rp8SghyKo1vvVtk+CwhuMkUQPRfiUp
qAkHEbdtKzca8rU2WszH3Cy8oHJXakErCjNgX4SMALDbG3W6Hac=
-----END RSA PRIVATE KEY-----`
    RSAPublicKey1024 = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQCnJO93pHgEsebeAyxUi6M4EQac
jV3inwYzdxo/XWx9IWu78souEaTikl4p3HeKgC3R2SnxSpDyhBUrm0vlnDkuc51w
QRHK1UlFHKM0nhEDU6WGJoDKq5sNj4s2g+Pcg6RzjADWzzYmsNlRIlsLWTRs8/w/
+dqPWv0bFOMRZQcbMQIDAQAB
-----END PUBLIC KEY-----`
    RSAPublicKey2048 = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAyifjJuo1UYB3Is4pLYk4
Q8CpQAXMNfsx7Klczc2uDOi5HKpUnScZKKaHfME/BJsqjvXmt8QO1+wLO2Poizf+
hlmw+iCU47sWdzkpuEdAu4GnsdHBK0CjLCxPSuR8on8986j0L/rdWbU9mZCXJJxG
kZF+qui2cGcIWh7B1bX+Hk5ijgJAhIzpYwEAqMIzDqsftpRw8nzGs+RmTOR6vB3J
VcFHinwRBfw6MAm8PZc/7ZQT9MbY7FBBixcdrUQBuSdYwd1h7Q1SoGjYLAoaW+QM
7Xwx2dR374zWrf9a12kmimOJHJE+v1IL5fgm/CDYeDVgR40ygZK8aLIf/TKvWoXn
pwIDAQAB
-----END PUBLIC KEY-----`
)

func TestRSAEncryptDecrypt(t *testing.T) {
    type args struct {
        src        string
        publicKey  string
        privateKey string
        opts       []security.OptionFunc
    }
    tests := []struct {
        name    string
        args    args
        want    string
        wantErr bool
    }{
        {
            args: args{
                src:        RSASrcStr1,
                publicKey:  RSAPublicKey1024,
                privateKey: RSAPrivateKey1024,
                opts:       nil,
            },
        },
        {
            args: args{
                src:        RSASrcStr1,
                publicKey:  RSAPublicKey1024,
                privateKey: RSAPrivateKey1024,
                opts:       nil,
            },
        },
        {
            args: args{
                src:        RSASrcStr2,
                publicKey:  RSAPublicKey1024,
                privateKey: RSAPrivateKey1024,
                opts:       nil,
            },
        },
        {
            args: args{
                src:        RSASrcStr3,
                publicKey:  RSAPublicKey1024,
                privateKey: RSAPrivateKey1024,
                opts:       nil,
            },
        },
        {
            args: args{
                src:        RSASrcStr4,
                publicKey:  RSAPublicKey1024,
                privateKey: RSAPrivateKey1024,
                opts:       nil,
            },
        },
        {
            args: args{
                src:        RSASrcStr4,
                publicKey:  RSAPublicKey2048,
                privateKey: RSAPrivateKey2048,
                opts:       nil,
            },
        },
        {
            args: args{
                src:        RSASrcStr5,
                publicKey:  RSAPublicKey1024,
                privateKey: RSAPrivateKey1024,
                opts:       nil,
            },
        },
        {
            args: args{
                src:        RSASrcStr5,
                publicKey:  RSAPublicKey2048,
                privateKey: RSAPrivateKey2048,
                opts:       nil,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            rsaEncrypt, err := security.RSAEncrypt(tt.args.src, tt.args.publicKey, tt.args.opts...)
            fmt.Println("RSAEncrypt:", rsaEncrypt, err)
            if err != nil {
                return
            }
            rsaDecrypt, err := security.RSADecrypt(rsaEncrypt, tt.args.privateKey, tt.args.opts...)
            fmt.Println("RSADecrypt:", rsaDecrypt, err)
            if err != nil {
                return
            }
        })
    }
}

func TestGenerateRsaKey(t *testing.T) {
    type args struct {
        bits int
    }
    tests := []struct {
        name        string
        args        args
        wantPrivKey string
        wantPubKey  string
        wantErr     bool
    }{
        {
            args: args{bits: 1024},
        },
        {
            args: args{bits: 2048},
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            gotPrivKey, gotPubKey, err := security.GenerateRsaKey(tt.args.bits)
            fmt.Println()
            fmt.Println(gotPrivKey)
            fmt.Println(gotPubKey)
            fmt.Println(err)
        })
    }
}

func TestRSASignWithSHA256RSASignVerifyWithSha256(t *testing.T) {
    type args struct {
        src        string
        publicKey  string
        privateKey string
        opts       []security.OptionFunc
    }
    tests := []struct {
        name    string
        args    args
        want    string
        wantErr bool
    }{
        {
            args:args{
                src:        RSASrcStr1,
                publicKey:  RSAPublicKey1024,
                privateKey: RSAPrivateKey1024,
                opts:       nil,
            },
        },
        {
            args:args{
                src:        RSASrcStr2,
                publicKey:  RSAPublicKey1024,
                privateKey: RSAPrivateKey1024,
                opts:       nil,
            },
        },
        {
            args:args{
                src:        RSASrcStr3,
                publicKey:  RSAPublicKey1024,
                privateKey: RSAPrivateKey1024,
                opts:       nil,
            },
        },
        {
            args:args{
                src:        RSASrcStr4,
                publicKey:  RSAPublicKey1024,
                privateKey: RSAPrivateKey1024,
                opts:       nil,
            },
        },
        {
            args:args{
                src:        RSASrcStr5,
                publicKey:  RSAPublicKey1024,
                privateKey: RSAPrivateKey1024,
                opts:       nil,
            },
        },
        {
            args:args{
                src:        RSASrcStr1,
                publicKey:  RSAPublicKey2048,
                privateKey: RSAPrivateKey2048,
                opts:       nil,
            },
        },
        {
            args:args{
                src:        RSASrcStr2,
                publicKey:  RSAPublicKey2048,
                privateKey: RSAPrivateKey2048,
                opts:       nil,
            },
        },
        {
            args:args{
                src:        RSASrcStr3,
                publicKey:  RSAPublicKey2048,
                privateKey: RSAPrivateKey2048,
                opts:       nil,
            },
        },
        {
            args:args{
                src:        RSASrcStr4,
                publicKey:  RSAPublicKey2048,
                privateKey: RSAPrivateKey2048,
                opts:       nil,
            },
        },
        {
            args:args{
                src:        RSASrcStr5,
                publicKey:  RSAPublicKey2048,
                privateKey: RSAPrivateKey2048,
                opts:       nil,
            },
        },
    }
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            rsaSignWithSHA256, err := security.RSASignWithSHA256(tt.args.src, tt.args.privateKey, tt.args.opts...)
            fmt.Println("RSASignWithSHA256", rsaSignWithSHA256, err)
            if err != nil {
                fmt.Println()
                fmt.Println("---")
                return
            }
            b, err := security.RSASignVerifyWithSha256(tt.args.src, rsaSignWithSHA256, tt.args.publicKey, tt.args.opts...)
            fmt.Println("RSASignVerifyWithSha256", b, err)
        })
    }
}

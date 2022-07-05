package security_test

import (
    "fmt"
    "github.com/leeprince/goinfra/security"
    "testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/28 下午10:45
 * @Desc:
 */

const (
    MD5SHAHMACBcryptSrcStr = "MD5SHAHMACBcryptSrcStr"
    // MD5SHAHMACBcryptSrcStr = "MD5SHAHMACSrcStr我爱中国！"
    MD5SHAHMACBcryptKeyStr = "MD5SHAHMACBcryptKeyStr"
)

func TestMd5(t *testing.T) {
    fmt.Println(security.MD5(MD5SHAHMACBcryptSrcStr))
    fmt.Println(security.MD5(MD5SHAHMACBcryptSrcStr))
}

func TestSha256(t *testing.T) {
    fmt.Println(security.SHA256(MD5SHAHMACBcryptSrcStr))
    fmt.Println(security.SHA256(MD5SHAHMACBcryptSrcStr))
}

func TestHmacHash(t *testing.T) {
    fmt.Println("--- HMACHashTypeMd5")
    fmt.Println(security.HMACHash(MD5SHAHMACBcryptSrcStr, MD5SHAHMACBcryptKeyStr, security.WithHMACHashType(security.HMACHashTypeMd5)))
    fmt.Println(security.HMACHash(MD5SHAHMACBcryptSrcStr, MD5SHAHMACBcryptKeyStr, security.WithHMACHashType(security.HMACHashTypeMd5)))
    
    fmt.Println("--- HMACHashTypeSha1")
    fmt.Println(security.HMACHash(MD5SHAHMACBcryptSrcStr, MD5SHAHMACBcryptKeyStr, security.WithHMACHashType(security.HMACHashTypeSha1)))
    fmt.Println(security.HMACHash(MD5SHAHMACBcryptSrcStr, MD5SHAHMACBcryptKeyStr, security.WithHMACHashType(security.HMACHashTypeSha1)))
    
    fmt.Println("--- HMACHashTypeSha256")
    fmt.Println(security.HMACHash(MD5SHAHMACBcryptSrcStr, MD5SHAHMACBcryptKeyStr, security.WithHMACHashType(security.HMACHashTypeSha256)))
    fmt.Println(security.HMACHash(MD5SHAHMACBcryptSrcStr, MD5SHAHMACBcryptKeyStr, security.WithHMACHashType(security.HMACHashTypeSha256)))
}

func TestBcryptBcryptVerify(t *testing.T) {
    crypt, err := security.Bcrypt(MD5SHAHMACBcryptSrcStr)
    fmt.Println("Bcrypt", crypt, err)
    if err != nil {
        return
    }
    fmt.Println(security.BcryptVerify(MD5SHAHMACBcryptSrcStr, crypt))
    
    crypt, err = security.Bcrypt(MD5SHAHMACBcryptSrcStr)
    fmt.Println("Bcrypt", crypt, err)
    if err != nil {
        return
    }
    fmt.Println(security.BcryptVerify(MD5SHAHMACBcryptSrcStr, crypt))
    
    crypt, err = security.Bcrypt(MD5SHAHMACBcryptSrcStr, security.WithOutputType(security.OutputTypeHex))
    fmt.Println("Bcrypt", crypt, err)
    if err != nil {
        return
    }
    fmt.Println(security.BcryptVerify(MD5SHAHMACBcryptSrcStr, crypt))
    
    crypt, err = security.Bcrypt(MD5SHAHMACBcryptSrcStr, security.WithOutputType(security.OutputTypeHex))
    fmt.Println("Bcrypt", crypt, err)
    if err != nil {
        return
    }
    fmt.Println(security.BcryptVerify(MD5SHAHMACBcryptSrcStr, crypt, security.WithInputType(security.OutputTypeHex)))
}
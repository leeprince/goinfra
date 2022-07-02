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
    srcStr = "leeprince"
    keyStr = "prince.lee"
)

func TestMd5(t *testing.T) {
    fmt.Println(security.Md5(srcStr, security.WithIsToHex(false)))
    fmt.Println(security.Md5(srcStr, security.WithIsToHex(true)))
}

func TestSha256(t *testing.T) {
    fmt.Println(security.Sha256(srcStr, security.WithIsToHex(false)))
    fmt.Println(security.Sha256(srcStr, security.WithIsToHex(true)))
}

func TestHmacHash(t *testing.T) {
    fmt.Println("--- HmacHashTypeMd5")
    fmt.Println(security.HmacHash(security.HmacHashTypeMd5, srcStr, keyStr, security.WithIsToHex(false)))
    fmt.Println(security.HmacHash(security.HmacHashTypeMd5, srcStr, keyStr, security.WithIsToHex(true)))
    
    fmt.Println("--- HmacHashTypeSha1")
    fmt.Println(security.HmacHash(security.HmacHashTypeSha1, srcStr, keyStr, security.WithIsToHex(false)))
    fmt.Println(security.HmacHash(security.HmacHashTypeSha1, srcStr, keyStr, security.WithIsToHex(true)))
    
    fmt.Println("--- HmacHashTypeSha256")
    fmt.Println(security.HmacHash(security.HmacHashTypeSha256, srcStr, keyStr, security.WithIsToHex(false)))
    fmt.Println(security.HmacHash(security.HmacHashTypeSha256, srcStr, keyStr, security.WithIsToHex(true)))
}
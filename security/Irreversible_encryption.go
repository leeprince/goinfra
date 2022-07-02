package security

import (
    "crypto/hmac"
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "encoding/hex"
    "github.com/leeprince/goinfra/code"
    "hash"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/28 下午10:26
 * @Desc:   不可逆加密: 对应函数称为：散列函数/hash函数/哈希函数，如：MD5、SHA、HMAC 等
 */

func Md5(src string, opts ...OptionFunc) string {
    h := md5.New()
    h.Write([]byte(src))
    
    opt := initOption(opts...)
    if opt.IsToHex {
        // fmt.Printf("%x\n", h.Sum(nil)) // 通过 fmt 转成十六进制
        // fmt.Printf("%s\n", hex.EncodeToString(h.Sum(nil))) // 通过 hex 转成十六进制
        return hex.EncodeToString(h.Sum(nil))
    }
    return string(h.Sum(nil))
}

func Sha256(src string, opts ...OptionFunc) string {
    h := sha256.New()
    h.Write([]byte(src))
    
    opt := initOption(opts...)
    if opt.IsToHex {
        return hex.EncodeToString(h.Sum(nil))
    }
    return string(h.Sum(nil))
}

type HmacHashType int32
const (
    HmacHashTypeMd5 HmacHashType = iota
    HmacHashTypeSha1
    HmacHashTypeSha256
)
// Hmac+Hash函数
func HmacHash(hashType HmacHashType, src, key string, opts ...OptionFunc) (string, error) {
    var h hash.Hash
    switch hashType {
    case HmacHashTypeMd5:
        h = hmac.New(md5.New, []byte(key))
    case HmacHashTypeSha1:
        h = hmac.New(sha1.New, []byte(key))
    case HmacHashTypeSha256:
        h = hmac.New(sha256.New, []byte(key))
    default:
        return "", code.BizErrNoExistType
    }
    
    h.Write([]byte(src))
    
    opt := initOption(opts...)
    if opt.IsToHex {
        return hex.EncodeToString(h.Sum(nil)), nil
    }
    return string(h.Sum(nil)), nil
}
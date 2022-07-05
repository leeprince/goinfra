package security

import (
    "crypto/hmac"
    "crypto/md5"
    "crypto/sha1"
    "crypto/sha256"
    "github.com/leeprince/goinfra/code"
    "golang.org/x/crypto/bcrypt"
    "hash"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/28 下午10:26
 * @Desc:   不可逆加密: 对应函数称为：散列函数/hash函数/哈希函数，如：MD5、SHA、HMAC 等
 */

// MD5:
//  encryptOpts.isToHex: 默认转十六进制
func MD5(src string, opts ...OptionFunc) string {
    h := md5.New()
    h.Write([]byte(src))
    
    opt := initOption(opts...)
    return output(h.Sum(nil), opt.outputType)
}

// SHA256:
//  encryptOpts.isToHex: 默认转十六进制
func SHA256(src string, opts ...OptionFunc) string {
    h := sha256.New()
    h.Write([]byte(src))
    
    opt := initOption(opts...)
    return output(h.Sum(nil), opt.outputType)
}

// Hmac+Hash函数
//  encryptOpts.isToHex: 默认转十六进制
func HMACHash(src, key string, opts ...OptionFunc) (string, error) {
    opt := initOption(opts...)
    
    var h hash.Hash
    switch opt.hmacHashType {
    case HMACHashTypeMd5:
        h = hmac.New(md5.New, []byte(key))
    case HMACHashTypeSha1:
        h = hmac.New(sha1.New, []byte(key))
    case HMACHashTypeSha256:
        h = hmac.New(sha256.New, []byte(key))
    default:
        return "", code.BizErrNoExistType
    }
    
    h.Write([]byte(src))
    
    return output(h.Sum(nil), opt.outputType), nil
}

func Bcrypt(src string, opts ...OptionFunc) (string, error){
    opt := initOption(opts...)
    
    cryptByte, err := bcrypt.GenerateFromPassword([]byte(src), opt.bcryptCost)
    if err != nil {
    
    }
    return output(cryptByte, opt.outputType), nil
}

func BcryptVerify(src, dest string, opts ...OptionFunc) bool {
    opt := initOption(opts...)
    destByte, err := input(dest, opt.inputType)
    if err == nil {
        err = bcrypt.CompareHashAndPassword(destByte, []byte(src))
        return err == nil
    }
    return false
}
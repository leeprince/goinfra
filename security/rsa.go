package security

import (
    "bytes"
    "crypto"
    "crypto/rand"
    "crypto/rsa"
    "crypto/sha256"
    "crypto/x509"
    "encoding/pem"
    "errors"
    "github.com/leeprince/goinfra/code"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/7/3 下午6:41
 * @Desc:   非对称加密，加密与解密的密钥不相同，如：RSA
 *              - 公钥加密;私钥解密
 *              - 公钥解密;私钥加密
 */

// RSA加密:公钥加密
//  key: 公钥
//  encryptOpts.isToHex: 默认是转十六进制
func RSAEncrypt(src, publicKey string, opts ...OptionFunc) (string, error) {
    // 解密pem格式的公钥
    block, _ := pem.Decode([]byte(publicKey))
    if block == nil {
        return "", code.BizErrEncrypt.WithError(errors.New("block == nil "))
    }
    // 解析公钥
    pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return "", code.BizErrEncrypt.WithError(err, "ParsePKIXPublicKey")
    }
    // 类型断⾔
    pub, ok := pubInterface.(*rsa.PublicKey)
    if !ok {
        return "", code.BizErrTypeAsserts
    }
    
    // EncryptPKCS1v15加密:兼容长文本加密
    cryptByte, err := compatibleEncryptPKCS1v15([]byte(src), pub)
    if err != nil {
        return "", err
    }
    
    opt := initOption(opts...)
    return output(cryptByte, opt.outputType), nil
}

// EncryptPKCS1v15加密:兼容长文本加密
//  - 兼容 `len(srcByte) > *rsa.PublicKey.Size()-11` 的情况。分组处理长文本，避免报`rsa.ErrMessageTooLong`的错误
func compatibleEncryptPKCS1v15(srcByte []byte, pubKey *rsa.PublicKey) (crypted []byte, err error) {
    srcSize := len(srcByte)
    keySize := pubKey.Size()
    
    // `srcSize <= keySize-11`的情况正常处理
    if srcSize <= keySize-11 {
        // 加密
        crypted, err = rsa.EncryptPKCS1v15(rand.Reader, pubKey, srcByte)
        if err != nil {
            return nil, code.BizErrEncrypt.WithError(err, "EncryptPKCS1v15")
        }
        return
    }
    
    // 兼容 `len(srcByte) > *rsa.PublicKey.Size()-11` 的情况处理长文本，否则会报`rsa.ErrMessageTooLong`的错误
    offset, once := 0, keySize-11
    buffer := bytes.Buffer{}
    for offset < srcSize {
        endindex := offset + once
        if endindex > srcSize {
            endindex = srcSize
        }
        byteOnce, err := rsa.EncryptPKCS1v15(rand.Reader, pubKey, srcByte[offset:endindex])
        if err != nil {
            return nil, code.BizErrEncrypt.WithError(err, "EncryptPKCS1v15 of long text")
        }
        buffer.Write(byteOnce)
        offset = endindex
    }
    crypted = buffer.Bytes()
    
    return
}

// RSA解密:私钥解密
//  crypt: 默认是十六进制字符串
//  key: 私钥
//  encryptOpts.isToHex: 默认是转十六进制
func RSADecrypt(crypt, privateKey string, opts ...OptionFunc) (string, error) {
    opt := initOption(opts...)
    srcByte, err := input(crypt, opt.inputType)
    if err != nil {
        return "", code.BizErrDecrypt.WithError(err)
    }
    if len(srcByte) == 0 {
        return "", code.BizErrDecrypt.WithError(code.BizErrLen)
    }
    
    // 解密
    block, _ := pem.Decode([]byte(privateKey))
    if block == nil {
        return "", code.BizErrDecrypt.WithError(errors.New("block == nil "))
    }
    // 解析PKCS1格式的私钥
    priv, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        return "", code.BizErrDecrypt.WithError(err, "ParsePKCS1PrivateKey")
    }
    
    // DecryptPKCS1v15解密:兼容长文本解密
    decryptByte, err := compatibleDecryptPKCS1v15(srcByte, priv)
    if err != nil {
        return "", err
    }
    
    return string(decryptByte), nil
}

// DecryptPKCS1v15解密:兼容长文本解密
//  - 兼容 `len(srcByte) > *rsa.PrivateKey.Size()` 的情况。分组处理长文本，避免报`rsa.ErrDecryption`的错误
func compatibleDecryptPKCS1v15(cryptByte []byte, privKey *rsa.PrivateKey, opts ...OptionFunc) (decryptByte []byte, err error) {
    srcSize := len(cryptByte)
    keySize := privKey.Size()
    
    // `len(srcByte) <= *rsa.PrivateKey.Size()`的情况正常处理
    if srcSize <= keySize {
        decryptByte, err = rsa.DecryptPKCS1v15(rand.Reader, privKey, cryptByte)
        if err != nil {
            return nil, code.BizErrDecrypt.WithError(err, "DecryptPKCS1v15")
        }
        return
    }
    
    // - 兼容 `len(srcByte) > *rsa.PrivateKey.Size()` 的情况。分组处理长文本，避免报`rsa.ErrDecryption`的错误
    offset := 0
    buffet := bytes.Buffer{}
    for offset < srcSize {
        endIndex := offset + keySize
        if endIndex > srcSize {
            endIndex = srcSize
        }
        byteOnce, err := rsa.DecryptPKCS1v15(rand.Reader, privKey, cryptByte[offset:endIndex])
        if err != nil {
            return nil, code.BizErrDecrypt.WithError(err, "DecryptPKCS1v15 long text")
        }
        buffet.Write(byteOnce)
        offset = endIndex
    }
    return buffet.Bytes(), nil
}

// 生成RSA公/私钥匙
//  - bits: 生成密钥的位数。如：1024、2048
func GenerateRsaKey(bits int) (privKey string, pubKey string, err error) {
    // 生成私钥
    privateKey, err := rsa.GenerateKey(rand.Reader, bits)
    if err != nil {
        err = code.BizErrGenerateData.WithError(err, "GenerateKey")
        return
    }
    derStream := x509.MarshalPKCS1PrivateKey(privateKey)
    block := &pem.Block{
        Type:  "RSA PRIVATE KEY",
        Bytes: derStream,
    }
    prvKeyByte := pem.EncodeToMemory(block)
    privKey = string(prvKeyByte)
    
    // 生成公钥
    publicKey := &privateKey.PublicKey
    derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
    if err != nil {
        err = code.BizErrGenerateData.WithError(err, "MarshalPKIXPublicKey")
        return
    }
    block = &pem.Block{
        Type:  "PUBLIC KEY",
        Bytes: derPkix,
    }
    pubKeyByte := pem.EncodeToMemory(block)
    pubKey = string(pubKeyByte)
    
    return
}

// RSA+SHA256 签名
func RSASignWithSHA256(src, privKey string, opts ...OptionFunc) (string, error) {
    h := sha256.New()
    h.Write([]byte(src))
    hashed := h.Sum(nil)
    block, _ := pem.Decode([]byte(privKey))
    if block == nil {
        return "", code.BizErrSign.WithError(errors.New("block == nil"))
    }
    privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
    if err != nil {
        return "", code.BizErrSign.WithError(err, "ParsePKCS1PrivateKey")
    }
    
    signByte, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hashed)
    if err != nil {
        return "", code.BizErrSign.WithError(err, "SignPKCS1v15")
    }
    
    opt := initOption(opts...)
    return output(signByte, opt.outputType), nil
}

// RSA+SHA256 签名验证: 公钥验证签名
func RSASignVerifyWithSha256(src, sign, publicKey string, opts ...OptionFunc) (bool, error) {
    opt := initOption(opts...)
    signByte, err := input(sign, opt.inputType)
    if err != nil {
        return false, code.BizErrVerifySign.WithError(err, "input")
    }
    if len(signByte) == 0 {
        return false, code.BizErrVerifySign.WithError(code.BizErrLen)
    }
    
    block, _ := pem.Decode([]byte(publicKey))
    if block == nil {
        return false, code.BizErrVerifySign.WithError(errors.New("block == nil"))
    }
    pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
    if err != nil {
        return false, code.BizErrVerifySign.WithError(err)
    }
    
    hashed := sha256.Sum256([]byte(src))
    err = rsa.VerifyPKCS1v15(pubInterface.(*rsa.PublicKey), crypto.SHA256, hashed[:], signByte)
    if err != nil {
        return false, code.BizErrVerifySign.WithError(err)
    }
    return true, nil
}

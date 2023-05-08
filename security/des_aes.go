package security

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"errors"
	"github.com/leeprince/goinfra/perror"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/29 上午12:37
 * @Desc:   对称加密:加密与解密的密钥相同，如：DES、AES
 */

// DES密码函数：密码分组函数对应密钥字节数进行匹配
const (
	cipherDESKeyLen       = 8
	tripleDESCipherKeyLen = 24
)

// --- DES ------------------------------------------------------------------------------------
// DES加密
//
//	key: 密钥字节数必须等于8||24。根据密钥字节数匹配具体的密码函数：密码分组函数
//	encryptOpts.isToHex: 默认转十六进制
func DESEncrypt(src, key string, opts ...OptionFunc) (string, error) {
	srcByte := []byte(src)
	keyByte := []byte(key)
	// 密码函数：密码分组函数
	blockFunc, err := desSwitchKeyLenGetCipherBlock(keyByte)
	if err != nil {
		return "", perror.BizErrEncrypt.WithError(err)
	}
	bs := blockFunc.BlockSize()
	srcByte = desZeroPadding(srcByte, bs)
	if len(srcByte)%bs != 0 {
		return "", perror.BizErrDecrypt.WithError(errors.New("len(srcByte)%bs != 0"))
	}
	cryptByte := make([]byte, len(srcByte))
	dst := cryptByte
	for len(srcByte) > 0 {
		blockFunc.Encrypt(dst, srcByte[:bs])
		srcByte = srcByte[bs:]
		dst = dst[bs:]
	}

	opt := initOption(opts...)
	return output(cryptByte, opt.outputType), nil
}

// DES解密
//
//	crypt: 默认是十六进制字符串
//	key: 密钥字节数必须等于8||24。。根据key的字节数对应到不同的`密码函数：密码分组函数`
//	encryptOpts.isToHex: 默认传入的decrypted是转十六进制
func DESDecrypt(crypt, key string, opts ...OptionFunc) (string, error) {
	opt := initOption(opts...)
	srcByte, err := input(crypt, opt.inputType)
	if err != nil {
		return "", perror.BizErrDecrypt.WithError(err)
	}
	if len(srcByte) == 0 {
		return "", perror.BizErrDecrypt.WithError(perror.BizErrLen)
	}

	keyByte := []byte(key)
	// 密码函数：密码分组函数
	blockFunc, err := desSwitchKeyLenGetCipherBlock(keyByte)
	if err != nil {
		return "", perror.BizErrDecrypt.WithError(err)
	}
	decryptByte := make([]byte, len(srcByte))
	dst := decryptByte
	bs := blockFunc.BlockSize()
	if len(srcByte)%bs != 0 {
		return "", perror.BizErrDecrypt.WithError(errors.New("len(srcByte)%bs != 0"))
	}
	for len(srcByte) > 0 {
		blockFunc.Decrypt(dst, srcByte[:bs])
		srcByte = srcByte[bs:]
		dst = dst[bs:]
	}
	decryptByte = desZeroUnPadding(decryptByte)

	return string(decryptByte), nil
}

// DES填充函数
func desZeroPadding(srcByte []byte, blockSize int) []byte {
	padding := blockSize - len(srcByte)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(srcByte, padtext...)
}

// DES去填充函数
func desZeroUnPadding(decryptByte []byte) []byte {
	return bytes.TrimFunc(decryptByte,
		func(r rune) bool {
			return r == rune(0)
		})
}

// 根据密钥字节数匹配具体的密码函数：密码分组函数
func desSwitchKeyLenGetCipherBlock(key []byte) (blockFunc cipher.Block, err error) {
	switch len(key) {
	case cipherDESKeyLen:
		blockFunc, err = des.NewCipher(key)
	case tripleDESCipherKeyLen:
		blockFunc, err = des.NewTripleDESCipher(key)
	default:
		err = perror.BizErrLen
	}
	return
}

// --- DES-end ------------------------------------------------------------------------------------

// --- AES ------------------------------------------------------------------------------------
// AES加密
//
//	src: 明文
//	key: 密钥字节数必须是：16（AES-128）|| 24（AES-192）|| 32（AES-256）。根据key的字节数对应到不同的`密码函数：密码分组函数`
//	可选项
//	    opt.aesBlockModeType: 默认AESBlockModeTypeCBC
//	    opt.aesIV: 默认填充"0000000000000000"
func AESEncrypt(src, key string, opts ...OptionFunc) (string, error) {
	opt := initOption(opts...)

	if src == "" {
		return "", perror.BizErrDataEmpty
	}

	keyByte := []byte(key)
	// 密码函数：密码分组函数
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", perror.BizErrEncrypt.WithError(err)
	}

	var cryptByte []byte
	switch opt.aesBlockModeType {
	case AESBlockModeTypeCBC:
		// 加密模式
		blockModeCBCFunc := cipher.NewCBCEncrypter(block, []byte(opt.aesIV))
		content := []byte(src)
		// 填充
		content = aesPKCS5Padding(content, block.BlockSize())
		cryptByte = make([]byte, len(content))
		// 执行加密
		blockModeCBCFunc.CryptBlocks(cryptByte, content)
	default:
		return "", perror.BizErrEncrypt.WithError(perror.BizErrNoExistType)
	}

	return output(cryptByte, opt.outputType), nil
}

// AES解密
//
//	crypt: 默认是十六进制字符串
//	key: 密钥字节数必须是：16（AES-128）|| 24（AES-192）|| 32（AES-256）
//	可选项
//	    opt.aesBlockModeType: 默认AESBlockModeTypeCBC
//	    opt.aesIV: 默认填充"0000000000000000
func AESDecrypt(crypt, key string, opts ...OptionFunc) (string, error) {
	opt := initOption(opts...)
	srcByte, err := input(crypt, opt.inputType)
	if err != nil {
		return "", perror.BizErrDecrypt.WithError(err)
	}
	if len(srcByte) == 0 {
		return "", perror.BizErrDecrypt.WithError(perror.BizErrLen)
	}

	keyByte := []byte(key)
	// 密码函数：密码分组函数
	block, err := aes.NewCipher(keyByte)
	if err != nil {
		return "", perror.BizErrDecrypt.WithError(err)
	}

	var decryptByte []byte
	switch opt.aesBlockModeType {
	case AESBlockModeTypeCBC:
		// 解密模式
		blockModeCBCFunc := cipher.NewCBCDecrypter(block, []byte(opt.aesIV))
		decryptByte = make([]byte, len(srcByte))
		blockModeCBCFunc.CryptBlocks(decryptByte, srcByte)
	default:
		return "", perror.BizErrEncrypt.WithError(perror.BizErrNoExistType)
	}

	// 取出
	return string(aesPKCS5UnPadding(decryptByte)), nil
}

// AES填充函数
func aesPKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// AES去填充函数
func aesPKCS5UnPadding(encrypt []byte) []byte {
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}

// --- AES-end ------------------------------------------------------------------------------------

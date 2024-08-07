package security

import (
	"bytes"
	"crypto/cipher"
	"github.com/leeprince/goinfra/perror"
	"github.com/tjfoc/gmsm/sm4"
)

// SM4加密
//   - key: 字节长度只能等于16
func SM4Encrypt(src, key string, opts ...OptionFunc) (string, error) {
	keyBytes := []byte(key)
	dataBytes := []byte(src)
	
	iv := make([]byte, sm4.BlockSize)
	
	block, err := sm4.NewCipher(keyBytes)
	if err != nil {
		return "", perror.BizErrSecurityEncrypt.WithError(err, "NewCipher")
	}
	blockSize := block.BlockSize()
	origData := pkcs5Padding(dataBytes, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, iv)
	cryptByte := make([]byte, len(origData))
	blockMode.CryptBlocks(cryptByte, origData)
	
	opt := initOption(opts...)
	return OutputFormat(cryptByte, opt.outputType), nil
}

func SM4Decrypt(crypt, key string, opts ...OptionFunc) (string, error) {
	opt := initOption(opts...)
	srcByte, err := InputFormat(crypt, opt.inputType)
	if err != nil {
		return "", perror.BizErrSecurityDecrypt.WithError(err)
	}
	if len(srcByte) == 0 {
		return "", perror.BizErrSecurityDecrypt.WithError(perror.BizErrLen)
	}
	
	keyByte := []byte(key)
	
	iv := make([]byte, sm4.BlockSize)
	block, err := sm4.NewCipher(keyByte)
	if err != nil {
		return "", err
	}
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(srcByte))
	blockMode.CryptBlocks(origData, srcByte)
	decryptByte := pkcs5UnPadding(origData)
	return string(decryptByte), nil
}

// pkcs5填充
func pkcs5Padding(src []byte, blockSize int) []byte {
	padding := blockSize - len(src)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(src, padtext...)
}

func pkcs5UnPadding(src []byte) []byte {
	length := len(src)
	if length == 0 {
		return nil
	}
	unpadding := int(src[length-1])
	return src[:(length - unpadding)]
}

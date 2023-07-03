package characterutil

import (
	"encoding/base64"
	"errors"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/22 15:40
 * @Desc:
 */

func StringToBase64Str(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64StrToString(base64Str string) (string, error) {
	s, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return "", errors.New("base64解码错误：" + err.Error())
	}
	return string(s), nil
}

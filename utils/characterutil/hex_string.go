package characterutil

import (
	"encoding/hex"
	"errors"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/22 15:23
 * @Desc:
 */

func StringToHexStr(s string) string {
	return hex.EncodeToString([]byte(s))
}

func HexStrToString(hexStr string) (string, error) {
	decodedBytes, err := hex.DecodeString(hexStr)
	if err != nil {
		return "", errors.New("解码错误：" + err.Error())
	}
	return string(decodedBytes), nil
}

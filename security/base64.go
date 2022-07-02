package security

import "encoding/base64"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/29 上午12:06
 * @Desc:
 */

func Base64Encode(src string) string {
    return string(base64.StdEncoding.EncodeToString([]byte(src)))
}

func Base64Decode(src string) (string, error) {
    a, err := base64.StdEncoding.DecodeString(src)
    if err != nil {
        return "", err
    }
    return string(a), nil
}
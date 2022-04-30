package httpcli

import (
    "net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/3/27 下午11:59
 * @Desc:
 */

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/5 下午4:18
 * @Desc:   http 客户端
 */

func Do(req *http.Request) (respByte []byte, err error) {
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return
    }
    return ToBytes(resp)
}
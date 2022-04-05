package httpcli

import (
    "fmt"
    "io/ioutil"
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
 * @Desc:
 */

func Do(req *http.Request) (respByte []byte, err error) {
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return
    }
    defer resp.Body.Close()
    
    respByte, err = ioutil.ReadAll(resp.Body)
    if err != nil {
        return
    }
    
    if resp.StatusCode  != http.StatusOK {
        err = fmt.Errorf("StatusCode: %d, Body: %s", resp.StatusCode, respByte)
        return
    }
    return
}
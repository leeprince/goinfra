package httpcli

import (
	"fmt"
	"io"
	"net/http"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/4/21 下午2:46
 * @Desc:   http 响应（*http.Response）处理
 */

func ResponseToBytes(resp *http.Response) (respByte []byte, err error) {
	defer resp.Body.Close()
	
	respByte, err = io.ReadAll(resp.Body)
	if err != nil {
		return
	}
	
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("StatusCode: %d, requestBody: %s", resp.StatusCode, respByte)
		return
	}
	return
}

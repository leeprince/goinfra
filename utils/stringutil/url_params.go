package stringutil

import (
	"net/url"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/13 14:13
 * @Desc:
 */

// 不直接解析 `optype=BookingOrderForAli&ExistOrderNumbers=1` 而是通过构造
/*
在前面拼接上 http://localhost? 并通过 url.ParseRequestURI 和 .Query() 会比使用 for 循环析查询字符串的性能更好。
这是因为 url.ParseRequestURI 函数会对 URL 进行解析和规范化，以确保 URL 的正确性和一致性。而使用for 循环解析查询字符串则需要手动处理各种边界情况，例如转义字符、空值等等，这会增加代码的复杂性和运行时间。
另外，使用 url.ParseRequestURI 和 .Query() 函数可以更好地利用 golang 内置的优化和缓存机制，从而提高代码的性能和稳定性。
因此，如果性能是一个关键因素，建议使用 urlRequestURI 和 .Query() 函数来解析查询字符串。
*/

// @param urlString 需要构造成含 host 的完整 url。如:"http://localhost?optype=BookingOrderForAli&ExistOrderNumbers=1"
func GetUrlStrParam(urlString, key string) (string, error) {
	u, err := url.ParseRequestURI(urlString)
	if err != nil {
		return "", err
	}
	query := u.Query()
	return query.Get(key), nil
}

// @param urlString 需要构造成含 host 的完整 url。如:"http://localhost?a=1&b=hello"
func GetUrlStrParams(urlString string, keys []string) (map[string]string, error) {
	u, err := url.ParseRequestURI(urlString)
	if err != nil {
		return nil, err
	}
	query := u.Query()

	data := make(map[string]string, len(keys))
	for _, k := range keys {
		data[k] = query.Get(k)
	}
	return data, nil
}

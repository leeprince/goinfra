package perror

import (
	"errors"
	"regexp"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/13 18:48
 * @Desc:
 */

var dmReg *regexp.Regexp

func init() {
	dmReg = regexp.MustCompile("(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]")
}

// 替换报错内容中包含IP地址
func ReplaceIPErr(err error) error {
	if err == nil {
		return nil
	}
	errMsg := err.Error()
	return errors.New(dmReg.ReplaceAllLiteralString(errMsg, "#.#.#.#"))
}

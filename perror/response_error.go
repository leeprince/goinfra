package perror

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"sync"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/13 18:48
 * @Desc:
 */

var dmReg *regexp.Regexp
var replaceIPErrOnce sync.Once

// 替换报错内容中包含IP地址
func ReplaceIPErr(err error) error {
	replaceIPErrOnce.Do(func() {
		dmReg = regexp.MustCompile("(?:[a-z0-9](?:[a-z0-9-]{0,61}[a-z0-9])?\\.)+[a-z0-9][a-z0-9-]{0,61}[a-z0-9]")
	})
	if err == nil {
		return nil
	}
	errMsg := err.Error()
	return errors.New(dmReg.ReplaceAllLiteralString(errMsg, "#.#.#.#"))
}

func ErrPanic(err error) {
	if err != nil {
		panic(fmt.Sprintf("ErrPanic:%+v", err))
	}
}

func ErrLogFatal(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

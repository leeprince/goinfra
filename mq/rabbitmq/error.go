package rabbitmq

import (
	"fmt"
	"github.com/leeprince/goinfra/plog"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/5/12 上午10:43
 * @Desc:   错误处理
 */

func failOnError(err error, msg string) {
	if err != nil {
		plog.Errorf("failOnError. msg:%s > err:%s", msg, err)
		return
	}
	plog.Infof("failOnError. msg:%s > but not err", msg)
}

func panicOnError(err error, msg string) {
	if err != nil {
		println(fmt.Sprintf("msg:%s, err:%s", msg, err.Error()))
		plog.Fatalf("%s: %v", msg, err)
	}
}

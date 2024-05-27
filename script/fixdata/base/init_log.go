package base

import (
	"github.com/leeprince/goinfra/plog"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2024/5/27 17:13
 * @Desc:
 */

func InitLog() {
	plog.NewDefaultLogger()
}

package runtimeutil

import "runtime"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/4/4 10:38
 * @Desc:
 */

func CallersFunc(n int) string {
	pc := make([]uintptr, 1)
	runtime.Callers(n, pc)
	return runtime.FuncForPC(pc[0]).Name()
}

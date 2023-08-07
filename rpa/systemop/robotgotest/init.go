package main

import "runtime"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/4 00:31
 * @Desc:
 */

var (
	ACTIVE_NAME = "sublime_text"
)

func init() {
	if runtime.GOOS == "windows" {
		ACTIVE_NAME = ACTIVE_NAME + ".exe"
	}
}

package robofigcrontest

import "testing"

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/21 10:40
 * @Desc:
 */

func Test_taskRunningDefault(t *testing.T) {
	taskRunningDefault()
}

func Test_taskRunningSkipIfStillRunning(t *testing.T) {
	taskRunningSkipIfStillRunning()
}

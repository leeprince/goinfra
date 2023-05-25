package timeutil

import (
	"time"
)

/**
 * @Author: prince.lee
 * @Date:   2022/3/24 17:12
 * @Desc:   时间
 */

// 是否使用毫秒为单位
//
//	小于秒的单位时间则使用毫秒
func UseMillisecondUnit(dur time.Duration) bool {
	return dur < time.Second || dur%time.Second != 0
}

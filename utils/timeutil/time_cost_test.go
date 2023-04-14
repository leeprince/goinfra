package timeutil

import (
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/7/2 下午3:21
 * @Desc:
 */

func TestNewTimeCost(t *testing.T) {
	newTimeCost := NewTimeCost("开始")

	time.Sleep(time.Millisecond * 20)
	newTimeCost.Duration("-")

	time.Sleep(time.Millisecond * 20)
	newTimeCost.Duration("-")

	time.Sleep(time.Millisecond * 20)
	newTimeCost.Stop("-")
}

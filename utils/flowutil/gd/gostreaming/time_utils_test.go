package gostreaming_test

/*
 * @Date: 2020-07-17 10:20:35
 * @LastEditors: aiden.deng (Zhenpeng Deng)
 * @LastEditTime: 2020-07-23 15:09:27
 */

import (
	"fmt"
	"testing"

	"gitlab.yewifi.com/risk-control/risk-common/pkg/gostreaming"
)

func TestTimeUtils_GetTodayLatestTime(t *testing.T) {
	latestTime, err := gostreaming.GetTodayLatestTime()
	if err != nil {
		t.Errorf("get error: %s", err)
		return
	}
	fmt.Println(latestTime)
	fmt.Println(latestTime.Unix())
	fmt.Println(latestTime.Format("2006-01-02 15:04:05"))

	// 如果当天是2020-07-17，那么会得到以下输出

	// 2020-07-17 23:59:59 +0800 CST
	// 1595001599
	// 2020-07-17 23:59:59
}

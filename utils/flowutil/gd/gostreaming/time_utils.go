package gostreaming

/*
 * @Date: 2020-07-17 10:17:29
 * @LastEditors: aiden.deng(Zhenpeng Deng)
 * @LastEditTime: 2020-07-17 10:24:21
 */

import (
	"fmt"
	"time"
)

// 获取当天 23:59:59 的时间戳
func GetTodayLatestTime() (*time.Time, error) {
	now := time.Now()
	nowDateStr := now.Format("2006-01-02")
	todayLatestTime, err := time.ParseInLocation("2006-01-02 15:04:05", fmt.Sprintf("%s 23:59:59", nowDateStr), time.Local)
	if err != nil {
		return nil, err
	}
	return &todayLatestTime, err
}

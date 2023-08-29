package timeutil

import (
	"fmt"
	"testing"
	"time"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/5/20 11:50
 * @Desc:
 */

func TestFormat(t *testing.T) {
	var ts string
	ts = Year()
	fmt.Println("Year", ts)

	ts = Month()
	fmt.Println("Month", ts)

	ts = Data()
	fmt.Println("Data", ts)

	ts = DataTime()
	fmt.Println("DataTime", ts)

	ts = DataTimeT(time.Now())
	fmt.Println("DataTimeT", ts)

	ts = DataTimeMicrosecond()
	fmt.Println("DataTimeMicrosecond", ts)

	ts = DataTimeMillisecond()
	fmt.Println("DataTimeMillisecond", ts)

	ts = DataTimeDataSecond()
	fmt.Println("DataTimeDataSecond", ts)

}

func TestMonthNumUnixTime(t *testing.T) {
	month := MonthNumUnixTime(time.Now().Unix())
	fmt.Println(month)
}

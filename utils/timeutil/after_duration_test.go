package timeutil

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/23 14:52
 * @Desc:
 */

func TestAfterSecondD(t *testing.T) {
	gotD, err := AfterSecondD(1)
	fmt.Println(gotD.String(), err)
}

func TestAfterMinuteD(t *testing.T) {
	gotD, err := AfterMinuteD(1)
	fmt.Println(gotD.String(), err)
}

func TestAferHoursD(t *testing.T) {
	gotD, err := AfterHoursD(1)
	fmt.Println(gotD.String(), err)
}

func TestAfterDayD(t *testing.T) {
	gotD, err := AfterDayD(1)
	fmt.Println(gotD.String(), err)
}

package testdata

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/26 14:49
 * @Desc:
 */

func TestArrParams(t *testing.T) {
	fmt.Println("-01", "-02", "-03")
	fmt.Println("-1", "-2", "-3")

	Warn("a1", "a2", "a3")

	fmt.Println("---------------")
	Warn()
}

func Warn(v ...interface{}) {
	fmt.Println(v...)

	// 错误写法，会发生报错：too many arguments in call to fmt.Println。
	// 应该使用：fmt.Println("aa", v)的写法
	//fmt.Println("aa", v...)

	fmt.Println("aa", v)

	v = append(v, "bb")
	fmt.Println(v...)
}

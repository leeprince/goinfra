package test

import (
	"fmt"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/10/26 11:59
 * @Desc:
 */

func TestType(t *testing.T) {
	var intf interface{}

	intf = string("leeprince")

	i, err := intf.(string)
	fmt.Println(i, err)

	ii := intf.(string)
	fmt.Println(ii)

	// -------
	fmt.Println("---------------")

	intf = 1
	i1, err := intf.(string)
	fmt.Println(i1, err)

	// 会抛出异常
	ii1 := intf.(string)
	fmt.Println(ii1)
}

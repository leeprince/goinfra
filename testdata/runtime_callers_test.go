package testdata

import (
	"fmt"
	"github.com/leeprince/goinfra/utils/runtimeutil"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/8/26 23:15
 * @Desc:
 */

func TestCallersFunc(t *testing.T) {
	var s string

	s = runtimeutil.CallersFunc(0)
	fmt.Printf("CallersFunc 0: %+v \n", s)

	s = runtimeutil.CallersFunc(1)
	fmt.Printf("CallersFunc 1: %+v \n", s)

	s = runtimeutil.CallersFunc(2)
	fmt.Printf("CallersFunc 2: %+v \n", s)

	s = runtimeutil.CallersFunc(3)
	fmt.Printf("CallersFunc 3: %+v \n", s)

	s = runtimeutil.CallersFunc(4)
	fmt.Printf("CallersFunc 3: %+v \n", s)

}

package testdata

import (
	"fmt"
	"github.com/leeprince/goinfra/utils/dumputil"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/9/13 10:48
 * @Desc:
 */

func TestParams(t *testing.T) {
	type TT struct {
		i *int
		s *string
	}

	i := 1
	s := "s1"

	tt := &TT{
		i: &i,
		s: &s,
	}
	fmt.Printf("tt 1: %+v\n", tt)
	dumputil.Println(tt)

	i2 := 2
	s2 := "s2"
	tt.i = &i2
	*tt.s = s2
	fmt.Printf("tt 2: %+v\n", tt)
	dumputil.Println(tt)

	// 直接复制tt后，tt2的所有修改也都会影响到tt，因为赋值给tt2的是指针地址。如果需要不影响，需要取出tt的值而非指针地址
	tt2 := tt
	i21 := 21
	s21 := "s21"
	tt2.i = &i21
	*tt2.s = s21
	dumputil.Println(tt)
	dumputil.Println(tt2)
	dumputil.Println(tt)

	// tt 赋值后不被新变量影响
	tt3 := *tt
	i22 := 22
	s22 := "s22"
	tt3.i = &i22
	*tt3.s = s22
	dumputil.Println(tt)
	dumputil.Println(tt3)

}

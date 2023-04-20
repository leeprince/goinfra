package stringer_test_test

import (
	"fmt"
	"github.com/leeprince/goinfra/utils/stringutil/stringer_test"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2022/6/27 下午2:39
 * @Desc:   测试 PrinceTypeInt
 */

// Type describes the type of the data Value holds.
type PrinceTypeInt int

const (
	// INVALID is used for a Value with no value set.
	INVALID PrinceTypeInt = iota
	// BOOL is a boolean Type Value.
	BOOL
	// INT64 is a 64-bit signed integral Type Value.
	INT64
	// FLOAT64 is a 64-bit floating point Type Value.
	FLOAT64
	// STRING is a string Type Value.
	STRING
	// BOOLSLICE is a slice of booleans Type Value.
	BOOLSLICE
	// INT64SLICE is a slice of 64-bit signed integral numbers Type Value.
	INT64SLICE
	// FLOAT64SLICE is a slice of 64-bit floating point numbers Type Value.
	FLOAT64SLICE
	// STRINGSLICE is a slice of strings Type Value.
	STRINGSLICE
)

func TestPrinceType(t *testing.T) {
	v := BOOL
	fmt.Println(v)

	v1 := stringer_test.BOOL
	fmt.Println(v1)
	v11 := stringer_test.BOOL.String()
	fmt.Println(v11)

	v2 := stringer_test.TestComentHaveOther
	fmt.Println(v2)
	v21 := stringer_test.TestComentHaveOther.String()
	fmt.Println(v21)

	fmt.Println(v)
	v3 := stringer_test.TestNotComment
	fmt.Println(v3)
	v31 := stringer_test.TestNotComment.String()
	fmt.Println(v31)
}

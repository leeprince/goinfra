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

func TestSlice(t *testing.T) {
	i := []int{1, 2, 3}
	Slice(i)
	fmt.Println(i)
	fmt.Println("----Slice")
	
	SliceModify(i)
	fmt.Println(i)
	fmt.Println("----Slice")
	
	ii := []int{1, 2}
	Slice1(&ii)
	fmt.Println(ii)
	fmt.Println("----Slice1")
	
	ii1 := 1
	ii2 := 2
	iii := []*int{&ii1, &ii2}
	Slice2(iii)
	fmt.Println(iii)
	// fmt.Printf("%+v \n", iii)
	// fmt.Printf("%#v \n", iii)
	for _, i3 := range iii {
		fmt.Println("Slice2:", *i3)
	}
	fmt.Println("----Slice2")
	
	// ---
	in := make([]int, 4)
	in = []int{1, 2, 3}
	SliceAppend(in)
	fmt.Println(in)
	fmt.Println("----SliceAppend")
	
	// ---
	incap := make([]int, 4)
	incap[0], incap[1], incap[2] = 0, 1, 2
	SliceCap(incap)
	fmt.Println(incap)
	fmt.Println("----SliceCap")
	
	// ---
	incap1 := []int{0, 1, 2, 0}
	fmt.Println(incap1)
	SliceCap(incap1)
	fmt.Println(incap1)
	fmt.Println("----SliceCap1")
	
	// ---
	incap2 := make([]int, 4)
	incap2[0], incap2[1], incap2[2] = 0, 1, 2
	SliceCapAppend(incap2)
	fmt.Println(incap2)
	fmt.Println("----SliceCapAppend")
	
	// ---
	incap3 := make([]int, 4)
	incap3 = []int{0, 1, 2}
	SliceCapAppend1(incap3)
	fmt.Println(incap3)
	fmt.Println("----SliceCapAppend1")
}
func Slice(i []int) {
	i = append(i, 4)
}
func SliceModify(i []int) {
	i[1] = 0
}
func Slice1(i *[]int) {
	*i = append(*i, 4)
}
func Slice2(i []*int) {
	i4 := 4
	i = append(i, &i4)
}
func SliceAppend(i []int) {
	i = append(i, 4)
}
func SliceCap(i []int) {
	i[3] = 3
}
func SliceCapAppend(i []int) {
	i = append(i, 3)
}
func SliceCapAppend1(i []int) {
	i = append(i, 3)
}

func TestMap(t *testing.T) {
	m := map[int]int{
		0: 0,
		1: 1,
		2: 2,
		3: 3,
	}
	Map(m)
	fmt.Println(m)
}
func Map(m map[int]int) {
	m[4] = 4
}

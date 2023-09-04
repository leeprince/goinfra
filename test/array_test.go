package test

import (
	"fmt"
	"github.com/leeprince/goinfra/utils/dumputil"
	"testing"
)

/**
 * @Author: prince.lee <leeprince@foxmail.com>
 * @Date:   2023/6/23 09:30
 * @Desc:
 */

func TestTwoArray(t *testing.T) {
	type person struct {
		i int
	}

	var personList []person
	var personListList [][]person

	for i := 0; i <= 20; i++ {
		personInfo := person{
			i: i,
		}

		personList = append(personList, personInfo)
		if i != 0 && i%5 == 0 {
			personListList = append(personListList, personList)

			// 初始化下一个数组
			// personList = nil
			personList = []person{}
		}
	}

	dumputil.Println("personListList:%+v", personListList)
}

// 因为golang的函数、方法的所有参数都是传值，所以就是是引用类型类型，需要方法中修改到外部变量时需要传递变量地址
func TestRequestParamHaveArray(t *testing.T) {
	var strArray []string
	var strArrayQute []string

	RequestParamHaveArray(strArray)
	fmt.Println(strArray)

	RequestParamHaveArrayQuote(&strArrayQute)
	fmt.Println(strArrayQute)

	fmt.Println("-------------------")
	str2Array := []string{}
	str2ArrayQute := make([]string, 0)

	RequestParamHaveArray(str2Array)
	fmt.Println(str2Array)

	RequestParamHaveArrayQuote(&str2ArrayQute)
	fmt.Println(str2ArrayQute)
	fmt.Println("-------------------")

	var str3Array []string
	str3ArrayQute := make([]string, 0)

	RequestParamHaveArray(str3Array)
	fmt.Println(str3Array)

	RequestParamHaveArrayQuote(&str3ArrayQute)
	fmt.Println(str3ArrayQute)
}

func RequestParamHaveArray(strArray []string) {
	strArray = append(strArray, "aaa")
}

func RequestParamHaveArrayQuote(strArray *[]string) {
	*strArray = append(*strArray, "aaa")
}

func TestArrCutting(t *testing.T) {
	a1 := 1
	a2 := 2
	a3 := 3
	arr := []*int{&a1, &a2, &a3}

	fmt.Println(arr)

	arrNew := arr[:2]
	fmt.Println(arrNew)
	fmt.Println(arr)

	fmt.Println(arr[:2])
	fmt.Println(arr)
}

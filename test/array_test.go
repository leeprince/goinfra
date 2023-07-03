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

func TestRequestParamHaveArray(t *testing.T) {
	var strArray []string
	
	RequestParamHaveArray(strArray)
	fmt.Println(strArray)
	
	RequestParamHaveArrayQuote(&strArray)
	fmt.Println(strArray)
	
	var str2Array []string
	str2Array = make([]string, 0)
	
	RequestParamHaveArray(str2Array)
	fmt.Println(str2Array)
	
	RequestParamHaveArrayQuote(&str2Array)
	fmt.Println(str2Array)
}

func RequestParamHaveArray(strArray []string) {
	strArray = append(strArray, "aaa")
}

func RequestParamHaveArrayQuote(strArray *[]string) {
	*strArray = append(*strArray, "aaa")
}
